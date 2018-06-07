---
title: "Curso de Python VII: Funciones"
description: "Las funciones son bloques de código que podemos utilizar para evitar repetir sentencias y mejorar la legibilidad de un script."
category: "informatica"
tags: "python curso"
type: "article"
postFolder: "Post/2"
image: "Python.png"
---

{{ page.description }} La sintaxis básica para definirlas es:

	def <identificador>():
		<sentencias>

Y para ejecutar las instrucciones dentro de ellas solo tenemos que invocarlas (suena a magia negra, pero es el término técnico),  que simplemente es usar su identificador con paréntesis `<identificador>()` (nada de sacrificios y esas cosas :B). Por cierto, es necesario definir una función antes de poder invocarla, así que lo mejor es que lo hagamos al principio del script para evitar problemas si llegamos a modificarlo. Estudiemos la ecuación cuadrática para entender mejor todo esto:

&nbsp;

<img src="{{ site.gDriveFolder }}/{{ page.postFolder }}/cuadratica.png" style="display: block; margin: 0 auto;" />

{% highlight python %}
# Valores necesarios
a = int(input("Indique el valor del coeficiente cuadrático (x^2): "))
b = int(input("Indique el valor del coeficiente lineal (x): "))
c = int(input("Indique el valor del término independiente: "))

# Proceso
n1 = b * -1
n2 = (b ** 2 - 4 * a * c) ** 0.5
n3 = 2 * a

# Resultado
x1 = (n1 + n2) / n3
x2 = (n1 - n2) / n3

print((x1, x2))
{% endhighlight %}

Resultado:

	# x^2 + 2x - 8 = 0
	Indique el valor del coeficiente cuadrático (x^2): 1
	Indique el valor del coeficiente lineal (x): 2
	Indique el valor del término independiente: -8
	(2.0, -4.0)  # (x - 2)(x + 4) = 0

No se ve tan complicado, pero imaginen que necesitemos resolver varias ecuaciones cuadráticas dentro de un script, el código sería algo así:

{% highlight python %}
# x^2 + 6x - 16 = 0
a = 1
b = 6
c = -16

# Proceso
n1 = b * -1
n2 = (b ** 2 - 4 * a * c) ** 0.5
n3 = 2 * a

# Resultado
x1 = (n1 + n2) / n3
x2 = (n1 - n2) / n3

print((x1, x2))

# Segunda ecuación

# x^2 + 8x + 18 = 0
a = 1
b = 8
c = 16

# Proceso
n1 = b * -1
n2 = (b ** 2 - 4 * a * c) ** 0.5
n3 = 2 * a

# Resultado
x1 = (n1 + n2) / n3
x2 = (n1 - n2) / n3

print((x1, x2))
{% endhighlight %}

Resultado:

	(2.0, -8.0)  # (x - 2)(x + 16) = 0
	(-4.0, -4.0)  # (x + 4)(x + 4) = 0

Tenemos dos problemas, el primero es que hemos perdido legibilidad porque las líneas repetidas hacen que el script se extienda; el segundo es que si necesitamos arreglar algunos detalles del proceso de la ecuación cuadrática, tendremos que hacer lo mismo en varios lugares. Para resolver ambos problemas, podemos utilizar funciones:

{% highlight python %}
def cuadratica():  # Definición
	# Proceso
	n1 = b * -1
	n2 = (b ** 2 - 4 * a * c) ** 0.5
	n3 = 2 * a

	# Resultado
	x1 = (n1 + n2) / n3
	x2 = (n1 - n2) / n3

	print((x1, x2))

# x^2 + 6x - 16 = 0
a = 1
b = 6
c = -16

cuadratica()  # Invocación

# Segunda ecuación

# x^2 + 8x + 18 = 0
a = 1
b = 8
c = 16

cuadratica()
{% endhighlight %}

Resultado (igual que el anterior):

	(2.0, -8.0)  # (x - 2)(x + 16) = 0
	(-4.0, -4.0)  # (x + 4)(x + 4) = 0

