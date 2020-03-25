---
publishdate: 2018-05-06T22:07:39-04:00
date: 2020-03-24T11:41:42-04:00
metadata:
  source-code: https://github.com/ntrrg/docker-hugo
  license: MIT
title: docker-hugo
description: Dockerized Hugo CLI.
tags:
  - containers
  - docker
  - hugo
toc: true
comments: true
---

[![Docker Build Status](https://img.shields.io/docker/build/ntrrg/hugo.svg)](https://hub.docker.com/r/ntrrg/hugo)

[Hugo]: https://gohugo.io

**docker-hugo** is a Dockerized [Hugo][] CLI.

| Tag | Dockerfile |
| --: | :-- |
| `latest`, `0.68.3` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.68.3/Dockerfile) |
| `0.67.1` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.67.1/Dockerfile) |
| `0.65.3` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.65.3/Dockerfile) |
| `0.65.2` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.65.2/Dockerfile) |
| `0.64.0` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.64.0/Dockerfile) |
| `0.63.2` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.63.2/Dockerfile) |
| `0.62.2` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.62.2/Dockerfile) |
| `0.61.0` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.61.0/Dockerfile) |
| `0.60.1` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.60.1/Dockerfile) |
| `0.59.1` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.59.1/Dockerfile) |
| `0.59.0` | [Dockerfile](https://github.com/ntrrg/docker-hugo/blob/0.59.0/Dockerfile) |

# Usage

```shell-session
$ docker run --rm -v /path/to/my/site/:/site/ \
  ntrrg/hugo [OPTIONS] [COMMAND]
```

Any command from the Hugo CLI might be used, for extra information use `docker run --rm ntrrg/hugo help`
or see the [official documentation](https://gohugo.io/commands/).

{{< note >}}
Since the Hugo binary from the container is called by `root`, it is
recommendable to add the `-u` Docker flag.

```shell-session
$ docker run --rm -v /path/to/my/site/:/site/ \
  -u $(id -u $USER) \
  -v ${TMPDIR:-/tmp/}:/tmp/ \
  ntrrg/hugo [OPTIONS] [COMMAND]
```
{{< /note >}}

## Examples

* Create a new Hugo skeleton

```shell-session
$ docker run --rm -v /path/to/my/site/:/site/ \
      ntrrg/hugo new site .
```

* Build a Hugo project

```shell-session
$ docker run --rm -v /path/to/my/site/:/site/ ntrrg/hugo
```

* Run the Hugo server

```shell-session
$ docker run --rm -i -t -p 1313:1313 \
      -v /path/to/my/site/:/site/ \
      ntrrg/hugo server -DEF --baseUrl=/ \
        --bind=0.0.0.0 --appendPort=false
```

* Run the Hugo server with a custom port

```shell-session
$ export PORT=8080
```

```shell-session
$ docker run --rm -i -t -p $PORT:$PORT \
      -v /path/to/my/site/:/site/ \
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

* [Hugo](https://gohugo.io)

* [Firefox Developer Edition](https://www.mozilla.org/en-US/firefox/developer/)

*Websocket for LiveReload using wrong port if Hugo binds to port 80.* <https://github.com/gohugoio/hugo/issues/2205>

