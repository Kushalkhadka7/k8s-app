# Kubernetes Dashboard

Dashboard is a web-based Kubernetes user interface. You can use Dashboard to deploy containerized applications to a Kubernetes cluster, troubleshoot your containerized application, and manage the cluster resources.

https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/

## Installation

- Create a service account with admin access.

  ```
   kubectl apply -f ../accounts/dashboard-rbac.yml
  ```

- Deploy kubernetes dashboard.

  ```
  kubectl apply -f deployement.yml

  OR

  kubectl apply -f https://raw.githubusercontent.com/ kubernetes/dashboard/v2.2.0/aio/deploy/recommended.yaml
  ```

It will deploy the resources on ` kubetnetes-dashboard` namespace.

- We need to get the secret token to login to the dashboard.

  ```
  kubectl -n kube-system describe sa dashboard-admin
  ```

- I will a secret token, use it to obtain the accesstoken used to login to the kubernetes dashboard.

  ```
  kiubectl -n kube-system describe secret <token secret>
  ```

- It will give the actual token, we can use it to login to the kubernetes dashboard.

  ```
  kc -n kubernetes-dashboard port-forward <pod_name>  <port>:8443
  ```

- Go to `localhost:<defined_port>` and use the token to login.

We need to create the service account because previously created service account doesnt have full access to the cluster which may prevent the dashboard access to certain resources.
