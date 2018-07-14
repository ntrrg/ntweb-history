---
title: NtDocutils
source: https://github.com/ntrrg/ntdocutils
description: Gestor de temas para Docutils.
license: MIT
kinds:
  - cli
  - documentación
techs:
  - python
  - docutils
  - pygments
draft: true
---

[![pypi](https://img.shields.io/pypi/v/NtDocutils.svg)](https://pypi.python.org/pypi/NtDocutils)

**NtDocutils** es un gestor de temas para [Docutils](http://docutils.sourceforge.net/).
Su función es ser un mediador para el frontend `rst2html5.py`, y con esto
habilitar la posibilidad de personalizar el archivo obtenido.

{{< toc >}}

# Instalción

**NtDocutils** depende de:

* Python >= 3.4 
* Docutils == 0.14 (instalado automáticamente)
* Pygments == 2.2.0 (instalado automáticamente)

## Desde el PyPI

```shell-session
$ pip install NtDocutils
```

## Desde el código fuente

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

# Uso

Básicamente de deben hacer dos cosas:

1\. Crear un archivo `.rst`:

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

2\. Procesar el archivo:

```shell-session
$ ntdocutils example.rst example.html
```

Y eso es todo, ya debería tener su archivo HTML con el estilo predeterminado.

<p align="center">
  <img alt="Default Theme" src="/uploads/ntdocutils/example.png"/>
</p>

Para usar un tema, solo debe instalarse y pasarse la opción `-T TEMA`, por
ejemplo, para usar el [tema MDL](https://ntrrg.github.io/ntdocutils-theme-mdl):

```shell-session
$ pip install ntdocutils-theme-mdl
```

```shell-session
$ ntdocutils -T mdl example.rst example.html
```

Y el resultado es:

<p align="center">
  <img alt="MDL Theme" src="/uploads/ntdocutils/mdl-example.png"/>
</p>

# Temas

* [MDL](https://ntrrg.github.io/ntdocutils-theme-mdl)

## Crear un tema

1\. Obtener la plantilla.

```shell-session
$ git clone --depth 1 https://github.com/ntrrg/ntdocutils-theme-template.git
```

2\. Configurar la plantilla.

```shell-session
$ mv ntdocutils-theme-template NOMBRE_REPOSITORIO
```

```shell-session
$ cd NOMBRE_REPOSITORIO
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

3\. Editar y probar la plantilla (vea el código fuente del [tema MDL](https://github.com/ntrrg/ntdocutils-theme-mdl/)
y úselo como ejemplo).

```shell-session
$ pip install -e .
```

```shell-session
$ cd docs
```

```shell-session
$ ntdocutils -T TEMA -S local demo.rst index.html
```

4\. Publicar el tema.

**Nota:** Es necesario tener una cuenta Python ([crear una cuenta](https://pypi.org/account/register/)).

```shell-session
$ rm -rf ntdocutils-theme-TEMA
```

```shell-session
$ cd ..
```

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

# Desinstalar

```shell-session
$ pip uninstall -y NtDocutils docutils Pygments
```

