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
- Setup local develpement enviornemt ([Setup](docs/setup/README.md)).
- Once the setup is completed our cluster will be ready, now we can deploy our application.
- Deploy the applications ([Deploy](#Deploy))
