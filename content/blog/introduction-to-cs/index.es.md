---
title: Introducción a la Informática
author: ntrrg
date: 2020-03-03T09:00:00-04:00
description: Fundamentos de la informática. Orientado a personas sin conocimiento previo, pero también puede ser interesante para los que ya están familiarizados con la tecnología.
tags:
  - tecnología
  - aprendizaje
  - fundamentos
  - programación
comments: true
toc: true
draft: true
---

El término informática fue usado académicamente por primera vez en el libro
*«Informatik: Automatische Informationsverarbeitung»* (Informática:
procesamiento automático de información) por Karl Steinbuch en 1957. Proviene
de la unión de las palabras alemanas *Informationen* (información) y
*Automatik* (automática).

Consiste en estudiar los métodos para el almacenamiento, procesamiento y
transmisión de la información, con el fin de sistematizar y/o automatizar
tareas de manera óptima.

No se debe confundir sistematización con automatización. La sistematización se
refiere a la organización de los procesos, que en todo caso podrían realizarse
manualmente. Por otro lado, la automatización se trata de que los procesos sean
realizados por máquinas o cualquier otro organismo sin voluntad (autómata).

Teniendo claro estos conceptos, puede decirse que la informática se ha ido
aplicando desde la antigüedad, solo que sin la participación de las máquinas
que hoy se conocen como computadoras. Por ejemplo:

* En la antigua Grecia, aproximadamente unos 500 años antes de Cristo, se usaba
  un herramienta llamada escítala para cifrar y descifrar mensajes.

* En varios lugares del mundo entre los años 500a.C y 300a.C, se usaba un
  instrumento llamado ábaco para realizar cálculos matemáticos.

* En la antigua Grecia, cerca del año 300a.C el matemático Euclides desarrolló
  un método para el cálculo del máximo común divisor (El Algoritmo de
  Euclides).

Todas estas aplicaciones tienen algo en común, deben seguirse metódicamente
unos pasos para obtener el resultado esperado.

En la actualidad, los autómatas son máquinas, debido a esto todo el proceso de
automatización está orientado a la representación de la información como
electricidad, que es su medio de comunicación. Esto puede cambiar en el futuro
si se crean otro tipo de autómatas, como computadoras biológicas, pero no sería
un problema para la informática pues aunque está bastante relacionada a las
máquinas, lo que cambiaría sería la forma en la que se representa la
información.

# Algoritmos (Software)

Son un conjunto de instrucciones para realizar una tarea. Estas deben ser
suficientemente explícitas para evitar ambigüedad y deben manejar casos
inesperados.

Aunque el público de un algoritmo puede ser cualquier entidad, tanto seres
vivos como autómatas, por lo general serán autómatas quienes lo lleven acabo,
pues estos pueden realizar dichas tareas con mayor rapidez.

Expresar algoritmos en lenguaje natural sin ambigüedad es una tarea compleja,
pues nuestro modo de comunicación es muy amplio y los autómatas todavía no son
capaces de entenderlo completamente, pero a su vez permite que sea mucho más
fácil explicar su funcionamiento otros humanos.

Por esta razón, es común que un algoritmo sea expresado en múltiples lenguajes
que estarán orientados a un público específico y tendrán diferentes niveles
descriptivos. Los niveles descriptivos de un algoritmo no están estrictamente
estandarizados, pero podrían agruparse en:

* Alto nivel: cuando se quiere hablar sobre qué hace y para qué es útil un
  algoritmo, el lenguaje natural es la mejor forma de hacerlo gracias a su alto
  nivel descriptivo y a que puede ser entendido por personas sin conocimientos
  técnicos.

```
Algoritmo para calcular el factorial de un número:

Para obtener el factorial de un número positivo n, se deben
multiplicar todos los números enteros desde 1 hasta n (inclusivo).
Por ejemplo:

  4! = 4 x 3 x 2 x 1 = 24
```

* Formal: cuando se quiere explicar su funcionamiento, el pseudocódigo es la
  mejor herramienta para hacerlo gracias a que usa términos técnicos
  estandarizados, pero aún así puede ser entendido por personas sin mucho
  conocimiento previo. Es ideal para la academia y el Software Libre.

```
```

* Implementación: cuando se quiere que un autómata ejecute el proceso, los
  lenguajes de programación permiten traducir algoritmos a lenguaje máquina.
  Existe una gran cantidad de lenguajes de programación, que al igual que el
  lenguaje natural, cada uno tiene una sintaxis específica.

{{< details summary="Go:" open=true >}}
```go
// Fact returns the factorial of n.
func Fact(n int) int {
  r := 1

  for ; n > 1; n-- {
    r *= n
  }

  return r
}
```
{{< /details >}}

{{< details summary="C:" >}}
```c
// Returns the factorial of n.
int fact(int n) {
  int r = 1;

  while (n > 1)
    r *= n--;

  return r;
}
```
{{< /details >}}

{{< details summary="JavaScript:" >}}
```js
/**
 * Calculates the factorial of the given number.
 *
 * @param {Number} n - Initial number.
 * @return {Number}
 */
function fact(n) {
  let r = 1

  while (n > 1)
    r *= n--

  return r
}
```
{{< /details >}}

Los algoritmos son escritos de manera generalizada, es decir, sus instrucciones
deben ser seguidas sin importar quién o qué los ejecute. Esto es beneficioso
porque asegura su comportamiento, pero también puede ser una desventaja ya que
asume que quienes lo ejecuten tienen las mismas características, por lo que
en algunas ocasiones, la implementación de un algoritmo puede variar según el
autómata que lo siga e incluso el lenguaje de programación que se use.

# Máquinas (Hardware)

* Relojes
* Calculadoras
* Electrodomésticos
* Vehículos
* Computadoras personales
* Teléfonos
* Tabletas
* Computadoras de una tarjeta
* Microcontroladores

## Representación de los datos

Números binarios

* Números
* Texto
* Colores
* Imágenes
* Sonidos
* Videos
* Sensores

# Atribuciones

**HarvardX.** *CS50's Introduction to Computer Science.* <https://courses.edx.org/courses/course-v1:HarvardX+CS50+X/course/>

**Khan Acedemy.** *Computer science.* <https://www.khanacademy.org/computing/computer-science>

**Khan Acedemy.** *AP®︎ Computer Science Principles.* <https://www.khanacademy.org/computing/ap-computer-science-principles/>

**Wikipedia.** *Informática.* <https://es.wikipedia.org/wiki/Inform%C3%A1tica>

**Wikipedia.** *Algorithm.* <https://en.wikipedia.org/wiki/Algorithm>

