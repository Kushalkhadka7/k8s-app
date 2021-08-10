# Based on Memory metrics

Auto scaling k8's pod using memory metrics.

## Installation

- Install HPA for each deployment.

```
kubectl apply -f admin-hpa.yml

kubectl apply -f admin-hpa.yml

kubectl apply -f admin-hpa.yml
```

- In `hpa.yml` file `scaleTargetRef` should be given properly.

```
 scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: <Deployment name for which we are using the HPA>
```
