# NFS Client Provisioning

NFS Client provisioner is an automatic provisining `Persisetne Volumes` when `Persistent Volume Claims` are created.

## Installation

- Create a nfs server and get its ip address.
- Replace the <Nfs server ip> in `deployment.yml` by the obtainerd IP address.

```
kubectl apply -f deployment.yml
```

It will create a NFS-Provisioner which mounts the exported path in NFS server to `/persistentvolumes` in then provisioner.

## Provision NFS server locally

- Create a directory in local system.
  ```
    mkdir /srv/nfs/data
  ```
- Install nfs server in system.

  ```
  sudo apt install nfs-kernel-server

  sudo systemctl enable nfs-kernel-server

  sudo systemctl start nfs-kernel-server
  ```

- Edit /etc/exports file.

  ```
  srv/nfs/kubedata \*(rw,sync,no_subtree_check,no_root_squash,    no_all_squash,insecure)
  ```

  ```
  sudo exportfs -rav
  ```

  ```
  sudo exportfs -v
  ```

- Login to the worker nodes and mount the path.
