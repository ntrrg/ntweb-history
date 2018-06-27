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
  *«¿Y no es así como se escriben todos los programas? 😒»*, la respuesta es
  no, existen otros paradigmas de programación que trabajan con enfoques muy
  diferentes a este).

* Compilado, todo el código escrito es traducido a lenguaje máquina antes de
  poder ejecutarse.

* Tipado estático, una vez que se define el tipo de una variable, este no puede
  ser modificado.

* Fuertemente tipado, no permite realizar operaciones entre datos de diferente
  tipo, deben hacerse cambios de tipo explícitamente.

* No es necesario liberar manualmente la memoria asignada, usa un [GC][] que se
  encarga de esto, lo que da facilidades en el manejo de memoria.

* Concurrencia y paralelismo de manera nativa (por medio de palabras reservadas
  y operadores, también tiene algunas bibliotecas que permiten aplicar técnicas
  de más bajo nivel).

* Minimalista. La mayoría de las utilidades que faltan en el lenguaje, fueron
  [excluidas intencionalmente](#funcionalidades-excluidas).

{{< toc >}}

# Funcionalidades excluidas

{{% loi %}}
* <https://golang.org/doc/faq#Design>
* <https://www.youtube.com/watch?v=k9Zbuuo51go>
{{% /loi %}}

* Genéricos, aunque es posible que en alguna futura versión se agregue, por
  ahora no se ha logrado obtener una solución que compense su complejidad con
  su utilidad. En su lugar pueden usarse las [interfaces](#interfaces), que
  ofrecen abstracción de una manera muy elegante.

* Conjuntos, por la falta de tipos genéricos en el lenguaje, representan
  agregar un nuevo tipo de dato al código base Go y debido a que es bastante
  sencillo implementarlos en la mayoría de los casos, no es muy relevante su
  existencia.

{{% go-playground "oZR8BSYAUVB" %}}
```go
x := make(map[int]struct{})

x[1] = struct{}{}
x[2] = struct{}{}
x[1] = struct{}{}

fmt.Println(len(x)) // 2
```
{{% /go-playground %}}

* `while` y `do-while`, solo hay una estructura de repetición (`for`) y aunque
  parezca limitado, es una ventaja para los programadores no tener que pensar
  en cuál estructura usar, tal vez suene a exagerar, pero en internet es muy
  fácil encontrar discusiones largas y repetitivas de varios lenguajes sobre
  cuál de todas las estructuras de repetición es la más rápida.

* `map`, `filter` y la familia de funciones favoritas de los programadores
  funcionales, por la falta de tipos genéricos sería necesario definir
  muchísimas funciones para cada tipo, pero además, ¿por qué llamar 100
  funciones para sumar los elementos de un arreglo si puede usarse una
  estructura de repetición muy sencilla?, si la reacción a esto es *«No me
  importa el rendimiento, quiero mis funciones 😒»* no hay problema, es muy
  fácil implementarlas, pero en este caso les recomendaría usar otro lenguaje.

{{% go-playground "oNGlnMctzXv" %}}
```go
func ForEach(s []int, f func(int, int, []int)) {
  for i, v := range s {
    f(v, i, s)
  }
}

func Map(s []int, f func(int) int) (ns []int) {
  for _, v := range s {
    ns = append(ns, f(v))
  }

  return ns
}

func Filter(s []int, f func(int) bool) (ns []int) {
  for _, v := range s {
    if f(v) {
      ns = append(ns, v)
    }
  }

  return ns
}

func Reduce(s []int, f func(int, int) int, a int) int {
  for _, v := range s {
    a = f(a, v)
  }

  return a
}
```
{{% /go-playground %}}

* Excepciones, usar estructuras de control (como `try {} catch {}`) para
  manejar los errores puede resultar en flujos complejos que dificultan a los
  programadores el proceso de identificación y solución de errores (debugging).
  En su lugar los errores se manejan por medio del mecanismo de retorno
  múltiple en las funciones y por otras funciones predefinidas en Go.

* Afirmaciones (asserts), en la mayoría de las ocasiones, los programadores
  usan las afirmaciones para generar errores en tiempo de ejecución y así, no
  tener que implementar de mejor manera el flujo del programa. Al igual que con
  las excepciones, existen otros métodos para manejar las afirmaciones en Go.

* Aritmética de punteros, no está permitida, ya que Go usa un GC.

* Hilos de procesos (threads), una de las tareas que suele agregar muchísima
  complejidad al código fuente es la programación multithreading, aunque claro,
  si se pretende programar una aplicación que se usará en servidores o en
  computadores personales con procesadores de múltiples núcleos y hacer toda la
  computación en un solo hilo, sería un descaro decir que Go es un lenguaje de
  alto rendimiento, pero la verdad es que no hacen falta, ya se que suena loco
  y probablemente piensen *«Claaaro, un programa con gran demanda de computo
  que corre en un hilo puede ser tan rápido como uno que corre en múltiples
  hilos.. 😒»*, pensamiento sarcástico que sería muy merecido, pero el hecho es
  que Go cuenta con gorutinas, que son funciones que se ejecutan
  independientemente del hilo principal y son automáticamente distribuidas
  entre más hilos para evitar el bloqueo de las operaciones, esto genera una
  abstracción de más alto nivel para este tipo de operaciones, por lo que el
  programador no debe lidear directamente con hilos (vean la sección de
  [gorutinas](#gorutinas)).

# Herramientas necesarias

Para empezar a programar solo hacen falta dos cosas:

<!--lint disable no-undefined-references no-shortcut-reference-link-->

[Instalar Go]: {{< relref "blog/install-go-1.10.es.md" >}}
[Install Go]: https://golang.org/doc/install

* El compilador (pueden ver las instrucciones para instalarlo en la
  [documentación oficial][Install Go] o en este [artículo][Instalar Go]
  sobre como instalar Go).

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

* [GoDoc](https://godoc.org/golang.org/x/tools/cmd/godoc) para generar la
  [documentación](#documentación) de los paquetes.

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

{{% go-playground "hR9ZBMz-Pst" %}}
```go
package main

import "fmt"

func main() {
  fmt.Println("hola, mundo")
}
```
{{% /go-playground %}}

El compilador ofrece dos métodos para ejecutarlo: el primero y más sencillo es
usando el comando `go run`.

```go
$ go run hola_mundo.go
hola, mundo
```

El segundo, es generar un archivo ejecutable a partir del archivo fuente y
después ejecutarlo (obvio no? 😅), el comando anterior hace esto mismo, solo
que crea un archivo temporal y lo ejecuta automáticamente.

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

{{% loi %}}
* <https://golang.org/ref/spec#Comments>
{{% /loi %}}

Los comentarios son texto ignorado por el compilador, su función principal es
documentar ciertas secciones de código que sean un poco difíciles de entender
a simple vista, pero en muchas ocasiones también son usados para ocultar
código de los ojos del compilador y ver como se comporta el programa. Existen
dos tipos de comentarios:

* De línea

{{% go-playground "4g5BEqD0RGU" %}}
```go
fmt.Println("hola, mundo") // Esto muestra "hola, mundo"

// Las sentencias comentadas no son procesadas por el compilador
// fmt.Println("chao, mundo")
```
{{% /go-playground %}}

* Generales

{{% go-playground "4HyigTWqiZ8" %}}
```go
/*
  Así se escribe un comentario general

  fmt.Println("hola, mundo")
  fmt.Println("chao, mundo")

  Este programa no hace nada..
*/
```
{{% /go-playground %}}

## Documentación

{{% loi %}}
* <https://blog.golang.org/godoc-documenting-go-code>
{{% /loi %}}

[GoDoc]: https://godoc.org
[Docutils]: http://docutils.sourceforge.net/

[GoDoc][] es una herramienta que permite usar los comentarios para generar
documentación, algo parecida a [Docutils][] en Python, solo que un poco más
sencilla, pues no requiere de un lenguaje de marcas para generar buena
documentación, sino que usa texto plano.

El objetivo principal de la documentación son las definiciones (`package`,
`const`, `var`, `type`, `func`, etc...) exportadas, GoDoc procesará solo
aquellas precedidas directamente por una o más líneas de comentarios.

`$GOPATH/src/local/arithmetic/aritmetic.go`:

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
  result := operands[0]

  for _, v := range operands[1:] {
    if v.Val() == AdditiveIdentity {
      continue
    }

    result += v.Val()
  }

  return result
}
```

Y para ver el resultado, se debe ejecutar el comando `godoc` con la ruta de
importación del paquete.

```shell-session
$ godoc local/arithmetic
use 'godoc cmd/local/arithmetic' for documentation on the arithmetic command 

PACKAGE DOCUMENTATION

package arithmetic
    import "local/arithmetic"

    Package arithmetic provides arithmetics operations for any type.

CONSTANTS

const (
    AdditiveIdentity       = 0
    MultiplicativeIdentity = 1
)
    Identity constants

FUNCTIONS

func Add(operands ...Operander) float64
    Add gets any number of Operander and returns their addition.

TYPES

type Operander interface {
    Val() float64
}
    Operander is the interface that wraps the arithmetic representation
    methods.

    Val returns the variable's arithmetic representation (float64).


```

O iniciar el servidor HTTP de GoDoc e ir a la ruta <http://localhost:6060/pkg/local/arithmetic>
con un navegador si se quiere ver la versión HTML.

```shell-session
$ godoc -http :6060
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

{{% loi %}}
* <https://blog.golang.org/examples>
{{% /loi %}}

Además de texto, GoDoc da la posibilidad de mostrar el funcionamiento con
ejemplos dinámicos, que pueden ser ejecutados e incluso modificados en la
interfaz web. Para usar esta gran utilidad se deben crear funciones de ejemplo
en archivos `*_test.go`, estas funciones deberán tener como nombre `Example`
si se quiere mostrar algún ejemplo que use varios elementos del paquete, o
`ExampleIDENTIFICADOR[_MÉTODO]` para tener como objetivo solo un elemento.

`$GOPATH/src/local/arithmetic/aritmetic.go`:

```go
package arithmetic

func Add(operands ...int) int {
  result := operands[0]

  for _, v := range operands[1:] {
    result += v
  }

  return result
}

func Sub(operands ...int) int {
  result := operands[0]

  for _, v := range operands[1:] {
    result -= v
  }

  return result
}
```

`$GOPATH/src/local/arithmetic/example_test.go`:

```go
package arithmetic_test

import (
  "fmt"

  a "local/arithmetic"
)

func Example() {
  r := a.Add(1, 2) - a.Sub(3, 4)
  fmt.Println(r)
  // Output: 4
}

func ExampleAdd() {
  r := a.Add(1, 2, 3, 4)
  fmt.Println(r)
  // Output: 10
}

func ExampleSub() {
  r := a.Sub(5, 3, 1)
  fmt.Println(r)
  // Output: 1
}
```

Para ver los ejemplos se debe iniciar el servidor HTTP de GoDoc e ir a la ruta
<http://localhost:6060/pkg/local/arithmetic> con un navegador.

```shell-session
$ godoc -http :6060
```

Cada función de ejemplo deberá mostrar por la salida estándar los resultados,
y por cada salida que se realice, deberá existir un comentario especial
`// Output: VALOR` que indica el valor esperado. Estas funciones son ejecutadas
por `go test`, por lo que no solo tienen un uso informativo, sino que también
ayudan a probar el código; si no se encuentra algún comentario especial, las
funciones serán compiladas, pero no ejecutadas.

```shell-session
$ go test -v local/arithmetic
=== RUN   Example
--- PASS: Example (0.00s)
=== RUN   ExampleAdd
--- PASS: ExampleAdd (0.00s)
=== RUN   ExampleSub
--- PASS: ExampleSub (0.00s)
PASS
ok  	local/arithmetic
```

Para los casos en que se necesiten múltiples funciones de ejemplo de un mismo
objetivo, solo hace falta agregar un sufijo que inicie con un guión bajo y una
letra.

`$GOPATH/src/local/arithmetic/multiexample_test.go`:

```go
package arithmetic_test

import (
  "fmt"

  a "local/arithmetic"
)

func ExampleAdd_two() {
  r := a.Add(1, 2)
  fmt.Println(r)
  // Output: 3
}

func ExampleAdd_five() {
  r := a.Add(1, 2, 3, 4, 5)
  fmt.Println(r)
  // Output: 15
}
```

```shell-session
$ go test -v local/arithmetic
=== RUN   Example
--- PASS: Example (0.00s)
=== RUN   ExampleAdd
--- PASS: ExampleAdd (0.00s)
=== RUN   ExampleSub
--- PASS: ExampleSub (0.00s)
=== RUN   ExampleAdd_two
--- PASS: ExampleAdd_two (0.00s)
=== RUN   ExampleAdd_five
--- PASS: ExampleAdd_five (0.00s)
PASS
ok  	local/arithmetic
```

Como un ejemplo es representado por una función, no es posible demostrar
algunas funcionalidades como la implementación de interfaces, por esta razón
existen los ejemplos de archivos, estos consisten en un archivo que contiene
exclusivamente una función de ejemplo y todas las definiciones a nivel de
paquete que sean necesarias.

```shell-session
$ rm -rf $GOPATH/src/local/arithmetic
```

`$GOPATH/src/local/arithmetic/aritmetic.go`:

```go
package arithmetic

type Operander interface {
  Val() float64
}

func Add(operands ...Operander) float64 {
  result := operands[0].Val()

  for _, v := range operands[1:] {
    result += v.Val()
  }

  return result
}
```

`$GOPATH/src/local/arithmetic/whole_file_example_test.go`:

```go
package arithmetic_test

import (
  "fmt"

  a "local/arithmetic"
)

type Operand string

func (o Operand) Val() float64 {
  return float64(len(o))
}

func ExampleAdd() {
  var x, y Operand = "a", "b"

  r := a.Add(x, y)
  fmt.Println(r)
  // Output: 2
}
```

```shell-session
$ go test -v local/arithmetic
=== RUN   ExampleAdd
--- PASS: ExampleAdd (0.00s)
PASS
ok  	local/arithmetic
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

{{% loi %}}
* <https://golang.org/ref/spec#Boolean_types>
{{% /loi %}}

[George Boole]: https://es.wikipedia.org/wiki/George_Boole

Nombrados así en honor a [George Boole][], también son conocidos como
lógicos, representan valores de verdad (verdadero o falso) que normalmente son
usados para controlar el flujo de los programas.

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

## Numéricos

Existen tres grupos de datos numéricos:

* Enteros
* Punto flotante
* Complejos

### Enteros

{{% loi %}}
<!--lint disable no-undefined-references no-shortcut-reference-link-->
* <https://golang.org/ref/spec#Numeric_types>
* <https://golang.org/ref/spec#Integer_literals>
* [Números binarios]({{< relref "blog/binary-numbers.es.md" >}})
* [Números octales]({{< relref "blog/octal-numbers.es.md" >}})
* [Números hexadecimales]({{< relref "blog/hex-numbers.es.md" >}})
* [Complemento a dos]({{< relref "blog/twos-complement.es.md" >}})
<!--lint enable no-undefined-references no-shortcut-reference-link-->
{{% /loi %}}

Representan los números del conjunto matemático con el mismo nombre, aunque
claro, con una cantidad finita de elementos, que puede ser controlada por el
espacio de memoria que se reserve, es decir, el programador tiene la capacidad
de especificar si quiere un número entero que ocupe `N` bits de memoria, donde
`N` puede ser `8`, `16`, `32` o `64` debido a las especificaciones de los CPUs.

Existen dos tipos de números enteros, o mejor dicho, dos métodos de
representación: el primero es la conversión binaria tradicional, pero solo
puede ser usado para procesar números positivos; el segundo es llamado
*Complemento a dos* y permite representar tanto números positivos como
negativos de una manera bastante ingeniosa, solo que se pierde una cantidad
considerable de números positivos.

Además de números decimales, es posible usar otras notaciones como  octales y
hexadecimales.

#### Representación sintáctica

```go
// Enteros sin signo

uint8  // 11111111
       // ⬑ 8 bits  -> 0 - 255

uint16 // 1111111111111111
       // ⬑ 16 bits -> 0 - 65535

uint32 // 11111111111111111111111111111111
       // ⬑ 32 bits -> 0 - 4294967295

uint64 // 1111111111111111111111111111111111111111111111111111111111111111
       // ⬑ 64 bits -> 0 - 18446744073709551615

// Enteros con signo

      // ⬐ Signo
int8  // 11111111
      //  ⬑ Números, 7 bits -> -128 - 127

      // ⬐ Signo
int16 // 1111111111111111
      //  ⬑ Números, 15 bits -> -32768 - 32767

      // ⬐ Signo
int32 // 11111111111111111111111111111111
      //  ⬑ Números, 31 bits -> -2147483648 - 2147483647

      // ⬐ Signo
int64 // 1111111111111111111111111111111111111111111111111111111111111111
      //  ⬑ Números, 63 bits -> -9223372036854775808 - 9223372036854775807

// Alias

byte // Equivale a uint8
rune // Equivale a uint32

// Según la arquitectura del sistema operativo

uint    // Equivale a uint32 o uint64
int     // Equivale a int32 o int64
uintptr // Permite almacenar direcciones de memoria
```

#### Ejemplos

```go
5   // Decimal
05  // Octal (tienen el prefijo `0`)
0x5 // Hexadecimal (tienen el prefijo `0x`)

// Con signo

+10  // ┐
+012 // │-> Optimistas 😄
+0xa // ┘

-10  // ┐
-012 // │-> Pesimistas 😞
-0xa // ┘
```

#### Valor cero

```go
0
```

### Punto flotante

{{% loi %}}
<!--lint disable no-undefined-references no-shortcut-reference-link-->
* <https://golang.org/ref/spec#Numeric_types>
* <https://golang.org/ref/spec#Floating-point_literals>
* <http://www.oxfordmathcenter.com/drupal7/node/43>
* [Números binarios]({{< relref "blog/binary-numbers.es.md" >}})
* [Representación de números de punto flotante]({{< relref "blog/ieee-754.es.md" >}})
<!--lint enable no-undefined-references no-shortcut-reference-link-->
{{% /loi %}}

Representan al conjunto matemático de los números fraccionarios, aunque
claro, con una cantidad finita de elementos, que puede ser controlada por el
espacio de memoria que se reserve, es decir, el programador tiene la capacidad
de especificar si quiere un número entero que ocupe `N` bits de memoria, donde
`N` puede ser `32` o `64` según el estándar IEEE 754, que también especifica su
representación.

Un número de punto flotante literal está compuesto por dos enteros separados
por un punto (`.`), una letra `e`/`E` y otro entero; todos los enteros deben
escribirse en base 10 y pueden tener signo (exceptuando el segundo).

#### Representación sintáctica

```go
        // ⬐ Signo  ⬐ Fracción, 23 bits
float32 // 11111111111111111111111111111111
        //  ⬑ Exponente, 8 bits

        // ⬐ Signo     ⬐ Fracción, 52 bits
float64 // 1111111111111111111111111111111111111111111111111111111111111111
        //  ⬑ Exponente, 11 bits
```

#### Ejemplos

```go
0.         // Nivel de bondad en nuestra raza
3.14       // 14/03/1988
-9.8       // El mundo al reves
59724.e20  // Madre tierra
59724e20   // Madre tierra sin punto
.91093e-30 // http://bit.ly/2Iv08BI
111.09+e87 // Straight flush
```

#### Valor cero

```go
0
```

### Complejos

{{% loi %}}
<!--lint disable no-undefined-references no-shortcut-reference-link-->
* <https://golang.org/ref/spec#Numeric_types>
* <https://golang.org/ref/spec#Imaginary_literals>
* <https://golang.org/ref/spec#Constant_expressions>
* <https://golang.org/ref/spec#Complex_numbers>
* <http://www.oxfordmathcenter.com/drupal7/node/43>
* [Números binarios]({{< relref "blog/binary-numbers.es.md" >}})
* [Representación de números de punto flotante]({{< relref "blog/ieee-754.es.md" >}})
<!--lint enable no-undefined-references no-shortcut-reference-link-->
{{% /loi %}}

Representan los números del conjunto matemático con el mismo nombre, aunque
claro, con una cantidad finita de elementos, que puede ser controlada por el
espacio de memoria que se reserve, es decir, el programador tiene la capacidad
de especificar si quiere un número entero que ocupe `N` bits de memoria, donde
`N` puede ser `64` o `128` pues están conformados por un par de números de
punto flotante, representando la parte real y la imaginaria cada uno.

Un número complejo literal está compuesto por dos números reales (enteros o de
punto flotante) separados por una cruz (`+`) o un guión (`-`), y el último
número debe tener la letra `i` al final. Dentro de Go existe una función,
[`complex`](#complex), que permite crear números complejos usando variables (no
solo constantes como en el caso de los literales); y otras dos, [`real`](#real)
e [`imag`](#imag), que hacen lo opuesto, pues permiten extraer la parte real e
imaginaria de un número complejo respectivamente (por si no es obvio el orden
😂).

**TODO:** ¿Cómo son implementados los números complejos en Go?.

**TODO:** Agregar referencias de uso.

#### Representación sintáctica

```go
complex32
complex64
```

#### Ejemplos

```go
1 + 2i
3 - 4.5i
7e8 + 9e-10i

1i      // ┐
2.3i    // │-> Parte imaginaria, `0 + IMAGINARIO`
45.6e7i // ┘
```

#### Valor cero

```go
0
```

## Arreglos

<https://tour.golang.org/moretypes/6>
<https://blog.golang.org/slices>

## Porciones

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
<https://github.com/golang/go/wiki/SliceTricks>
<https://blog.golang.org/slices>
<https://research.swtch.com/godata>

## Cadenas

<!--lint disable no-undefined-references no-shortcut-reference-link-->

* [Codificación de texto]({{< relref "blog/text-encoding.es.md" >}})

<!--lint enable no-undefined-references no-shortcut-reference-link-->

* <https://golang.org/ref/spec#String_types>

* <https://golang.org/ref/spec#String_literals>

* <https://golang.org/ref/spec#Rune_literals>

* <https://blog.golang.org/slices>

* <https://blog.golang.org/strings>

* <https://research.swtch.com/godata>

Son un conjunto de bytes, cada uno de estos bytes puede representar o ser parte
de una runa (un punto de código Unicode codificado en UTF-8), que no es más
que un caracter dentro de la cadena; aunque los bytes y las runas sean datos
numéricos (`uint8` y `uint32` respectivamente), Go puede interpretarlos como
texto, es decir, si se intenta convertir el número `77` en una cadena
(`string(77)`), Go seleccionará el punto de código Unicode `U+004d` (`77` es
`4d` en números hexadecimales) y obtendrá la letra `M`.

Se usan las comillas (`"`) y los acentos graves (<code>\`</code>) para la
definición de cadenas literales, y a diferencia de otros lenguajes, el
apóstrofo (`'`) se usa para representar runas literales, no cadenas.

```go
"Yo soy una cadena interpretada\ny puedo procesar secuencias de escape 😎"
// Yo soy una cadena interpretada
// y puedo procesar secuencias de escape 😎
```

```go
`Yo soy una cadena sin formato\ny no puedo procesar secuencias de escape 😔`
// Yo soy una cadena sin formato\ny no puedo procesar secuencias de escape 😔

`Pero puedo tener varias líneas,
quien es mejor ahora 😒`
// Pero puedo tener varias líneas,
// quien es mejor ahora 😒
```

Como su unidad es el byte y no la runa, es posible que cadenas como `Hola` y `😂` tengan la misma longitud.

```go
len("Hola") // 4
// "Hola" es una cadena compuesta por cuatro bytes, cada uno
// representa una runa.
// 'H' ->  72 -> U+0048 -> 1001000
// 'o' -> 111 -> U+006f -> 1101111
// 'l' -> 108 -> U+006c -> 1101100
// 'a' ->  92 -> U+0061 -> 1100001

len("😂") // 4
// "😂" es una cadena compuesta por cuatro bytes, todos
// representan una runa
// '😂' -> 128514 -> U+1f602 -> 11110000 10011111 10011000 10000010
```

Y ya que son un tipo de dato compuesto (técnicamente `[]byte`), soportan
operaciones de porciones para acceder a sus elementos.

```go
"Hola"[0]   // 72, primer elemento (0)
"Hola"[1:3] // "ol", elementos del 1 (inclusivo) al 3 (exclusivo)
"Hola"[:2]  // "Ho", elementos del 0 (inclusivo) al 2 (exclusivo)
"Hola"[2:]  // "la", elementos del 2 (inclusivo) al último elemento (len("Hola"))
"Hola"[3]   // 97, último elemento (len("Hola") - 1)
"Hola"[:]   // "Hola", todos los elementos

"😂"[1] // 159, extrae uno de los bytes que componen la runa
```

Pero son inmutables.

```go
"Hola"[2] = 'L' // cannot assign to "Hola"[2]
```

También si no se tiene en cuenta el hecho de que algunos caracteres están
compuestos por más de un byte, iterar sobre ellos resultaría diferente a lo
esperado.

```go
x := "😂"

for i := 0; i < len(x); i++ {
  fmt.Println(x[i])
}
```

```shell-session
240
159
152
130
```

Por esto se recomienda usar `range`, que extrae runa a runa.

```go
for _,  v := range "😂" {
  fmt.Println(v)
}
```

```shell-session
128514
```

O para los casos en los que `range` tampoco cumpla con las expectativas (no se
quiera iterar sobre la cadena), se puede usar [`unicode/utf8.DecodeRuneInString`](https://golang.org/pkg/unicode/utf8/#DecodeRuneInString).

```go
package main

import (
  "fmt"
  "unicode/utf8"
)

func main() {
  x := "😂"

  // Sin iteración, extrae solo la primera runa
  fmt.Println(utf8.DecodeRuneInString(x))

  // Equivale a usar range
  for i := 0; i < len(x); {
    v, w := utf8.DecodeRuneInString(x[i:])
    fmt.Println(v)
    i += w
  }
}
```

```shell-session
128514 4
128514
```

### Representación sintáctica

```go
string
```

### Ejemplos

```go
'M'  //     74 -> U+004d  -> 1001101 (7 bits)
'😄' // 128516 -> U+1f604 -> 11110000 10011111 10011000 10000100 (4 bytes)

"C"
"Cadena de caracteres"

`Cadena
de
caracteres
multilineal`
```

#### Secuencias de escape

```go
"\a" // Bell character
"\b" // Backspace
"\t" // Horizontal tab
"\n" // Line feed
"\v" // Vertical tab
"\f" // Form feed
"\r" // Carriage return
"\"" // Quotation mark
"\\" // Backslash

'\a' // 7
'\b' // 8
'\t' // 9
'\n' // 10
'\v' // 11
'\f' // 12
'\r' // 13
'\'' // 39
'\\' // 92
```

También es posible escribir puntos de código Unicode o su representación en
bytes.

```go
// Unicode

  // Versión corta (u y 4 dígitos)

'\u004d' // 77
"\u004d" // "M"

  // Versión larga (U y 8 dígitos)

'\U0000004d' // 77
"\U0000004d" // "M"
'\U00f1f064' // 128516
"\U00f1f064" // "😄"

// Bytes (UTF-8)

  // Octales (3 dígitos)

'\115'                // 77
"\115"                // "M"
// '\360\237\230\204' // No soporta más de un byte escapado
"\360\237\230\204"    // "😄"

  // Hexadecimales (x y 2 dígitos)

'\x4d'                // 77
"\x4d"                // "M"
// '\xf0\x9f\x98\x84' // No soporta más de un byte escapado
"\xf0\x9f\x98\x84"    // "😄"
```

### Valor cero

```go
""
```

## Mapas

<https://tour.golang.org/moretypes/19>
<https://tour.golang.org/moretypes/20>
<https://tour.golang.org/moretypes/21>
<https://tour.golang.org/moretypes/22>

## Estructuras

<https://tour.golang.org/moretypes/2>
<https://tour.golang.org/moretypes/3>

## Punteros

<https://tour.golang.org/moretypes/1>
<https://tour.golang.org/moretypes/4>
<https://tour.golang.org/moretypes/5>

## Cambios de tipos

<https://tour.golang.org/basics/13>

<https://golang.org/ref/spec#Conversions>

# Constantes

* <https://golang.org/ref/spec#Constants>
* <https://golang.org/ref/spec#Constant_expressions>
* <https://golang.org/ref/spec#Complex_numbers>

<https://blog.golang.org/constants>

```go
const (
  x = 2
  y = 3i
)

x + y // (2+3i)
```

# Estructuras de repetición

## `for`

# Funciones

## Predefinidas

### `complex`

### `real`

### `imag`

# Manejo de errores

<https://golang.org/ref/spec#Handling_panics>

<https://blog.golang.org/error-handling-and-go>

# Compilador

{{% loi %}}
* <https://golang.org/pkg/go/build/>
{{% /loi %}}

GOPATH

GOROOT

GOTPMDIR

## Condiciones de compilación

{{% loi %}}
* <https://golang.org/pkg/go/build/#hdr-Build_Constraints>

* <https://www.youtube.com/watch?v=COCUqAwAbD0&t=0s&index=31&list=PL5MnW0XCND7IjWv810mg4H81BxYN8BPQh>
{{% /loi %}}

Permiten establecer condiciones para el compilador, como usar el archivo para
ciertas arquitecturas o sistemas operativos, deben aparecer entre las primeras líneas, incluso antes de `package`. Para usarlas, solo hace falta un comentario
como este `// +build CONDICION [...]`

# Buenas prácticas

# Atribuciones

**Go Team.** *Documentation* <https://golang.org/doc/>

**Ariel Mashraki.** *An overview of Go syntax and features.* <https://github.com/a8m/go-lang-cheat-sheet>

**Thomas Finley.** *Two's Complement.* <https://www.cs.cornell.edu/~tomf/notes/cps104/twoscomp.html>

**Autores de Wikipedia.** *List of Unicode characters.* <https://en.wikipedia.org/wiki/List_of_Unicode_characters>

