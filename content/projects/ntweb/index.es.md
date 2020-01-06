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

# API

Casi todas las páginas HTML tienen JSON como formato alternativo, que puede ser
obtenido agregando `index.json` al final de su URL. Por ejemplo, para esta
página, la URL de su formato JSON sería `https://nt.web.ve/es/projects/ntweb/index.json`.

Cada página tiene los siguientes atributos:

`url` (cadena):
: URL del recurso.

`kind` (cadena):
: Tipo de recurso. Puede ser `home`, `section`, `taxonomyTerm`, `taxonomy` o
`page`.

`type` (cadena):
: Tipo de contenido. Puede ser `blog`, `gallery`, `projects`, `tag` o `page`.

`lang` (cadena):
: Idioma del recurso.

`title` (cadena):
: Título del recurso.

`params` (cadena):
: Parámetros del frontmatter del recurso. Puede variar segur el tipo del
contenido.

`content` (cadena codificada en base64):
: Contenido Markdown procesado del recurso.

`data` (objeto):
: Datos específicos del recurso. Para la página principal, contiene todos los
elementos de primer nivel; para colecciones, contiene sus elementos e
información sobre la paginación; y para recursos comunes, es un objeto vacío.

`altLang` (arreglo de objetos):
: Traducciones del recurso. Cada objeto tiene las propiedades `lang` y `url`.

`altMediaType` (arreglo de objetos):
: Formatos alternativos del recurso. Cada objeto tiene las propiedades
`mediaType` y `url`.

```shell-session
$ wget -qO - https://nt.web.ve/es/projects/ntweb/index.json | jq
{
  "url": "https://nt.web.ve/es/projects/ntweb/",
  "kind": "page",
  "type": "projects",
  "lang": "es",
  "title": "ntWeb",
  "params": {
    "comments": true,
    "description": "Un pequeño sitio web con grandes intenciones.",
    "draft": false,
    "iscjklanguage": false,
    "metadata": {
      "license": "MIT",
      "source-code": "https://github.com/ntrrg/ntweb"
    },
    "tags": [
      "go",
      "go-templates",
      "html",
      "css",
      "javascript",
      "sitio-web",
      "json-api",
      "hugo",
      "mage",
      "docker",
      "github-actions",
      "netlify"
    ],
    "title": "ntWeb",
    "toc": true
  },
  "content": "...",
  "data": {},
  "altLang": [
    {
      "lang": "en",
      "url": "https://nt.web.ve/en/projects/ntweb/"
    }
  ],
  "altMediaType": [
    {
      "mediaType": "text/html",
      "url": "https://nt.web.ve/es/projects/ntweb/"
    }
  ]
}
```

## Endpoints

### Principal

<https://nt.web.ve/es/index.json>

Obtiene todos los elementos de primer nivel. Ver [API](#api) para mas
información sobre las propiedades en común.

`data.sections` (arreglo de objetos):
: Secciones. Cada objeto tiene las propiedades `url`, `title` y `pages`. La
propiedad `pages` es la cantidad de páginas dentro de la sección.

`data.taxonomies` (arreglo de objetos):
: Taxonomías. Cada objeto tiene las propiedades `url`, `title` y `terms`. La
propiedad `terms` es la cantidad de términos dentro de la taxonomía.

`data.pages` (arreglo de objetos):
: Páginas de primer nivel. Cada objeto tiene las propiedades `url` y `title`.

### Colecciones

<https://nt.web.ve/es/:section/index.json>

<https://nt.web.ve/es/:section/page/:pageNumber/index.json>

<https://nt.web.ve/es/tags/:tag/index.json>

<https://nt.web.ve/es/tags/:tag/page/:pageNumber/index.json>

{{< note "Parámetros" >}}

`section`:
: Nombre de la sección. Puede ser `blog`, `gallery` o `projects`.

`tag`:
: Nombre de la etiqueta. Debe ser [una etiqueta existente](/es/tags).

`pageNumber`:
: Número de página. La primera página es obtenida sin `page/:pageNumber/`.

{{< /note >}}

Obtiene la lista de elementos de una colección específica. Ver [API](#api) para
mas información sobre las propiedades en común.

`data.pages` (arreglo de objetos):
: Lista de elementos. Cada página tiene las propiedades `url`, `title`,
`publishDate`, `date`, `description` y `tags`. La propiedad `tags` es un
arreglo de etiquetas asignadas al elemento.

`data.pagination` (objeto):
: Información de paginación. Contiene las propiedades `prefix`, `first`,
`prev`, `next` y `last`.

### Elementos

<https://nt.web.ve/es/:section/:title/index.json>

{{< note "Parámetros" >}}

`section`:
: Sección. Puede ser `blog`, `gallery` o `projects`.

`title`:
: Título codificado en URL del recurso.

{{< /note >}}

Obtiene un elemento. Ver [API](#api) para mas información sobre las propiedades
en común.

### Índice de búsqueda

<https://nt.web.ve/es/search-index/index.json>

Obtiene todos los elementos que pueden ser procesados por motores de búsqueda.
Es un arreglo de objetos, cada objeto tiene las propiedades `url`, `title`,
`description` y `content`. El valor de la propiedad `content` esta codificado
con base64.

**Nota:** este recurso no tiene las propiedades comunes de otros recursos

# Modo sin conexión

**Requerimientos:**

* Hugo >= 0.61.0

Descargar el código fuente

```shell-session
$ # Paquete
$ wget https://github.com/ntrrg/ntweb/archive/master.tar.gz
$
$ # Repositorio Git
$ git clone --depth 1 https://github.com/ntrrg/ntweb.git
```

Ejecutar el servidor de Hugo en la carpeta raíz del proyecto:

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

