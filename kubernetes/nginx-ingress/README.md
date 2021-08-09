# Nginx ingress

Nginx ingress can be deployed by two ways.

- [Nginx ingress](#nginx-ingress)
  - [Manual](#manual)
  - [Helm](#helm)

## Manual

- Deployed manually using the manifest files.
- Clone the nginx ingress git repo and deploy each resources required

## Helm

```
# Add nignx ingress repo to helm.
helm repo add <nginx ingress repo url>

# Update helm repo.
helm repo update
```

```
# Deploy nignx ingress.

helm install nginx ingress-nginx/ingress-nginx

```

This will deploy nginx ingress pods along with sevices, deployment and replicaset.

After this we need to create nginx-contorller rule for redirection.

```
kubectl apply -f < admin | auth | manager >.yml

```

Run the above command for each app, it will create rules to redirect the request for nginx-ingres-contorller.

Once all of the above steps are completed

```
kubectl get svc -o wide

```

There will be nginx ingress service along with the assigned **external ip** and **port**

Add the external ip in the `/etc/hosts` file in local system and assign it to a domain name for developnment use.

```
eg: 172.16.16.240   app.admin.com
```

Now use the url `app.admin.com` to access the app.

**Note:** `External IP` should be assigned to domain name otherwise it won't work.
