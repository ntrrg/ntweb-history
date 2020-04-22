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
que la ubicación geográfica de los centros de datos de Google donde correrá el
recurso.

Los participantes del proyecto son usuarios a los que se les pueden asignar
roles específicos que les permitirán interactuar con los recursos del proyecto.
También es posible crear cuentas de servicio, estas pueden ser usadas por
aplicaciones para acceder y hasta manejar los recursos del proyecto.

# Control de Indetidades y Acceso (IAM)

Entre lo roles predefinidos se encuentran:

* *Observador (Viewer):*

* *Editor (Editor):*

* *Propietario (Owner):*

# Interfaces de Usuario

## Consola

## CLIs (`gcloud`, `gsutil`, `bq`)

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
$ gcloud config get-value KEY
$ gcloud config set KEY VALUE
$ gcloud config set project "$GCP_PROJECT"
$ gcloud config set GROUP/region "$GCP_REGION"
$ gcloud config set GROUP/zone "$GCP_ZONE"
```

### IAM

```shell-session
$ gcloud auth list

$ gcloud auth login [ACCOUNT]
```

Service accounts

```shell-session
$ gcloud iam service-accounts create NAME [--display-name DISPLAY_NAME]

$ gcloud projects add-iam-policy-binding PROJECT_ID \
    --member serviceAccount:ACCOUNT@PROJECT_ID.iam.gserviceaccount.com \
    --role ROLE

$ gcloud iam service-accounts disable NAME

$ gcloud iam service-accounts enable NAME

$ # Get access token

$ gcloud iam service-accounts keys create key.json --iam-account ACCOUNT

$ gcloud auth activate-service-account --key-file key.json

$ gcloud auth print-access-token
```

### Tools

#### Git repositories

```shell-session
$ gcloud source repos create NAME

$ gcloud source repos clone NAME
```

### Computing

#### Compute Engine (GCE)

```shell-session
$ gcloud compute project-info describe --project PROJECT_ID

$ gcloud compute instances list
```

Virtual machines

```shell-session
$ gcloud compute images list [--project IMAGE_PROJECT] [--no-standard-images]

$ gcloud compute instances create NAME [--preemptible] \
    [--image-project IMAGE_PROJECT --image IMAGE] \
    [--machine-type TYPE] [--zone ZONE] [--tags TAGS] \
    [--subnet SUBNETWORK] \
    [--metadata-from-file startup-script=FILE]

