<!--
Copyright 2026 Defense Unicorns
SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial
-->

# Controller and admission migration status

This branch contains Go reconciliation and native admission against Core 1.14,
with the existing Keycloak API as its identity boundary. The standalone Dex
runtime and combined development deployment recipe have separate reviews. No
code in this module imports or builds Dex.

## Implemented boundary

The controller implements `Package`, network, Istio, SSO, monitor, probe, CA and
reload lifecycle. Native policies handle suitable object-local decisions; Go
retains complex topology and JavaScript-compatible exemption behavior. Protected
grant union, source UID ownership, revision fencing, drift recovery, leadership
and shared webhook trust remain part of the controller contract.

Public S256 client pairing adds authenticated `Package` provenance only for
explicitly paired clients. Ordinary Keycloak client serialization retains its
original representation. Those identity-neutral operator fields do not require
the separate Dex service.

## Validation scope

The split runs its own Go race suite and chart/admission checks. The integrated
draft's earlier 1,878-case suite includes identity packages, so that total does
not describe this module. Likewise, earlier full Core/browser results used the
combined native/Dex profile and remain integration evidence, not an independent
Go-plus-Keycloak live pass.

Core's original Keycloak chart, bundle identity schemas and existing tests are
byte-identical to the pinned base on this branch. The exact identity
administration override lives in `values.keycloak.yaml` and selects only the Go
controller's principal through an existing chart value.

## Split-source result

The independent module passes 1,788 Go race-tested cases across 23 tested
packages, with zero failures or skipped test cases. The standalone controller
build and vet pass. Both complete and narrowed admission renders pass the
existing contract checker. All 113 pristine Keycloak chart assertions pass.
Original provider/test files remain byte-identical. These are new split-source
results, separate from prior integrated live runs.

## Rollout gates

You must qualify a fresh installation with the original identity provider before
enabling this branch for routine development. Verify actual SSO, Fleet/probe
credential modes, leader replacement, registered admission, exemption
revocation, certificate rotation, drift and guarded retirement. Do not assign
two active writers to the same phase.

The combined development bootstrap and the historical resource measurements
remain separate. This source split does not claim new memory, startup, density
or cost results. It does not migrate other Pepr modules or KFC.

## Related documentation

The following contracts describe the source:

- [Controller guide](README.md) lists modes, build commands and identity configuration.
- [Native admission](native-admission/README.md) defines authority, routing and recovery.
