// Copyright 2026 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package devidentity

import (
	"testing"

	authv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	k8stesting "k8s.io/client-go/testing"
)

func TestFleetRequiresReviewedAudienceAndExactBoundServiceAccount(t *testing.T) {
	client := fake.NewSimpleClientset(&corev1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: "fleet", Namespace: "fleet", UID: "sa-1"}}, &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "fleet-pod", Namespace: "fleet", UID: "pod-1"}})
	authority := &KubeAuthority{Core: client.CoreV1(), Auth: client.AuthenticationV1(), FleetAudience: "http://identity.internal/realms/uds", FleetClientEnabled: true, FleetNamespace: "fleet", FleetServiceAccount: "fleet"}
	status := authv1.TokenReviewStatus{Authenticated: true, Audiences: []string{authority.FleetAudience}, User: authv1.UserInfo{Username: "system:serviceaccount:fleet:fleet", UID: "sa-1", Extra: map[string]authv1.ExtraValue{"authentication.kubernetes.io/pod-name": {"fleet-pod"}, "authentication.kubernetes.io/pod-uid": {"pod-1"}}}}
	client.PrependReactor("create", "tokenreviews", func(action k8stesting.Action) (bool, runtime.Object, error) {
		review := action.(k8stesting.CreateAction).GetObject().(*authv1.TokenReview)
		if review.Spec.Token != "real-reviewed-assertion" || len(review.Spec.Audiences) != 1 || review.Spec.Audiences[0] != authority.FleetAudience {
			t.Fatal("TokenReview authority request changed")
		}
		return true, &authv1.TokenReview{Status: status}, nil
	})
	principal, err := authority.Assertion(t.Context(), "real-reviewed-assertion")
	if err != nil || principal.Role != "fleet" {
		t.Fatal(err)
	}
	status.Audiences = []string{"other-audience"}
	if _, err := authority.Assertion(t.Context(), "real-reviewed-assertion"); err == nil {
		t.Fatal("mismatched audience authorized")
	}
	status.Audiences = []string{authority.FleetAudience}
	status.User.Username = "system:serviceaccount:untrusted:fleet"
	if _, err := authority.Assertion(t.Context(), "real-reviewed-assertion"); err == nil {
		t.Fatal("different namespace authorized")
	}
	status.User.Username = "system:serviceaccount:fleet:fleet"
	if err := client.CoreV1().Pods("fleet").Delete(t.Context(), "fleet-pod", metav1.DeleteOptions{}); err != nil {
		t.Fatal(err)
	}
	if err := authority.Validate(t.Context(), principal); err == nil {
		t.Fatal("bound Pod deletion retained API authority")
	}
}

func TestOperatorCredentialsDoNotGrantMasterAuthority(t *testing.T) {
	m, _, _, _ := managementFixture(t)
	principal, err := m.Authority.Operator(t.Context(), "uds-operator", "real-operator-secret")
	if err != nil {
		t.Fatal(err)
	}
	if principal.Role != "operator" {
		t.Fatal("wrong operator role")
	}
	if _, err := m.Authority.Operator(t.Context(), "other-client", "real-operator-secret"); err == nil {
		t.Fatal("secret reused by unrelated client")
	}
	if _, err := m.Authority.Operator(t.Context(), "uds-operator", "wrong-secret"); err == nil {
		t.Fatal("wrong operator secret accepted")
	}
}

func TestProjectedNativeControllerTokenGetsOnlyOperatorAuthority(t *testing.T) {
	client := fake.NewSimpleClientset(&corev1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: "uds-controller", Namespace: "uds-system", UID: "operator-sa-uid"}})
	authority := &KubeAuthority{Core: client.CoreV1(), Auth: client.AuthenticationV1(), FleetAudience: "http://keycloak-http.keycloak.svc.cluster.local/realms/uds", FleetNamespace: "uds-fleet-command", FleetServiceAccount: "uds-fleet-command-sa", OperatorNamespace: "uds-system", OperatorServiceAccount: "uds-controller"}
	client.PrependReactor("create", "tokenreviews", func(action k8stesting.Action) (bool, runtime.Object, error) {
		return true, &authv1.TokenReview{Status: authv1.TokenReviewStatus{Authenticated: true, Audiences: []string{authority.FleetAudience}, User: authv1.UserInfo{Username: "system:serviceaccount:uds-system:uds-controller", UID: "operator-sa-uid"}}}, nil
	})
	principal, err := authority.Assertion(t.Context(), "projected-controller-assertion")
	if err != nil {
		t.Fatal(err)
	}
	if principal.Role != "operator" || principal.CredentialType != "kubernetes" {
		t.Fatal("projected native controller obtained wrong scope")
	}
	if clientPermitted(principal, ClientRecord{Realm: "master", Data: map[string]any{"clientId": "ordinary"}}) {
		t.Fatal("operator assertion obtained master authority")
	}
}