$ gcloud compute instances create-with-container NAME [--preemptible] \
    --image-project cos-cloud --image IMAGE \
    --container-image CONTAINER_IMAGE \
    [--machine-type TYPE] [--zone ZONE] [--tags TAGS] \
    [--container-restart-policy never|on-failure|always] \
    [--container-env KEY=VALUE] [--container-env-file PATH] \
    [--container-mount-host-path host-path=PATH,mount-path=PATH[,mode=(rw|ro)] \
    [--container-mount-tmpfs=mount-path=MOUNTPATH] \
    [--container-stdin] [--container-tty] [--container-privileged] \
    [--container-arg ARGS] [--container-command ENTRYPOINT]

$ gcloud compute instances describe NAME

$ gcloud compute ssh NAME [--zone ZONE]

$ gcloud compute target-instances create NAME --instance INSTANCE
```

Unmanaged groups

```shell-session
$ gcloud compute instance-groups unmanaged create NAME \
    [--zone ZONE] --instances INSTANCES
```

Managed groups

```shell-session
$ gcloud compute instance-templates create NAME \
    [--machine-type TYPE] [--tags TAGS] \
    --metadata-from-file startup-script=FILE

$ gcloud compute target-pools create NAME [--region REGION]

$ gcloud compute instance-groups managed create NAME [--zone ZONE] \
    --base-instance-name PREFIX --template TEMPLATE \
    --size NUMBER_OF_INSTANCES --target-pool POOL
```

#### Kubernetes Engine (GKE)

```shell-session
$ gcloud container clusters create NAME [--zone ZONE | --region REGION] \
    [--cluster-version VERSION] [--node-version VERSION] \
    [--num-nodes NODES] [--node-locations ZONES] \
    [--machine-type TYPE] [--tags TAGS] [--preemptible] \
    [--metadata-from-file startup-script=FILE] \
    [--disk-type pd-standard|pd-ssd] [--disk-size SIZE] \
    [--network NETWORK] [--subnet SUBNETWORK] \
    [--enable-stackdriver-kubernetes]
    
$ gcloud auth configure-docker

$ gcloud container clusters get-credentials NAME \
    (--zone ZONE | --region REGION)
```

```shellsession
$ kubectl create serviceaccount NAME [--namespace NAMESPACE]

$ kubectl create clusterrolebinding NAME \
    --clusterrole ROLE --user [NAMESPACE:]ACCOUNT

$ kubectl create ns NAME

$ kubectl get nodes

$ kubectl get services

$ kubectl get deployments

$ kubectl get replicasets

$ kubectl get pods [--show-labels] [-l LABELS]

$ kubectl explain OBJECT

$ kubectl create -f PATH [-n NAMESPACE]

$ kubectl apply -f PATH [-n NAMESPACE]

$ kubectl edit deployments NAME

$ # Rollouts

$ kubectl rollout history OBJECT

$ kubectl rollout status OBJECT

$ kubectl rollout pause OBJECT

$ kubectl rollout resume OBJECT

$ kubectl rollout undo OBJECT

$ # Manual management

$ kubectl create deployments NAME --image CONTAINER_IMAGE

$ kubectl exec POD --stdin --tty -c CONTAINER COMMAND

$ kubectl label pods POD LABELS

$ kubectl expose deployments NAME \
    --protocol (TCP|UDP) --port PORT [--target-port CONTAINER_PORT] \
    [--type ClusterIP|NodePort|LoadBalancer]

$ kubectl scale deployments NAME --replicas REPLICAS
```

```shell-sesion
$ kubectl create serviceaccount tiller --namespace kube-system

$ kubectl create clusterrolebinding tiller-admin-binding \
    --clusterrole=cluster-admin --serviceaccount=kube-system:tiller

$ helm init --service-account tiller

$ helm repo update

$ helm install REPO --name NAME --version VERSION -f PATH --wait
```

```shell-sesion
$ terraform init

$ terraform plan

$ terraform apply
```

#### App Engine (GAE)

```shell-session
$ git clone https://github.com/GoogleCloudPlatform/golang-samples.git
$ cd golang-samples/appengine/go11x/helloworld
$ gcloud components install app-engine-go

$ gcloud app deploy [FILE] [--bucket BUCKET] [--ignore-file FILE]

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
    [--source DIR] --stage-bucket BUCKET \
    --runtime (nodejs8|nodejs10|python37|go111|go113) \
    [--entry-point FUNCTION_NAME_IN_SOURCE] \
    [--clear-env-vars | --env-vars-file FILE | --set-env-vars KEY=VALUE] \
    [--remove-env-vars KEYS] [--update-env-vars KEY=VALUE] \
    [--memory (128|256|512|1024|2048)MB] [--timeout SECONDS] \
    [--max-instances INSTANCES | --clear-max-instances] \
    [--allow-unauthenticated] \
    [--trigger-bucket BUCKET | --trigger-topic PUBSUB_TOPIC |
    (--trigger-event EVENT --trigger-resource RESOURCE) |
    --trigger-http ]

$ gcloud functions call NAME [--region REGION] --data JSON_DATA

$ gcloud functions logs read NAME [--region REGION] \
    [--start-time START_TIME] [--end-time END_TIME] \
    [--min-log-level debug|info|error] \
    [--filter EXPRESSION] [--page-size SIZE] [--limit LIMIT] [--sort-by FIELD]
```

### Networking

#### IP addresses

```shell-session
$ gcloud compute addresses create NAME \
    [--ip-version (IPV4|IPV6)] \
    [--addresses ADDRESSES] [--prefix-length PREFIX] \
    [--network NETWORK | --subnet SUBNETWORK] \
    (--region REGION | --global)
```

#### Subnets

```shell-session
$ gcloud compute networks create NAME --subnet-mode custom

$ gcloud compute networks subnets create NAME \
  --network NETWORK --range IP_RANGE --region REGION
```

#### Firewall rules

```shell-session
$ gcloud compute firewall-rules list

$ gcloud compute firewall-rules create NAME \
    [--network NETWORK] [--direction INGRESS|EGRESS] \
    [--source-ranges IP_RANGE] [--source-tags TAGS] [--target-tags TAGS] \
    (--allow (tcp|udp)[:PORT[-PORT]]] |
    --action ALLOW|DENY --rules (tcp|udp)[:PORT[-PORT]])
```

#### Forwarding rules

```shell-session
$ gcloud compute forwarding-rules list
```

Protocol Forwarding (single VM)

```shell-session
$ gcloud compute forwarding-rules create NAME \
    [--address IP_ADDRESS] [--ip-protocol (TCP|UDP)] \
    --ports PORTS --target-instance INSTANCE \
    (--region REGION | --global)
