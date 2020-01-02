---
title: ntWeb
description: A small site with great intentions.
metadata:
  source-code: https://github.com/ntrrg/ntweb
  license: MIT
tags:
  - go
  - go-templates
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
the favicons), a blog article downloads ~215 KB (including its images, the
favicons and disabling the comments). It is functional even without JavaScript
and can be visualized on TUI web browsers.

{{< figure src="images/elinks-home-en.png" class="block" caption="Home page on TUI" >}}

{{< figure src="images/elinks-blog-es.png" class="block" caption="Blog on TUI" >}}

{{< figure src="images/elinks-article-es.png" class="block" caption="Article on TUI" >}}

Its content may be fetched as JSON, so creating new complex frontends is
trivial without the need for content duplication. Since it is a static API, it
has good performance because of the lack of computing complexity (just reading
a file), it is more resistant to cyber-attacks because there is no database
server and it can be hosted in any service supporting static files.

```shell-session
$ wget -qO - https://nt.web.ve/en/projects/index.json | jq
{
  "url": "https://nt.web.ve/en/projects/",
  "kind": "section",
  "type": "projects",
  "lang": "en",
  "title": "Projects",
  "params": {
    "draft": false,
    "iscjklanguage": false,
    "title": "Projects"
  },
  "content": "",
  "data": {
    "pages": [
      {
        "url": "https://nt.web.ve/en/projects/docker-hugo/",
        "title": "docker-hugo",
        "author": null,
        "publishdate": "2018-05-06T22:07:39-04:00",
        "date": "2019-12-11T17:00:10-04:00",
        "description": "Dockerized Hugo CLI.",
        "tags": [
          "cli",
          "containers",
          "docker",
          "hugo"
        ]
      },
      {
        "url": "https://nt.web.ve/en/projects/sdb/",
        "title": "sdb",
        "author": null,
        "publishdate": "2019-11-28T10:31:22-04:00",
        "date": "2019-12-04T05:50:01-04:00",
        "description": "Simple and embeddable database with full text search support.",
        "tags": [
          "go",
          "make",
          "database",
          "key-value-store",
          "library",
          "golangci",
          "badgerdb",
          "blevesearch"
        ]
      },
      {
        "url": "https://nt.web.ve/en/projects/ntgo/",
        "title": "ntgo",
        "author": null,
        "publishdate": "2018-07-08T21:29:39-04:00",
        "date": "2019-10-29T10:46:23-04:00",
        "description": "A collection of Go packages.",
        "tags": null
      },
      {
        "url": "https://nt.web.ve/en/projects/pish/",
        "title": "pish",
        "author": null,
        "publishdate": "2019-09-24T10:14:00-04:00",
        "date": "2019-10-23T17:01:12-04:00",
        "description": "Environments replicator.",
        "tags": null
      },
      {
        "url": "https://nt.web.ve/en/projects/ntdocutils/",
        "title": "ntDocutils",
        "author": null,
        "publishdate": "2017-03-06T01:49:00-04:00",
        "date": "2018-06-21T13:27:00-04:00",
        "description": "Docutils theme manager.",
        "tags": [
          "python",
          "cli",
          "documentation",
          "docutils",
          "pygments"
        ]
      },
      {
        "url": "https://nt.web.ve/en/projects/ntweb/",
        "title": "ntWeb",
        "author": null,
        "publishdate": "0001-01-01T00:00:00Z",
        "date": "0001-01-01T00:00:00Z",
        "description": "A small site with great intentions.",
        "tags": [
          "go",
          "go-templates",
          "html5",
          "css3",
          "javascript",
          "website",
          "json-api",
          "hugo",
          "mage",
          "docker",
          "github-actions",
          "netlify"
        ]
      }
    ],
    "paginationPrefix": "https://nt.web.ve/en/projects/page/",
    "fist": "1",
    "prev": "",
    "next": "",
    "last": "1"
  },
  "altLang": [
    {
      "lang": "es",
      "url": "https://nt.web.ve/es/projects/"
    }
  ],
  "altMediaType": [
    {
      "mediaType": "text/html",
      "url": "https://nt.web.ve/en/projects/"
    },
    {
      "mediaType": "application/rss+xml",
      "url": "https://nt.web.ve/en/projects/index.xml"
    }
  ]
}
```

# Usage

**Requirements:**

* Hugo >= 0.61.0

Get the source code and run the following command in the project root
directory:

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

