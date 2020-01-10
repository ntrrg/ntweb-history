---
title: Go (Golang)
date: 2018-09-23T16:40:00-04:00
image: images/go.png
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
toc: true
---

[Go license]: https://golang.org/LICENSE

Es un lenguaje minimalista y de alto rendimiento.  Su fase de diseño inició en
el año 2007 por parte de un equipo de ingenieros de Google, conformado en ese
tiempo por Ken Thompson, Rob Pike y Robert Griesemer; luego de tener una base
estable, se unieron los ingenieros Russ Cox e Ian Lance Taylor. Para inicios
del 2012 se liberó la primera versión estable, de código abierto y ditribuido
bajo una licencia [BSD-style][Go license].

Algunas de sus características más resaltantes son:

[GC]: https://es.wikipedia.org/wiki/Recolector_de_basura

* Imperativo, los programas se escriben como una serie de instrucciones que la
  computadora debe seguir para resolver un problema (leyendo esto se puede
  pensar *«¿Y no es así cómo se escriben todos los programas? 😒»*, 
  la respuesta es no, existen otros paradigmas de programación que trabajan con
  enfoques muy diferentes a este).

* Compilado, todo el código escrito es traducido a lenguaje máquina antes de
  poder ejecutarse, esto significa que no hace falta instalar Go en la máquina
  que se usará el binario generado.

* Tipado estático, una vez que se define el tipo de una variable, este no puede
  ser modificado.

* Fuertemente tipado, no permite realizar operaciones entre datos de diferente
  tipo, deben hacerse cambios de tipo explícitamente.

* No es necesario liberar manualmente la memoria asignada, usa un [GC][] que se
  encarga de esto, pero también ofrece facilidades para el manejo de memoria.

* Concurrencia y paralelismo de manera nativa (por medio de palabras reservadas
  y operadores, también tiene algunas bibliotecas que permiten aplicar técnicas
  de más bajo nivel).

