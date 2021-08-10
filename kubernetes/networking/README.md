# Networking

Calico and Flannel are overlay netrork that needs to be deployed on the k8's cluster that is running in bare metal.
If these network are not deployed `pods` won't be able to communicate with one another.

## Installation

- Calico

  ```
  kubectl apply -f calico.yml
  ```

- Flannel
  ```
  kubectl apply -f calico.yml
  ```

Any one of them should be installed not both.
