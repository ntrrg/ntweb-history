# Git

<img src="https://git-scm.com/images/logo@2x.png" class="logo"/>
<!-- <img src="/home/ntrrg/Documents/Notes/Git/images/logo.png" class="logo"/> -->

Es un sistema de control de versiones, esto quiere decir que se encarga de llevar el registro de todos los cambios de uno o más archivos (normalmente proyectos), un super `Ctrl` + `Z`. Fue creado por Linus Torvalds y su primera versión usable fue liberada el 7 de abril del 2005; después de dos meses y hasta la actualidad, ha sido mantenido por Junio Hamano y otra gran cantidad de personas de todas partes del mundo. Es un proyecto de código abierto (https://github.com/git/git) y está licenciado bajo los términos de la GPLv2. Su última versión estable es la número 2.15.0.

## Tabla de contenido

- [Características](#Caracter%C3%ADsticas)
- [Instalación](#Instalaci%C3%B3n)
- [Inicio rápido](#Inicio-r%C3%A1pido)
- [Configuración](#Configuraci%C3%B3n)
- [Buenas prácticas](#Buenas-pr%C3%A1cticas)
- [Desinstalación](#Desinstalaci%C3%B3n)
- [Recursos útiles](#Recursos-%C3%BAtiles)
- [Reconocimientos](#Reconocimientos)

## Características

* **Liviano:** está escrito en C, lo que le da un gran rendimiento.

* **Etiquetas:** se pueden identificar momentos específicos del estado del proyecto para el manejo de liberaciones de código (`v0.1.0`, `v1.0.0`, `v2.1.3`).

* **Ramas:** se pueden crear diferentes líneas de desarrollo sobre un mismo proyecto, facilitando el trabajo simultaneo entre muchos programadores.

* **Adaptable:** soporta prácticamente cualquier flujo de trabajo, lo que permite usarlo en cualquier tipo de proyecto.

* **Distribuido:** cada proyecto es independiente, por lo que no hace falta estar conectado a un servidor para usar Git, incluso cada persona que tenga el proyecto podría cumplir la función de servidor para que alguien más pueda obtenerlo o actualizarlo.

## Instalación

```shell
sudo su || su
apt update
apt upgrade
apt install git
```

## Inicio rápido

## Configuración

Es posible establecer configuraciones para alterar el comportamiento tradicional de Git, existen tres niveles:

1. **Sistema:** (`--system`) aplica a todos los proyectos del sistema, se guarda en `/etc/gitconfig`.
2. **Usuario:** (`--global`) aplica a todos los proyectos que maneje el usuario, se guarda en `~/.gitconfig`.
3. **Proyecto:** (`--local`) aplica solo al proyecto, es el nivel predeterminado, se guarda en `.git/config`.

Git lee las configuraciones de manera descendente (**Sistema** -> **Usuario** -> **Proyecto**), esto permite reescribir ciertas opciones fácilmente

```
git config [<nivel>] <configuración>  # Crear
git config [<nivel>] --unset <configuración>  # Eliminar
```

### Alias

**Sintaxis:**

```
git config [<nivel>] alias.<alias> ("<subcomando>" | "!<programa externo>")
```

**Ejemplos:**

Subcomando de Git:

```shell
git config alias.cm "commit"
```

```shell
git cm -m "initial commit"
# Equivale a: git commit -m "initial commit"
```

Subcomando de Git con opciones:

```shell
git config alias.l "log --oneline"
```

```shell
git l
# Equivale a: git log
```

Programa externo:

```shell
git config alias.ignore "!subl .gitignore"
```

```shell
git ignore
# Equivale a: subl .gitignore
```

### Autor

**Sintaxis:**

```
git config [<nivel>] user.name "<usuario>"
git config [<nivel>] user.email "<correo>"
git config [<nivel>] user.signingkey "<GPG key ID>"
```

**Ejemplos:**

```shell
git config user.name "Miguel Angel Rivera Notararigo"
git config user.email "ntrrgx@gmail.com"
git config user.signingkey "3AA5C34371567BD2"
```

### Colores

**Sintaxis:**

```
git config [<nivel>] color.ui (auto | always | false)
```

**Ejemplos:**

```shell
git config color.ui false
```

### Editor

**Sintaxis:**

```
git config [<nivel>] core.editor "<editor>"
```

**Ejemplos:**

```shell
git config core.editor "subl -w"
```

## Buenas prácticas

* Cada programador debe realizar sus cambios en una rama personalizada.

* Verificar que los archivos no contengan espacios en blanco innecesarios, se puede usar `git diff --check` para esto.

* Cada confirmación debe representar una acción (arreglar un typo, mejorar la documentación de algunos archivos, renombrar una función, etc).

* El mensaje de una confirmación puede variar según los criterios de cada proyecto, pero normalmente cumple las siguientes normas:

  1. La primera línea debe ser una descripción general de los cambios, no mayor a 50 caracteres, en minúsculas y en infinitivo.

  2. Puede tener una descripción más detallada separada por una línea en blanco y que cada línea no supere los 73 caracteres.

  ```
  mejorar documentación inline de la API

  Se agregó documentación a todos los controladores nuevos y se mejoró la
  de los controladores existentes.

  Autores
  -------

  * Miguel Angel Rivera Notararigo <ntrrgx@gmail.com>
  * Nombre Apellido <user@server.dom>
  ```

## Desinstalación

```shell
sudo su || su
apt purge git
```

## Recursos útiles

*tryGit* - https://try.github.io/

*Pro Git 2nd Edition* - https://git-scm.com/book/en/v2

*Pro Git 2nd Edition en español* (en progreso) - https://git-scm.com/book/es/v2

*Pro Git en español* - https://git-scm.com/book/es/v1

## Reconocimientos

El artículo fue escrito con [Boostnote](https://boostnote.io/).

**Wikipedia.** *Git.* https://en.wikipedia.org/wiki/Git

**Git.** *About.* https://git-scm.com/about

<style type="text/css">
  .logo {
    float: right;
    margin: 1em;
  }
<style>
