---
title: ntWeb
description: Un pequeño sitio web con grandes intenciones.
metadata:
  source-code: https://github.com/ntrrg/ntweb
  license: MIT
tags:
  - go
  - go-templates
  - html
  - css
  - javascript
  - sitio-web
  - json-api
  - hugo
  - mage
  - docker
  - github-actions
  - netlify
toc: true
comments: true
---

[![GitHub Actions](https://github.com/ntrrg/ntweb/workflows/Go/badge.svg)](https://github.com/ntrrg/ntweb/actions?query=workflow:Go)
[![GitHub Actions](https://github.com/ntrrg/ntweb/workflows/Hugo/badge.svg)](https://github.com/ntrrg/ntweb/actions?query=workflow:Hugo)
[![Netlify Status](https://api.netlify.com/api/v1/badges/2f18cd17-5e78-45fa-a95d-0ae120ffc603/deploy-status)](https://app.netlify.com/sites/ntweb/deploys)

**ntWeb** es un pequeño sitio web con grandes intenciones. Provee un
portafolio, una plataforma de blog, URLs personalizadas para paquetes Go y una
API JSON.

Probablemente cualquiera podría pensar que este proyecto es muy sencillo y que
su autor no tiene buenas capacidades en desarrollo frontend, y de hecho es
cierto en algunos aspectos porque no es mi especialidad, pero su sencillez es
intencional (y está escrito con tecnologías web modernas).

[Hugo]: https://gohugo.io
[Go]: https://golang.org
[Mage]: https://magefile.org

El sitio web es creado con [Hugo][] y parte de su contenido es generado con
[Go][] y [Mage][] (los proyectos y las URLs personalizadas de los paquetes Go).

Para mostrar la pagina principal solo se descargan ~28 KB (incluyendo los
favicons), para un artículo del blog se descargan ~215 KB (incluyendo sus
imágenes, los favicons y desactivando los comentarios). Es usable incluso sin
JavaScript y puede visualizarse en navegadores web TUI.

{{< figure src="images/elinks-home-es.png" class="block" caption="Página principal en TUI" >}}

{{< figure src="images/elinks-blog-es.png" class="block" caption="Blog en TUI" >}}

{{< figure src="images/elinks-article-es.png" class="block" caption="Artículo en TUI" >}}

Todo su contenido también puede obtenerse como JSON, por lo que es posible
crear nuevos Frontends más complejos sin necesidad de tener que repetir el
contenido. Por ser una API estática se obtiene un buen rendimiento pues no se
realiza ningún computo extra al de leer el archivo, es bastante seguro porque
no hay servidor de base de datos que pueda ser atacado y puede hospedarse en
cualquier servicio que soporte archivos estáticos.

```shell-session
$ wget -qO - https://nt.web.ve/es/blog/index.json | jq
```

# Uso

**Requerimientos:**

* Hugo >= 0.61.0

Descargar el código fuente y ejecutar el siguente comando en la carpeta raíz
del proyecto:

```shell-session
$ hugo server
```

Ir a <http://localhost:1313/> con un navegador web.

## Docker

```shell-session
$ docker run --rm -p 1313:80 ntrrg/ntweb
```

Ir a <http://localhost:1313/> con un navegador web.

# Atribuciones

Trabando en este proyecto uso/usé:

* [Debian](https://www.debian.org/)

* [XFCE](https://xfce.org/)

* [st](https://st.suckless.org/)

* [Zsh](http://www.zsh.org/)

* [GNU Screen](https://www.gnu.org/software/screen)

* [Git](https://git-scm.com/)

* [EditorConfig](http://editorconfig.org/)

* [Vim](https://www.vim.org/)

* [GNU make](https://www.gnu.org/software/make/)

* [Hugo](https://gohugo.io)

* [Chrome](https://www.google.com/chrome/browser/desktop/index.html)

* [GitHub](https://github.com)

* [Gitlab](https://gitlab.com/)

* [Gogs](https://gogs.io/)

* [Travis CI](https://travis-ci.org)

* [Drone](https://drone.io/)

* [Docker](https://docker.com)

* [Netlify](https://www.netlify.com/)

* [github-markdown-css](https://github.com/sindresorhus/github-markdown-css)

* [Normalize.css](https://necolas.github.io/normalize.css/)

* [Google Tag Manager](https://www.google.com/analytics/tag-manager/)

* [Forestry](https://forestry.io) 

* [FontAwesome](https://fontawesome.com/) 

* [Disqus](https://disqus.com/) 

* [MathJax](https://www.mathjax.org/) 

* [Mage](https://magefile.org/)

* [Termux](https://termux.com)

* [Firefox Developer Edition](https://www.mozilla.org/en-US/firefox/developer/)

* [GitHub Actions](https://github.com/features/actions)

* [Lunr](https://lunrjs.com)

