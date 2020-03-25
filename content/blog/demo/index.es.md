---
title: Página de ejemplo
author: ntrrg
publishdate: 2028-07-05T18:35:00-04:00
date: 2028-07-05T18:35:00-04:00
image: images/ntrrg.png
description: Esta es una página de ejemplo para ver los estilos de Markdown.
tags:
  - tag1
  - tag2
  - tag3
  - tag4
  - tag5
math: true
mermaid: true
toc: true
draft: true
---

# Parrafos

Cualquier línea con líneas vacías antes y después de ella es un párrafo,
las líneas consecuentes son unidas.

Necesita una línea vacía para un nuevo párrafo.

# Separadores

---

# Títulos (h1) 

## Títulos (h2) 

### Títulos (h3) 

#### Títulos (h4) 

##### Títulos (h5) 

###### Títulos (h6) 

# Decoración de texto

**Este es texto en negrita**

__Este es texto en negrita__

*Este es texto en cursiva*

_Este es texto en cursiva_

~~Este es texto tachado~~

<https://nt.web.ve>

[Este es un enlace](https://nt.web.ve)

[Este es un enlace con título](https://nt.web.ve "Este es el título!").

[ntweb]: https://nt.web.ve

[ntweb][]

[Mi sitio web][ntweb]

# Listas

* Cree una lista iniciando con `+`, `-`, o `*`
* Las sublistas se crean agregando una indentación de 2 espacios
  * Esta es una sublista
* Y todo se vuelve normal otra vez

1. Esta es
2. una lista
3. ordenada

* [ ] Esta es
* [x] una lista de tareas

Esta
: es una lista de definiciones.

Término:
: definición, puede agregar los `:` al término.

# Citas

Este párrafo tiene una nota de pie[^1].

[^1]: Y aquí está la nota de pie.

> Las citas de bloque
> son escritas así.
>
> Ellas pueden usar varios párrafos si lo prefiere.
>
> Y **Markdown**!.
>
> -- El Autor

# Tablas

| Título  |   Otro título   |
| ------- | --------------- |
| texto   | texto           |
| texto   | texto           |
| texto   | texto           |

| Título  |   Otro título   |
| :-----: | :-------------: |
|  texto  |      texto      |
|  texto  |      texto      |
|  texto  |      texto      |

| Título  |   Otro título   |
| ------: | --------------: |
|   texto |           texto |
|   texto |           texto |
|   texto |           texto |

# Imágenes

!["Texto altenativo"](images/ntrrg.png "Este es el título")

# Fórmulas matemáticas

Esta es una fracción inteligente 1/2, este texto tiene una fórmula de línea
\\(\sum\_{n=1}^{\infty} 2^{-n} = 1\\) y esta es una fórmula de bloque:

$$
\sum\_{n=1}^{\infty} 2^{-n} = 1
$$

# Código

Este es `código` de línea.

```go {linenos=true,hl_lines=["1", "5-7"],linenostart=0}
package main

import "fmt"

func main() {
  fmt.Println("hello, world")
}
```

# Shortcodes

## Teclado

{{< kbd "Ctrl" >}} + {{< kbd "Alt" >}} + {{< kbd "Supr" >}}

## Imágenes

### De línea

Imágenes de línea con {{< img src="images/hugo.png" style="height: 1em;" >}}

### Extendidas

{{< img src="images/merida.jpg" class="wide" >}}

### De bloque (centradas)

{{< img src="images/merida.jpg" class="block" >}}

### Alineadas a la izquierda

{{< img src="images/ntrrg.png" class="align-left" >}}

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

### Alineadas a la derecha

{{< img src="images/ntrrg.png" class="align-right" >}}

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

## Figuras

### Extendidas

{{< figure src="images/merida.jpg" class="wide" title="Título" caption="Cuerpo de la figura. Soporta **Markdown**." >}}

### De bloque (centradas)

{{< figure src="images/merida.jpg" class="block" title="Título" caption="Cuerpo de la figura. Soporta **Markdown**." >}}

### Alineadas a la izquierda

{{< figure src="images/ntrrg.png" class="align-left" title="Título" caption="Cuerpo de la figura. Soporta **Markdown**." >}}

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

### Alineadas a la derecha

{{< figure src="images/ntrrg.png" class="align-right" title="Título" caption="Cuerpo de la figura. Soporta **Markdown**." >}}

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

## Go Playground

{{< go-playground >}}
```go
package main

import "fmt"

func main() {
  fmt.Println("hello, world")
}
```
{{< /go-playground >}}

## Notas

{{< note >}}
Esta es una nota
{{< /note >}}

{{< note "Mi título" >}}
Esta es una **nota** con título personalizado.
{{< /note >}}

## Enlaces de interés

{{< loi >}}
* <https://nt.web.ve>
* <https://nt.web.ve/en/>
* <https://nt.web.ve/es/>
{{< /loi >}}

## Detalles

{{< details >}}
**Lorem ipsum** dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
{{< /details >}}

{{< details summary="Texto personalizado" open=true >}}
**Lorem ipsum** dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
{{< /details >}}

## Gráficos

{{< mermaid "Título del gráfico. Soporta **Markdown**." >}}
```mermaid
graph TD
  A[Christmas] -->|Get money| B(Go shopping)
  B --> C{Let me think}
  C -->|One| D[Laptop]
  C -->|Two| E[Phone]
  C -->|Three| F[fa:fa-car Car]
```
{{< /mermaid >}}

## Snippets

{{< snippet "files/hello.go" "go" >}}

{{< snippet "files/hello.go" "go {linenos=true,hl_lines=[\"1\", \"5-7\"],linenostart=0}" >}}

{{< import "files/imports.es.md" >}}

## Tarjetas

{{< card "/blog/demo/" >}}

