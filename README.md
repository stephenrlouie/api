# optikon

an edge app manager prototype

- [asana](https://app.asana.com/0/605226027291146/board)
- [box folder](https://cisco.app.box.com/folder/48111186509)
- [long-term app manager roadmap](https://cisco.box.com/v/cie-am-roadmap)


### what is optikon?

- A platform that allows you to
    1. view a set of kubernetes clusters (edge cluster management)
    2. deploy a Helm chart to multiple clusters at once (one-touch edge app deployment)
    3. automatically route inbound cluster traffic to the nearest cluster (edge load balancing)
- A proof of concept edge app architecture for Kubecon Europe '18
- The "version 0" implementation of the Cisco Intelligent Edge - Application Manager


### what is optikon not?

- An edge device manager (does not do baremetal provisioning)
- An edge platform manager (assumes one or more k8s clusters already online, does not do installations for cluster logging or monitoring)


### optikon architecture

![archdiagram](https://wwwin-github.cisco.com/edge/optikon/blob/master/docs/arch.png)


#### arch explanation
- A central kubernetes cluster is the "eye in the sky." It manages other ("edge") kubernetes clusters.
- For POC purposes, these clusters could live anywhere. example: VM clusters on a laptop. or all clusters co-located in the same lab.

**Central Cluster**
- The central cluster hosts the optikon API and UI. Endpoints are `/cluster` and `/app`.
- The `/cluster` endpoint passes through to the Kubernetes Cluster-Registry API server. cluster-registry [is an existing thing](https://github.com/kubernetes/cluster-registry).
- The `/app` endpoint passes through to the helm forwarder.
- The helm forwarder is a tiny API client that lets users run helm charts on mutiple k8s clusters at once, and also get the status of their helm releases. we write that takes in a helm chart / some inputs and a list of cluster names to deploy onto. It knows about all the helm receivers on every "edge" cluster, and forwards the helm app info to each of the clusters listed. (*note* - this is a "hack" in advance of Helm v3, which plans to have multicluster management)
- Central cluster runs a CoreDNS nameserver. DNS names from the helm charts are added by the Optikon API. DNS name mapped to a list of anycast IPs for every helm app. See the [DNS doc](https://wwwin-github.cisco.com/edge/optikon/blob/master/coredns/dns.md) for details.

**"Edge" Clusters**
- This arch shows just one edge cluster, but a good demo would show at least two additional clusters
- Cluster registry manages this cluster
- Helm receiver is a tiny api server- gets requests to deploy a helm chart onto itself. Helm receiver can talk to the helm tiller (helm server) to do a `helm install`,  `helm list`.
- One edge app corresponds to one helm chart.
- CoreDNS also runs on every Edge cluster. This "edge CoreDNS" can talk to the centrally-running CoreDNS. Again, see the [DNS explanation doc](https://wwwin-github.cisco.com/edge/optikon/blob/master/coredns/dns.md) for details / open questions.


### repo structure

- `api/` = optikon user API, to interact with `/cluster`s and `/app`s
- `cluster/` = view and interact with multiple k8s clusters
- `dns/` = central CoreDNS and edge CoreDNS setup
- `docs/` = docs
- `helm/`= helm forwarder and receiver
- `ui/` = optikon user interface (on top of the `api`)
