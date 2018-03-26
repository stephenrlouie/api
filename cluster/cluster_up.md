## initial setup notes

22 march 2018

### single node VM cluster C1

https://github.com/kubernetes-incubator/kubespray#vagrant
https://github.com/kubernetes-incubator/kubespray/blob/master/docs/vagrant.md
https://kubernetes.io/docs/getting-started-guides/kubespray/


^ It was straightforward to get this to work (and I didn't even find the getting started guide until after the fact). My general steps/tips...
1. VBox and guest additions are latest (5.2.8)
2. vagrant is latest (2.0.3)
3. ansible is latest (2.4.3.0)
4. i'm on python3 (3.6.4)
5. I had to install "netaddr" for python3 on my host, to get the playbook working for the VM: `pip install netaddr`
6. I cloned Kubespray, and went into the `Vagrantfile` (in root) and made the VM cluster single node centos, and changed the subnet so that the ip is `172.16.7.101`. (avoid docker interference).


7. **IMPORTANT** - I also had to increase the Memory to 4096 and have 3 CPUs because `KubeDNS` was throwing a `not enough memory` error and never deployed.

```
kube-dns-79d99cdcd5-5jzkt
Node didn't have enough resource: cpu, requested: 260, used: 650, capacity: 800
```

8. do not be VPN-ed!!!


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

(FROM MY MAC / quick and dirty)

1. from the VM, copy the contents of `/etc/kubernetes/admin.conf` into `admin.conf` on physical host
2. on physical host, `kubectl --kubeconfig ./admin.conf get nodes`

https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/

#### kube dashboard

I thought I set up the host-only/NAT VM stuff correctly but going to where the Dash was supposed to be running didn't work: `http://172.16.7.101:8443`

So instead I [proxied the kube API](https://kubernetes.io/docs/tasks/access-kubernetes-api/http-proxy-access-api/#using-kubectl-to-start-a-proxy-server) to localhost and that worked. Hacky steps:

1. From my mac, `kubectl proxy`. Should see a message: `Starting to serve on 127.0.0.1:8001`
2. In some other window follow [these instructions](https://github.com/kubernetes/dashboard/wiki/Access-control#admin-privileges) to do an admin `ClusterRoleBinding` for kube dashboard. basically blanket permissions without having to add one of the many `tokens`. (Which you can see by doing `kubectl -n kube-system get secret`).
3. Then navigate to the proxied dashboard at 127.0.0.1:8001/ui
4. Don't enter any tokens or kubeconfig info! There should be a `SKIP` button after you created the admin-level ClusterRoleBinding. press SKIP.
5. You should have full dashboard access to every namespace, including `kube-system` and `kube-public`.


#### my mac managing two clusters

right now i'm manually splicing `/etc/kubernetes/admin.conf` to make one superfile, which lives on my physical host

and exporting `$KUBECONFIG`, and doing a `use-context` to switch between both clusters.



### install helm

Helm has a client (`helm` CLI) and a server (tiller) deployed onto the k8s cluster.
configure my mac kubectl to talk to VM cluster / there is a context for it

make sure helm is installed on my mac: `brew install kubernetes-helm`

then -- `helm init --kube-context=k8s1`