&nbsp;

Aunque arreglamos todo, en el proceso rompí costumbres muy importantes para evitar complicar el ejemplo; iré mejorándolo poco a poco para que entiendan con facilidad todas las partes de una función. Entonces, compliquemos la sintaxis:

	def <identificador>():
		<sentencias>

		[return <expresión>]  # Nuevo

El valor retornado por una función normalmente es el resultado de su ejecución, puede usarse en cualquier lugar que se espere una expresión. En el ejemplo, usamos `print()` para mostrar el resultado obtenido de la ecuación, pero esto lo limita a escribirse por la salida estándar evitando que pueda usarse como una expresión, si lo arreglamos debería verse como:

{% highlight python %}
def cuadratica():
	# Proceso
	n1 = b * -1
	n2 = (b ** 2 - 4 * a * c) ** 0.5
	n3 = 2 * a

	# Resultado
	x1 = (n1 + n2) / n3
	x2 = (n1 - n2) / n3

	return x1, x2  # Cambia print() por return

# x^2 + 6x - 16 = 0
a = 1
b = 6
c = -16

print(cuadratica())

# Segunda ecuación

# x^2 + 8x + 18 = 0
a = 1
b = 8
c = 16

e2 = cuadratica()

print(e2)
{% endhighlight %}

&nbsp;

Más sintaxis:

	def <identificador>([<parámetro>[, ...]]):  # Nuevo
		<sentencias>

		[return <expresión>]

	<identificador>([<valor>[, ...]])  # Invocación

En programación las variables tienen un ámbito, que determinará si sus valores podrán ser obtenidos en ciertos lugares del algoritmo. En el caso de Python, las que se declaren en el cuerpo del script serán de ámbito global y por ello, estarán disponibles en cualquier lugar; los parámetros y las variables que se definan en una función serán de ámbito local, por lo que solo estarán disponibles donde se declararon.

Los parámetros son identificadores comunes que recibirán valores externos, pueden usarse para alterar el comportamiento o el resultado de la función además de darle independencia. En los ejemplos anteriores utilizábamos variables globales, que realmente no representa un error pero si deja una brecha para que se generen algunos como:

- Si alguna de las variables globales `a`, `b` o `c` no existieran, la función no lograría ejecutarse.

- Intentar redefinir una variable global dentro de una función después de acceder a su valor, generará un error porque Python la definirá como una variable local en todo el bloque y como ya hemos visto, no puede usarse una variable sin antes ser definida.

		def cuadratica():
			n1 = b * -1

			b = 2  # Se define una "b" local, por lo que la sentencia anterior intentará leer una "b" local que no ha sido definida y no la global como en el ejemplo

Para evitarnos estos inconvenientes, lo mejor sería agregarle parámetros a `cuadratica()` e intentar no usar variables globales dentro de ella; después de arreglar el algoritmo debería verse así:

{% highlight python %}
def cuadratica(a, b, c):
	# Proceso
	n1 = b * -1
	n2 = (b ** 2 - 4 * a * c) ** 0.5
	n3 = 2 * a

	# Resultado
	x1 = (n1 + n2) / n3
	x2 = (n1 - n2) / n3

	return x1, x2  # Cambia print() por return

# x^2 + 6x - 16 = 0
cuadratico = 1
lineal = 6
independiente = -16

print(cuadratica(cuadratico, lineal, independiente))

# Segunda ecuación

# x^2 + 8x + 18 = 0
e2 = cuadratica(1, 4 + 4, 16)

print(e2)
{% endhighlight %}

Definida de esta manera, la función solo necesitará recibir tres expresiones durante la invocación y sus parámetros tomarán los valores en el orden que fueron declarados; se puede alterar el orden referenciando los identificadores, por ejemplo:

	cuadratica(b = 4 + 4, c = 16, a = 1)

