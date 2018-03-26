## cluster registry attempted setup

*note* -- I am pretty sure (almost positive at this point) that cluster-registry will not work (in its current version) on non Cloud (aka non GCE) k8s clusters.
AKA won't work on VM clusters, baremetal clusters, in the lab. 
This is because the `Service` that `crinit` tries to set up is of type `LoadBalancer`:

Example from my own deployment:
```
{
  "kind": "Service",
  "apiVersion": "v1",
  "metadata": {
    "name": "tryreg",
    "namespace": "clusterregistry",
    "selfLink": "/api/v1/namespaces/clusterregistry/services/tryreg",
    "uid": "1a1acf9b-311f-11e8-acdf-525400225b53",
    "resourceVersion": "10958",
    "creationTimestamp": "2018-03-26T17:57:08Z",
    "labels": {
      "app": "clusterregistry"
    }
  },
  "spec": {
    "ports": [
      {
        "name": "https",
        "protocol": "TCP",
        "port": 443,
        "targetPort": "https",
        "nodePort": 31600
      }
    ],
    "selector": {
      "app": "clusterregistry",
      "module": "clusterregistry-apiserver"
    },
    "clusterIP": "10.233.9.28",
    "type": "LoadBalancer",
    "sessionAffinity": "None",
    "externalTrafficPolicy": "Cluster"
  },
  "status": {
    "loadBalancer": {}
  }
}
```

And according to [the Kubernetes Docs](https://kubernetes.io/docs/concepts/services-networking/service/#type-loadbalancer):

> On cloud providers which support external load balancers, setting the type field to "LoadBalancer" will provision a load balancer for your Service.  


#### cluster registry on cluster1

Tried to use cluster-registry's own `crinit` tool which spins up a cluster registry API server on one of the clusters you want to manage. And from that cluster registry you can manage multiple clusters. But having issues where the init process never completes.

`crinit standalone init cr-megan --host-cluster-context=admin-cluster1.local`

NOTE: standalone cluster registry is getting stuck right here.
This has happened on every VM cluster i've tried -- including minikube and "raw" kubespray vagrant VMs, centos etc..
It also gets stuck deploying into a lab cluster.

example:

```
➜  cluster-registry git:(master) ✗ crinit standalone init cr-lab --host-cluster-context=kubernetes-admin@kubernetes
CREATING NAMESPACE...
Creating a namespace clusterregistry for the cluster registry... done
CREATING SERVICE...
Creating cluster registry API server service...................................
```

And 20 minutes later it's still trying to make the service ^^.

DEBUGGING FROM SOURCE, WITH PRINTS
`bazel build //cmd/crinit`

`sudo cp bazel-bin/cmd/crinit/darwin_amd64_stripped/crinit /usr/local/bin`



### cluster registry alternative? raw client-go

If we can't get the cluster-registry api server working in a VM cluster, we could always use the raw client go API tools
and hack a "multicluster observation API" to render in a UI.

And I ran the "outside the cluster" demo, modifying the kubeconfig variable to point to my spliced-together kubeconfig.

https://github.com/kubernetes/client-go/tree/master/examples/out-of-cluster-client-configuration

This is *extremely hacky* though and not ideal.

So from my mac I have 3 contexts (2 VM clusters, 1 for the lab `38` cluster), and running the client-go demo I get:

```
➜  clientgo git:(clientgo) ✗ ./app
There are 12 pods in the cluster
Pod example-xxxxx in namespace default not found
```
