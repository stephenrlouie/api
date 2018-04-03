# optikon/cluster

cluster-registry info, samples. for the `/cluster` endpoint of the Optikon API / `cr-forward` mechanism

## Resources

- [cluster-registry repo](https://github.com/kubernetes/cluster-registry)
- [initial motivation](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/multicluster/cluster-registry/api-design.md#motivating-use-cases)
- [purpose of cluster-registry](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/multicluster/cluster-registry/api-design.md#purpose)


## cluster-registry + optikon

- Optikon is a proof of concept for managing multiple clusters and deploying apps to those clusters.
- We want to follow kubernetes community trends around multicluster management and orchestration.
- cluster-registry is designed to be a common abstraction for tools to do operations on multiple clusters - optikon is exactly that kind of tool.
- the optikon api has a `/cluster` endpoint to work with multiple clusters -- this (in part) points to the cluster registry API.

## details: using the cluster-registry in the optikon API

So given what the cluster-registry can do (store info about clusters), this is how I propose we use it in the POC:
- the `POST` and `DELETE` cluster endpoints of the optikon API can be a direct passthrough to the cluster-registry API server. **Note** that `POST/cluster` and `DELETE/cluster` won't actually init or torch the cluster. It's just CRUD on an object representing that cluster.
- for `GET` clusters we want the ability to actually get the nodes, running state, maybe the running workloads, etc. of that cluster -- cluster-registry doesn't currently support this. so the optikon API server will not only have to interact with the centrally-running cluster registry, but *also* be able to open up client connection to the k8s api for each individual cluster. so the call flow might go:
    1. client calls `GET/cluster/<id>`
    2. optikon API server does a `cluster-registry GET/cluster<id>`
    3. optikon API server ^ gets cluster name, API endpoint, and *auth info* -- all available in the `cluster` object. This auth info will have to include the `token` for the k8s `apiserver` on the cluster in question.
    4. optikon API server calls the k8s api, the equivalent of [this](https://kubernetes.io/docs/tasks/administer-cluster/access-cluster-api/#without-kubectl-proxy).

All the ops of the `/charts` endpoint will have to work in a similar way to the `GET/cluster` endpoint, where in the optikon API server, you'll have to get all the clusters you want to operate on, get the helm tiller endpoint, etc. Also we may have to do some magic to make helm tiller externally routable, across all the clusters. Ie. we have to get to the edge cluster's tiller from the optikon api server running in the central cluster.  

## To keep in mind

- `./crinit init standalone` doesn't work (has to be an [aggregated API server](https://kubernetes.io/docs/concepts/api-extension/apiserver-aggregation/)). Aka use `./crinit init aggregated...`
- The `status` field for the `cluster` object is [not yet implemented](https://github.com/kubernetes/cluster-registry/issues/28)
- the whole `cluster-registry` internals are being transformed as we speak, from an API server to a [Custom Resource Definition](https://kubernetes.io/docs/concepts/api-extension/custom-resources/#custom-resources). [doc here.](https://docs.google.com/document/d/1FhoxzYUJ4qiyZ4qA8LE5zgCO4jUihiDUlbttvLlnRmk/edit#heading=h.c38m7d4yucj5)
- `cluster-registry` is only a few months old (currently at version `0.0.3`, first release was in Dec. '17)

## cluster-registry setup (vagrant)

### manual cluster-registry setup (central cluster)

1. download crinit
2. create PV
3. sudo su
4. crinit aggregated $current-context


### adding an edge cluster

5. create edgecluster.yaml
6. kubectl create -f edgecluster.yaml
