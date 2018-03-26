## optikon/cmd

command-run for optikon POC, spins up the `/api` server


#### config fields

- `central-cluster` - string
- `edge-clusters `- string array

#### before deploying optikon 

1. spin up 2 or more clusters
2. decide who the "central" cluster is - the rest are "edge" cluster(s).
3. set up Kubeconfig file so that "central" has cluster access to the other cluster(s).


#### what `optikon.go` does

1. takes in the `central-cluster` and `edge-cluster info`
2. performs Setup tasks on Central cluster
    a) installs helm tiller
    b) deploys cluster registry (raw)
    c) deploys prometheus helm chart
    d) deploy
3. perform Setup tasks on Edge clusters, in parallel
    a) install helm tiller and keep track of where Helm API servers are  
    b)
4. using central cluster-registry info and list of Helm tiller servers, spin up `api`
