## initial setup notes

22 march 2018 

### single node VM cluster C1

https://github.com/kubernetes-incubator/kubespray#vagrant
https://github.com/kubernetes-incubator/kubespray/blob/master/docs/vagrant.md
https://kubernetes.io/docs/getting-started-guides/kubespray/


^ It was straightforward to get this to work (and I didn't even find the getting started guide until after the fact). I ended up doing:
1. VBox and guest additions are latest (5.2.8)
2. vagrant is latest (2.0.3)
3. ansible is latest (2.4.3.0)
4. i'm on python3 (3.6.4)
5. I had to install "netaddr" for python3 on my host, to get the playbook working for the VM: `pip install netaddr`
6. I cloned Kubespray, and went into the `Vagrantfile` (in root) and made the VM cluster single node centos, and changed the subnet so that the ip is `172.16.7.101`. (avoid docker interference)
7. do not be VPN-ed!!

The whole `vagrant up` command, which includes a full docker/k8s installation, takes 10 minutes or less.

When all is said and done, should see:
```
PLAY RECAP *********************************************************************
k8s-01                     : ok=356  changed=112  unreachable=0    failed=0
```

And ssh-ing into VM:
```
[vagrant@k8s-01 ~]$ kubectl get nodes
NAME      STATUS    ROLES         AGE       VERSION
k8s-01    Ready     master,node   2m        v1.9.3+coreos.0
```

### configure external access to this cluster


#### FROM MY MAC
(quick and dirty)

1. from the VM, copy the contents of `/etc/kubernetes/admin.conf` into `admin.conf` on physical host
2. on physical host, `kubectl --kubeconfig ./admin.conf get nodes`

https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/

#### my mac manging both clusters

right now i'm manually splicing `/etc/kubernetes/admin.conf` to make one superfile, which lives on my physical host

and exporting `$KUBECONFIG`, and doing a `use-context` to switch between both clusters.


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


### install helm

Helm has a client (`helm` CLI) and a server (tiller) deployed onto the k8s cluster.
configure my mac kubectl to talk to VM cluster / there is a context for it

make sure helm is installed on my mac: `brew install kubernetes-helm`

then -- `helm init --kube-context=k8s1`

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
