# Hubbl Workflow Sketch

## Pre-requisites
- Docker
- Minikube (or any k8s cluster)
- Golang 1.16+
- Kubectl
- Helm
- Git

### Docker

The steps to install `docker` can be found [https://docs.docker.com/desktop/mac/install/](here).

### Minikube

I personally like to use `minikube`. So for that, the installation guide can be found [https://minikube.sigs.k8s.io/docs/start/](here).

### Golang

The steps to install `Golang` can be found [https://go.dev/doc/install](here).

### Kubectl

The steps to install `kubectl` can be found [https://kubernetes.io/docs/tasks/tools/install-kubectl-macos/](here).

### Helm

The steps to install `helm` can be found [https://helm.sh/docs/intro/install/](here).

### Git

The steps to install `git` can be found [https://git-scm.com/book/en/v2/Getting-Started-Installing-Git](here).

## Setting up Temporal infra

### Start minikube

```
$ minikube start --kubernetes-version=v1.18.20 --cpus 6 --memory 5945
```

We are setting 6 cpus and 6gb of memory to keep the `Temporal` performance good.

### Creating temporal.io cluster

```
$ git clone git@github.com:temporalio/helm-charts.git
$ cd helm-charts
$ helm dependencies update
$ helm install \
    --set server.replicaCount=1 \
    --set cassandra.config.cluster_size=1 \
    --set prometheus.enabled=false \
    --set grafana.enabled=false \
    --set elasticsearch.enabled=false \
    temporaltest . --timeout 15m
```

We are going to use `Cassandra` as database for temporal. Other options are: `Mysql` and `Postgres`.

This configuration consumes limited resources and it is useful for small scale tests (such as using minikube).


To check if the cluster is working fine, you can run

```
$ kubectl get pods
NAME                                       READY   STATUS    RESTARTS   AGE
temporaltest-admintools-68c998fbfb-gthtd   1/1     Running   0          19m
temporaltest-cassandra-0                   1/1     Running   0          19m
temporaltest-frontend-68c6887f44-5gcr2     1/1     Running   4          19m
temporaltest-history-f8c8c46fc-nps9k       1/1     Running   3          19m
temporaltest-matching-56fd7ccdd6-4rhm6     1/1     Running   3          19m
temporaltest-web-79fd8b6c57-lsljw          1/1     Running   0          19m
temporaltest-worker-796b4fc9c6-zzzg2       1/1     Running   4          19m
```

And finally, to check `Temporal` working, you can shell into `admin-tool`

```
$ kubectl exec -it services/temporaltest-admintools /bin/bash
```

and check if everything is fine

```
bash-5.1# tctl namespace list
```

Another way is to view the `Temporal` web-ui, forwarding your machine's local port to the web service

``` 
$ kubectl port-forward services/temporaltest-web 8088:8088
```

After that, just go to a web-browser on `localhost:8088`.

## Running the workflow

The first step to run the workflow locally, is to start the worker. Since we are using `minikube` to instantiate all `Temporal` features, we need to expose one the services called `frontend`.


The `frontend` service from temporal is responsible to receive all requests to start workers, run workflows and so on.

To do that, we can simply run

```
$ kubectl port-forward services/temporaltest-frontend-headless 7233:7233
```

You need to create a namespace (if you didn't do yet) to use with `Temporal`. You can achieve this doing the following:

- connect to the admintools: `$ kubectl exec -it services/temporaltest-admintools /bin/bash`
- register the namespace: `tctl namespace register default`

