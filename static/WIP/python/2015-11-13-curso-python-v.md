---
title: "Curso de Python V: Estructuras de decisión"
description: "Es una de las partes más importantes de cualquier algoritmo, pues nos permite adaptarlo a las circunstancias"
category: "informatica"
tags: "python curso"
type: "article"
postFolder: "Post/2"
image: "Python.png"
---

A diferencia de otros lenguajes, Python solo tiene una estructura de decisión llamada `if`. {{ page.description }}; para esto, se usan datos lógicos y operadores de comparación. Veamos un ejemplo:

Desde que nos levantamos, estamos estudiando nuestro ambiente para determinar que es lo que debemos hacer, cuando intentamos decidir que comer antes de irnos a estudiar, trabajar o simplemente quedarnos a vaguear un rato en la casa :B creamos ciertas posibilidades que serán evaluadas según su capacidad de cumplirse, una arepa, hay tiempo para prepararla? si lo hay, la hacemos, pero si no buscamos algo más rápido.

	si tiempo mayor que 10
		hacer arepa

	si no
		buscar algo rápido

Si lo convertimos a código Python sería algo como:

{% highlight python %}
if tiempo > 10:
	comida = "Arepa"

else:
	comida = "Pan"
{% endhighlight %}

La sintaxis es bastante sencilla, pero faltan algunos detalles:

	if <condición>:
		<sentencias>

	[elif <condición>:
		# Se ejecuta si la condición anterior no se cumplió y se cumple la condición especificada en este bloque
		<sentencias>]

	[else:
		# Se ejecuta si no se cumple ninguna de las condiciones
		<sentencias>]

Las palabras reservadas `if` y `elif` poseen condiciones, que son expresiones evaluadas y su resultado se convierte en un valor lógico (si no lo es) para determinar si se ejecutarán las sentencias que están dentro del bloque. Su sintaxis es:

	<expresión>[ <operadorComparación> <otraExpresión>[ ...]]

Ejemplo:

	if tiempo > 10:  # Se evalúan las expresiones

	tiempo > 10  # El operador > retorna True si tiempo es mayor que 10

	if True:  # La condición después de ser evaluada

En caso de que la condición no obtenga un valor lógico, sería algo así:

	if 5 - 5:

	5 - 5

	if bool(0):  # Se convierte a un valor lógico

	if False:  # No se ejecutará lo que esté en el bloque

La palabra reservada `else` se encarga de ejecutar las sentencias en caso de que ninguna de las condiciones especificadas antes de él se cumplan. Como pueden ver, después de las tres instrucciones se agregan los `:` para especificar que comienza el bloque y desde ahí en adelante, todas las sentencias con una tabulación adelante formarán parte de él, aunque es posible definirlas en una línea si solo poseen una sentencia:

{% highlight python %}
if tiempo > 10: comida = "Arepa"

else: comida = "Pan"
{% endhighlight %}

La única obligatoria de las tres es `if`, entonces podrán haber bloques con solo `if`, pero ninguno solo con `elif` o `else`.

&nbsp;

##Operadores de comparación

- `==`:

	Retorna `True` si ambas expresiones son iguales.

		<expresión> == <otraExpresión>

		1 == 1  # True
		2 == 1  # False
		3 == 3.0  # True
		4 == "4"  # False
		5 == [5]  # False

		# Caso especial
		0 == False  # True
		1 == True  # True

- `!=`:

	Retorna `True` si ambas expresiones son diferentes.

		<expresión> != <otraExpresión>

		1 != 2  # True
		2 != 2  # False
		3 != 3.0  # False
		4 != "4"  # True
		5 != [5]  # True

		# Caso especial
		0 != False  # False
		1 != True  # False

- `<`:

	Retorna `True` si la primera expresión es menor que la segunda.

		<expresión> < <otraExpresión>

		1 < 2  # True
		2 < 2  # False
		3 < "3"  # Genera error, deben ser del mismo tipo
		"Hola" < "Mundo"  # True
		"51" < "6"  # True
		"A" < "a"  # True, las minúsculas son mayores que las mayúsculas
		[1] < ["A"]  # Genera error porque 1 y "A" no son del mismo tipo

	Las cadenas, tuplas y listas se ordenan lexicográficamente, es decir, básicamente se compara elemento a elemento y si alguno es menor finaliza la comprobación, por esto `"51" < "6"` retorna `True`. Como los conjuntos y diccionarios no son ordenables, no es posible usar este operador con ellos.

- `<=`:

	Retorna `True` si la primera expresión es menor o igual que la segunda.

		<expresión> <= <otraExpresión>

		1 <= 1  # True
		"Hola" <= "Hola"  # True

- `>`:

	Retorna `True` si la primera expresión es mayor que la segunda.

		<expresión> > <otraExpresión>

		1 > 0  # True
		2 > 2  # False
		"B" > "A"  # True
		"C" > "c"  # False

- `>=`:

	Retorna `True` si la primera expresión es mayor o igual que la segunda.

		<expresión> >= <otraExpresión>

		1 >= 1  # True
		"A" >= "A"  # True

