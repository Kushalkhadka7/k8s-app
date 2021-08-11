# Pod Disruptions

It will help to run the defined amount of pods, while node used for maintainance.If node2 is used for maintainance and the pod disription is set to 2 then it will make sure there are 2 pods running event though node2 is not available. It will schedule these 2 pods on another node.

https://kubernetes.io/docs/concepts/workloads/pods/disruptions/

## Installation

    kubectl apply -f admin.yml

    kubectl apply -f auth.yml

    kubectl apply -f manager.yml
