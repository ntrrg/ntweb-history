---
title: Balanceo de carga con Google Cloud Platform
author: ntrrg
date: 2020-04-09T14:00:00-04:00
description: 
tags:
  - tecnología
  - guías
  - redes
  - cloud
  - gcp
  - sysadmin
  - devops
series:
  - Google Cloud Platform
comments: true
toc: true
draft: true
---

```shell-session
$ export GCP_PROJECT="..."
$ export GCP_REGION="us-central1"
$ export GCP_ZONE="$GCP_REGION-a"

$ gcloud config set project "$GCP_PROJECT"
$ gcloud config set compute/region "$GCP_REGION"
$ gcloud config set compute/zone "$GCP_ZONE"
$ gcloud config list [--all]
```
