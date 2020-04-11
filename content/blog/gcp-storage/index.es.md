---
title: Almacenamiento en Google Cloud Platform
author: ntrrg
date: 2020-04-09T14:05:00-04:00
description: Gestión de reden en GCP
tags:
  - tecnología
  - guías
  - cloud
  - gcp
  - redes
  - sysadmin
  - devops
series:
  - Google Cloud Platform
comments: true
toc: true
draft: true
---

```shell-session
$ gsutil mb gs://BUCKET_NAME
$ gsutil ls PATH
$ gsutil cp SRC DEST
$ gsutil acl ch -u AllUsers:R PATH
$ gsutil acl ch -d AllUsers PATH
$ gsutil rm PATH
```
