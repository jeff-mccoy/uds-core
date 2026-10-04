<!--
Copyright 2026 Defense Unicorns
SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial
-->

# Native development Core qualification ledger

The integrated experimental development runtime passes its final Core gates. Matched measurements show a large process-memory reduction and a slower Package lifecycle at the current 100m CPU limit. Private frontend and isolated guest qualification are being completed. The fork does not migrate Pepr's TypeScript authoring API or qualify advanced production identity.

## Pinned reference and artifacts

The reference is Core 1.14.0 commit `c15677633cffa17e99656a7d7ccd1f98d7cb6e4d`, the March Go prototype `56b292f4e2b545dca7ce9ca7d6637804ecb3a87f`, and Dex v2.45.1 commit `11d2eeb52b42e1980e14cb91e69dd9e3faab2076`.

The C8 Go binary hash is `1ed37e6920ab255041a055ff3cec7bdea1d7c83593e005255414f47e2a33547d`. Identity id10 is `fc2b974dfc71c37cc150ef2242f103267b6d1a09570fe5e6f28570f21e5a95f1`, and Dex id4 is `b4ac41206514b95ee1ebd427983106fdf998cd967efc2e4a970970d1e2e0fa1a`. The native policy hash is `78bb1412dfcd14590b82fa15c84ab7077384261e018f9e4e139bb43c870e9291`. Local derivative images remain unsigned.

## Final runtime evidence

These results apply to the captured C8/id10/Dex4 artifacts. Historical failed checkpoints remain distinct from repaired results.

| Check | Actual result |
| --- | --- |
| Original Core TypeScript units and values | 1,092 and 176 passed respectively. |
| C8 Go race suite | 1,827 tests/subtests across 25 tested packages passed; no failure or named skip. |
| Typed source capture | Six suites passed 366 source tests; all 401 captured calls remained byte-identical. |
| Native admission | 136 real API cases passed, including revision routing, topology, ownership rollover, failed desired state, repair, finalization and exact cleanup. |
| Global callback counters | Forty ordinary nonambient requests made zero Go calls. Complex and ambient cases reached required handlers. |
| Main Core suite | All 24 files ran: 142 passed, zero failed, one original procMount skip. Twenty-two source files match exactly; the two documented adaptations remain explicit. |
| Trust bundle suite | Two unchanged assertions passed against real Google HTTPS with valid and empty CA bundles. Original global configuration was restored with UID/version guards. |
| Development browser and concurrent changes | All 21 passed without retry or skip during ten persisted, audited Go-derived client updates. All changes stayed inside the browser window and the fixed 20-second deadline; maximum convergence 8.6 seconds. |
| Actual periodic recovery | The five-minute event checked 34 Packages without invalidating clean epochs. Foreground updates took 2.7–12.4 seconds. Four actual drift repairs took 3.8–8.2 seconds. |
| Identity lifecycle and restart | Real signed issuers, revocation, sessions, directory, clients and refresh checks pass with strict TLS. Immediate restart verification finished before original-token expiry. |
| Fleet switch and full suite | Disabled grants and bearer access reject; re-enable requires fresh grants. Restart retains the fresh grant and rejects the old one. Six original full-suite assertions pass within their setup deadline. |
| Exemption outage and serving CA rotation | Missing/invalid snapshot recovery passes; 79 API admissions and 158 verified TLS connections pass through actual two-replica CA rotation. |
| Chart and bootstrap helper | All 120 Keycloak chart checks and 113 helper race cases pass, including actual embedded UDS and precise terminal-Pod kubelet retirement. |
| Reproducibility | Exported source without Git metadata reproduces exact release identity and Dex binaries. Policy, image payload and rendered-object hashes are checked. |

The two main-suite adaptations select actual native operator log labels and remove independently authored positive Probe fixtures after their assertions. The default Pepr log selector and original allow assertions stay intact. The full qualification stack includes actual optional monitoring, logging, Portal, Falco and backup fixtures; it is not a slim resource sample.

