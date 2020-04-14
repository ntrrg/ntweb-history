---
title: Introducción a Google Cloud Platform
author: ntrrg
date: 2020-04-09T12:30:00-04:00
description: Computación en la Nube con la ayuda de uno de los gigantes de Internet. Conceptos básicos y una reseña histórica de Google Cloud Platform.
tags:
  - tecnología
  - aprendizaje
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

# Conceptos básicos

El elemento principal de la plataforma son los proyectos, un proyecto es un
conjunto de participantes, recursos y servicios que darán vida a una o varias
aplicaciones. Es posible tener muchas aplicaciones dentro de un solo proyecto,
pero se recomienda mantener esta cantidad distribuida entre diferentes
proyectos para facilitar su mantenimiento e impulsar la productividad de los
miembros del equipo.

Los recursos y servicios son todas las utilidades que ofrece la plataforma,
entre ellas máquinas virtuales, orquestadores de contenedores, bases de datos,
dispositivos de almacenamiento, herramientas de monitoreo y muchísimas cosas
más. También incluyen APIs como la de Google Maps y las de inteligencia
artificial para el reconocimiento de voz y procesamiento de imágenes.

Ya que la Computación en la Nube promueve el alcance global de las
aplicaciones, es muy común que al momento de solicitar recursos se deba indicar
la región y la zona en la que se quieren desplegar. Estos atributos no son más
que la ubicación geográfica de los centros de datos de Google.

Los participantes del proyecto son usuarios a los que se les pueden asignar
roles específicos que les permitirán interactuar con los recursos y servicios
del proyecto.

# Interfaces de Usuario

## Consola

## CLIs (`gcloud`, `gsutil`)

```shell-session
$ gcloud init [--console-only]

$ gcloud info

$ gcloud components list

$ gcloud components install beta
$ gcloud beta GROUP

$ gcloud components install alpha
$ gcloud alpha GROUP
```

Configuration

```shell-session
$ export GCP_PROJECT="..."
$ export GCP_REGION="us-central1"
$ export GCP_ZONE="$GCP_REGION-a"

$ gcloud config list [--all] [KEY]
$ gcloud config set KEY VALUE
$ gcloud config set project "$GCP_PROJECT"
$ gcloud config set GROUP/region "$GCP_REGION"
$ gcloud config set GROUP/zone "$GCP_ZONE"
```

### IAM

```shell-session
$ gcloud auth list
```

### Compute power

#### Compute Engine (GCE)

```shell-session
$ gcloud compute project-info describe --project PROJECT_ID

$ gcloud compute instances list
```

Virtual machines

```shell-session
$ gcloud compute instances create NAME \
    [--image-project debian-cloud --image-family debian-(9|10)] \
    [--machine-type TYPE] [--zone ZONE] [--tags TAGS] \
    [--subnet SUBNET_NAME] \
    [--metadata-from-file startup-script=FILE]

$ gcloud compute instances describe NAME

$ gcloud compute ssh NAME [--zone ZONE]

$ gcloud compute target-instances create NAME --instance INSTANCE_NAME
```

Unmanaged groups

```shell-session
$ gcloud compute instance-groups unmanaged create NAME \
    [--zone ZONE] --instances INSTANCE_NAMES
```

Managed groups

```shell-session
$ gcloud compute instance-templates create NAME \
    [--machine-type TYPE] [--tags TAGS] \
    --metadata-from-file startup-script=FILE

$ gcloud compute target-pools create NAME [--region REGION]

$ gcloud compute instance-groups managed create NAME [--zone ZONE] \
    --base-instance-name PREFIX --template TEMPLATE_NAME \
    --size NUMBER_OF_INSTANCES --target-pool POOL_NAME
```

#### Kubernetes Engine (GKE)

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

#### App Engine

```shell-session
$ git clone https://github.com/GoogleCloudPlatform/golang-samples.git
$ cd golang-samples/appengine/go11x/helloworld
$ gcloud components install app-engine-go

$ gcloud app deploy [FILE] [--bucket BUCKET_NAME] [--ignore-file FILE]

$ gcloud app browse

$ gcloud app logs read [--service default] [--limit LIMIT] \
  [--logs stderr,stdout,crash.log,nginx.request,request_log] \
  [--level critical|error|warning|info|debug|any]

$ gcloud app logs tail [--service default] \
  [--logs stderr,stdout,crash.log,nginx.request,request_log] \
  [--level critical|error|warning|info|debug|any]
```

#### Functions (GCF)

