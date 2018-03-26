## aggregated cluster-registry

(apparently standalone just has a bug and that cluster-registry SHOULD work on minikube, baremetal etc)

(i'm on k8s 1.9 do NOT be on 1.6 or earlier for this)

#### Background Links

https://github.com/kubernetes/cluster-registry/blob/master/docs/userguide.md#aggregated-api-server

https://github.com/kubernetes-incubator/apiserver-builder/blob/master/docs/concepts/aggregation.md

https://github.com/kubernetes-incubator/apiserver-builder/blob/master/docs/concepts/auth.md#requestheader-authentication

https://kubernetes.io/docs/tasks/access-kubernetes-api/configure-aggregation-layer/#enable-apiserver-flags

#### Install cluster registry on the "central" cluster

install crinit on linux

`./crinit aggregated init aggro --host-cluster-context=admin-cluster.local`

should be able to `kubectl get clusters --context $kubectl config current-context)` and see:

```
No resources found.
```

YAY i can get clusters



#### Add a sample cluster to make sure it's working

export CONTEXT=$(kubectl config current-context)
kubectl apply -f mock-cluster.yaml  --context $CONTEXT

kubectl get clusters --> there should be 1 there, called `test-cluster`

or curl http://localhost:8001/apis/clusterregistry.k8s.io/v1alpha1/clusters

### Add a new edge cluster to the central registry

bring up a minikube cluster or whatever's fastest
get the API server endpoint
update minikube-cluster.yaml