* Minimalista, la mayoría de las utilidades que faltan en el lenguaje fueron
  [excluidas intencionalmente](#funcionalidades-excluidas).

# Funcionalidades excluidas

{{< loi >}}
* <https://golang.org/doc/faq#Design>
* <https://www.youtube.com/watch?v=k9Zbuuo51go>
{{< /loi >}}

* Genéricos. Aunque es posible que en alguna futura versión se agregue, por
  ahora no se ha logrado obtener una solución que compense su complejidad con
  su utilidad. En su lugar pueden usarse las [interfaces](#interfaces), que
  ofrecen abstracción de una manera muy elegante.

* Conjuntos. Por la falta de tipos genéricos en el lenguaje, representan
  agregar un nuevo tipo de dato al código base de Go y debido a que es bastante
  sencillo implementarlos en la mayoría de los casos, no es muy relevante su
  existencia.

{{< go-playground >}}
```go
x := make(map[int]struct{})

x[1] = struct{}{}
x[2] = struct{}{}
x[1] = struct{}{}

len(x) // 2
```

--- PLAYGROUND ---

```go
package main

import "fmt"

func main() {
  x := make(map[int]struct{})

  x[1] = struct{}{}
  x[2] = struct{}{}
  x[1] = struct{}{}

  fmt.Println(len(x))
}
```
{{< /go-playground >}}

* `while` y `do-while`. Solo hay una estructura de repetición (`for`) y aunque
  parezca limitado, es una ventaja para los programadores no tener que pensar
  en cuál estructura usar, tal vez suene a exagerar, pero en internet es muy
  fácil encontrar discusiones largas y repetitivas de varios lenguajes sobre
  cuál de todas las estructuras de repetición es la más rápida.

* `map`, `filter` y la familia de funciones favoritas de los programadores
  funcionales. Por la falta de tipos genéricos sería necesario definir
  muchísimas funciones para cada tipo, pero además, ¿por qué llamar 100
  funciones para sumar los elementos de un arreglo si puede usarse una
  estructura de repetición muy sencilla?, si la reacción a esto es *«No me
  importa el rendimiento, quiero mis funciones 😒»* no hay problema, es muy
  fácil implementarlas, pero en este caso les recomendaría usar otro lenguaje.

{{< go-playground id="oNGlnMctzXv" >}}
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
{{< /go-playground >}}

* Excepciones. Usar estructuras de control (como `try {} catch {}`) para
  manejar los errores puede resultar en flujos complejos que dificultan a los
  programadores el proceso de identificación y solución de errores (debugging).
  En su lugar los errores se manejan por medio del mecanismo de retorno
  múltiple en las funciones y por otras funciones predefinidas en Go.

* Afirmaciones (asserts), en la mayoría de las ocasiones, los programadores
  usan las afirmaciones para generar errores en tiempo de ejecución y así, no
  tener que implementar de mejor manera el flujo del programa. Para un
  comportamiento similar se puede usar la función `panic`, pero su uso para
  simplificar el manejo de errores no está recomendado.

* Aritmética de punteros. No está permitida, por la forma en que se maneja la
  memoria.

* Hilos de procesos (threads), una de las tareas que suele agregar muchísima
  complejidad al código fuente es la programación multithreading, aunque claro,
  si se pretende programar una aplicación que se usará en computadoras potentes
  (como servidores o computadores personales con procesadores de múltiples
  núcleos) y hacer toda la computación en un solo hilo, sería un descaro decir
  que Go es un lenguaje de alto rendimiento, pero la verdad es que no hacen
  falta, ya se que suena loco y probablemente se pueda pensar *«Claaaro, un
  programa con gran demanda de cómputo que corre en un hilo puede ser tan
  rápido como uno que corre en múltiples hilos.. 😒»*, pensamiento sarcástico
  que sería muy merecido, pero el hecho es que Go cuenta con goroutines, que
  son funciones que se ejecutan independientemente del hilo principal y son
  automáticamente distribuidas entre más hilos para evitar el bloqueo de las
  operaciones, esto genera una abstracción de más alto nivel para este tipo de
  tareas, por lo que el programador no debe lidear directamente con hilos (vear
  la sección de [Goroutines](#goroutines)).

# Herramientas necesarias

Para empezar a programar solo hacen falta dos cosas:

<!--lint disable no-undefined-references no-shortcut-reference-link-->

[Instalar Go]: ./../install-go-1.13/index.es.md
[Install Go]: https://golang.org/doc/install

* El compilador (se pueden ver las instrucciones para instalarlo en la
  [documentación oficial][Install Go] o en este [artículo][Instalar Go]).

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
especial que le dice al compilador que la intención del paquete no es servir
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

{{< go-playground id="hR9ZBMz-Pst" >}}
```go
package main

import "fmt"

func main() {
  fmt.Println("hola, mundo")
}
```
{{< /go-playground >}}

El compilador ofrece dos métodos para ejecutarlo: el primero y más sencillo es
usando el comando `go run`.

```shell-session
$ go run hola_mundo.go
hola, mundo
```

El segundo, es generar un archivo ejecutable a partir del archivo fuente y
después ejecutarlo (obvio no? 😅), el comando anterior hace esto mismo, solo
que crea un archivo temporal y lo ejecuta automáticamente.

```shell-session
$ go build hola_mundo.go
$ ./hola_mundo
hola, mundo
```

Aunque en algunos casos baste con un archivo para crear un paquete útil, en
otras ocasiones la cantidad de código tiende a expandirse y tener muchas líneas
en un solo lugar puede generar algunos problemas, por lo que es recomendable
leer la sección sobre [Modularización](#paquetes). Para obtener más información
sobre el comando `go` y como usarlo con múltiples archivos, se debe leer la
sección del [Compilador](#compilador).

# Comentarios

{{< loi >}}
* <https://golang.org/ref/spec#Comments>
{{< /loi >}}

Los comentarios son texto ignorado por el compilador, su función principal es
documentar ciertas secciones de código que sean un poco difíciles de entender
a simple vista, pero en muchas ocasiones también son usados para ocultar
código de los ojos del compilador y ver como se comporta el programa. Existen
dos tipos de comentarios:

* De línea

{{< go-playground id="4g5BEqD0RGU" >}}
```go
fmt.Println("hola, mundo") // Esto muestra "hola, mundo"

// Las sentencias comentadas no son procesadas por el compilador
// fmt.Println("chao, mundo")
```
{{< /go-playground >}}

* Generales

{{< go-playground id="4HyigTWqiZ8" >}}
```go
/*
  Así se escribe un comentario general

  fmt.Println("hola, mundo")
  fmt.Println("chao, mundo")

  Este programa no hace nada..
*/
```
{{< /go-playground >}}

## Documentación

{{< loi >}}
* <https://blog.golang.org/godoc-documenting-go-code>
{{< /loi >}}

[GoDoc]: https://godoc.org
[Docutils]: http://docutils.sourceforge.net/

[GoDoc][] es una herramienta que permite usar los comentarios para generar
documentación, algo parecida a [Docutils][] en Python, solo que un poco más
sencilla, pues no requiere de un lenguaje de marcas para generar buena
documentación, sino que usa texto plano.

El objetivo principal de la documentación son las definiciones (`package`,
`const`, `var`, `type`, `func`, etc...) exportadas, GoDoc procesará solo
aquellas precedidas directamente por una o más líneas de comentarios.

`$GOPATH/src/local/arithmetic/arithmetic.go`:

```go
// Package arithmetic provides arithmetic operations for any type.
package arithmetic

// Identity constants
const (
  AdditiveIdentity       = 0
  MultiplicativeIdentity = 1
)

// Operander is the interface that wraps the arithmetic representation
// methods.
//
// Val returns the variable's arithmetic representation (float64).
type Operander interface {
  Val() float64
}

// Add gets any number of Operander and returns their addition.
func Add(operands ...Operander) float64 {
  result := operands[0].Val()

  for _, v := range operands[1:] {
    if v.Val() == AdditiveIdentity {
      continue
    }

    result += v.Val()
  }

  return result
}
```

Para ver el resultado se debe iniciar GoDoc e ir a la ruta <http://localhost:6060/pkg/local/arithmetic>
con un navegador web.

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

* Tiene el formato `IDENTIFICADOR(USUARIO): DESCRIPCIÓN.`, será agregado a la
  lista de notas del paquete. `IDENTIFICADOR` puede ser cualquier combinación
  de más de dos letras mayúsculas. El idenficador `BUG` tiene el comportamiento
  especial de crear una lista de bugs en la página del paquete.

Cuando se tiene un paquete con múltiple archivos, cada uno de ellos tendrá la
sentencia `package NOMBRE`, pero esto no quiere decir que sea necesario repetir
el comentario del paquete en cada archivo, en realidad basta con que uno de los
archivos lo tenga, por esto, si la documentación es algo extensa, se
recomienda crear un archivo `doc.go` que contenga solo en nombre del paquete y
su comentario de documentación.

```go
/*
Package arithmetic provides arithmetic operations for any type.

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

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin
euismod egestas elit sed viverra. Nunc tincidunt lacinia orci in
mattis. Praesent cursus neque et dapibus faucibus. Maecenas at
sem ut arcu ornare commodo. Morbi laoreet diam sit amet est
ultricies imperdiet. Proin ullamcorper ac massa a accumsan.
Praesent quis bibendum tellus. Sed id velit libero. Fusce dapibus
purus neque, sit amet sollicitudin odio porttitor posuere. Mauris
eu dui elementum, fermentum ante vitae, porttitor nunc. Duis mi
elit, viverra at turpis vitae, sollicitudin aliquet velit.
Pellentesque nisl turpis, pulvinar et consectetur et, iaculis vel
leo. Suspendisse euismod sem at vehicula fermentum. Duis viverra
eget ante a accumsan.

Aenean dui lectus, ultrices at elit id, pellentesque faucibus
dolor. Duis blandit vulputate est, eget sollicitudin ipsum
pellentesque quis. Cras sed nibh sed sapien suscipit tincidunt
venenatis id eros. Praesent laoreet, erat quis hendrerit
dignissim, justo diam semper elit, sit amet commodo lacus ipsum
eget nisl. In a mi tellus. In hac habitasse platea dictumst.
Aliquam et neque a quam mollis molestie. Etiam tempor arcu quis
justo molestie congue.
*/
package arithmetic
```

Además de texto, GoDoc da la posibilidad de mostrar el funcionamiento con
ejemplos dinámicos, que pueden ser ejecutados e incluso modificados en la
interfaz web. Para más información ver la sección de [Ejempos](#ejemplos-pruebas).

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

{{< loi >}}
* <https://golang.org/ref/spec#Boolean_types>
{{< /loi >}}

[George Boole]: https://es.wikipedia.org/wiki/George_Boole

Nombrados así en honor a [George Boole][], también son conocidos como
lógicos, representan valores de verdad (verdadero o falso) que normalmente son
usados para controlar el flujo de los programas.

**TODO:** ¿Cómo son implementados?

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

{{< loi >}}
<!--lint disable no-undefined-references no-shortcut-reference-link-->
* <https://golang.org/ref/spec#Numeric_types>
* <https://golang.org/ref/spec#Integer_literals>
* [Números binarios]({{< relref "/blog/binary-numbers.es.md" >}})
* [Números octales]({{< relref "/blog/octal-numbers.es.md" >}})
* [Números hexadecimales]({{< relref "/blog/hex-numbers.es.md" >}})
* [Complemento a dos]({{< relref "/blog/twos-complement.es.md" >}})
<!--lint enable no-undefined-references no-shortcut-reference-link-->
{{< /loi >}}

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

```
10101010 -> 170
⬑ 8 bits -> 0 - 255

1010101010101010 -> 43690
⬑ 16 bits -> 0 - 65535

10101010101010101010101010101010 -> 2863311530
⬑ 32 bits -> 0 - 4294967295

1010101010101010101010101010101010101010101010101010101010101010 -> 12297829382473034410
⬑ 64 bits -> 0 - 18446744073709551615
```

```
⬐ Signo
10101010 -> -86
 ⬑ Números, 7 bits -> -128 - 127

⬐ Signo
0101010101010101 -> 21845
 ⬑ Números, 15 bits -> -32768 - 32767

⬐ Signo
10101010101010101010101010101010 -> -1431655766
 ⬑ Números, 31 bits -> -2147483648 - 2147483647

⬐ Signo
0101010101010101010101010101010101010101010101010101010101010101 -> 6148914691236517205
 ⬑ Números, 63 bits -> -9223372036854775808 - 9223372036854775807
```

Además de números decimales, es posible usar otras notaciones como  octales y
hexadecimales para expresar enteros literales.

#### Representación sintáctica

```go
// Enteros sin signo

uint8
uint16
uint32
uint64

// Enteros con signo

int8
int16
int32
int64

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

{{< loi >}}
<!--lint disable no-undefined-references no-shortcut-reference-link-->
* <https://golang.org/ref/spec#Numeric_types>
* <https://golang.org/ref/spec#Floating-point_literals>
* <http://www.oxfordmathcenter.com/drupal7/node/43>
* [Números binarios]({{< relref "/blog/binary-numbers.es.md" >}})
* [Representación de números de punto flotante]({{< relref "/blog/ieee-754.es.md" >}})
<!--lint enable no-undefined-references no-shortcut-reference-link-->
{{< /loi >}}

Representan al conjunto matemático de los números fraccionarios, aunque
claro, con una cantidad finita de elementos, que puede ser controlada por el
espacio de memoria que se reserve, es decir, el programador tiene la capacidad
de especificar si quiere un número entero que ocupe `N` bits de memoria, donde
`N` puede ser `32` o `64` según el estándar IEEE 754, que también especifica su
representación.

```
⬐ Signo  ⬐ Fracción, 23 bits
10101010101010101010101010101010
 ⬑ Exponente, 8 bits

⬐ Signo     ⬐ Fracción, 52 bits
1010101010101010101010101010101010101010101010101010101010101010
 ⬑ Exponente, 11 bits
```

Un número de punto flotante literal está compuesto por dos enteros separados
por un punto (`.`), una letra `e`/`E` y otro entero; todos los enteros deben
escribirse en base 10 y pueden tener signo (exceptuando el segundo).

**TODO:** ¿Cómo son implementados?

#### Representación sintáctica

```go
float32
float64
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

{{< loi >}}
* <https://golang.org/ref/spec#Numeric_types>
* <https://golang.org/ref/spec#Imaginary_literals>
* <https://golang.org/ref/spec#Constant_expressions>
* <https://golang.org/ref/spec#Complex_numbers>
{{< /loi >}}

Representan los números del conjunto matemático con el mismo nombre, aunque
claro, con una cantidad finita de elementos, que puede ser controlada por el
espacio de memoria que se reserve, es decir, el programador tiene la capacidad
de especificar si quiere un número entero que ocupe `N` bits de memoria, donde
`N` puede ser `64` o `128` pues están conformados por un par de números de
punto flotante, representando la parte real y la imaginaria cada uno.

```
⬐ Parte real, 32 bits           ⬐ Parte imaginaria, 32 bits
1010101010101010101010101010101010101010101010101010101010101010

⬐ Parte real, 64 bits                                           ⬐ Parte imaginaria, 64 bits
10101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010
```

Un número complejo literal está compuesto por dos números reales (enteros o de
punto flotante) separados por una cruz (`+`) o un guión (`-`), y el último
número debe tener la letra `i` al final. Dentro de Go existe una función,
[`complex`](#complex), que permite crear números complejos usando variables (no
solo constantes como en el caso de los literales); y otras dos, [`real`](#real)
e [`imag`](#imag), que hacen lo opuesto, pues permiten extraer la parte real e
imaginaria de un número complejo respectivamente (por si no es obvio el orden
😂).

**TODO:** Agregar referencias de uso

#### Representación sintáctica

```go
complex64
complex128
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

{{< loi >}}
* <https://tour.golang.org/moretypes/6>
* <https://golang.org/ref/spec#Array_types>
* <https://golang.org/ref/spec#Composite_literals>
* <https://golang.org/ref/spec#Length_and_capacity>
* <https://blog.golang.org/go-slices-usage-and-internals>
* <https://blog.golang.org/slices>
* <https://research.swtch.com/godata>
{{< /loi >}}

Son un conjunto de elementos de algún tipo de dato asignado arbitrariamente, la
cantidad debe ser una constante y no puede cambiar después de su creación.

Todos los elementos están enumerados e inician en la posición `0`, a estas
posiciones se les llama *índices* y se usa la notación `x[i]` para acceder a
sus elementos, donde `x` es un arreglo e `i` el índice. También soportan
operaciones de porciones, que consisten en tomar un subconjunto de elementos
del arreglo, para esto se usa una notación parecida, `x[i:j]`, donde `x` es un
arreglo, `i` el índice inicial inclusivo y `j` el índice final exclusivo, pero
en este caso el tipo de dato obtenido no es un arreglo, sino una porción.

```
    ┌─┬─┬─┬─┬─┐
x = │1│3│5│7│9│
    └─┴─┴─┴─┴─┘
     0 1 2 3 4

x[0]   -> 1
x[2]   -> 5
x[4]   -> 9
x[0:2] -> [1 3]
x[3:5] -> [7 9]
x[:3]  -> [1 3 5]
x[2:]  -> [5 7 9]
x[:]   -> [1 3 5 7 9]

x[0] = 0
x[4] = 0
x[:] -> [0 3 5 7 0]
```

Internamente, no son más que un bloque de memoria reservado que tiene a todos
sus elementos uno después de otro, es decir, si se crea un arreglo de bytes con
los cuatro primeros números pares, el espacio de memoria ocupado por el arreglo
sera 4 bytes (16 bits normalmente) y sus elementos se ubicarán en estos bytes
según sus indices.

```
    ┌─┬─┬─┬─┐
x = │2│4│6│8│ -> 1 byte x 4 elementos -> 4 bytes
    └─┴─┴─┴─┘
     0 1 2 3

Ubicación en la memoria: 0x10313020

x[0] -> 0 * 1 byte -> 0x10313020 + 0 -> 0x10313020 -> 00000010 -> 2
x[1] -> 1 * 1 byte -> 0x10313020 + 1 -> 0x10313021 -> 00000100 -> 4
x[2] -> 2 * 1 byte -> 0x10313020 + 2 -> 0x10313022 -> 00000110 -> 6
x[3] -> 3 * 1 byte -> 0x10313020 + 3 -> 0x10313023 -> 00001000 -> 8
```

Del mismo modo pasa con los primeros cuatro números pares después del límite de
un byte, la única diferencia es que ocuparán el doble de memoria.

```
    ┌───┬───┬───┬───┐
x = │256│258│260│262│ -> 2 bytes (uint16) x 4 elementos -> 8 bytes
    └───┴───┴───┴───┘
      0   1   2   3

Ubicación en la memoria: 0x10313020

x[0] -> 0 * 2 bytes -> 0x10313020 + 0 -> 0x10313020 -> 0000000100000000 -> 256
x[1] -> 1 * 2 bytes -> 0x10313020 + 2 -> 0x10313022 -> 0000000100000010 -> 258
x[2] -> 2 * 2 bytes -> 0x10313020 + 4 -> 0x10313024 -> 0000000100000100 -> 260
x[3] -> 3 * 2 bytes -> 0x10313020 + 6 -> 0x10313026 -> 0000000100000110 -> 262
```

Para obtener la cantidad de elementos de un arreglo se debe usar la función
`len(ARREGLO)` que retorna un número entero del tipo `int`.

{{< go-playground id="vpsI0bAQlYS" >}}
```go
x := [3]int{1, 2, 3}

len(x)) // 3
```
{{< /go-playground >}}

### Representación sintáctica

```
[CANTIDAD]TIPO
```

### Ejemplos

```go
[5]byte{1, 2, 3, 4, 5}   // [1 2 3 4 5]
[...]byte{1, 2, 3, 4, 5} // Igual que el de arriba, solo que obtiene
                         // la cantidad de elementos automáticamente

[3]bool{} // [false false false]
          // Inicializa todos los elementos con su valor 0

[3]bool{true} // [true false false]
              // Se pueden indicar solo los primero valores y los
              // demás serán inicializados con valor 0.

[5]byte{2: 'M'} // [0 0 77 0 0]
                // Se pueden asignar valores a índices específicos,
                // los demás serán inicializados con su valor 0

[...]byte{2: 'M', 'A', 4: 'R', 'N'} // [0 0 77 64 0 82 78]
                                    // Si se especifíca un índice, los
                                    // siguientes elementos sin índice
                                    // sumarán uno al valor anterior

[...]string{    // Se pueden usar varias líneas para mejorar la
  "Miguel",     // legibilidad
  "Angel",
  "Rivera",
  "Notararigo", // Pero incluso el último elemento deberá tener una
}               // coma

[...]struct{ X, Y float64 }{
  struct{ X, Y float64 }{5, 10},

  {15, 30}, // Se puede omitir el tipo de dato en los elementos
}
```

### Valor cero

```go
[VALOR_CERO_0 ... VALOR_CERO_N]
```

## Porciones

{{< loi >}}
* <https://tour.golang.org/moretypes/7>
* <https://tour.golang.org/moretypes/8>
* <https://tour.golang.org/moretypes/9>
* <https://tour.golang.org/moretypes/10>
* <https://tour.golang.org/moretypes/11>
* <https://tour.golang.org/moretypes/12>
* <https://tour.golang.org/moretypes/13>
* <https://tour.golang.org/moretypes/14>
* <https://tour.golang.org/moretypes/15>
* <https://golang.org/ref/spec#Slice_types>
* <https://golang.org/ref/spec#Composite_literals>
* <https://golang.org/ref/spec#Length_and_capacity>
* <https://golang.org/ref/spec#Making_slices_maps_and_channels>
* <https://golang.org/ref/spec#Appending_and_copying_slices>
* <https://blog.golang.org/go-slices-usage-and-internals>
* <https://blog.golang.org/slices>
* <https://research.swtch.com/godata>
* <https://github.com/golang/go/wiki/SliceTricks>
{{< /loi >}}

Al igual que los arreglos, son un conjunto de elementos de un tipo de dato
asignado arbitrariamente, pero con algunas diferencias importantes, entre las
cuales destaca la posibilidad de alterar su tamaño después de crearse, por lo
que generalmente son más comunes en el código fuente. Sus elementos también
están enumerados como los arreglos y soportan operaciones de porciones (ya se
que por algo se llaman porciones, pero es bueno aclararlo 😅).

Otra diferencia con los arreglos, es la forma en la que son implementadas
internamente por el lenguaje, pues en lugar de representar bloques de memoria,
son estructuras de datos que contienen un puntero a un elemento de un arreglo;
una longitud, que determina la cantidad de elementos que pertenecen a la
porción después del referenciado por el puntero; y una capacidad, que es la
máxima longitud que puede tener la porción, calculada por la cantidad de
elementos desde el referenciado por el puntero hasta el final del arreglo.

```
    ┌─┬─┬─┬─┬─┐
x = │1│3│5│7│9│
    └─┴─┴─┴─┴─┘
     0 1 2 3 4

y = x[:2]

     ┌─────┬───┬───┐    ┌─┬─┐ ┌─┬─┬─┐ 
y -> │&x[0]│ 2 │ 5 │ -> │1│3│ │5│7│9│ 
     └─────┴───┴───┘    └─┴─┘ └─┴─┴─┘ 
       ptr  lon cap      0 1   2 3 4

y[:]  -> [1 3]
y[:2] -> [1 3]
y[:5] -> [1 3 5 7 9]
y[:6] -> Error, sobrepasa la capacidad
y[2]  -> Error, sobrepasa la longitud

z = x[1:4]

     ┌─────┬───┬───┐    ┌─┬─┬─┐ ┌─┐
z -> │&x[1]│ 3 │ 4 │ -> │3│5│7│ │9│
     └─────┴───┴───┘    └─┴─┴─┘ └─┘
       ptr  lon cap      0 1 2   3

z[:]  -> [3 5 7]
z[:2] -> [3 5]
z[:4] -> [3 5 7 9]
z[:5] -> Error, sobrepasa la capacidad
y[3]  -> Error, sobrepasa la longitud

a = x[3:]

     ┌─────┬───┬───┐    ┌─┬─┐
a -> │&x[3]│ 2 │ 2 │ -> │7│9│
     └─────┴───┴───┘    └─┴─┘
       ptr  lon cap      0 1

a[:]  -> [7 9]
a[:2] -> [7 9]
a[:3] -> Error, sobrepasa la capacidad
a[2]  -> Error, sobrepasa la longitud
```

Es posible limitar su capacidad agregando un tercer índice a la sintaxis de
porciones (`x[i:j:k]`), y al igual que el segundo índice, es exclusivo.

```
b = x[:3:4] -> Solo el primer índice es opcional con esta sintaxis

     ┌─────┬───┬───┐    ┌─┬─┬─┐ ┌─┐
b -> │&x[0]│ 3 │ 4 │ -> │1│3│5│ │7│
     └─────┴───┴───┘    └─┴─┴─┘ └─┘
       ptr  lon cap      0 1 2   3

b[:]  -> [1 3 5]
b[:2] -> [1 3]
b[:4] -> [1 3 5 7]
b[:5] -> Error, sobrepasa la capacidad
```

Ya que las porciones solo tienen una referencia a un arreglo, pasarlas como
argumentos es una operación muy ligera, pero esto quiere decir que cualquier
modificación que se haga a los valores de una porción, afectará a las demás con
el mismo arreglo.

```
    ┌─┬─┬─┬─┐
x = │2│4│6│8│
    └─┴─┴─┴─┘
     0 1 2 3

y = [:3]
z = [1:]

x -> [2 4 6 8]
y -> [2 4 6]
z -> [4 6 8]

x[1] = 3

x -> [2 3 6 8]
y -> [2 3 6]
z -> [3 6 8]

y[0] = 1

x -> [1 3 6 8]
y -> [1 3 6]
z -> [3 6 8]

z[2] = 9

x -> [1 3 6 9]
y -> [1 3 6]
z -> [3 6 9]
```

Para obtener la longitud y la capacidad de una porción se deben usar las
funciones `len(PORCIÓN)` y `cap(PORCIÓN)`, ambas retornan un número entero del
tipo `int`.

{{< go-playground id="l9D0hIL8Mpl" >}}
```go
x := [5]int{1, 2, 3, 4, 5}
y := x[1:4]

len(y) // 3
cap(y) // 4
```
{{< /go-playground >}}

Es posible inicializar una porción sin valores literales, para esto se puede
usar la función `make`, que recibe tres argumentos: el tipo de porción, la
longitud y opcionalmente la capacidad.

{{< go-playground id="QqtBDs72WGQ" >}}
```go
x := make([]bool, 3)
// [false false false]

y := make([]byte, 3, 5)
// [0 0 0]

z := y[:cap(y)]
// [0 0 0 0 0]
```
{{< /go-playground >}}

Existen dos funciones que ayudan con el trabajo cotidiano de las porciones, la
primera es `append`, que permite agregar elementos al final de una porción,
recibe como argumentos una porción de un tipo específico y una lista de
datos del mismo tipo, retorna una nueva porción que dependiendo de la
capacidad, reutilizará el arreglo referenciado por la porción pasada como
argumento o creará uno nuevo que pueda almacenar los elementos.

{{< go-playground id="gaW_r9YvadO" >}}
```go
a := []byte{1, 2, 3, 4, 5}
b := a[:3]
c := a[2:]

// [1 2 3 4 5]
// [1 2 3] 3 5
// [3 4 5] 3 3

x := append(b, 6)

// La capacidad de b es 5 y su longitud 3, esto quiere decir que
// todavía quedan 2 (5-3) índices reusables en el arreglo referenciado,
// por lo que se agregará el nuevo valor en el índice después de la
// porción (3)

// [1 2 3 6 5]
// [1 2 3] 3 5
// [3 6 5] 3 3
// [1 2 3 6] 4 5

y := append(c, 7, 8)

// La capacidad de c es 3 y su longitud 3, esto quiere decir que
// no quedan índices reusables en el arreglo referenciado, por lo que
// se creará uno nuevo que logre almacenar los valores, pero con una
// capacidad un poco difícil de predecir pues por ahora no hay un
// comportamiento definido en la especificación del lenguaje y puede
// variar entre sus implementaciones

// [1 2 3 6 5]
// [1 2 3] 3 5
// [3 6 5] 3 3
// [1 2 3 6] 4 5
// [3 6 5 7 8] 5
```
{{< /go-playground >}}

La segunda función es `copy`, se encarga de copiar elementos de una porción a
otra, recibe dos porciones del mismo tipo como argumento y la primera es a la
que se copiarán los elementos, al finalizar retorna la cantidad de elementos
copiados, que es determinada por la mínima longitud entre ambas porciones.

{{< go-playground id="zmWI34jS_Pv" >}}
```go
x := make([]int, 2)
y := []int{1, 2, 3, 4}

copy(x, y)     // 2
fmt.Println(x) // [1 2]

a := []byte{'a', 'b', 'c', 'o', 'u'}
b := "aei"

copy(a, b)     // 3
fmt.Println(a) // "aeiou"

n := []bool{true, true, false, false, true}
m := []bool{false, true}

copy(n[1:3], m) // 2
fmt.Println(n)  // [true false true false true]
```
{{< /go-playground >}}

Ya que las porciones hacen referencia a arreglos, aunque una porción solo tenga
algunos elementos, mantendrá completo en memoria su arreglo referenciado, es
decir, aunque exista una porción con solo dos elementos, si el arreglo
referenciado tiene mil elementos, estos mil elementos se mantendrán en memoria
hasta que todas sus porciones sean liberadas, por esto, cuando se pretende
tener una porción que pase por gran parte del programa y no importe el
contenido completo de su arreglo, es recomendable copiar los elementos de la
porción a una nueva con un arreglo propio.

### Representación sintáctica

```
[]TIPO
```

### Ejemplos

```go
[]byte{1, 2, 3, 4, 5} // [1 2 3 4 5]

[]byte{2: 'M'} // [0 0 77]
               // Se pueden asignar valores a índices específicos,
               // los demás serán inicializados con su valor 0

[]byte{2: 'M', 'A', 4: 'R', 'N'} // [0 0 77 64 0 82 78]
                                 // Si se especifíca un índice, los
                                 // siguientes elementos sin índice
                                 // sumarán uno al valor anterior

[]string{       // Se pueden usar varias líneas para mejorar la
  "Miguel",     // legibilidad
  "Angel",
  "Rivera",
  "Notararigo", // Pero incluso el último elemento deberá tener una
}               // coma

[]struct{ X, Y float64 }{
  struct{ X, Y float64 }{5, 10},

  {15, 30}, // Se puede omitir el tipo de dato en los elementos
}
```

### Valor cero

```go
nil
```

## Cadenas

{{< loi >}}
<!--lint disable no-undefined-references no-shortcut-reference-link-->
* <https://golang.org/ref/spec#String_types>
* <https://golang.org/ref/spec#String_literals>
* <https://golang.org/ref/spec#Rune_literals>
* <https://blog.golang.org/slices>
* <https://blog.golang.org/strings>
* <https://research.swtch.com/godata>
* [Codificación de texto]({{< relref "/blog/text-encoding.es.md" >}})
<!--lint enable no-undefined-references no-shortcut-reference-link-->
{{< /loi >}}

Son un conjunto de bytes, cada uno de estos bytes puede representar o ser parte
de una runa (un punto de código Unicode codificado en UTF-8), que no es más
que un caracter/símbolo para el ojo humano; aunque los bytes y las runas sean
datos numéricos (`uint8` y `uint32` respectivamente), Go puede interpretarlos
como texto, es decir, si se intenta representar el número `77` como una cadena,
Go seleccionará el punto de código Unicode `U+004d` (`77` es `4d` en números
hexadecimales), lo codificará con UTF-8 y obtendrá la letra `M`.

Para la definición de cadenas literales interpretadas se usan las comillas
(`"`) y para las cadenas sin formato los acentos graves (<code>\`</code>); a
diferencia de otros lenguajes, el apóstrofo (`'`) se usa para representar runas
literales, no cadenas.

{{< go-playground id="M0lvf5r9D8p" >}}
```go
"Soy una cadena interpretada\ny puedo procesar secuencias de escape 😎"
// Soy una cadena interpretada
// y puedo procesar secuencias de escape 😎

`Soy una cadena sin formato\ny no puedo procesar secuencias de escape 😔

Pero puedo tener varias líneas,
quien es mejor ahora 😒`
// Soy una cadena sin formato\ny no puedo procesar secuencias de escape 😔
//
// Pero puedo tener varias líneas,
// quien es mejor ahora 😒
```
{{< /go-playground >}}

Las cadenas interpretadas y las runas tienen la capacidad de procesar
secuencias de escape, estas secuencias son caracteres precedidos por una barra
invertida (`\`) que les permite alterar su comportamiento.

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

// Unicode

  // Versión corta (u y 4 dígitos)

"\u004d" // "M"
'\u004d' // 77

  // Versión larga (U y 8 dígitos)

"\U0000004d" // "M"
'\U0000004d' // 77
"\U00f1f064" // "😄"
'\U00f1f064' // 128516

// Bytes (UTF-8)

  // Octales (3 dígitos)

"\115"                // "M"
'\115'                // 77
"\360\237\230\204"    // "😄"
// '\360\237\230\204' // No soporta más de un byte escapado

  // Hexadecimales (x y 2 dígitos)

"\x4d"                // "M"
'\x4d'                // 77
"\xf0\x9f\x98\x84"    // "😄"
// '\xf0\x9f\x98\x84' // No soporta más de un byte escapado
```

Internamente, Go implementa las cadenas como porciones de bytes (`[]byte`), por
lo que cuentan con casi todas las cualidades de las porciones, solo que son
inmutables y por esta misma razón no tienen capacidad.

{{< go-playground id="yHrBgqgfqE9" >}}
```go
x := "Hola"

x[2] = 'L' // Error
cap(x)     // Error
```
{{< /go-playground >}}

Como su unidad es el byte y no la runa, es posible que cadenas como `Hola` y
`😂` tengan la misma longitud.

{{< go-playground id="oCaft33c5jj" >}}
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
{{< /go-playground >}}

Por lo que al iterar sobre ellas no se obtendrán caracteres/símbolos sino su
representación en UTF-8.

{{< go-playground id="y0O2H_Y91Tc" >}}
```go
x := "😂"

for i := 0; i < len(x); i++ {
  fmt.Println(x[i])
}

// 240 -> 11110000
// 159 -> 10011111
// 152 -> 10011000
// 130 -> 10000010
```
{{< /go-playground >}}

Para evitar este comportamiento se puede usar `range`, que extrae runa a runa.

{{< go-playground id="CcnClPYtrEn" >}}
```go
for _,  v := range "😂" {
  fmt.Println(v)
}

// 128514
```
{{< /go-playground >}}

O [`unicode/utf8.DecodeRuneInString`](https://golang.org/pkg/unicode/utf8/#DecodeRuneInString)
en los casos que no se quiera iterar sobre la cadena.

{{< go-playground id="cStYBcRb9ZX" >}}
```go
x := "😂"

// Sin iteración, extrae solo la primera runa y retorna la cantidad de
// bytes que se leyeron.
utf8.DecodeRuneInString(x) // 128514 4

// Equivale a usar range
for i := 0; i < len(x); {
  v, w := utf8.DecodeRuneInString(x[i:])
  fmt.Println(v)
  i += w
}

// 128514
```
{{< /go-playground >}}

### Representación sintáctica

```go
string
```

### Ejemplos

```go
'M'  // 74 -> U+004d -> 1001101 (7 bits)
'😄' // 128516 -> U+1f604 -> 11110000 10011111 10011000 10000100 (4 bytes)

"C"
"Cadena de caracteres"

`Cadena
de
caracteres
multilineal`
```

### Valor cero

```go
""
```

## Mapas

{{< loi >}}
* <https://tour.golang.org/moretypes/19>
* <https://tour.golang.org/moretypes/20>
* <https://tour.golang.org/moretypes/21>
* <https://tour.golang.org/moretypes/22>
* <https://golang.org/ref/spec#Map_types>
* <https://golang.org/ref/spec#Composite_literals>
* <https://golang.org/ref/spec#Length_and_capacity>
* <https://golang.org/ref/spec#Deletion_of_map_elements>
* <https://golang.org/ref/spec#Making_slices_maps_and_channels>
* <https://blog.golang.org/go-maps-in-action>
* <https://golang.org/ref/spec#Comparison_operators>
{{< /loi >}}

**Nota:** cada vez que mencione a los arreglos, también hago referencia a los
demás tipos que derivan de ellos, como las porciones y las cadenas.

Son una estructura de datos que permite acceder a valores por medio de índices
del tipo especificado (que no sea función, porción o mapa, pues no son valores
comparables) durante su definición, a estos índices se les llaman claves, y a
diferencia de los arreglos, el orden de sus elementos es irrelevante.

La especificación del lenguaje no regula la manera en que son implementados,
por lo que cada compilador puede tener su propia forma de manejarlos, lo único
que debe mantenerse es que sean tipos de datos referenciales, como las
porciones o los punteros.

Para crear mapas se pueden usar valores literales.

{{< go-playground id="FGRrpitkgtQ" >}}
```go
x := map[string]int {
  "cero": 0,
  "uno":  1,
  "dos":  2,
  "tres": 3,
}
```
{{< /go-playground >}}

O la función `make`, que permite crear mapas vacíos, recibe como argumentos un
tipo de mapa y opcionalmente una capacidad aproximada, que a diferencia de las
porciones no representa un límite, sino la cantidad de elementos a las que se
les reservará memoria automáticamente, esto evitará futuras tareas de
reservación de memoria por lo que mejorará el rendimiento, pero estos espacios
no serán contados en su longitud hasta que reciban algún valor, cosa que puede
comprobarse usando la función `len(MAPA)`, que retorna la cantidad de elementos
dentro del mapa y la representa con un número entero del tipo `int`.

{{< go-playground id="p1eBzG_B9_G" >}}
```go
x := make(map[string]bool, 10)

x["go"] = true
x["javascript"] = true
x["python"] = true
x["php"] = true
x["c#"] = false

len(x) // 5
```
{{< /go-playground >}}

Al igual que los arreglos, para acceder a sus valores se usan los corchetes
(`[]`). Intentar acceder a una clave que no existe retornará el valor cero del
tipo de dato que pueda recibir el mapa, para verificar la existencia de una
clave se debe realizar una doble asignación, la primera variable recibirá el
valor almacenado, y la segunda variable un booleano que será `true` si la clave
existe o `false` en caso contrario.

{{< go-playground id="61Is7Ve_W4G" >}}
```go
x := map[string][]int{
  "pares": {2, 4, 6, 8},
  "impares": {1, 3, 5, 7, 9},
}

y := x["impares"]   // [1 3 5 7 9]
z, ok := x["pares"] // [2 4 6 8] true

a := x["fraccionales"] // []
b, ok := x["enteros"]  // [] false
```
{{< /go-playground >}}


La creación de nuevos pares clave-valor y la modificación de valores existentes
son tareas bastante sencillas, que consisten en simplemente referenciar la
clave que se quiere crear/modificar y asignarle un valor.

{{< go-playground id="BCPhbpeY_K3" >}}
```go
x := map[bool][]interface{}{
  true: []interface{}{0, "True", []int{1, 2}},
}

x[false] = []interface{}{0, "", []int(nil)}     // Asignación
x[true] = []interface{}{1, "True", []int{1, 2}} // Modificación
```
{{< /go-playground >}}

Ya que sus claves no ofrecen ninguna garantía de orden, usar `range` o
simplemente mostrarlos como una cadena podría resultar en un comportamiento
impredecible.

{{< go-playground id="89nUjKLW7nn" >}}
```go
x := map[string]struct{ X, Y float64 }{
  "l1": struct{ X, Y float64 }{5, 10},
  "l2": {15, 30},
  "l3": {25, 50},
  "l4": {35, 70},
  "l5": {45, 90},
}

// Orden aleatorio

fmt.Println(x)
fmt.Println(x)

for k, v := range x {
  fmt.Println(k, v)
}

// Orden predecible gracias al patrón en las claves

for i := 1; i <= len(x); i++ {
  k := fmt.Sprintf("l%v", i)

  fmt.Println(k, x[k])
}
```
{{< /go-playground >}}

Es posible eliminar elementos de los mapas con la función `delete`, que recibe
como argumentos un mapa y la clave del elemento a ser eliminado.

{{< go-playground id="tN0s8GaicHo" >}}
```go
x := map[int]string{
  0: "cero",
  1: "uno",
  2: "dos",
  1<<30: "infinito",
}

delete(x, 1<<30)
```
{{< /go-playground >}}

### Representación sintáctica

```go
map[TIPO_CLAVE]TIPO_VALOR
```

### Ejemplos

```go
map[string]int{
  "Miguel": 6,
  "Angel": 5,
  "Rivera": 6,
  "Notararigo": 10,
}

map[string]struct{ X, Y float64 }{
  "Lugar 1": struct{ X, Y float64 }{5, 10},

  "Lugar 2": {15, 30}, // Se puede omitir el tipo de dato en los
                       // elementos
}
```

### Valor cero

```go
nil
```

## Estructuras

<https://tour.golang.org/moretypes/2>
<https://tour.golang.org/moretypes/3>
* <https://research.swtch.com/godata>

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

# Pruebas

## Ejemplos (pruebas)

{{< loi >}}
* <https://blog.golang.org/examples>
{{< /loi >}}

Además de texto, GoDoc da la posibilidad de mostrar el funcionamiento con
ejemplos dinámicos, que pueden ser ejecutados e incluso modificados en la
interfaz web. Para usar esta gran utilidad se deben crear funciones de ejemplo
en archivos `*_test.go`, estas funciones deberán tener como nombre `Example`
si se quiere mostrar algún ejemplo que use varios elementos del paquete, o
`ExampleIDENTIFICADOR`/`ExampleIDENTIFICADOR_MÉTODO` para tener como objetivo
solo un elemento.

`$GOPATH/src/local/arithmetic/arithmetic.go`:

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

{{< go-playground id="F89MWsdAyLS" >}}
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
{{< /go-playground >}}

Para ver los ejemplos se debe iniciar el servidor HTTP de GoDoc e ir a la ruta
<http://localhost:6060/pkg/local/arithmetic> con un navegador.

```shell-session
$ godoc -http :6060
```

Cada función de ejemplo deberá mostrar por la salida estándar los resultados,
y al final de cada una deberá existir un comentario especial `// Output: VALOR`
que indica los valores esperados, si se necesitan múltiples líneas, simplemente
se agregan como comentarios justo después del comentario especial. Estas
funciones son ejecutadas por `go test`, por lo que no solo tienen un uso
informativo, sino que también ayudan a probar el código; si no se encuentra
algún comentario especial, las funciones serán compiladas, pero no ejecutadas.

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

{{< go-playground id="cKBokfh3L9v" >}}
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
{{< /go-playground >}}

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

{{< go-playground id="8D3QO97NKE-" >}}
`$GOPATH/src/local/arithmetic/arithmetic.go`:

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
{{< /go-playground >}}

```shell-session
$ go test -v local/arithmetic
=== RUN   ExampleAdd
--- PASS: ExampleAdd (0.00s)
PASS
ok  	local/arithmetic
```

# Go tool

## Compilador

{{< loi >}}
* <https://golang.org/pkg/go/build/>
{{< /loi >}}

GOPATH

GOROOT

GOTPMDIR

## Condiciones de compilación

{{< loi >}}
* <https://golang.org/pkg/go/build/#hdr-Build_Constraints>

* <https://www.youtube.com/watch?v=COCUqAwAbD0&t=0s&index=31&list=PL5MnW0XCND7IjWv810mg4H81BxYN8BPQh>
{{< /loi >}}

Permiten establecer condiciones para el compilador, como usar el archivo para
ciertas arquitecturas o sistemas operativos, deben aparecer entre las primeras líneas, incluso antes de `package`. Para usarlas, solo hace falta un comentario
como este `// +build CONDICION [...]`

# Buenas prácticas

# Filosofía, proverbios y citas

{{< loi >}}
* <https://www.youtube.com/watch?v=PAAkCSZUG1c>
{{< /loi >}}

> Don't communicate by sharing memory, share memory by communicating.

<!-- -->

> Concurrency is not parallelism.

<!-- -->

> Channels orchestrate; mutexes serialize.

<!-- -->

> The bigger the interface, the weaker the abstraction.

<!-- -->

> Make the zero value usefull.

<!-- -->

> interface{} says nothing.

<!-- -->

> Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.

<!-- -->

> A little copying is better than a little dependency.

<!-- -->

> Syscall must always be guarded with build tags.

<!-- -->

> Cgo must always be guarded with build tags.

<!-- -->

> Cgo is not Go.

<!-- -->

> With the unsafe package there are no guarantees.

<!-- -->

> Clear is better than clever.

<!-- -->

> Reflection is never clear.

<!-- -->

> Errors are values.

<!-- -->

> Don't just check errors, handle them gracefully.

<!-- -->

> Design the architectura, name the components, document the details.

<!-- -->

> Documentation is for users.

# Atribuciones

**Go Team.** *Documentation* <https://golang.org/doc/>

**Ariel Mashraki.** *An overview of Go syntax and features.* <https://github.com/a8m/go-lang-cheat-sheet>

**Thomas Finley.** *Two's Complement.* <https://www.cs.cornell.edu/~tomf/notes/cps104/twoscomp.html>

**Autores de Wikipedia.** *List of Unicode characters.* <https://en.wikipedia.org/wiki/List_of_Unicode_characters>

