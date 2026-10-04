<!--
Copyright 2026 Defense Unicorns
SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial
-->

# Native development Core qualification ledger

This ledger records actual source and running-artifact checks for the experimental native development profile. The profile is undergoing final qualification. It does not migrate Pepr's TypeScript authoring API or qualify advanced production identity.

## Pinned reference

The implementation starts from Core 1.14.0 commit `c15677633cffa17e99656a7d7ccd1f98d7cb6e4d`, the March Go prototype `56b292f4e2b545dca7ce9ca7d6637804ecb3a87f`, and Dex v2.45.1 commit `11d2eeb52b42e1980e14cb91e69dd9e3faab2076`.

## Current evidence

The following results describe their captured checkpoint. Later integration changes require their own qualification.

| Check | Result |
| --- | --- |
| Original Core TypeScript units and values | 1,092 and 176 passed respectively. |
| C6 Go race tests | 1,786 test/subtest passes across 25 tested packages; no named skip or failure. |
| Typed source capture | Six suites passed 366 source tests. All 401 captured calls remained byte-identical after type and formatting fixes. |
| Actual C6 admission | 136 passed, including routing, topology, ownership rollover, failed desired state, repair, finalization and cleanup. |
| Global callback counters | Forty ordinary nonambient requests made zero Go calls. Complex and ambient cases reached their required handlers. |
| Main C6 Core suite | All 24 files ran: 136 assertions passed, six Fleet assertions were blocked by the original 120-second setup deadline, and one original procMount skip remained. |
| Dedicated Fleet on C6/id9 | Six unchanged assertions passed with real bound Kubernetes credentials and strict TLS. |
| Full development browser on C6/id9 | All 21 passed without retries or skips. Two real client updates completed inside that window; a third exceeded the bounded update deadline. |
| Id9 lifecycle and fresh restart | Actual signed issuers, revocation, sessions, directory, clients and refresh checks passed; original token verification completed before expiry. |
| Native identity storage repair | Six actual renders preserve original Java behavior and remove unused native volumes; all 120 Keycloak chart tests pass. Empty-storage replay remains required. |

The two main-suite adaptations select the actual native operator's log labels and remove independently authored positive Probe fixtures after their assertions. The default Pepr log selector and the original allow assertions remain intact. Other main files match the pinned source.

## Remaining integration gates

Live observation found an unconditional five-minute recovery timer invalidating all 33 Package epochs. Full serialized reconciliation of unchanged resources delayed foreground work. C7 adds clean periodic checks and scoped observation of actual gateway, ownership, namespace and shared-configuration drift. It must pass integrated race, full Core, bounded update, fresh installation and isolated guest checks before routine use.

Keep failed attempts in the evidence. A successful dedicated test does not turn a failed full-suite setup into a pass. Rendered charts, low idle RSS and Ready Pods do not establish tenant access, guest capacity or complete policy parity.

## Current integrated result

The actual C8/id10/Dexid4 main run passes all 24 files: 142 assertions pass, no assertion or setup fails, and the existing source procMount skip remains. The test configuration is unchanged; the two documented adaptations remain explicit. Final native admission passes all 136 checks with independent restoration. The full unchanged development browser passes all 21 without a retry or skip while ten real Go-derived client updates persist and produce audit records inside the exact browser window. Each update keeps the fixed 20-second deadline; the maximum convergence is 8.6 seconds.

The actual periodic event checks 34 Packages without invalidating clean recovery epochs. The subsequent foreground updates complete in 2.7–12.4 seconds. Four real drift repairs complete in 3.8–8.2 seconds, and fixture deletion is verified by recorded resource identity. These measurements include their stated boundaries; the tiny cached workflow duration is not the end-to-end API or queue duration.

## Private gateway boundary

The connected development package selects `ClusterIP` for both original Istio gateways. The stock comparison control uses the same override. The admin and tenant gateways remain distinct and retain their real routes, TLS and authorization behavior. This choice qualifies a private inner-cluster profile; it does not qualify production LoadBalancer provisioning.

Hosted access requires an authenticated outer broker bound to the accepted owner, generation, attempt and physical guest. That broker and Ark's hosted UDS control are not implemented or qualified by this Core fork. The owned test harness can relay to private gateway endpoints and verify real Host/SNI, TLS, redirects, signed identity and app responses, but it is not evidence that an Ark assignment transaction exists. Do not expose private API credentials or infer public readiness from a Ready Gateway Pod.

Cold installation, actual private tenant access, matched native performance, final isolated guest capacity and warm assignment remain completion gates. Preserve the earlier unused-volume, empty-chain, field-manager and missing-load-balancer failures separately from repaired results.
