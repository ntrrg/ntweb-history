---
title: Cómo hacer contribuciones a Go
date: 2018-06-05T21:39:06-04:00
description: Contribuir a Go sin morir en el intento.
image: /uploads/gopher.png
categories:
  - tecnología
tags:
  - guías
  - programación
  - lenguajes-de-programación
  - go
  - contribuir-a-un-proyecto
---

Además de intimidante, hacer una contribución por primera vez en cualquier
proyecto open source suele ser complicado (aunque no lo sea), o por lo menos
para mí suele ser así. Más cuando se trata del código fuente de un lenguaje
de programación, que normalmente tiene a un equipo de genios trabajando sobre
él y sus flujos de trabajo suelen ser complejos.

[Go Contributing]: https://golang.org/doc/contribute.html

Para el caso de Go, hay una [página][Go Contributing] de la documentación
oficial que se enfoca en explicar los pasos para realizar contribuciones y
hacer de este trabajo una tarea más amigable, por lo que gran parte de las
instrucciones que se encuentran aquí son extraídas o están inspiradas en
ella.

{{< toc >}}

# Requisitos

{{% note "Proceso de contribución" %}}

<!--lint disable no-undefined-references no-shortcut-reference-link-->

[Go Release Cycle]: https://github.com/golang/go/wiki/Go-Release-Cycle
[Ciclo de desarrollo de Go]: {{< relref "blog/go-release-cycle.es.md" >}}

<!--lint enable no-undefined-references no-shortcut-reference-link-->

Antes de empezar a trabajar, es bueno tener una idea de como se desarrolla el
proceso de contribución para no sentir que todo funciona por arte de magia y
que hay crear cuentas en sitios específicos solo porque la guía lo dice.

<!--lint disable no-undefined-references-->

Cada versión estable de Go tiene un ciclo de desarrollo de seis meses, de los
cuales los primeros tres tienen como meta recibir todos los cambios posibles, y
durante los tres meses restantes solo se recibirán arreglos de fallas o
mejoras en la documentación; si quieren más detalles sobre el ciclo de
desarrollo, vayan a esta [página][Go Release Cycle] o lean este
[artículo][Ciclo de desarrollo de Go].

<!--lint enable no-undefined-references-->

El repositorio oficial del proyecto está alojado en
<https://go.googlesource.com/go>, la rama `master` es la que recibe todo el
código que será utilizado para generar la próxima versión estable; antes de
llegar ahí, el código debe ser auditado por miembros de la comunidad, esto se
hace desde <https://go-review.googlesource.com>; una vez auditado, dependiendo
de la fecha y de cuanto afecten al lenguaje, los nuevos cambios son mezclados
y esperan para ser liberados al final del ciclo de desarrollo activo; el estado
actual del proyecto puede verse desde <https://build.golang.org>.

{{% /note %}}

* [Una cuenta de Google](https://www.google.com/accounts/NewAccount).

* Iniciar sesión en <https://go.googlesource.com>, hacer clic en *Generate
  Password* y seguir las instrucciones.

* [Registrarse en el Gerrit de Go](https://go-review.googlesource.com/login/).

* Firmar el [Contributor License Agreement (CLA)](https://developers.google.com/open-source/cla/individual).

* Git.

* La última versión estable de Go instalada.

* Clonar el repositorio oficial o descargar el código de alguno de los
  subrepositorios si se pretende hacer el cambio en uno de ellos.

**Go:**

```shell-session
$ git clone https://go.googlesource.com/go
```

**Subrepositorio:**

```shell-session
$ go get -d -v golang.org/x/oauth2/...
```

---

Después de cumplir con todos los requisitos, se deben instalar dos
herramientas que facilitan algunas de las tareas durante el desarrollo:

**git-codereview:**

Permite usar Git para enviar los cambios al Gerrit de Go.

```shell-session
$ go get -u -v golang.org/x/review/git-codereview
```

En la sección [Realizar cambios](#realizar-cambios) muestro los comandos
básicos, pero si quieren conocer más sobre esta herramienta pueden ver su
ayuda ejecutando el siguiente comando

```shell-session
$ git codereview help
```

**go-contrib-init:**

Verifica el entorno de desarrollo y determina si hacen falta algunas
configuraciones. Solo es útil para el código fuente de Go.

```shell-session
$ go get -u -v golang.org/x/tools/cmd/go-contrib-init
```

# Realizar cambios

1\. Ir al código fuente de Go o al del subrepositorio

**Go:**

```shell-session
$ cd path/to/go
```

**Subrepositorio:**

```shell-session
$ cd "$GOPATH/src/golang.org/x/oauth2"
```

2\. Verificar que la rama activa sea `master`

```shell-session
$ git checkout master
$ git pull origin master
```

3\. Hacer los cambios.

```shell-session
$ EDITOR src/net/http/server.go # Go
$ EDITOR oauth2.go              # Subrepositorio
```

**Código:**

Todo archivo creado debe tener el texto de licencia

```go
// Copyright YYYY The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
```

Toda funcionalidad agregada o modificada debe ser probada

```shell-session
$ # Go
$ unset GOROOT
$ export GOROOT_BOOTSTRAP="$(go env GOROOT)"
$ go-contrib-init
$ cd src
$
$ # Para ejecutar solo algunas pruebas
$ export OLD_GOPATH="$(go env GOPATH)"
$ export GOPATH="$(dirname "$PWD")"
$ ./make.bash
$ "$GOPATH/bin/go" test -run '^TestHostHandlers$' -v net/http
$
$ # Para ejecutar todas las pruebas
$ ./all.bash

$ # Subrepositorio
$ go test -run '^TestAuthCodeURL$' -v ./... # Específica
$ go test -v ./...                          # Todas
```

**Documentación:**

Es suficiente con solo modificar los archivos.

4\. Agregar los cambios al área de confirmación de Git

```shell-session
$ git add .
```

5\. Confirmar los cambios

```shell-session
$ git codereview gofmt
$ git codereview change RAMA
```

`RAMA` será el nombre de la rama local donde se agregarán los cambios. Este
comando solicitará mensaje de la confirmación, que debería ser algo como:

```
net/http: descripción corta que complete "This change modifies Go to ___."

Descripción detallada que responda "¿Por qué se hace la
modificación?" y que no sobrepase los 72 caracteres.

Fixes: #123
```

Si se debe hacer una modificación después de confirmar, simplemente se
agregan los cambios al área de confirmación y se confirman sin especificar el
nombre de la rama

```shell-session
$ git add .
$ git codereview gofmt
$ git codereview change
```

6\. Verificar que el código esté actualizado

```shell-session
$ git codereview sync
```

Si hay cambios es recomendable ejecutar nuevamente las pruebas.

7\. Publicar los cambios para que sean auditados

```shell-session
$ git codereview mail
```

# Atribuciones

**Go Team.** *Contributing Guide.* <https://golang.org/doc/contribute.html>

