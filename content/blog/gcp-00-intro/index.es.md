---
title: Introducción a Google Cloud Platform
author: ntrrg
date: 2020-04-09T12:30:00-04:00
description: Computación en la Nube con la ayuda de uno de los gigantes de Internet. Conceptos básicos y una reseña histórica de GCP.
tags:
  - tecnología
  - aprendizaje
  - fundamentos
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

Consola y gcloud.

```shell-session
$ gcloud init

$ gcloud components list
$ gcloud components install beta
$ gcloud beta GROUP
$ gcloud components install alpha
$ gcloud alpha GROUP
```

IAM

```shell-session
$ gcloud auth list
```

CLI configuration

```shell-session
$ export GCP_PROJECT="..."
$ export GCP_REGION="us-central1"
$ export GCP_ZONE="$GCP_REGION-a"

$ gcloud config list [--all] [KEY]
$ gcloud config set KEY VALUE
$ gcloud config set project "$GCP_PROJECT"
$ gcloud config set compute/region "$GCP_REGION"
$ gcloud config set compute/zone "$GCP_ZONE"
```

# Control de Indetidades y Acceso (IAM)

Entre lo roles predefinidos se encuentran:

* *Observador (Viewer):*

* *Editor (Editor):*

* *Propietario (Owner):*

# Atribuciones

**Google Cloud Team.** *Google Cloud Training.* <https://google.qwiklabs.com>

