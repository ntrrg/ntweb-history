---
title: Go (Golang) 1.10
date: 2018-05-16T12:23:25-04:00
image: /uploads/gopher.png
description: Es un lenguaje de código abierto, minimalista y de alto rendimiento; su fuerte es la concurrencia.
categories:
  - tecnología
tags:
  - referencias
  - programación
  - lenguajes-de-programación
  - go
  - programación-de-sistemas
  - programación-web
  - backend
draft: true
---

**Tabla de contenido:**

{{< toc >}}

[Go license]: https://golang.org/LICENSE

Fue diseñado en el año 2007 por Ken Thompson, Rob Pike y Robert Griesemer en
Google. Es de código abierto y es ditribuido bajo una licencia
[BSD-style][Go license]. Algunas de sus características son:

[GC]: https://es.wikipedia.org/wiki/Recolector_de_basura

* Imperativo, los programas se escriben como una serie de instrucciones que la
  computadora debe seguir para resolver un problema (si leyendo esto piensan
  *«¿Y no es así como se escriben todos los programas? 😒»*, la respuesta
  es no, existen otros paradigmas de programación que trabajan con enfoques
  muy diferentes a este).

* Compilado, todo el código escrito es traducido a lenguaje máquina antes de
  poder ejecutarse.

* Tipado estático, una vez que se define el tipo de una variable, este no
  puede ser modificado.

* Fuertemente tipado, no permite realizar operaciones entre datos de diferente
  tipo, deben hacerse cambios de tipo explícitamente.

* No es necesario liberar manualmente la memoria asignada, usa un [GC][] que se
  encarga de esto, lo que da facilidades en el manejo de memoria.

* Concurrencia y paralelismo de manera nativa (por medio de palabras reservadas
  y operadores, también tiene algunas bibliotecas que permiten aplicar técnicas
  de más bajo nivel).

