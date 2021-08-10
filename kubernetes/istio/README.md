# Istio

## Overview

Istio is an open source service mesh that layers transparently onto existing distributed applications. Istio’s powerful features provide a uniform and more efficient way to secure, connect, and monitor services.Istio is the path to load balancing, service-to-service communication.
https://istio.io/latest/about/service-mesh/

## Installation

Istio can be installed using helm but it is depreciated and istio enforces to use `istioctl` to deploy istio to kubernetes cluster.

- Install istioctl binary

  ```
  # Download istio.

  curl -L https://istio.io/downloadIstio | sh -

  # Cd to isito directory.

  # Copy the binary to /usr/local/bin.

  sudo mv istioctl /usr/local/bin

  # Verify installation.

  istioctl --version

  # Output.
  client version: 1.10.3
  control plane version: 1.10.3
  data plane version: 1.10.3 (9 proxies)
  ```

- Now we can use `istioclt` to deploy resources in our kubernetes cluster.
- Istio provide different profiles to use while installing the istio in cluster.

```
istioctl profile list

# Output

Istio configuration profiles:
    default
    demo
    empty
    external
    minimal
    openshift
    preview
    remote
```

- Select one of the profile to deploy istio to cluster.For development purpose it's better to use `demo`,

```
# Install istio components.

istioctl install --set profile=demo
```

- It will deploy different istio resources along with `istio-ingress`.
- It will deploy everthing in `istio-system` namespace.

```
# Verify the installation.

kubectl get all -o wide -n istio-system
```

- After this we need to crete the virtual sevice and gateway.

```
kubectl apply -f virtual-service.yml -n istio-system

kubectl apply -f gateway.yml istio-system
```

- After this update each serivices port

```
# eg:

apiVersion: v1
kind: Service
metadata:
  name: admin
  labels:
    app: admin
    service: admin
spec:
  selector:
    app: admin
  ports:
    - name: http
      port: 3000
```

Now the application will be accessible.
