# Horizontal Pod AutoScaling

The Horizontal Pod Autoscaler automatically scales the number of Pods in a replication controller, deployment, replica set or stateful set based on observed CPU utilization (or, with custom metrics support, on some other application-provided metrics).

https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/

Before the installation `metric server` should be deployed to the kubernetes cluster to collect the metirces.

## Installation

- [CPU](#prerequisites)
- [Memory](#installation)

Aftet the installatin we need to change the deployment files for each app and add resource field in the deployment file.

```
resources:
    limits:
      memory: '2000Mi'
    requests:
      memory: '50Mi'
```
