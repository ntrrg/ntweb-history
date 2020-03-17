---
publishdate: 2018-05-06T22:07:39-04:00
date: 2020-02-25T05:50:25-04:00
metadata:
  source-code: https://github.com/ntrrg/docker-hugo
  license: MIT
title: docker-hugo
description: CLI de Hugo en Docker.
tags:
  - contenedores
  - docker
  - hugo
toc: true
comments: true
---

[![Docker Build Status](https://img.shields.io/docker/build/ntrrg/hugo.svg)](https://hub.docker.com/r/ntrrg/hugo)

[Hugo]: https://gohugo.io

**docker-hugo** es el CLI de [Hugo][] en Docker.

| Etiqueta | Dockerfile |
| --: | :-- |
| `latest`, `0.65.3` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.65.3/Dockerfile) |
| `0.65.2` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.65.2/Dockerfile) |
| `0.64.0` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.64.0/Dockerfile) |
| `0.63.2` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.63.2/Dockerfile) |
| `0.62.2` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.62.2/Dockerfile) |
| `0.61.0` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.61.0/Dockerfile) |
| `0.60.1` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.60.1/Dockerfile) |
| `0.59.1` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.59.1/Dockerfile) |
| `0.59.0` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.59.0/Dockerfile) |

# Uso

```shell-session
$ docker run --rm -v /ruta/a/mi/sitio/:/site/ \
  ntrrg/hugo [OPCIONES] [COMANDO]
```

Puede usarse cualquier comando del CLI de Hugo, para más información ejecutar `docker run --rm ntrrg/hugo help`
o ver la [documentación oficial](https://gohugo.io/commands/).

{{< note >}}
Como el binario de Hugo del contenedor es ejecutado por `root`, es recomendable
agregar la opción `-u` de Docker.

```shell-session
$ docker run --rm -v /ruta/a/mi/sitio/:/site/ \
  -u $(id -u $USER) \
  -v ${TMPDIR:-/tmp/}:/tmp/ \
  ntrrg/hugo [OPCIONES] [COMANDO]
```
{{< /note >}}

## Ejemplos

* Crear el esqueleto de un projecto Hugo

```shell-session
$ docker run --rm -v /ruta/a/mi/sitio/:/site/ \
      ntrrg/hugo new site .
```

* Construir un proyecto Hugo

```shell-session
$ docker run --rm -v /ruta/a/mi/sitio/:/site/ ntrrg/hugo
```

* Ejecutar el servidor de Hugo

```shell-session
$ docker run --rm -i -t -p 1313:1313 \
      -v /ruta/a/mi/sitio/:/site/ \
      ntrrg/hugo server -DEF --bind=0.0.0.0 \
        --baseUrl=/ --appendPort=false
```

* Ejecutar el servidor de Hugo en un puerto personalizado

```shell-session
$ export PORT=8080
```

```shell-session
$ docker run --rm -i -t -p $PORT:$PORT \
      -v /path/to/my/site:/site \
      ntrrg/hugo server -DEF --bind=0.0.0.0 --port=$PORT \
        --baseUrl=/ --appendPort=false
```

# Atribuciones

Trabajando en este proyecto uso/usé:

* [Debian](https://www.debian.org/)

* [XFCE](https://xfce.org/)

* [Vim](https://www.vim.org/)

* [Chrome](https://www.google.com/chrome/browser/desktop/index.html)

* [st](https://st.suckless.org/)

* [Zsh](http://www.zsh.org/)

* [GNU Screen](https://www.gnu.org/software/screen)

* [Git](https://git-scm.com/)

* [EditorConfig](http://editorconfig.org/)

* [Docker](https://docker.com)

* [Github](https://github.com)

* [Hugo](https://gohugo.io)

* [Firefox Developer Edition](https://www.mozilla.org/en-US/firefox/developer/)

*Websocket for LiveReload using wrong port if Hugo binds to port 80.* <https://github.com/gohugoio/hugo/issues/2205>


