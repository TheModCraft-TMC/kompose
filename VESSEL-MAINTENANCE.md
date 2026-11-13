# Vessel maintained Kompose branch

Based on upstream `kubernetes/kompose` tag `v1.38.0`, commit
`a8f5d1cbdcfe3b47cb5c936fb1d57966390064a8`. The upstream module identity is
retained; consumers select this reviewed source fork using a versioned Go module
replacement. This is a source migration, not a replacement of an unrelated module.

Changes:

- Use stable `autoscaling/v2` HPA Go types, matching generated wire metadata.
  A regression registers only the stable API and checks the generated object.
- Use `fsouza/go-dockerclient` 1.13.3 and the split Moby SDK. The tidied fork
  graph no longer selects `github.com/docker/docker` or `github.com/docker/distribution`.
- Retain Vessel's Kubernetes 0.35.4 baseline. The stable HPA implementation also
  passed the complete upstream suite with Kubernetes API/apimachinery 0.37.1.
- Require Go 1.26.11 and compression 1.20.3; use Vessel's patched x/net 0.62.0.
- Correct a port-test diagnostic format rejected by current Go vet.

Run `GOWORK=off go test -race -p 4 -count=1 ./...` and `go mod tidy -diff`
before tagging. Keep source ancestry and dependency checksums. This module alone
does not remove old Docker/Distribution requirements from Vessel's other modules.
No live Kubernetes/OpenShift deployment is certified by unit tests.