Aunque se debe tener cuidado con esto, las referencias no permiten saltarse parámetros, por ejemplo `cuadratica(1, 16, b = 4 + 4)` no nos dejará asignar `4 + 4` a `b` y `16` a `c`, sino que `b` recibirá ambos valores y se generará un error.

También es posible prestablecer los parámetros durante la definición de las funciones (al final), es útil en los casos que no se indiquen suficientes expresiones o que alguno de los parámetros reciba casi siempre el mismo valor. Si en la mayoría de las ecuaciones que usemos, el coeficiente cuadrático vale `1` sería buena idea definirla de la siguiente manera:

{% highlight python %}
def cuadratica(b, c, a = 1):
	# Proceso
	n1 = b * -1
	n2 = (b ** 2 - 4 * a * c) ** 0.5
	n3 = 2 * a

	# Resultado
	x1 = (n1 + n2) / n3
	x2 = (n1 - n2) / n3

	return x1, x2

# x^2 + 6x - 16 = 0
lineal = 6
independiente = -16

print(cuadratica(lineal, independiente))

# Segunda ecuación

# x^2 + 8x + 18 = 0
e2 = cuadratica(4 + 4, 16)

print(e2)

# Tercera ecuación

# 2x^2 - 5x - 3 = 0
print(cuadratica(-5, -3, 2))
{% endhighlight %}

Resultado:

	(2.0, -8.0)  # (x - 2)(x + 16) = 0
	(-4.0, -4.0)  # (x + 4)(x + 4) = 0
	(3.0, -0.5)  # (x - 3)(2x + 1) = 0

&nbsp;

Hasta aquí podemos decir que hemos creado una función como mandan los dioses binarios, después de la reconstrucción se ve así:

{% highlight python %}
def cuadratica(b, c, a = 1):
	# Proceso
	n1 = b * -1
	n2 = (b ** 2 - 4 * a * c) ** 0.5
	n3 = 2 * a

	# Resultado
	x1 = (n1 + n2) / n3
	x2 = (n1 - n2) / n3

	return x1, x2
{% endhighlight %}

Y la sintaxis que respeta es:

	def <identificador>([<parámetro>[, ...]][[,]<parámetro> = <valor>[, ...]]):
		<sentencias>

		[return <expresión>]

	<identificador>([<valor>[, ...]][[, ]<parámetro> = <valor>[, ...]])

&nbsp;

Pero todavía nos falta un último detalle, el empaquetamiento de parámetros: consiste en utilizar listas y diccionarios para guardar expresiones extras que se usen en la invocación; creemos una función para obtener el promedio de algunos números números:

{% highlight python %}
def promedio(n1, n2, n3):
	return (n1 + n2 + n3) / 3

print(promedio(1, 2, 3))
{% endhighlight %}

Resultado:

	2.0

Con esto logramos obtener el promedio de tres números, pero si quisiéramos usar la función para una cantidad indeterminada de números? no sería posible con parámetros comunes pues tenemos que ser muy precisos cuando los definimos, para eso existe el empaquetamiento:

{% highlight python %}
def promedio(*numeros):
	suma = 0
	cantidad = 0

	for numero in numeros:
		suma += numero
		cantidad += 1

	return suma / cantidad

print(promedio(1, 2, 3))
print(promedio(1, 2, 3, 4, 5))
{% endhighlight %}

Resultado:

	2.0
	3.0

Ahora la función no tendrá que invocarse con un número fijo de expresiones, sino que guardará todas las que se indiquen en `numeros` como una lista gracias al operador `*`; pero aunque resolvimos un problema, ahora tenemos otro: es posible llamar a la función sin parámetros pues `*numeros` puede estar vacía, pero como todo tiene solución podemos forzar la lectura de al menos un valor usando parámetros tradicionales:

{% highlight python %}
def promedio(suma, *numeros):
	cantidad = 1

	for numero in numeros:
		suma += numero
		cantidad += 1

	return suma / cantidad

print(promedio(1, 2, 3))
print(promedio(1, 2, 3, 4, 5))
{% endhighlight %}