```shell-session
$ gcloud functions describe NAME

$ gcloud functions event-types list

$ gcloud functions deploy NAME [--region REGION] \
  [--source DIR] --stage-bucket BUCKET_NAME \
  --runtime (nodejs8|nodejs10|python37|go111|go113) \
  [--entry-point FUNCTION_NAME_IN_SOURCE] \
  [--clear-env-vars | --env-vars-file FILE | --set-env-vars KEY=VALUE] \
  [--remove-env-vars KEYS] [--update-env-vars KEY=VALUE] \
  [--memory (128|256|512|1024|2048)MB] [--timeout SECONDS] \
  [--max-instances INSTANCES | --clear-max-instances] \
  [--allow-unauthenticated] \
  [--trigger-bucket BUCKET_NAME | --trigger-topic PUBSUB_TOPIC |
  (--trigger-event EVENT_NAME --trigger-resource RESOURCE_NAME) |
  --trigger-http ]

$ gcloud functions call NAME [--region REGION] --data JSON_DATA

$ gcloud functions logs read NAME [--region REGION] \
  [--start-time START_TIME] [--end-time END_TIME] \
  [--min-log-level debug|info|error] \
  [--filter EXPRESSION] [--page-size SIZE] [--limit LIMIT] [--sort-by FIELD]
```

### Networking

IP addresses

```shell-session
$ gcloud compute addresses create NAME \
    [--ip-version (IPV4|IPV6)] \
    [--addresses ADDRESSES] [--prefix-length PREFIX] \
    [--network NETWORK_NAME | --subnet SUBNET_NAME] \
    (--region REGION | --global)
```

Subnets

```shell-session
$ gcloud compute networks create NAME --subnet-mode custom

$ gcloud compute networks subnets create NAME \
  --network NETWORK_NAME --range IP_RANGE --region REGION
```

Firewall rules

```shell-session
$ gcloud compute firewall-rules list

$ gcloud compute firewall-rules create NAME \
    [--network NETWORK_NAME] [--direction INGRESS|EGRESS] \
    [--source-ranges IP_RANGE] [--source-tags TAGS] [--target-tags TAGS] \
    (--allow (tcp|udp)[:PORT[-PORT]]] |
    --action ALLOW|DENY --rules (tcp|udp)[:PORT[-PORT]])
```

Forwarding rules

```shell-session
$ gcloud compute forwarding-rules list

$ # Protocol Forwarding (single VM)

$ gcloud compute forwarding-rules create NAME \
    [--address IP_ADDRESS] [--ip-protocol (TCP|UDP)] \
    --ports PORTS --target-instance INSTANCE_NAME \
    (--region REGION | --global)

$ # L3 Network Balancer (VMs pool)

$ gcloud compute forwarding-rules create NAME \
    [--address IP_ADDRESS] [--ip-protocol (TCP|UDP)] \
    --ports PORTS --target-pool POOL_NAME \
    --region REGION

$ # L7 Global HTTP/HTTPS/HTTP2 Load Balancer (VMs pool)

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

$   # HTTP

$ gcloud compute target-http-proxies create NAME --url-map URLMAP_NAME

$ gcloud compute forwarding-rules create NAME \
    [--address IP_ADDRESS] \
    --ports PORTS --target-http-proxy PROXY_NAME \
    --global

$   # HTTPS

$ gcloud compute ssl-certificates create NAME \
    --private-key FILE --certificate FILE

$ gcloud compute target-https-proxies create NAME \
    --url-map URLMAP_NAME --ssl-certificates CERTS_NAME

$ gcloud compute forwarding-rules create NAME \
    [--address IP_ADDRESS] \
    --ports PORTS --target-https-proxy PROXY_NAME \
    --global
```

### Storage (GCS)

```shell-session
$ gsutil mb gs://BUCKET_NAME
$ gsutil ls PATH
$ gsutil cp SRC DEST
$ gsutil acl ch -u AllUsers:R PATH
$ gsutil acl ch -d AllUsers PATH
$ gsutil rm PATH
```

# Control de Indetidades y Acceso (IAM)

Entre lo roles predefinidos se encuentran:

* *Observador (Viewer):*

* *Editor (Editor):*

* *Propietario (Owner):*

### Big Data

#### Pub/Sub

```shell-session
$ gcloud pubsub topics list \
  [--filter EXPRESSION] [--page-size SIZE] [--limit LIMIT] [--sort-by FIELD]

$ gcloud pubsub topics create NAME

$ gcloud pubsub topics list-subscriptions NAME \
  [--filter EXPRESSION] [--page-size SIZE] [--limit LIMIT] [--sort-by FIELD]

$ gcloud pubsub topics publish NAME --message DATA

$ gcloud pubsub topics delete NAME

$ gcloud pubsub subscriptions list \
  [--filter EXPRESSION] [--page-size SIZE] [--limit LIMIT] [--sort-by FIELD]

$ gcloud pubsub subscriptions create --topic TOPIC_NAME NAME \
  [--expiration-period EXPIRATION_PERIOD] [--retain-acked-messages] \
  [--message-retention-duration MESSAGE_RETENTION_DURATION]

$ gcloud pubsub subscriptions pull NAME [--auto-ack] \
  [--filter EXPRESSION] [--page-size SIZE] [--limit LIMIT] [--sort-by FIELD]

$ gcloud pubsub subscriptions delete NAME
```

#### BigQuery

```shell-session
$ bq
```

# Atribuciones

**Google Cloud Team.** *Google Cloud Training.* <https://google.qwiklabs.com>

