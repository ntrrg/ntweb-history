---
title: Gestión de redes en Google Cloud Platform
author: ntrrg
date: 2020-04-09T14:00:00-04:00
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
$ # IP addresses

$ gcloud compute addresses create NAME \
    [--ip-version (IPV4|IPV6)] \
    [--addresses ADDRESSES] [--prefix-length PREFIX] \
    [--network NETWORK_NAME | --subnet SUBNET_NAME] \
    (--region REGION | --global)

$ # Subnets

$ gcloud compute networks create NAME --subnet-mode custom

$ gcloud compute networks subnets create NAME \
  --network NETWORK_NAME --range IP_RANGE --region REGION

$ # Firewall rules

$ gcloud compute firewall-rules list

$ gcloud compute firewall-rules create NAME \
    [--network NETWORK_NAME] [--direction INGRESS|EGRESS] \
    [--source-ranges IP_RANGE] [--source-tags TAGS] [--target-tags TAGS] \
    (--allow (tcp|udp)[:PORT[-PORT]]] |
    --action ALLOW|DENY --rules (tcp|udp)[:PORT[-PORT]])

$ # Forwarding rules

$ gcloud compute forwarding-rules list

$ # Protocol Forwarding (single VM)

$ gcloud compute forwarding-rules create NAME \
    [--address IP_ADDRESS] [--ip-protocol (TCP|UDP)] \
    --ports PORTS --target-instance INSTANCE_NAME \
    (--region REGION | --global)

$ # L3 Load Balancer (VMs pool)

$ gcloud compute forwarding-rules create NAME \
    [--address IP_ADDRESS] [--ip-protocol (TCP|UDP)] \
    --ports PORTS --target-pool POOL_NAME \
    (--region REGION | --global)

$ # L7 Load Balancer (VMs pool)

$ gcloud compute instance-groups (managed | unmanaged) \
    set-named-ports GROUP_NAME --named-ports (http|https|http2):PORT \
    [--zone ZONE | --region REGION]

$ gcloud compute (http-health-checks | https-health-checks) create NAME

$ gcloud compute backend-services create NAME \
    --protocol (HTTP|HTTPS|HTTP2) \
    (--http-health-checks | --https-health-checks) HEALTH_CHECK_NAME \
    --global

$ gcloud compute backend-services add-backend NAME \
    --instance-group GROUP_NAME --instance-group-zone ZONE \
    --global

$ gcloud compute url-maps create NAME --default-service BACKEND_NAME

$ gcloud compute url-maps add-path-matcher URLMAP_NAME \
  --default-service BACKEND_NAME \
  --path-matcher-name PATHMATCHER_NAME --path-rules="PATH=BACKEND_NAME"

$ # HTTP

$ gcloud compute target-http-proxies create NAME --url-map URLMAP_NAME

$ gcloud compute forwarding-rules create NAME \
    [--address IP_ADDRESS] \
    --ports PORTS --target-http-proxy PROXY_NAME \
    --global

$ # HTTPS

$ gcloud compute ssl-certificates create NAME \
    --private-key FILE --certificate FILE

$ gcloud compute target-https-proxies create NAME \
    --url-map URLMAP_NAME --ssl-certificates CERTS_NAME

$ gcloud compute forwarding-rules create NAME \
    [--address IP_ADDRESS] \
    --ports PORTS --target-https-proxy PROXY_NAME \
    --global
```

