---
title: Go (Golang)
date: 2018-06-16T00:48:00-04:00
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
---

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

{{< toc >}}

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

* [Extensiones para editores de texto][].

* [Herramientas para mejorar el código][].

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
una [función](#funciones) identificada con el mismo nombre para comunicarle al
compilador cuál es el código que debe ser invocado al usar el ejecutable.

```go
func main() {                 // ┐
  fmt.Println("hello, world") // │-> Bloque de código
}                             // ┘
```

[hello, world]: https://es.wikipedia.org/wiki/Hola_mundo

En resumen, todo archivo escrito en Go tendrá la siguiente estructura:

1. Nombre del paquete.
2. Llamados a paquetes externos (opcional).
3. Cuerpo del paquete.

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

El compilador ofrece dos métodos para ejecutarlo, el primero y más sencillo es
usando el comando `go run`:

```go
$ go run hola_mundo.go
hola, mundo
```

El segundo, es generar un archivo ejecutable a partir del archivo fuente y
después ejecutarlo (obvio no? 😅), el comando anterior hace esto mismo, solo
que crea un archivo temporal y lo ejecuta automáticamente

```go
$ go build hola_mundo.go
$ ./hola_mundo
hola, mundo
```

Aunque en algunos casos baste con un archivo para crear un paquete útil, en
otras ocasiones la cantidad de código tiende a expandirse y tener muchas
líneas en un solo lugar puede generar algunos problemas, por lo que es
recomendable que lean la sección sobre [modularización](#paquetes). Para
obtener más información sobre el comando `go` y como usarlo con múltiples
archivos, pueden ir a la sección del [compilador](#compilador).

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

[GoDoc]: https://godoc.org
[Docutils]: http://docutils.sourceforge.net/

[GoDoc][] es una herramienta que permite usar los comentarios para generar
documentación, algo parecida a [Docutils][] en Python, solo que un poco más
sencilla, pues no requiere de un lenguaje de marcas para generar buena
documentación, sino que usa texto plano.

El objetivo principal de la documentación son las definiciones (`package`,
`const`, `var`, `type`, `func`, etc...) exportadas, GoDoc procesará solo
aquellas precedidas directamente por una o más líneas de comentarios.

```go
// Package arithmetic provides arithmetics operations for any type.
package arithmetic

// Identity constants
const (
  AdditiveIdentity       = 0
  MultiplicativeIdentity = 1
)

// Operander is the interface that wraps the arithmetic representation methods.
//
// Val returns the variable's arithmetic representation (float64).
type Operander interface {
  Val() float64
}

// Add gets any number of Operander and returns their addition.
func Add(operands ...Operander) float64 {
  result := float64(0)

  for _, v := range operands {
    if v.Val() == AdditiveIdentity {
      continue
    }

    result += v.Val()
  }

  return result
}
```

Es común (y una buena práctica) que cada comentario inicie con el
identificador del elemento que se quiere documentar, con la excepción de: el
nombre del paquete, que debería iniciar con la palabra `Package` y luego sí
el nombre del paquete; y también las constantes y variables agrupadas, que
suele ser suficiente con documentar el grupo y no cada una de ellas.

Aunque solo se use texto plano, GoDoc puede dar formato especial a algún texto
si:

* Tiene el formato de un URL, será convertido a un enlace HTML.

* Tiene una indentación, será convertido a un bloque de código.

* Tiene el formato `BUG(USUARIO): DESCRIPCIÓN.`, será agregado a la lista de
  bugs conocidos del paquete.

Cuando se tiene un paquete con múltiple archivos, cada uno de ellos tendrá la
sentencia `package NOMBRE`, pero esto no quiere decir que sea necesario repetir
el comentario del paquete en cada archivo, en realidad basta con que uno de los
archivos lo tenga, por esto, si la documentación es algo extensa, se
recomienda crear un archivo `doc.go` que contenga solo en nombre del paquete y
su comentario de documentación.

```go
/*
Package arithmetic provides arithmetics operations for any type.

This is a long description of the Arithmetic package.

	type Operand string

	func (o Operand) Val() float64 {
		return float64(len(o))
	}

	func main() {
		var x, y Operand = "a", "b"

		r := Add(x, y)
		fmt.Println(r)
	}

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin euismod egestas
elit sed viverra. Nunc tincidunt lacinia orci in mattis. Praesent cursus neque
et dapibus faucibus. Maecenas at sem ut arcu ornare commodo. Morbi laoreet
diam sit amet est ultricies imperdiet. Proin ullamcorper ac massa a accumsan.
Praesent quis bibendum tellus. Sed id velit libero. Fusce dapibus purus neque,
sit amet sollicitudin odio porttitor posuere. Mauris eu dui elementum,
fermentum ante vitae, porttitor nunc. Duis mi elit, viverra at turpis vitae,
sollicitudin aliquet velit. Pellentesque nisl turpis, pulvinar et consectetur
et, iaculis vel leo. Suspendisse euismod sem at vehicula fermentum. Duis
viverra eget ante a accumsan.

Aenean dui lectus, ultrices at elit id, pellentesque faucibus dolor. Duis
blandit vulputate est, eget sollicitudin ipsum pellentesque quis. Cras sed nibh
sed sapien suscipit tincidunt venenatis id eros. Praesent laoreet, erat quis
hendrerit dignissim, justo diam semper elit, sit amet commodo lacus ipsum eget
nisl. In a mi tellus. In hac habitasse platea dictumst. Aliquam et neque a quam
mollis molestie. Etiam tempor arcu quis justo molestie congue.
*/
package arithmetic
```

### Ejemplos

Además de texto, GoDoc da la posibilidad de mostrar el funcionamiento con
ejemplos dinámicos, que pueden ser ejecutados e incluso modificados en la
interfaz web. Para usar esta gran utilidad se deben crear funciones de ejemplo
en archivos `*_test.go`, estas funciones deberán tener como nombre `Example`
si se quiere mostrar algún ejemplo que use varios elementos del paquete, o
`ExampleIDENTIFICADOR[_MÉTODO]` para tener como objetivo solo un elemento.

`arithmetic.go`:

```go
package arithmetic

func Add(operands ...int) (result int) {
  for _, v := range operands {
    result += v
  }

  return result
}

func Sub(operands ...int) (result int) {
  for _, v := range operands {
    result -= v
  }

  return result
}
```

`example_test.go`:

```go
package arithmetic_test

import "fmt"

func Example {
  r := Add(1, 2) - Sub(3, 4)
  fmt.Println(r)
  // Output: 4
}

func ExampleAdd { 
  r := Add(1, 2, 3, 4)
  fmt.Println(r)
  // Output: 10
}

func ExampleSub { 
  r := Sub(5, 3, 1)
  fmt.Println(r)
  // Output: 1
}
```

Cada función de ejemplo deberá mostrar por la salida estándar los
resultados, y por cada salida que se realice, deberá existir un comentario
especial `// Output: VALOR` que indica el valor esperado. Estas funciones son
ejecutadas por `go test`, por lo que no solo tienen un uso informativo, sino
que también ayudan a probar el código; si no se encuentra algún comentario
especial, las funciones serán compiladas, pero no ejecutadas.

Para los casos en que se necesiten múltiples funciones de ejemplo de un mismo
objetivo, solo hace falta agregar un sufijo que inicie con un guión bajo y una
letra.

`example_test.go`:

```go
package arithmetic_test

import "fmt"

func ExampleAdd_two { 
  r := Add(1, 2)
  fmt.Println(r)
  // Output: 3
}

func ExampleAdd_five { 
  r := Add(1, 2, 3, 4, 5)
  fmt.Println(r)
  // Output: 15
}
```

Como un ejemplo es representado por una función, no es posible demostrar
algunas funcionalidades como la implementación de interfaces, por esta razón
existen los ejemplos de archivos, estos consisten en un archivo que contiene
exclusivamente una función de ejemplo y todas las definiciones a nivel de
paquete que sean necesarias.

`arithmetic.go`:

```go
package arithmetic

type Operander interface {
  Val() float64
}

func Add(operands ...Operander) float64 {
  result := float64(0)

  for _, v := range operands {
    result += v.Val()
  }

  return result
}
```

`whole_file_example_test.go`:

```go
package arithmetic_test

type Operand string

func (o Operand) Val() float64 {
  return float64(len(o))
}

func ExampleAdd() {
  var x, y Operand = "a", "b"

  r := Add(x, y)
  fmt.Println(r)
  // Output: 2
}
```

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

# Compilador

GOPATH

GOROOT

GOTPMDIR

# Buenas prácticas

# Atribuciones

**Go Team.** *The Go Programming Language.* <https://golang.org/>

**Ariel Mashraki.** *An overview of Go syntax and features.* <https://github.com/a8m/go-lang-cheat-sheet>

**Thomas Finley.** *Two's Complement.* <https://www.cs.cornell.edu/~tomf/notes/cps104/twoscomp.html>

**Andrew Gerrand.** *Godoc: documenting Go code.* <https://blog.golang.org/godoc-documenting-go-code>

**Andrew Gerrand.** *Testable Examples in Go.* <https://blog.golang.org/examples>
