---
title: NtDocutils
source: https://github.com/ntrrg/ntdocutils
description: Docutils theme manager.
image: /images/logo.png
license: MIT
kinds:
  - cli
  - documentation
techs:
  - python
  - docutils
  - pygments
---

[![pypi](https://img.shields.io/pypi/v/NtDocutils.svg)](https://pypi.python.org/pypi/NtDocutils)

**NtDocutils** is a theme manager for [Docutils](http://docutils.sourceforge.net/).
It acts as a wrapper for the `rst2html5.py` front end.

{{< toc >}}

# Install

**NtDocutils** requires:

* Python >= 3.4 
* Docutils == 0.14 (auto installed)
* Pygments == 2.2.0 (auto installed)

## From PyPI

```shell-session
$ pip install NtDocutils
```

## From source

```shell-session
$ wget https://github.com/ntrrg/ntdocutils/archive/v1.0.0.tar.gz
```

```shell-session
$ tar -xvf v1.0.0.tar.gz
```

```shell-session
$ cd ntdocutils-1.0.0
```

```shell-session
$ python3 setup.py
```

# Usage

Basically, you have to do two things:

1\. Create a `.rst` file:

`example.rst`:

```rest
==========
My Article
==========

:Author: Vultur Gryphus
:Contact: info@vultur.org.ve

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do
eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip
ex ea commodo consequat. Duis aute irure dolor in reprehenderit in
voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur
sint occaecat cupidatat non proident, sunt in culpa qui officia
deserunt mollit anim id est laborum.
```

2\. Process your file:

```shell-session
$ ntdocutils example.rst example.html
```

And that's it, you already have a HTML file, just like Docutils.

<p align="center">
  <img alt="Default Theme" src="/uploads/ntdocutils/example.png"/>
</p>

To use a theme, just install it and pass the `-T THEME` flag, for example, to
use the [MDL](https://getmdl.io) theme.

```shell-session
$ pip install ntdocutils-theme-mdl
```

```shell-session
$ ntdocutils -T mdl example.rst example.html
```

And this is the result:

<p align="center">
  <img alt="MDL Theme" src="/uploads/ntdocutils/mdl-example.png"/>
</p>

# Themes

* [MDL](https://ntrrg.github.io/ntdocutils-theme-mdl)

## Create a theme

1\. Get the template.

```shell-session
$ git clone --depth 1 https://github.com/ntrrg/ntdocutils-theme-template.git
```

2\. Set up the template.

```shell-session
$ mv ntdocutils-theme-template REPOSITORY_NAME
```

```shell-session
$ cd REPOSITORY_NAME
```

```shell-session
$ EDITOR config.sh
```

`config.sh`:

```sh
NAME="test"
VERSION="1.0.0"
DESCRIPTION="This is a test theme."
URL="https://github.com/ntrrg/ntdocutils-theme-test"
AUTHOR="Miguel Angel Rivera Notararigo"
EMAIL="ntrrgx@gmail.com"
SERVER="https://ntrrg.github.io/ntdocutils-theme-test/ntdocutils-theme-test"
```

```shell-session
$ ./setup.sh
```

3\. Edit  and test the template (see the [MDL theme](https://github.com/ntrrg/ntdocutils-theme-mdl/)
code and use it as example).

```shell-session
$ EDITOR
```

```shell-session
$ pip install -e .
```

```shell-session
$ (cd docs && ntdocutils -T THEME_NAME -S local demo.rst index.html)
```

4\. Publish the theme.

**Note:** a Python account is needed ([create an account](https://pypi.org/account/register/)).

```shell-session
$ pip install setuptools twine
```

```shell-session
$ rm -rf dist
```

```shell-session
$ python setup.py sdist bdist_well
```

```shell-session
$ twine upload dist/*
```

# Uninstall

```shell-session
$ pip uninstall -y NtDocutils docutils Pygments
```