- `is`:

	Python genera un identificador para cada valor que crea, este operador retorna `True` si ambas expresiones tienen el mismo identificador.

		<expresión> is <otraExpresión>

		1 is 1  # True
		"Abc" is "Abc"  # True
		"Hola" is 11  # False
		(1, 2) is (1, 2)  # False
		1 is [1, 2][0]  # True

	Los datos compuestos (menos las cadenas) siempre tendrán diferentes identificadores aunque sean iguales, pero sus elementos si los mantendrán.

- `is not`:

	Retorna `True` si ambas expresiones tienen diferentes identificadores.

		<expresión> is not <otraExpresión>

		1 is not "1"  # True
		(1, 2) is not (1, 2)  # True
		(1, 2)[0] is not (1, 2)[0]  # False

&nbsp;

##Operadores lógicos

- `not`:

	Invierte el valor lógico obtenido de una expresión.

		not <expresión>

		not True  # False
		not False  # True
		not 1  # False
		not 0  # True
		not "Hola"  # False
		not ""  # True

- `and`:

	Retorna `True` si el valor lógico de ambas expresiones es `True`.

		<expresión> and <otraExpresión>

		True and True  # True
		True and False  # False
		False and True  # False
		False and False  # False

- `or`:

	Retorna `True` si el valor lógico de al menos una expresión es `True`.

		<expresión> or <otraExpresión>

		True or True  # True
		True or False  # True
		False or True  # True
		False or False  # False

- `xor`:

	Retorna `True` solo si el valor lógico de una de las expresiones es `True`.

		<expresión> xor <expresión>

		True xor True  # False
		True xor False  # True
		False xor True  # True
		False xor False  # False

Con estos operadores pueden evaluarse varias condiciones en un `if` o `elif`:

{% highlight python %}
numero = input("Ingrese un número: ")

if numero > 5 and numero < 10:  # Se ejecuta solo si el número está entre 5 y 10
	pass

elif numero <= 5 or numero >= 10:  # Se ejecuta si el número no está entre 5 y 10
	pass
{% endhighlight %}

Por cierto, la palabra reservada `pass` nos deja crear bloques vacíos.

&nbsp;

##Operador ternario:

Es una sentencia que nos permite simplificar el uso de `if`, pero limita su funcionaldad a una sentencia. La sintaxis para usarlo es:

	<expresión> if <condición> else <otraExpresión>

Su aplicación dependerá de la imaginación del programador, pero voy a mostrarles algunas para que tengan una idea:

{% highlight python %}
# Asignación
opcion = input("Cuadrado o raiz? (c/r)")
resultado = 4 ** 2 if opcion == "c" else 4 ** (1/2)

# Es equivalente a:
opcion = input("Cuadrado o raiz? (c/r)")

if opcion == "c":
	resultado = 4 ** 2

else:
	resultado = 4 ** (1/2)


# Salidas adaptadas
cantidad = input("Indique la cantidad de personas: ")
print("Bienvenido") if cantidad == "1" else print("Bienvenidos")
print("Bienvenido" if cantidad == "1" else "Bienvenidos")  # Optimizado
print("Bienvenido" + ("" if cantidad == "1" else "s"))  # Más!

# Es equivalente a:
cantidad = input("Indique la cantidad de personas: ")

if cantidad == "1":
	print("Bienvenido")

else:
	print("Bienvenidos")
{% endhighlight %}

&nbsp;

##Ejercicios

Ya conocemos suficiente para empezar a crear scripts básicos y como la mejor forma de aprender a programar es practicando, [aquí]({{ site.gDriveFolder }}/{{ page.postFolder }}/Ejercicios.pdf) les dejo una lista de ejercicios.

&nbsp;

##Referencias

Raúl Gonzáles Duque. Python para todos

Eugenia Bahit. Curso: Python para Principiantes

Python Software Foundation. Overview - Python 2.7.10 documentation. [https://docs.python.org/2.7/](https://docs.python.org/2.7/)

Python Software Foundation. Overview - Python 3.5.0 documentation. [https://docs.python.org/3.5/](https://docs.python.org/3.5/)

creamostuweb [https://www.blogger.com/profile/05907170252728070066](https://www.blogger.com/profile/05907170252728070066). Programacion en python. [http://pycol.blogspot.com/](http://pycol.blogspot.com/)

Bartolomé Sintes Marco. Introducción a la programación con Python. [http://www.mclibre.org/consultar/python/index.html](http://www.mclibre.org/consultar/python/index.html)

Javier Montero. El Club del Autodidacta. [http://elclubdelautodidacta.es/wp/](http://elclubdelautodidacta.es/wp/)

Raúl González Duque. Documentanción en Python. [http://mundogeek.net/archivos/2008/07/07/documentacion-en-python/](http://mundogeek.net/archivos/2008/07/07/documentacion-en-python/)

stack exchange inc. Stack Overflow. [http://stackoverflow.com/](http://stackoverflow.com/)

&nbsp;

<div class="pager">
	<a href="{% post_url 2015-11-11-curso-python-iv %}" data-force="true"><button title="Curso de Python IV: Entrada y salida" class="button"><i class="fa fa-chevron-left fa-fw"></i></button></a>

	<a href="{% post_url 2015-11-18-curso-python-vi %}" data-force="true"><button title="Curso de Python VI: Estructuras de repetición" class="button"><i class="fa fa-chevron-right fa-fw"></i></button></a>
</div>