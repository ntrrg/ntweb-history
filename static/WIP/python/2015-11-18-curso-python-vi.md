---
title: "Curso de Python VI: Estructuras de repetición"
description: "Además de condicionar el flujo de nuestros algoritmos, también podemos agregarles bucles (bloques que se repiten) que le darán un poco más de complejidad y capacidad."
category: "informatica"
tags: "python curso"
type: "article"
postFolder: "Post/2"
image: "Python.png"
---

{{ page.description }} Python dentro de su sintaxis tiene dos estructuras que nos permitirán definirlos:

&nbsp;

###Estructura `while`:

Ejecutará las sentencias dentro del bloque mientras se cumpla la condición (o condiciones) con la que se definió, su sintaxis es:

	while <condición>:
		<sentencias>

	[else:
		<sentencias>]

Intentaré explicarles esto con un ejemplo:

{% highlight python %}
# Script que saluda hasta que le digamos que salga

salir = ""  # Hace falta definir las variables antes de compararlas

while salir != "s":  # Mientras salir no valga "s" (Sí)
	print("Hola")

	salir = input("Salir? (s): ")  # Actualización de la variable de la condición
{% endhighlight %}

Resultado:

	Hola
	Salir? (s): n
	Hola
	Salir? (s): n
	Hola
	Salir? (s): s

Vemos algo conocido, los `:` y la tabulación, son parte fundamental de Python y de aquí en adelante se usarán para todos los bloques de código. Lo nuevo sería la actualización de la variable de la condición, que es realmente importante ya que sin esto, nuestro bucle nunca dejaría de ejecutarse y probablemente el uso del CPU aumente drásticamente si no hay alguna instrucción que frene al script (como `input`); en caso de no usar esta técnica, Python ofrece dos palabras reservadas para manejar el flujo de los bucles:

- `continue`: Evita que el código después de ella se ejecute y salte a la otra repetición si la condición se cumple. Ejemplo:

{% highlight python %}
repeticion = 0

while repeticion < 10:  # 10 repeticiones, del 0 al 9
	repeticion += 1
	numero = int(input("Ingrese un número: "))

	if numero % 2 == 1:  # Número impar
		continue

	print("Es un número par!")
{% endhighlight %}

Resultado:

	Ingrese un número: 1
	Ingrese un número: 2
	Es un número par!
	Ingrese un número: 3
	Ingrese un número: 4
	Es un número par!
	Ingrese un número: 5
	Ingrese un número: 6
	Es un número par!
	Ingrese un número: 7
	Ingrese un número: 8
	Es un número par!
	Ingrese un número: 9
	Ingrese un número: 10
	Es un número par!

No hace falta que creemos un bloque `else` para la salida, pues en caso de que se cumpla la condición, se saltará el resto del código.

- `break`: Finaliza el bucle. Ejemplo:

{% highlight python %}
# Los números del 1 al 10

numero = 0

while True:  # Esto crea un bucle infinito
	numero += 1

	print(numero)

	if numero == 10:
		break
{% endhighlight %}

Resultado:

	1
	2
	3
	4
	5
	6
	7
	8
	9
	10

&nbsp;

Esta estructura también tiene un bloque `else` que se ejecuta cuando la condición es o se convierte en `False`, si el bucle es finalizado con `break` no se ejecutará. Ejemplos:

- Cuando se convierte en `False`:

{% highlight python %}
# Los números del 10 al 1

numero = 10

while numero > 0:
	print(numero)

	numero -= 1

else:  # Se ejecuta porque la condición se convierte en False
	print("Salida del bloque else")
{% endhighlight %}

Resultado:

	10
	9
	8
	7
	6
	5
	4
	3
	2
	1
	Salida del bloque else

- Cuando es `False`:

{% highlight python %}
# Los números del 10 al 1

numero = 0  # Inicialización en 0

while numero > 0:
	print(numero)

	numero -= 1

else:  # Se ejecuta porque la condición nunca fue True
	print("Salida del bloque else")
{% endhighlight %}

Resultado:

	Salida del bloque else

- Cuando se rompe el bucle:

{% highlight python %}
# Los números del 10 al 5

numero = 10

while numero > 0:
	print(numero)

	if numero == 5:
		break

	numero -= 1

else:  # No se ejecuta porque se rompe el bucle
	print("Salida del bloque else")
{% endhighlight %}

Resultado:

	10
	9
	8
	7
	6
	5

&nbsp;

###Estructura `for`:

Su principal función es recorrer elementos iterables (cadenas, tuplas, listas, etc..), en cada repetición accederá a un elemento y nos dejará tratarlo como si fuese una variable totalmente diferente. Para usarlo tenemos que respetar la sintaxis:

	for <identificador> in <iterable>:
		<sentencias>

	[else:
		<sentencias>]

Les mostraré ejemplos de como trabajar con el y de algunos usos alternativos:

- Recorrer iterables:

{% highlight python %}
tupla = (1, "Hola", 2, "Mundo")

for variable in tupla:
	print(variable)
{% endhighlight %}

Resultado:

	1
	Hola
	2
	Mundo

Es posible crear rangos numéricos con la ayuda de la clase `range`, como estos rangos son iterables, podremos darle el siguiente uso al `for` gracias a ellos:

{% highlight python %}
for numero in range(5):
	print(numero)
{% endhighlight %}

Resultado:

	0
	1
	2
	3
	4

- Recorrer diccionarios:

**Método 1. Claves:**

Los diccionarios tienen un comportamiento especial, `<identificador>` contendrá la clave y no el valor. No importará el orden en que fueron definidos los elementos del diccionario, por lo que probablemente su salida en cada ejecución será diferente.