* Minimalista. La mayoría de las utilidades que faltan en el lenguaje, fueron
  [excluidas intencionalmente](#).

# Herramientas necesarias

Para empezar a programar solo hacen falta dos cosas:

<!--lint disable no-undefined-references no-shortcut-reference-link-->

[Instalar Go]: {{< relref "blog/install-go-1.10.es.md" >}}
[Install Go]: https://golang.org/doc/install

* El compilador (las instrucciones para instalarlo pueden verlas
  [aquí][Instalar Go] o en la [documentación oficial][Install Go]).

<!--lint enable no-undefined-references no-shortcut-reference-link-->

* Un editor de texto.

Aunque yo no soy muy fanático de usar muchos plugins, extensiones y cosas así
porque con esto es más que suficiente para desarrollar tranquilamente, existen
muchas herramientas que ayudan a mejorar la productividad e integran bastantes
utilidades en el flujo de trabajo sin mucha fricción, algunas de las que
conozco son:

[Playground]: https://play.golang.org/
[Better Go Playground]: https://chrome.google.com/webstore/detail/odfhkelcmblecfdnboahphiafolojmpl
[Extensiones para editores de texto]: https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins
[Herramientas para mejorar el código]: https://github.com/golang/go/wiki/CodeTools

* Gophertool, que es una extensión muy sencilla de Chrome y viene con la
  instalación de Go, específicamente en la carpeta `misc/chrome/gophertool`.
  Fue creada para ayudar a los programadores del lenguaje con algunos accesos
  rápidos, pero para simple mortales como yo, permite buscar en la
  documentación de la biblioteca estándar un poco más rápido.

* El [Playground][Playground] que permite probar código directamente en el
  navegador y [Better Go Playground][], que es una extensión de Chrome que lo
  hace más amigable.

* [Extensiones para editores de texto][]

* [Herramientas para mejorar el código][]

# Archivos Go

Todos los archivos escritos en Go forman parte de un paquete, que es la unidad
de distribución de código en este lenguaje, por esto, todos los archivos
deben iniciar con una línea que contenga `package NOMBRE`, donde `NOMBRE` es
un valor asignado por el desarrollador y es el identificador con el que otros
podrán utilizarlo dentro de sus programas. Cuando se pretende desarrollar
algún comando o alguna aplicación se usa `package main`, `main` es un nombre
especial que le dice al compilador que la intención del archivo no es servir
como biblioteca, sino como un ejecutable.

```go
package main // -> Nombre del paquete
```

Después de una línea en blanco, se hace el llamado a los paquetes que se
usarán en el programa (si hace falta ¿no?, no es que sea obligatorio usar al
menos un paquete 😂), por ejemplo, si se quiere escribir algo en pantalla se
debe importar el paquete `fmt`.

```go
import "fmt" // -> Paquetes importados
```

Luego se escriben todas las instrucciones que el programador quiera darle a la
computadora, en el caso de usar `main` como nombre del paquete, se debe crear
un bloque de código identificado con el mismo nombre para comunicarle al
compilador cuál es el código que debe ser invocado al usar el ejecutable.

```go
func main() {                 // ┐
  fmt.Println("hello, world") // │-> Bloque de código
}                             // ┘
```

A estos bloques se les llaman funciones (por eso el `func` al inicio, que viene
de *«function»*) y su principal utilidad es modularizar y reutilizar el
código, muy parecidas a los paquetes, solo que a una escala menor; tienen
cierta sintaxis específica, pero por ahora basta con saber que:

* Se usa la palabra reservada `func` para iniciar la declaración.

* Separado por un espacio en blanco se escribe el nombre de la función
  (`main` en este caso) y unos paréntesis (`()`).

* Se escribe el código a ejecutar dentro de llaves (`{}`).

[hello, world]: https://es.wikipedia.org/wiki/Hola_mundo

En resumen, todo archivo escrito en Go tendrá la siguiente estructura:

1. Nombre del paquete.
2. Llamados a paquetes externos (opcional).
3. Cuerpo del programa.

Siguiendo estas reglas, el programa más famoso ([hello, world][]) escrito en
Go se vería algo así:

`hola_mundo.go`:

```go
package main

import "fmt"

func main() {
  fmt.Println("hola, mundo")
}
```

# Comentarios

Los comentarios son texto ignorado por el compilador, su función principal es
documentar ciertas secciones de código que sean un poco difíciles de entender
a simple vista, pero en muchas ocasiones también son usados para ocultar
código de los ojos del compilador y ver como se comporta el programa. Existen
dos tipos de comentarios:

* De línea

```go
fmt.Println("hola, mundo") // Esto muestra `hola, mundo`

// Las sentencias comentadas no son procesadas por el compilador
// fmt.Println("chao, mundo")
```

* Generales

```go
/* Así se escribe un comentario general

fmt.Println("hola, mundo")
fmt.Println("chao, mundo")

Este programa no hace nada..
*/
```

## Documentación

# Tipos de datos

Son clasificaciones que permiten decirle al compilador como pretenden usarse
los datos que pasan por el programa. Go cuenta con una estructura muy bien
definida en cuanto a tipos de datos, pero también permite crear nuevos según
las necesidades del programador.

Todos los tipos de datos cuentan con un valor cero, que no quiere decir que
sean literalmente `0`, sino que los identifica como "vacío" en su contexto,
por ejemplo: cuando se habla de números, su valor cero sería `0`; cuando se
habla de personas, su valor cero sería `nadie`; cuando se habla de objetos, su
valor cero sería `nada`; y así dependiendo del contexto, por supuesto, estos
son **ejemplos**, no es que Go tenga un tipo de dato `person` o algo así 😂.

## Booleanos

Nombrados así en honor a George Boole, también son conocidos como lógicos,
representan valores de verdad (verdadero o falso) que normalmente son usados
para controlar el flujo de los programas.

### Representación sintáctica

```go
bool
```

### Ejemplos

```go
true
false
```

### Valor cero

```go
false
```

## Cadenas

### Representación sintáctica

```go
string
```

### Ejemplos

```go
"C"
"Cadena de caracteres"

`Cadena
de
caracteres
multilineal`

"Cadena con\nsalto de línea"
`Cadena sin\nsecuencias de escape`
```

### Valor cero

```go
""
```

## Numéricos

Existen tres grupos de datos numéricos:

* Enteros
* Punto flotante
* Complejos

### Enteros

Representan los números del conjunto matemático con el mismo nombre, aunque
claro, con una cantidad finita de elementos, que puede ser controlada por el
espacio de memoria que se reserve, es decir, el programador tiene la capacidad
de especificar si quiere un número entero que ocupe `N` bits de memoria, donde
`N` puede ser `8`, `16`, `32` o `64` (**TODO:** ¿por qué estos números?).

<!--lint disable no-undefined-references no-shortcut-reference-link-->

[Complemento a dos]: {{< relref "blog/twos-complement.es.md" >}}

Existen dos tipos de números enteros, o mejor dicho, dos métodos de
representación: el primero es la conversión binaria tradicional, pero solo
puede ser usado para procesar números positivos; el segundo es llamado
[Complemento a dos][] y permite representar tanto números positivos como
negativos de una manera bastante ingeniosa, solo que se pierde una cantidad
considerable de números positivos.

<!--lint enable no-undefined-references no-shortcut-reference-link-->

Además de números decimales, es posible usar octales y hexadecimales.

#### Representación sintáctica

```go
// Enteros sin signo
uint8  // 0 - 255
uint16 // 0 - 65535
uint32 // 0 - 4294967295
uint64 // 0 - 18446744073709551615

// Enteros con signo
int8  // -128 - 127
int16 // -32768 - 32767
int32 // -2147483648 - 2147483647
int64 // -9223372036854775808 - 9223372036854775807

uintptr // Permite almacenar direcciones de memoria

byte // Equivale a uint8, puede almacenar un caracter ASCII
rune // Equivale a uint32, puede almacenar un caracter UTF-8

// Dependiendo de la arquitectura del procesador
uint // Equivale a uint32 o uint64
int  // Equivale a int32 o int64
```

#### Ejemplos

```go
5   // Decimal
05  // Octal (tienen el prefijo `0`)
0x5 // Hexadecimal (tienen el prefijo `0x`)

10
012
0xA

'a'  // 97, caracter ASCII
'😂' // 128514, caracter UTF-8
```

#### Valor cero

```go
0
```

### Punto flotante

Representan al conjunto matemático de los números fraccionarios, aunque
claro, con una cantidad finita de elementos, que puede ser controlada por el
espacio de memoria que se reserve, es decir, el programador tiene la capacidad
de especificar si quiere un número entero que ocupe `N` bits de memoria, donde
`N` puede ser `32` o `64` (**TODO:** ¿por qué estos números?).

Para su representación siguen el estándar técnico IEEE-754.

#### Representación sintáctica

```go
float32 // 0 - 4294967295
float64 // 0 - 18446744073709551615
```

#### Ejemplos

```go
5   // Decimal
05  // Octal (tienen el prefijo `0`)
0x5 // Hexadecimal (tienen el prefijo `0x`)

10
012
0xA

'a'  // 97, caracter ASCII
'😂' // 128514, caracter UTF-8
```

#### Valor cero

```go
0
```

### Complejos

.. Numerics
.. Casting

..     .3. Tipos de variable
..         .1. Simples
..             .1. Numéricos
..         .2. Compuestos
..             .1. Cadenas
..             .2. Listas
..             .3. Tuplas
..             .4. Conjuntos
..             .5. Diccionarios
..     .4. Cambio de tipo

<https://tour.golang.org/basics/11>
<https://tour.golang.org/basics/13>

Estructuras:

<https://tour.golang.org/moretypes/2>
<https://tour.golang.org/moretypes/3>

Arreglos:

<https://tour.golang.org/moretypes/6>

Porciones:

<https://tour.golang.org/moretypes/7>
<https://tour.golang.org/moretypes/8>
<https://tour.golang.org/moretypes/9>
<https://tour.golang.org/moretypes/10>
<https://tour.golang.org/moretypes/11>
<https://tour.golang.org/moretypes/12>
<https://tour.golang.org/moretypes/13>
<https://tour.golang.org/moretypes/14>
<https://tour.golang.org/moretypes/15>
<https://blog.golang.org/go-slices-usage-and-internals>

Mapas:

<https://tour.golang.org/moretypes/19>
<https://tour.golang.org/moretypes/20>
<https://tour.golang.org/moretypes/21>
<https://tour.golang.org/moretypes/22>

Punteros:

<https://tour.golang.org/moretypes/1>
<https://tour.golang.org/moretypes/4>
<https://tour.golang.org/moretypes/5>

# Buenas prácticas

# Atribuciones

**Go Team.** *The Go Programming Language.* <https://golang.org/>

**Ariel Mashraki.** *An overview of Go syntax and features.* <https://github.com/a8m/go-lang-cheat-sheet>

**Thomas Finley.** *Two's Complement.* <https://www.cs.cornell.edu/~tomf/notes/cps104/twoscomp.html>

