---
title: ntWeb
description: A small site with great intentions.
metadata:
  source-code: https://github.com/ntrrg/ntweb
  license: MIT
tags:
  - go
  - html
  - css
  - javascript
  - website
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

**ntWeb** is a small site with great intentions. It provides a portfolio, a
blogging platform, Go packages vanity URLs and a JSON API.

Probably anyone could think of this as a really simple project and its author
is a bad frontend developer, and actually, it is true in some way because it is
not my profile, but its simplicity is intentional (and it is written with
modern web technologies).

[Hugo]: https://gohugo.io
[Go]: https://golang.org
[Mage]: https://magefile.org

The website is created by [Hugo][] and some parts of the content are generated
by [Go][] and [Mage][] (specifically the projects and the Go packages vanity
URLs).

For been printed in the web browser, the home page downloads ~28 KB (including
the favicons), a blog article downloads ~215 KB (including its images and the
favicons). It is functional even without JavaScript and can be visualized on
TUI web browsers.

{{< figure src="images/elinks-home-en.png" class="block" caption="Home page on TUI" >}}

{{< figure src="images/elinks-blog-es.png" class="block" caption="Blog on TUI" >}}

{{< figure src="images/elinks-article-es.png" class="block" caption="Article on TUI" >}}

Its content may be fetched as JSON, so creating new complex frontends is
trivial without the need for content duplication. Since it is a static API, it
has good performance because of the lack of computing complexity (just reading
a file), it is more resistant to cyber-attacks because there is no database
server and it can be hosted in any service supporting static files.

# API

Almost any HTML page has a JSON output format, it could be fetched by appending
`index.json` to its URL, so the JSON URL of this page should be
`https://nt.web.ve/en/projects/ntweb/index.json`.

Every page has the following properties:

`url` (string):
: Resource URL.

`kind` (string):
: Resource type. It may be one of `home`, `section`, `taxonomyTerm`, `taxonomy`
or `page`.

`type` (string):
: Content type. It may be one of `blog`, `gallery`, `projects`, `tag` or
`page`.

`lang` (string):
: Resource language.

`title` (string):
: Resource title.

`image` (string):
: Resource image URL.

`params` (string):
: Resource frontmatter parameters. This may be different from page types.

`content` (base64 encoded string):
: Resource rendered Markdown content. Notice that this contains a UTF-8 string
and JavaScript strings are UTF-16.

`data` (object):
: Resource specific data. For the main page, this contains all the sections,
taxonomies and top-level pages; for collections this contains its elements and
pagination information; and for single elements this is an empty object.

`altLang` (array of objects):
: Resource translations. Every object has the `lang` and `url` properties.

`altMediaType` (array of objects):
: Resource alternative formats. Every object has the `mediaType` and `url`
properties.

```shell-session
$ wget -qO - https://nt.web.ve/en/projects/ntweb/index.json | jq
{
  "url": "https://nt.web.ve/en/projects/ntweb/",
  "kind": "page",
  "type": "projects",
  "lang": "en",
  "title": "ntWeb",
  "params": {
    "comments": true,
    "description": "A small site with great intentions.",
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
      "website",
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
      "lang": "es",
      "url": "https://nt.web.ve/es/projects/ntweb/"
    }
  ],
  "altMediaType": [
    {
      "mediaType": "text/html",
      "url": "https://nt.web.ve/en/projects/ntweb/"
    }
  ]
}
```

## Endpoints

### Main

<https://nt.web.ve/en/index.json>

Retrieves all the top-level elements. See [API](#api) for more details about
common properties.

`data.sections` (array of objects):
: Website sections. Every object has the `url`, `title` and `pages`
properties. The `pages` property is the count of pages inside the section.

`data.taxonomies` (array of objects):
: Website taxonomies. Every object has the `url`, `title` and `terms`
properties. The `terms` property is the count of terms inside the taxonomy.

`data.pages` (array of objects):
: Top-level pages. Every object has the `url` and `title` properties.

### Collections

<https://nt.web.ve/en/:section/index.json>

<https://nt.web.ve/en/:section/page/:pageNumber/index.json>

<https://nt.web.ve/en/tags/:tag/index.json>

<https://nt.web.ve/en/tags/:tag/page/:pageNumber/index.json>

{{< note "Parameters" >}}

`section`:
: Section name. This could be one of `blog`, `gallery` or `projects`.

`tag`:
: Tag name. This should be [an existent tag](/en/tags).

`pageNumber`:
: Page number. The first page is retrieved without `page/:pageNumber/`.

{{< /note >}}

Retrieves the list of elements from the given collection. See [API](#api) for
more details about common properties.

`data.pages` (array of objects):
: List of elements. Every object has the `url`, `title`, `image`,
`publishDate`, `date`, `description` and `tags` properties. The `tags` property
is an array of tags assigned to the element.

`data.pagination` (object):
: Pagination object. It contains the `prefix`, `first`, `prev`, `next` and
`last` properties.

### Elements

<https://nt.web.ve/en/:section/:title/index.json>

{{< note "Parameters" >}}

`section`:
: Section name. This could be one of `blog`, `gallery` or `projects`.

`title`:
: URL encoded element title.

{{< /note >}}

Retrieves a single element. See [API](#api) for more details about common
properties.

### Search index

<https://nt.web.ve/en/search-index/index.json>

Retrieves all the indexable elements for search engines. It is an array of
objects, and every object has the `url`, `type`, `title`, `description`,
`content` and `tags` properties. The value of the `content` property is base64
encoded, notice that this contains a UTF-8 string and JavaScript strings are
UTF-16.

**Note:** this doesn't have the common properties.

# Offline mode

**Requirements:**

* Hugo >= 0.65

Get the source code

```shell-session
$ # Package
$ wget https://github.com/ntrrg/ntweb/archive/master.tar.gz
$
$ # Git repository
$ git clone --depth 1 https://github.com/ntrrg/ntweb.git
```

Run the Hugo server in the project root directory

```shell-session
$ hugo server
```

Go to <http://localhost:1313/> with a browser.

## Docker

```shell-session
$ docker run --rm -p 1313:80 ntrrg/ntweb
```

Go to <http://localhost:1313/> with a browser.

# Acknowledgment

Working on this project I use/used:

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

