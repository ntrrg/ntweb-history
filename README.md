This site was migrated to Hugo and is hosted now by Netlify.

<https://nt.web.ve>

## Offline version

### Docker

[![Docker Build Status](https://img.shields.io/docker/build/ntrrg/site.svg)](https://store.docker.com/community/images/ntrrg/site)
[![MicroBadger Size](https://img.shields.io/microbadger/image-size/ntrrg/site.svg)](https://microbadger.com/images/ntrrg/site)

```shell-session
$ docker run --rm -p 1234:80 ntrrg/site
```

Go to <http://localhost:1234/en/> with a browser.

### Hugo

**Note:** use Hugo 0.46 or above.

```shell-session
$ wget https://github.com/ntrrg/ntrrg.github.io/archive/master.tar.gz
```

```shell-session
$ tar -xvf master.tar.gz
```

```shell-session
$ cd ntrrg.github.io-master/
```

```shell-session
$ hugo server
```

Go to <http://localhost:1313/en/> with a browser.

## Editing

* Lint markdown files:

```shell-session
$ make lint
```

* Run the development server:

```shell-session
$ make run
```

## Acknowledgment

Working on this site I use/used:

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

