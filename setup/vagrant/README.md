# Setup Kubernetes Cluster using kubeadm with Vagrant

## Overview

The cluster is setup using kubeadm, in the virtual enviornment that can be run locally for development. It uses **Centos7** machines as nodes. To provision the virtual machines we are using vagrant, which is a tool used for managing virtual machines.(https://www.vagrantup.com/docs/cli/provision).

## Setup

### Prerequisites

- [Docker](https://docs.docker.com/install/)
- [Git](https://git-scm.com/downloads)
- Virtual box installed

### Installation

- Once the virtual box is installed in the system, clone this reop and navigate to this folder.
- Install vagrant in the system.
- Once vagrnant is install open the **VagrantFile** and in the **v.memory** and **v.cpus** section make sure to update the appropriate value for current use.
- After all of the above steps, we can make our K8's cluster.

```
# Uses the shell files to make the cluster.

vagrant up
```

- This will make the cluster all the configurations in the shell files eg. master.sh and worker.sh
- We can change the configuration in these files as per our need, by default it will create a cluster with `1` master node and `2` worker nodes.
- After this our cluster will be ready for use.

#### To stop the env

```
# It will pause the provisining.

vagrant suspend
```