{% highlight python %}
diccionario = {"Clave 1": 1, "Clave 2": 2}

for variable in diccionario:
	print("Valor en", variable + ":", diccionario[variable])
{% endhighlight %}

Resultado:

	Valor en Clave 2: 2
	Valor en Clave 1: 1

&nbsp;

**Método 2. Claves y valores:**

Por medio del método `.items()` de los diccionarios podemos generar una lista compuesta de tuplas que contienen cada clave y su valor, suena enredado pero les mostraré paso a paso el proceso y verán que es sencillo.

{% highlight python %}
diccionario = {"Clave 1": 1, "Clave 2": 2, "Clave 3": 3}

print(diccionario.items())
{% endhighlight %}

Resultado:

	[("Clave 1", 1), ("Clave 3", 3), ("Clave 2", 2)]  # Una lista con tuplas

Intentemos recorrer la lista con `for`:

{% highlight python %}
diccionario = {"Clave 1": 1, "Clave 2": 2, "Clave 3": 3}

for elemento in diccionario.items():
	print(elemento)
{% endhighlight %}

Resultado:

	("Clave 1", 1)
	("Clave 3", 3)
	("Clave 2", 2)

Como obtenemos una lista de tuplas simétricas (todas tienen la misma cantidad de elementos), podemos utilizar la asignación múltiple que consiste en asignar a varios identificadores los valores de un elemento compuesto. Ejemplo:

	a, b = (1, 2)  # a = 1 y b = 2

Para hacer esto, solo tenemos que agregar un identificador más al `for` y el resultado final del algoritmo sería algo como:

{% highlight python %}
diccionario = {"Clave 1": 1, "Clave 2": 2, "Clave 3": 3}

for clave, valor in diccionario.items():
	print("Valor en", clave + ":", valor)
{% endhighlight %}

Resultado:

	Valor en Clave 1: 1
	Valor en Clave 3: 3
	Valor en Clave 2: 2

&nbsp;

**Método 3. Claves serializadas:**

Se pueden definir claves estratégicamente para agrupar datos, por ejemplo una lista de personas con sus datos:

{% highlight python %}
personas = {"nombre0": "Miguel", "nombre1": "Angel", "apellido0": "Rivera", "apellido1": "Notararigo"}

print("Personas:")

for variable in range(2):
	print(personas["nombre" + str(variable)], personas["apellido" + str(variable)])
{% endhighlight %}

Resultado:

	Personas:
	Miguel Rivera
	Angel Notararigo

&nbsp;

**Método 4. Ordenamiento de claves:**

Al convertir un diccionario en una lista, se obtienen todas sus claves. Ejemplo:

{% highlight python %}
diccionario = {"Cotufas": "La chuchería de los dioses", "Algoritmo": "Un charquero de instrucciones", "Programar": "Uno de los placeres de la vida"}

claves = list(diccionario)
print(claves)
{% endhighlight %}

Resultado:

	['Algoritmo', 'Programar', 'Cotufas']

Con las claves dentro de una lista, podemos usar un método llamado `.sort()` que permite ordenar sus elementos:

{% highlight python %}
claves.sort()
print(claves)
{% endhighlight %}

	['Algoritmo', 'Cotufas', 'Programar']

Y con esto ya podemos obtener los valores del diccionario ordenado por sus claves, en este ejemplo les muestro como quedaría el algoritmo usando este método:

{% highlight python %}
diccionario = {"Cotufas": "La chuchería de los dioses", "Algoritmo": "Un charquero de instrucciones", "Programar": "Uno de los placeres de la vida"}

claves = list(diccionario)  # Obtener las claves
claves.sort()  # Ordenarlas

print("Diccionario:")

for palabra in claves:
	print(palabra + ":", diccionario[palabra])
{% endhighlight %}

Resultado:

	Diccionario:
	Algoritmo: Un charquero de instrucciones
	Cotufas: La chuchería de los dioses
	Programar: Uno de los placeres de la vida

&nbsp;

Aunque les haya mostrado solo cuatro métodos, no significa que existan solo esos porque realmente los últimos dos los descubrí practicando. Las aplicaciones dependerán de su imaginación, por lo que no pueden dejar que alguien limite sus capacidades.

&nbsp;

El bloque `else` se ejecutará solo si el iterable usado se recorre completamente, ejemplos:

- Se recorre el iterable:

{% highlight python %}
for numero in range(5):
	print(numero)

else:  # Se ejecuta porque recorre el bloque completo
	print("Salida del bloque else")
{% endhighlight %}

Resultado:

	0
	1
	2
	3
	4
	Salida del bloque else

- Se rompe el bucle:

{% highlight python %}
for numero in range(5):
	print(numero)

	if numero == 3:
		break

else:  # No se ejecuta porque el bucle se rompe
	print("Salida del bloque else")
{% endhighlight %}

Resultado:

	0
	1
	2
	3

&nbsp;

##Ejercicios

[Aquí]({{ site.gDriveFolder }}/{{ page.postFolder }}/Ejercicios-ii.pdf) les dejo otra pequeña lista de ejercicios para practicar las estructuras de repetición.

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
	<a href="{% post_url 2015-11-13-curso-python-v %}" data-force="true"><button title="Curso de Python V: Estructuras de decisión" class="button"><i class="fa fa-chevron-left fa-fw"></i></button></a>

	<a href="{% post_url 2015-12-02-curso-python-vii %}" data-force="true"><button title="Curso de Python VII: Funciones" class="button"><i class="fa fa-chevron-right fa-fw"></i></button></a>
</div>