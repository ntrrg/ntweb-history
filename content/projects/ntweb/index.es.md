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
crear nuevos frontends más complejos sin necesidad de tener que repetir el
contenido. Por ser una API estática se obtiene un buen rendimiento pues no se
realiza ningún computo extra al de leer el archivo, es bastante segura porque
no hay servidor de base de datos que pueda ser atacado y puede hospedarse en
cualquier servicio que soporte archivos estáticos.

```shell-session
$ wget -qO - https://nt.web.ve/es/blog/index.json | jq
{
  "url": "https://nt.web.ve/es/blog/",
  "kind": "section",
  "type": "blog",
  "lang": "es",
  "title": "Blog",
  "params": {
    "draft": false,
    "iscjklanguage": false,
    "title": "Blog"
  },
  "content": "",
  "data": {
    "pages": [
      {
        "url": "https://nt.web.ve/es/blog/por-qu%C3%A9-usar-contenedores/",
        "title": "¿Por qué usar contenedores?",
        "author": "ntrrg",
        "publishdate": "2019-05-14T22:54:00-07:00",
        "date": "2019-11-14T11:35:00-04:00",
        "description": "La forma más fácil de implemetar aplicaciones para programadores y administradores de sistemas.",
        "tags": [
          "tecnología",
          "entornos-de-desarrollo",
          "contenedores",
          "backend",
          "devops",
          "sysadmin"
        ]
      },
      {
        "url": "https://nt.web.ve/es/blog/instalar-go-1.13/",
        "title": "Instalar Go 1.13",
        "author": "ntrrg",
        "publishdate": "2019-11-12T10:10:00-04:00",
        "date": "2019-11-12T10:10:00-04:00",
        "description": "Instalar Go es bastante sencillo, con solo seguir unas pocas instrucciones cualquiera puede hacerlo.",
        "tags": [
          "tecnología",
          "guías",
          "instalaciones",
          "lenguajes-de-programación",
          "go"
        ]
      },
      {
        "url": "https://nt.web.ve/es/blog/instalar-go-1.12/",
        "title": "Instalar Go 1.12",
        "author": "ntrrg",
        "publishdate": "2019-06-01T10:10:00-07:00",
        "date": "2019-11-12T09:00:00-04:00",
        "description": "Instalar Go es bastante sencillo, con solo seguir unas pocas instrucciones cualquiera puede hacerlo.",
        "tags": [
          "tecnología",
          "guías",
          "instalaciones",
          "lenguajes-de-programación",
          "go"
        ]
      },
      {
        "url": "https://nt.web.ve/es/blog/instalar-go-1.11/",
        "title": "Instalar Go 1.11",
        "author": "ntrrg",
        "publishdate": "2019-06-01T10:00:00-07:00",
        "date": "2019-11-12T06:30:00-04:00",
        "description": "Instalar Go es bastante sencillo, con solo seguir unas pocas instrucciones cualquiera puede hacerlo.",
        "tags": [
          "tecnología",
          "guías",
          "instalaciones",
          "lenguajes-de-programación",
          "go"
        ]
      },
      {
        "url": "https://nt.web.ve/es/blog/lista-de-anime/",
        "title": "Lista de Anime",
        "author": null,
        "publishdate": "2019-05-31T11:30:00-07:00",
        "date": "2019-05-31T11:30:00-07:00",
        "description": "Una lista de anime que ya vi, estoy viendo o tal vez veré.",
        "tags": [
          "anime"
        ]
      },
      {
        "url": "https://nt.web.ve/es/blog/nombre-de-punto-de-acceso-apn/",
        "title": "Nombre de Punto de Acceso (APN)",
        "author": null,
        "publishdate": "2019-05-31T10:30:00-07:00",
        "date": "2019-05-31T10:30:00-07:00",
        "description": "En caso de No Internet en su dispositivo móvil, rompa el vidrio.",
        "tags": [
          "redes",
          "android"
        ]
      },
      {
        "url": "https://nt.web.ve/es/blog/go-golang/",
        "title": "Go (Golang)",
        "author": null,
        "publishdate": "2018-09-23T16:40:00-04:00",
        "date": "2018-09-23T16:40:00-04:00",
        "description": "Es un lenguaje de código abierto, minimalista y de alto rendimiento; su fuerte es la concurrencia.",
        "tags": [
          "referencias",
          "programación",
          "lenguajes-de-programación",
          "go",
          "programación-de-sistemas",
          "programación-web",
          "backend"
        ]
      },
      {
        "url": "https://nt.web.ve/es/blog/instalar-go-1.10/",
        "title": "Instalar Go 1.10",
        "author": "ntrrg",
        "publishdate": "2018-09-09T18:15:47-04:00",
        "date": "2018-09-09T18:15:47-04:00",
        "description": "Instalar Go es bastante sencillo, con solo seguir unas pocas instrucciones cualquiera puede hacerlo.",
        "tags": [
          "tecnología",
          "guías",
          "instalaciones",
          "lenguajes-de-programación",
          "go"
        ]
      },
      {
        "url": "https://nt.web.ve/es/blog/c%C3%B3mo-hacer-contribuciones-a-go/",
        "title": "Cómo hacer contribuciones a Go",
        "author": null,
        "publishdate": "2018-07-06T12:42:47-04:00",
        "date": "2018-07-06T12:42:47-04:00",
        "description": "Cómo contribuir a Go sin morir en el intento.",
        "tags": [
          "guías",
          "programación",
          "lenguajes-de-programación",
          "go",
          "contribuir-a-un-proyecto"
        ]
      },
      {
        "url": "https://nt.web.ve/es/blog/representaci%C3%B3n-de-n%C3%BAmeros-de-punto-flotante/",
        "title": "Representación de números de punto flotante",
        "author": null,
        "publishdate": "2018-06-25T23:15:14-04:00",
        "date": "2018-06-25T23:15:14-04:00",
        "description": "La conversión de números de punto flotante decimales a binarios es algo compleja, para llevarla acabo se debe seguir el estándar IEEE 754.",
        "tags": [
          "aritmética",
          "ciencia-de-la-computación",
          "números-binarios"
        ]
      }
    ],
    "paginationPrefix": "https://nt.web.ve/es/blog/page/",
    "fist": "1",
    "prev": "",
    "next": "2",
    "last": "2"
  },
  "altLang": [
    {
      "lang": "en",
      "url": "https://nt.web.ve/en/blog/"
    }
  ],
  "altMediaType": [
    {
      "mediaType": "text/html",
      "url": "https://nt.web.ve/es/blog/"
    },
    {
      "mediaType": "application/rss+xml",
      "url": "https://nt.web.ve/es/blog/index.xml"
    }
  ]
}
```

# Uso

**Requerimientos:**

* Hugo >= 0.61.0

Descargar el código fuente y ejecutar el siguiente comando en la carpeta raíz
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

