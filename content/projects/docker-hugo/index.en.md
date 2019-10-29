---
title: docker-hugo
description: Dockerized Hugo CLI.
metadata:
  source-code: https://github.com/ntrrg/docker-hugo
  license: MIT
kinds:
  - cli
  - container
techs:
  - docker
  - hugo
---

[![Docker Build Status](https://img.shields.io/docker/build/ntrrg/hugo.svg)](https://hub.docker.com/r/ntrrg/hugo)

[Hugo]: https://gohugo.io

**docker-hugo** is a Dockerized [Hugo][] CLI.

{{< toc >}}

<br/>

| Tag | Dockerfile |
| --: | :-- |
| `latest`, `0.59.0` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.59.0/Dockerfile) |

# Usage

```shell-session
$ docker run --rm -v /path/to/my/site:/site \
  ntrrg/hugo [OPTIONS] [COMMAND]
```

Any command from the Hugo CLI might be used, for extra information use `docker run --rm ntrrg/hugo help`
or see the [official documentation](https://gohugo.io/commands/).

{{% note %}}
Since the Hugo binary from the container is called by `root`, it is
recommendable to add the `-u` Docker flag.

```shell-session
$ docker run --rm -v /path/to/my/site:/site -u $(id -u $USER) \
  ntrrg/hugo [OPTIONS] [COMMAND]
```
{{% /note %}}

## Examples

* Create new Hugo skeleton

```shell-session
$ docker run --rm -v /path/to/my/site:/site ntrrg/hugo new site .
```

* Build Hugo project

```shell-session
$ docker run --rm -v /path/to/my/site:/site ntrrg/hugo
```

* Run Hugo server

```shell-session
$ docker run --rm -itp 1313:1313 -v /path/to/my/site:/site \
  ntrrg/hugo server -DEF --bind=0.0.0.0 \
    --baseUrl=/ --appendPort=false
```

* Run Hugo server with custom port

```shell-session
$ export PORT=8080
```

```shell-session
$ docker run --rm -itp $PORT:$PORT -v /path/to/my/site:/site \
  ntrrg/hugo server -DEF --bind=0.0.0.0 --port=$PORT \
    --baseUrl=/ --appendPort=false
```

# Acknowledgment

Working on this project I use/used:

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

* [Hugo][]

*Websocket for LiveReload using wrong port if Hugo binds to port 80.* <https://github.com/gohugoio/hugo/issues/2205>

