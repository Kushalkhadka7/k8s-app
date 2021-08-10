# Storage

## Persistent Volume claim

```
kubectl apply -f pvc.yml
```

## Storage Class

```
kubectl apply -f defaultStorageClass.yml

kubectl apply -f storageClass.yml
```

- Default storage class needs to be created because if we forgot to mention the storage class on pvc it will use the default storage class.

- Now whenever we use PVC on the deployment it will automatically create PV.
