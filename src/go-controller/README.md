<!--
Copyright 2026 Defense Unicorns
SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial
-->

# UDS Go controller

This development implementation runs the UDS `Package` lifecycle in Go and moves suitable admission behavior into Kubernetes. It extends the Dash Days prototype to the current Core 1.14 contract, including reload, drift recovery, durable cleanup, leader election, and shared webhook trust. The complete development profile also uses a restricted Dex identity provider.

The profile remains experimental while full cluster qualification, cold installation, failure recovery, and performance comparisons continue. Passing selected tests does not establish production readiness. See [migration status](status.md) for the implementation boundary and outstanding qualification.

## Responsibilities

The controller watches current UDS resources and their dependencies. Its elected leader reconciles owned resources and publishes exemption parameters; every replica serves admission from synchronized caches.

| Responsibility | Source |
| --- | --- |
| Keyed `Package` reconciliation, finalizers and retry | `internal/controller/udspackage/` |
| Network policies and API/Node discovery | `internal/controller/network/`, `internal/controller/dependencies/` |
| Istio namespace configuration, ingress and egress | `internal/controller/istio/` |
| UDP exposure through Envoy Gateway | `internal/controller/envoygateway/` |
| Authorization policies and SSO | `internal/controller/authpolicy/`, `internal/controller/sso/` |
| Monitor, uptime probe and CA bundle reconciliation | `internal/controller/monitoring/`, `internal/controller/probes/`, `internal/controller/cabundle/` |
| Secret and ConfigMap reload | `internal/controller/reload/` |
| Native exemption publication | `internal/controller/admission/` |
| Code admission, topology mutation and serving trust | `webhook/`, `internal/admission/` |
| Development identity and protocol probes | `cmd/dev-identity/`, `cmd/identity-probe/`, `internal/devidentity/` |

The [native admission contract](native-admission/README.md) explains policy ownership, JavaScript regex compatibility, grant union and revocation, protected recovery, and certificate rotation. Diagnostic annotations never authorize an exemption.

## Ownership modes

You can use native ownership or the prototype's staged ownership controls. Do not run two active writers for the same phase.

| Setting | Behavior |
| --- | --- |
| `UDS_CONTROLLER_MODE=native` | Go owns all implemented reconciliation phases. |
| `UDS_OPERATOR_*_ENABLED=false` without native mode | The corresponding Pepr phase is disabled and Go owns that phase. |
| Other staged flag values | Go leaves that phase to Pepr. |

The per-phase environment names remain in [feature flags](internal/featureflags/flags.go) and [development flags](dev-flags.env). The Go controller also owns reload and current cluster dependency observation in native mode. Disable and drain the original watcher before assigning that ownership.

## Build and check

Run Go commands from this module. Core's root Node dependencies provide source fixtures and authoring tools; the deployed native controller does not run Node.js.

```bash
cd src/go-controller
GOWORK=off go build -o uds-controller .
GOWORK=off go test -race ./...
GOWORK=off go vet ./...
```

The module pins its Go toolchain in [go.mod](go.mod). Keep immutable source checkpoints and actual image binary hashes with performance results so a build of earlier source cannot masquerade as the current implementation.

## Generate API clients

The generator reads Core's actual `Package` and `ClusterConfig` CRD schemas. `hack/api-layout.json` supplies stable Go names; it does not duplicate the field schema. Do not edit generated API or client output directly.

From the Core repository root, regenerate API types and clients:

```bash
node src/go-controller/hack/generate-api.mjs
```

The entry point runs the client and deepcopy generator through `hack/update-codegen.sh`. Regenerate twice when verifying reproducibility and require identical output. `ClusterConfig` remains cluster-scoped.

## Prepare a development package

The [native package recipe](../../packages/native-dev/prepare.mjs) composes the original Istio and Authservice components with the native controller, Kubernetes policies, and development identity. Its input lock records the exact Core commit, original signed archives, distroless base, and local derivative image digests.

From the repository root, prepare a connected candidate for an explicit isolated cluster:

```bash
node packages/native-dev/prepare.mjs --connected --candidate --render --package
```

Connected mode imports the emitted immutable image archive into that cluster's containerd before deployment. It does not install a development registry or registry agent. Set an explicit absolute `KUBECONFIG` and the trusted bootstrap CLI path. The locally derived archive is unsigned; verified original input signatures do not sign the derivative.

A candidate retains the tenant authoring fence after its readiness gates. The regular recipe releases that fence only after the complete native stack qualifies. Keep candidate and activation artifacts distinct in test receipts.

The package sequence starts a dormant Go backend, verifies it, registers complete callbacks, installs identity, then activates native policy. It checks the real API before exposing tenant gateways. Rendering alone does not establish a successful cold installation.

## Development identity

The Dex path requires `devMode=true` and the explicit password-only profile. It preserves public and administration issuer hosts and real signed OIDC tokens. The compatibility API manages persisted users, groups and clients, and validates bound Kubernetes ServiceAccount credentials.

Unsupported required authentication features fail closed. Do not select this development provider to qualify MFA, X.509 login, SAML, social identity providers, arbitrary protocol mappers, or production password policy. Those capabilities remain outside this provider's contract. The Go SSO implementation can still reconcile the original Keycloak provider.

Reference the [development identity guide](../../docs/dev/native-development-identity.md) for configuration, supported operations, token verification and restart qualification.

## Related documentation

Use these contracts when reviewing or extending the implementation:

- [Migration status](status.md) records current ownership and qualification limits.
- [Native admission](native-admission/README.md) defines exemption authority, policy routing and controller recovery.
- [Development identity](../../docs/dev/native-development-identity.md) defines the Dex compatibility boundary.
- [Native package recipe](../../packages/native-dev/zarf.yaml) defines bootstrap and activation order.
