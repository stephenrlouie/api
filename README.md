## optikon-api


The Optikon API makes it easy to watch multiple Kubernetes clusters, and to orchestrate [Helm](https://www.helm.sh/) charts on multiple clusters.
It was originally designed as an Edge Computing application manager, for easy monitoring and deployment across clusters.

## Components

The Optikon API wraps two APIs:
1. the Kubernetes [Cluster Registry](https://github.com/kubernetes/cluster-registry/blob/master/docs/userguide.md), deployed onto one "central" cluster.
2. the Helm client -- and facilitates Helm chart CRUD operations for multiple Helm tillers, allowing one-touch deployment of a Helm chart to multiple clusters at once.

See the `docs/` folder in this repo for an architecture diagram.

## Installation

These examples shows how to run the API using mock handlers. See the [optikon-vagrant](https://github.com/optikon/optikon-vagrant) docs for information on deploying the Optikon API in kubernetes, next to a live Cluster Registry API.

### Run in Docker

1. Deploy the API:

```
docker run intelligentedgeadmin/optikon-api:latest --scheme http --port 9000 --mock-base-path ./api/v0/mock/api`
```

2. You should now be able to `curl http://127.0.0.1:9000/v0/clusters` and get back a list of clusters:

```
~> curl http://127.0.0.1:9000/v0/clusters
[{"metadata":{"annotations":{"Health":"Green","Lat":"55.6899699",...
```


### Run from source

1. clone this repo

2. `make` (from repo root)

3. `./bin/optikon-api --scheme http --port 9000 --mock-base-path ./api/v0/mock/api`
