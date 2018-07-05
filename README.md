[![Docker Build Status](https://img.shields.io/docker/build/ntrrg/site.svg)](https://store.docker.com/community/images/ntrrg/site)
[![](https://images.microbadger.com/badges/image/ntrrg/site.svg)](https://microbadger.com/images/ntrrg/site "Get your own image badge on microbadger.com")

This site was migrated to Hugo and is hosted now by Netlify.

<https://nt.web.ve>

## Usage

### Docker

```shell-session
docker run -p 1234:80 ntrrg/site
```

Go to <http://localhost:1234/en/> with a browser.

### Hugo

**Note:** use Hugo 0.42.1.

```shell-session
git clone https://github.com/ntrrg/ntrrg.github.io.git
```

```shell-session
cd ntrrg.github.io/
```

```shell-session
hugo server
```

Go to <http://localhost:1313/en/> with a browser.

## Development

* Lint markdown files:

```shell-session
docker run -itv "$PWD":/files/ ntrrg/md-linter:watch
```

* Run the development server:

```shell-session
docker run -itp 1313:1313 -v "$PWD":/site/ ntrrg/hugo:0.42.1
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

