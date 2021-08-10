# MetalLB

MetalLB is a load-balancer implementation for bare metal Kubernetes clusters, using standard routing protocols.

## Installation

```
# Namespace.

kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.10.2/manifests/namespace.yaml

# MetalLB

kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.10.2/manifests/metallb.yaml

OR

kubectl apply -f metallb.yml -n metallb-system

```

Add configmap for metalLb.

```
kubectl apply -f config.yml -n metallb-system
```

It will deploy MetalLB in `metallb-system` namespace.

### This installation is requeired because in bare metal if we expose service as loadbalancer an `externalIP` is needed to be assigned, which is done by `metallb`.
