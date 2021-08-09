# Deploy Manager to Kubernetes

## Installation

- Make sure the namespace in which we want to deploy this app (By default it will deploy in default namespace).

### Deploy config map

```
# Create the configmap used by the app.

kubectl apply -f config.yml
```

### Deploy secrets.

#### Create the secrets used by the app (These secrets should be in base64 encoding otherwise it wont work).

```
echo <secret> | base64  # copy the output and replace in secret.yml

# Create secrets used by app.

kubectl apply -f secret.yml
```

### Deploy app.

```
kubectl apply -f deployment.yml
```

### Deploy service.

```
kubectl apply -f service.yml
```

### Verify installation

```
kubectl get all -n <namespace> -o wide

```

This will dispaly all the deployed pods, services, deployments along with replicasets in the namespace.

Or they can be verified individually.

```
kubectl get svc

kubectl get configmap

kubectl get secret

kubectl get pods

kubectl get deploy

kubectl get replicaset
```

Once this is verified the application is deployed to the cluster, but it won't be accessible from the local system, since we haven't exposed the service.

```
# App will be accessible by localhost:<hostport>

kubectl port-forward <pod_name> <host_port>:<container:port> --namespace <namespace in which app is deployed>
```

```
# Find out the node ip address in which the pod is running.

<node ip address>:<port in which the app is running>
```

```
# Expose the service as load balancer, the loadbalancing solution will append an external ip from which the app will be accessible.
```
