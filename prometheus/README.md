### manual prom deployment for 1 cluster

(the goal = automate this for every managed cluster, with the helm forwarder)

1.
```
kubectl create clusterrolebinding permissive-binding --clusterrole=cluster-admin --user=admin --user=kubelet --group=system:serviceaccounts;
```

2. Prom helm chart (by default) needs _two_ persistentVolumes. You can change the default to use an emptyDir instead but this speeds it up. `create-f` two. (see `pv-giraffe` and `pv-unicorn` in this folder for mine)

3. Make sure helm's installed on the cluster, if it's not, `helm init`

4. `helm install stable/prometheus`

Note that it shouldn't matter which namespace helm tiller / the PVs are in, you should be able to deploy prom in the default namespace or wherever.
If `tiller` is on some other namespace you might have to do `--tiller-namespace=<the-right-namespace>` every time you do a helm command.

5. If you want to see the prom UI do:

```

export POD_NAME=$(kubectl get pods --namespace default -l "app=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}");
kubectl --namespace default port-forward $POD_NAME 9090
```

6. Navigate to `localhost:9090/targets` and you should be up and running. 
