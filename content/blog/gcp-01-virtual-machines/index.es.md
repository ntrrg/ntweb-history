---
title: Máquinas virtuales en Google Cloud Platform
author: ntrrg
date: 2020-04-09T12:35:00-04:00
description: Máquinas virtuales en GCP.
tags:
  - tecnología
  - aprendizaje
  - fundamentos
  - cloud
  - gcp
  - máquinas-virtuales
  - sysadmin
  - devops
series:
  - Google Cloud Platform
comments: true
toc: true
draft: true
---

Compute Engine (GCE)

```shell-session
$ gcloud compute project-info describe --project PROJECT_ID

$ gcloud compute instances list

$ # Virtual machines

$ gcloud compute instances create NAME \
    [--image-project debian-cloud --image-family debian-9] \
    [--machine-type TYPE] [--zone ZONE] [--tags TAGS] \
    [--subnet SUBNET_NAME] \
    [--metadata-from-file startup-script=FILE]

$ gcloud compute instances describe NAME

$ gcloud compute ssh NAME [--zone ZONE]

$ gcloud compute target-instances create NAME --instance INSTANCE_NAME

$ # Unmanaged groups

$ gcloud compute instance-groups unmanaged create NAME \
    [--zone ZONE] --instances INSTANCE_NAMES

$ # Managed groups

$ gcloud compute instance-templates create NAME \
    [--machine-type TYPE] [--tags TAGS] \
    --metadata-from-file startup-script=FILE

$ gcloud compute target-pools create NAME [--region REGION]

$ gcloud compute instance-groups managed create NAME [--zone ZONE] \
    --base-instance-name PREFIX --template TEMPLATE_NAME \
    --size NUMBER_OF_INSTANCES --target-pool POOL_NAME
```

Kubernetes Engine (GKE)

```shell-session
$ gcloud container clusters create NAME (--zone ZONE | --region REGION)

$ gcloud container clusters get-credentials NAME \
    (--zone ZONE | --region REGION)

$ kubectl get service

$ kubectl create deployment NAME --image CONTAINER_IMAGE

$ kubectl expose deployment NAME \
    --protocol (TCP|UDP) --port PORT [--target-port CONTAINER_PORT] \
    [--type ClusterIP|LoadBalancer]
```

