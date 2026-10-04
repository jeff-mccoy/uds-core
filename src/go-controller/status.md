<!--
Copyright 2026 Defense Unicorns
SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial
-->

# Native Core implementation and qualification status

The development implementation replaces the UDS Core Pepr runtime with Go reconciliation and Kubernetes admission policy. It also provides an opt-in password-only Dex identity provider. This is a development profile with qualification in progress, not a production migration or a generic Go replacement for Pepr's TypeScript authoring API.

The March Dash Days snapshot established feasibility for selected responsibilities. This source extends that snapshot against pinned Core 1.14 behavior. Earlier statements that reload, complete policy validation, current CRD generation, dependency observation, and unit coverage were absent no longer describe this branch.

## Capability ownership

The native profile uses the following implementation boundaries. The original Core deployment remains the comparison control.

| Capability | Native implementation | Qualification boundary |
| --- | --- | --- |
| `Package` lifecycle | Go keyed reconciliation, live refetch, retries, finalizers, durable ownership journal | Resource and reload integration tests pass; remaining full-stack qualification continues. |
| Network and mesh lifecycle | Go network policies, current API/Node discovery, ingress/egress and UDP routes | Ordinary network tests pass; real sidecar, UDP and passthrough paths have independent positive/negative receipts. |
| Pod reload | Go Secret/ConfigMap data observation and owning workload restarts | Unchanged reload tests pass, including source event diagnostics. |
| Simple policy/default behavior | Native validating and mutating admission policies | Actual API checks cover defaults, markers, namespace boundaries, union grants, revocation and missing/invalid inputs. |
| Complex validation and mutation | Go code admission with a complete fallback during qualification | Source differential fixtures and unchanged isolated policy tests pass; callback narrowing and failure recovery remain active qualification work. |
| Exemptions | One protected compiled union plus Go JavaScript regex evaluation where required | Native grant lists use OR semantics, preserve source UID, and revoke stale authority. Tenant markers do not grant exemptions. |
| Monitor/probe/CA behavior | Go lifecycle plus native object-local admission | Full monitoring, logging, alert and trust integration runs continue. |
| Identity for ordinary development | Dex with persisted Go compatibility service | Real signed protocol, lifecycle and restart checks pass; full application default-scope compatibility remains under qualification. |
| Advanced production identity | Original Keycloak provider | Dex rejects unsupported required capabilities; it does not emulate their assurance. |
| TypeScript module authoring and KFC | Existing Pepr/KFC libraries | This Core implementation does not replace their consumer-facing APIs or establish external migration parity. |

## Evidence available in this branch

Keep source-unit tests, differential fixtures, actual API tests and complete application journeys separate. They answer different questions.

| Check | Current result | Meaning |
| --- | --- | --- |
| Pinned Core TypeScript unit suite | 1092 pass | Reference source checks run unchanged. |
| Core values suite | 176 pass | Existing configuration checks run unchanged. |
| Integrated C6 checkpoint under the race detector | 1786 test/subtest passes across 25 tested packages | Fresh execution; no named failures or skips. Later source changes require new qualification. |
| Package admission differential | 131 calls from 130 source cases | Preserves decision and source diagnostic text. |
| Policy helper differential | 182 calls from 148 source cases | Preserves captured policy behavior. |
| JavaScript regex corpus | 1008 cases | Detects semantic differences that successful RE2 compilation alone cannot rule out. |
| Generated resource name corpus | 21 cases | Covers explicit name, generated-name preference and empty-name boundaries. |
| Scoped native API qualification | 94 pass | Includes 54 explicit native denials; this earlier run retained the combined fallback stack. |
| Isolated unchanged policy suite | 28 pass, one original explicit skip | Pepr admission and watcher are both drained, their callbacks absent. A separate faithful API case covers the skipped procMount denial. |
| Real Fleet ServiceAccount integration | Six unchanged tests pass | Validates actual bound credentials, audience, ownership prefix, rename and built-in client restrictions. |
| Identity protocol and restart | Strict signed token, lifecycle and persistence checks pass | Keeps issuer, audience, nonce, expiration and TLS verification enabled. |

