# K8's Infrastructure

## Table of Contents

- [K8's Infrastructure](#k8s-infrastructure)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Development Setup](#development-setup)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)

## Overview

K8's app arichitecture consist of 3 worker node cluster, contorlled by 2 master nodes. The main idea to use k8's for container orchesteration is that it is :

- Flexible
- Open source
- Scalable
- Resilience for container-based workloads
- Preparation of workloads to move to the cloud

## Development Setup

### Prerequisites

- [Docker](https://docs.docker.com/install/)
- [Git](https://git-scm.com/downloads)

### Installation

- Git clone the repository.
- [Setup development enviornment](docs/setup/README.md)
- Once the setup is completed our cluster will be ready, now we can deploy our application.
- [Run the app locally](apps/README.md)
- Deploy the apps to `kubernetes cluster`.
  - [Admin](kubernetes/admin/README.md)
  - [Auth](kubernetes/auth/README.md)
  - [Manager](kubernetes/manager/README.md)
- Deploy NginxIngress or Istio.
  - [NginxIngress](kubernetes/nginx-ingress/README.md)
  - [Istio](kubernetes/istio/README.md)
- If there is no external Ip assigned to the service which are of type load balancer, the deploy metallb.
  - [Metallb](kubernetes/metallb/README.md)
- [Kubernetes Dashboard](kubernetes/dashboard/README.md)
- [Monitoring](kubernetes/promethousandgrafana/README.md)
- [Logging](kubernetes/efk/README.md)
- [DymanicNFSProvisining](kubernetes/NFS/README.md)

> Note: Rest of the other deployments can be used according to our use.