## Matched Node, Deno and native measurements

The controls share identical benchmark contracts, source network fixtures, eight verified lifecycle trials and 300 admission requests per case at concurrency one and eight. Resource observations join stable running CRI IDs across six quiet samples in nine Core namespaces. Cache and mutation history differ, so the table does not establish equivalent scheduler reservations, VM density or provider cost.

| Profile | Core RSS MiB | Core working set MiB | Admission/operator RSS MiB | Create/update/delete median seconds |
| --- | ---: | ---: | ---: | --- |
| Original Pepr/Node, two admission plus one watcher | 1419.1 | 1530.5 | 511.3 | 0.642/0.633/0.478 |
| Original Pepr/Node, one admission plus one watcher | 1286.6 | 1389.4 | 379.1 | 0.569/0.551/0.609 |
| Full Pepr/Deno, two admission plus one watcher | 1496.3 | 1653.8 | 562.5 | 0.494/0.560/0.462 |
| Native policies, Go and Dex | 436.4 | 520.2 | 51.9 | 3.555/3.493/1.374 |

The native container retains 100m CPU and 128Mi memory limits. Default-Pod admission p50/p95 is 9.36/11.35 ms at concurrency one and 8.17/63.82 ms at concurrency eight; the original Node topology is 5.37/7.24 and 11.60/87.43 ms respectively. Do not claim an unconditional native speed improvement. Full Deno retains failed raw IPC compatibility evidence and does not save memory in this trial.

## Fresh private startup and activation

The connected development profile keeps both original gateways distinct and selects ClusterIP for each, matching the stock comparison override. Actual TLS, routes and authorization remain enabled; production LoadBalancer defaults are unchanged.

Fresh fenced installation passes all 19 bootstrap stages without ever installing Pepr in 150.62 seconds from empty storage with local cached inputs. Regular activation replay passes in 160.27 seconds and removes the tenant fence through the final readiness gate, with no manual policy deletion. The final candidate package hash is `db9cac899307d6620d2c481e19a4e8b05977e04fa4cdd4fe9a68080be5a926c7`; regular activation is `df644929a58ed454fdd3e8616d16b21f41bb8004b21e255a2559379e71751fb8`; the preload is `913fb064e6f55046f959cea6b1e0198c79fae3c85068f294e57d73f6eb0d210b`. Acquisition and build duration are separate costs.

## Isolated guest and warm sample

The same current package passes 22 native authority-object checks and eight real admission/SSO topology canaries in a 4 GiB, six-vCPU Kata guest on a task-owned 6 GiB outer node. One actual controller-produced Agent Sandbox pool member reaches Core readiness in 204.06 seconds. A single warm claim takes 0.3082 seconds and retains physical Pod UID, container identity, zero restarts, private API CA and signing keys through real owner transfer. Assertions precede pool retirement; the claimed VM remains for frontend validation.

The observed VM/shim RSS is approximately 3.74 GiB and outer charge 5.46 GiB. Retain the 4 GiB guest and 6 GiB outer envelope. This is one sample, not a sustained refill result or latency SLO. The initial manual enrollment, cgroup-wrapper and readiness-grace failures stay preserved; no fabricated Ready status or extended SDK deadline is used.

## Hosted access boundary and remaining work

Strict private frontend application/group checks and exact claimed-guest retirement are still being completed. Task-owned relays preserve original gateway TLS, Host/SNI, signed identity and routes; they do not implement an Ark owner/assignment transaction.

Hosted access needs an authenticated broker bound to the accepted owner, generation, attempt and physical guest. This fork does not implement or qualify that broker or restore Ark's hosted UDS control. Keep the hosted chart gated until its lifecycle and execution isolation are qualified. Rendered charts, low process memory, Ready Pods and a warm claim alone do not establish complete hosted qualification.

Performance tuning must preserve the final passing baseline and be measured separately. Keep the prior unused-volume, empty-chain, field-manager, load-balancer and terminal-Pod failures as historical evidence rather than relabeling them as passes.