```

L3 Network Balancer (VMs pool)

```shell-session
$ gcloud compute forwarding-rules create NAME \
    [--address IP_ADDRESS] [--ip-protocol (TCP|UDP)] \
    --ports PORTS --target-pool POOL \
    --region REGION
```

L7 Global HTTP/HTTPS/HTTP2 Load Balancer (VMs pool)

```shell-session
$ gcloud compute instance-groups (managed | unmanaged) \
    set-named-ports GROUP --named-ports (http|https|http2):PORT \
    [--zone ZONE | --region REGION]

$ gcloud compute (http-health-checks | https-health-checks) create NAME

$ gcloud compute backend-services create NAME \
    --protocol (HTTP|HTTPS|HTTP2) \
    (--http-health-checks | --https-health-checks) HEALTH_CHECK \
    --global

$ gcloud compute backend-services add-backend NAME \
    --instance-group GROUP --instance-group-zone ZONE \
    --global

$ gcloud compute url-maps create NAME --default-service BACKEND

$ gcloud compute url-maps add-path-matcher URLMAP \
  --default-service BACKEND \
  --path-matcher-name PATHMATCHER --path-rules="PATH=BACKEND"

$ # HTTP

$ gcloud compute target-http-proxies create NAME --url-map URLMAP

$ gcloud compute forwarding-rules create NAME \
    [--address IP_ADDRESS] \
    --ports PORTS --target-http-proxy PROXY \
    --global

$ # HTTPS

$ gcloud compute ssl-certificates create NAME \
    --private-key FILE --certificate FILE

$ gcloud compute target-https-proxies create NAME \
    --url-map URLMAP --ssl-certificates CERTS

$ gcloud compute forwarding-rules create NAME \
    [--address IP_ADDRESS] \
    --ports PORTS --target-https-proxy PROXY \
    --global
```

#### VPC Network Peering

```shell-session
$ gcloud compute networks peerings create NAME --auto-create-routes \
    --network NETWORK --peer-network PEER_NETWORK [--peer-project PROJECT]
```

### Storage

#### Storage (GCS)

```shell-session
$ gsutil mb gs://NAME
$ gsutil ls PATH
$ gsutil cp SRC DEST
$ gsutil acl ch -u AllUsers:R PATH
$ gsutil acl ch -d AllUsers PATH
$ gsutil rm PATH
```

#### Datastore

```shell-session
$ gcloud datastore indexes create PATH

$ gcloud datastore indexes cleanup PATH
```

#### Cloud SQL

```shell-session
$ gcloud sql instances list

$ gcloud sql instances clone SOURCE DESTINATION

$ gcloud sql instances create NAME [--region REGION] \
    [--authorized-networks NETWORKS] [--availability-type regional|zonal] \
    [--backup] [--backup-location LOCATION] [--backup-start-time] \
    [--database-version VERSION] [--database-flags FLAGS] \
    [--replication synchronous|asynchronous] \
    [--replica-type READ|FAILOVER | --master-instance-name MASTER_DB] \
    [--failover-replica-name NEW_REPLICA] \
    [--tier MACHINE_TYPE | --cpu MAX_CORES --memory MEMORY] \
    [--storage-size SIZE] [--storage-type SSD|HDD] \
    [--root-password PASSWORD]

$ gcloud sql instances restart NAME

$ gcloud sql backups create --instance INSTANCE --description DESCRIPTION

$ gcloud sql backups restore ID --restore-instance INSTANCE

$ gcloud sql connect INSTANCE [--user USER] [--database DATABASE]
```

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

$ gcloud pubsub subscriptions create --topic TOPIC NAME \
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

### Security

### CFT Scorecard

```shell-session
$ gcloud services enable cloudasset.googleapis.com

$ gcloud asset export {--organization | --folder | --project} VALUE \
    --output-path=PATH/resource_inventory.json --content-type=resource

$ gcloud asset export {--organization | --folder | --project} VALUE \
    --output-path=PATH/iam_inventory.json --content-type=iam-policy

$ git clone https://github.com/forseti-security/policy-library.git

$ cp policy-library/samples/storage_blacklist_public.yaml \
    policy-library/policies/constraints/

$ wget -O cft 'https://storage.googleapis.com/cft-cli/latest/cft-linux-amd64'

$ chmod +x cft

$ ./cft scorecard --policy-path=policy-library/ \
    {--bucket | --dir-path} VALUE
```

# Atribuciones

**Google Cloud Team.** *Google Cloud Training.* <https://google.qwiklabs.com>

