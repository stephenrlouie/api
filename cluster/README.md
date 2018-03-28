## optikon/cluster

Cluster registry API spec: https://github.com/kubernetes/cluster-registry/tree/master/api

### setup instructions

Imagine that this is the "central cluster."

1. Vagrant up using the `../vagrant/cluster` Vagrantfile subbed into the [kubespray playbook](https://github.com/kubernetes-incubator/kubespray). This should take 10 minutes or less for a single node VM cluster.

2. [Optional] configure your physical host to use `kubectl` to talk to the VM cluster. The hacky way I do this right now is by copying `/etc/kubernetes/admin.conf` from the VM into some file on my desktop, then running:

```
export KUBECONFIG=~/Desktop/admin.conf
```

3. Inside the VM, install the `crinit tool` by following [these instructions](https://github.com/kubernetes/cluster-registry/blob/master/docs/userguide.md#deploying-a-cluster-registry).

4. Deploy an **aggregated** cluster-registry API server onto the single node cluster. (Standalone is buggy.)

```
./crinit aggregated init aggro --host-cluster-context=admin-cluster.local
```

Where "aggro" is the cluster registry context.


5. From anywhere (inside VM, outside, etc.) you should be able to `kubectl get clusters` and see "No resources found." This means that the custom API object is now recognized.

6. Add a cluster to the registry using the example in this folder:

```
export CONTEXT=$(kubectl config current-context)
kubectl apply -f minikube-cluster.yaml  --context $CONTEXT
```

7. Since I `kubectl` from outside the "central cluster" VM, I do `kubectl proxy` to proxy the k8s API server onto my physical host. Probably not a sustainable way to do this.

8. `kubectl create -f dashboard-admin.yaml` (in this folder). This lets you open up the kube dashboard at `http://localhost:8001/ui` and click "skip" without having to enter creds or tokens.


9. You should then be able to curl / browser this: `http://localhost:8001/apis/clusterregistry.k8s.io/v1alpha1/clusters`  
And see an output similar to `GET-clusters-sample.json`