The full browser qualification repaired default-scope claim differences, the account document, and per-client publication. Current id9 passed all 21 development browser journeys without a retry or skip, plus real protocol lifecycle and fresh restart verification before token expiry. Two actual client updates completed inside that browser window. A third exceeded its bounded 20-second controller deadline; the complete concurrency loop remains failed evidence.

The complete main Core suite executes all 24 source files with the real egress, UDP, monitoring, logging, backup and runtime-security fixtures. C6 passed 136 of 143 assertions. Fleet's original 120-second setup deadline blocked its six assertions; the remaining skip is Core's existing procMount case. A separate dedicated Fleet run passed all six. Twenty-two main files match the pinned source. The two explicit adaptations select actual native operator log labels and clean up independently authored positive Probe fixtures; neither changes the allow or collection assertions.

Actual C6 admission qualification passed 136 cases covering native routing, topology creation and updates, Package UID rollover, failed desired-state preservation, authenticated repair and finalization. Forty ordinary nonambient Pod/Service requests made zero Go calls; complex cases and ambient topology reached their required handlers. Independent cleanup restored all eleven canonical bindings and removed all qualification namespaces.

Qualification found a five-minute recovery timer invalidating every Package epoch. Two exact broadcasts queued 33 Packages and forced full serialized workflows, explaining the remaining Fleet and concurrency delays. The next checkpoint retains stable Ready epochs for periodic checks and adds missing scoped drift observation. It must preserve recovery after real configuration, ownership, namespace, gateway and shared-Secret changes. Its full suite and bounded concurrency repeat are still required.

Fresh C6 installation reached a real Ready Dex provider but timed out because the chart emitted unused Java data and theme PVCs. The native-only chart repair emits no such PVCs or Java migration exemption; original Java renders remain equivalent, and all 120 Keycloak chart tests pass. A new empty-storage install must prove the complete repaired package. Passing a health endpoint or a render does not establish that contract.

## Recovery and ownership

Go uses a native Lease to elect one reconciliation authority. Each replica maintains its admission caches and shared serving identity. The source fences work after leadership loss and propagates retryable cleanup errors instead of reporting successful deletion prematurely.

Owned resource updates preserve Kubernetes resource versions and reject foreign `Package` UIDs. The ownership journal retains enough SSO, waypoint and probe state to retry finalization after process recreation. Namespace cleanup preserves foreign ownership and explicit unmanaged namespaces.

The webhook replicas share a protected serving Secret. Their readiness gate checks certificate and registration convergence. Rotation publishes old and new roots before changing the leaf, then retires old trust after the overlap interval. Source race and real TLS-server tests cover this protocol. Earlier live qualification passed missing/invalid snapshot recovery, two-replica recreation, and 79 API admissions with 158 independently verified TLS connections during shared CA rotation. Those receipts describe their captured image; the final integrated artifact retains its own recovery qualification gate.

Reference the [admission recovery contract](native-admission/README.md#controller-recovery-boundary) for the protected bootstrap shape and mandatory independent policies.

## Remaining qualification

Complete the following gates before marking this profile ready for routine development use:

- Run the remaining unchanged Core cluster and browser suites with required fixtures and explicit feature gates.
- Complete actual controller outage, replica, certificate rotation, leadership and rollback checks.
- Verify a fresh native package installation without ever installing Pepr, including infrastructure publication and tenant fences.
- Narrow code callbacks only after proving safe native delegation under propagation delays and grant changes.
- Measure equivalent stock Node, narrowed runtime and complete native profiles with settled running-container inventories.
- Qualify the intended isolated guest runtime, storage and warm assignment boundary with the final package.

A framework language change, broad Pepr/KFC migration, production identity transition, and hosted Ark feature activation require their own ownership and qualification decisions.

## Related documentation

Use the implementation and contract pages for current configuration:

- [Controller guide](README.md) describes modes, build commands and generated API ownership.
- [Native admission](native-admission/README.md) defines policy semantics, grant publication and serving trust.
- [Development identity](../../docs/dev/native-development-identity.md) documents the restricted identity provider.
- [Native package](../../packages/native-dev/zarf.yaml) records the concrete installation sequence.
- [Qualification ledger](../../docs/dev/native-core-qualification.md) records subsequent artifact-level results and remaining gates.
