# zero-trust-proxy

> High-throughput L7 NetOps & SecOps reverse proxy with eBPF socket redirection, SPIFFE/SPIRE workload attestation, and automated mTLS 1.3.

## Architecture

- **eBPF Socket Layer (sockmap)**: Bypasses user-space TCP copy overhead for local container proxies.
- **Identity Enforcement**: Validates X.509 SVID credentials from SPIFFE SPIRE agent.
- **Observability**: Exposes p99 latency histograms, active mTLS sessions, and TLS cipher negotiation metrics.

<!-- commit: feat: initial zero-trust proxy scaffolding and go.mod -->

<!-- commit: feat(ebpf): integrate sockmap bypass for local container ingress -->

<!-- commit: feat(spiffe): implement SVID token validation and auto-rotation -->