Durante la invocación se puede hacer algo parecido, solo que tiene un funcionamiento inverso (desempaquetamiento) y para usarlo también se necesita el operador `*`:

{% highlight python %}
# print(promedio(1, 2, 3))

tupla = (1, 2, 3)

print(promedio(*tupla))

# print(promedio(1, 2, 3, 4, 5))

lista = [1, 2, 3, 4, 5]

print(promedio(*lista))
{% endhighlight %}

&nbsp;

En el caso del desempaquetamiento de diccionarios, sus claves harán referencia a los nombres de los parámetros y se debe usar el operador `**`, con `cuadratica()` podríamos invocarla de la siguiente manera:

{% highlight python %}
# 2x^2 - 5x - 3 = 0
diccionario = {"a": 2, "b": -5, "c": -3}

print(cuadratica(**diccionario))
{% endhighlight %}

Y al momento de empaquetarlos, guardarán todas las expresiones extras que estén acompañadas por identificadores, por ejemplo:

{% highlight python %}
def definiciones(**diccionario):
	claves = list(diccionario)  # Obtener las claves
	claves.sort()  # Ordenarlas

	print("Diccionario:")

	for palabra in claves:
		print(palabra + ":", diccionario[palabra])
	

definiciones(Cotufas = "La chuchería de los dioses", Algoritmo = "Un charquero de instrucciones", Programar = "Uno de los placeres de la vida")
{% endhighlight %}

Resultado:

	Diccionario:
	Algoritmo: Un charquero de instrucciones
	Cotufas: La chuchería de los dioses
	Programar: Uno de los placeres de la vida

&nbsp;

Y ahora si al fin, la sintaxis completa para las funciones sería:

	def <identificador>([<parámetro>[, ...]][[, ]*<lista>][[, ]**<diccionario>][[, ]<parámetro> = <valor>[, ...]]):
		<sentencias>

		[return <expresión>]

	<identificador>([<valor>[, ...]][[, ]*<lista>][<parámetro> = <valor>[, ...]][[, ]**<diccionario>])

&nbsp;

Por cierto, si se usan variables para asignar valores a los parámetros y en el cuerpo de la función se modifican, estos valores se mantendrán solo si las variables usadas son mutables (listas, diccionarios y conjuntos). Ejemplos:

- Variables inmutables:

{% highlight python %}
def cambiar(x):
	x += " :B"

cadena = "Ntrrg"

cambiar(cadena)

print(cadena)
{% endhighlight %}

Resultado:

	Ntrrg

- Variables mutables:

{% highlight python %}
def cambiar(x):
	x += [3]

lista = [1, 2]

cambiar(lista)

print(lista)
{% endhighlight %}

Resultado:

	[1, 2, 3]

- Redefinición: Aunque la variable sea mutable el valor no se verá afectado.

{% highlight python %}
def cambiar(x):
	x = [1, 2, 3]

lista = [1, 2]

cambiar(lista)

print(lista)
{% endhighlight %}

Resultado:

	[1, 2]

&nbsp;

{% highlight python %}
def cambiar(x):
	x = x + [3]

lista = [1, 2]

cambiar(lista)

print(lista)
{% endhighlight %}

Resultado:

	[1, 2]

Esto puede parecer contradictorio porque les había dicho que escribir `x = x + [3]` era equivalente a `x += [3]` y realmente es así, solo que cuando se trabaja con datos mutables, la asignación aumentada agrega `[3]` al final del espacio de memoria donde está `x` y no busca uno nuevo (como lo hace con los inmutables).

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
	<a href="{% post_url 2015-11-18-curso-python-vi %}" data-force="true"><button title="Curso de Python VI: Estructuras de repetición" class="button"><i class="fa fa-chevron-left fa-fw"></i></button></a>

	<a href="{% post_url 2015-12-06-curso-python-viii %}" data-force="true"><button title="Curso de Python VIII: POO" class="button"><i class="fa fa-chevron-right fa-fw"></i></button></a>
</div>