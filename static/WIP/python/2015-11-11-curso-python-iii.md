---
title: "Curso de Python III: Operadores"
description: "Al igual que en las matemáticas, los lenguajes de programación usan operadores para simbolizar ciertas actividades. Si se acuerdan de sumar y restar, es probable que esto sea rápido."
category: "informatica"
tags: "python curso"
type: "article"
postFolder: "Post/2"
image: "Python.png"
---

Python es un lenguaje de tipado fuerte, es decir, si intentamos usar datos de un tipo donde se espera otro, nos insultará :B pero tranquilos, aquí les mostraré cuales son los que espera Python que le demos. {{ page.description }}

&nbsp;

###Aritméticos:

- `+`:

	Suma, concatenación (cadenas).

{% highlight python %}
# Numéricos
print(2 + 5)

# Lógicos
print(True + True)
print(True + False)
print(False + False)

# Cadenas
print("Konichiwa" + "Sekai")  # Concatenación

# Tuplas
print((1, 2, 3) + (4, 5, 6))

# Listas
print([1, 2, 3] + [4, 5, 6])
{% endhighlight %}

Resultado:

	7
	2
	1
	0
	KonichiwaSekai
	(1, 2, 3, 4, 5, 6)
	[1, 2, 3, 4, 5, 6]

- `-`:

	Resta, diferencia (cojuntos).

{% highlight python %}
# Numéricos
print(3 - 1)

# Lógicos
print(True - True)
print(True - False)
print(False - True)
print(False - False)

# Conjuntos
print({1, 2, 3} - {1, 2})
{% endhighlight %}

Resultado:

	2
	0
	1
	-1
	0
	{3}

- `*`:

	Multiplicación.

{% highlight python %}
# Numéricos
print(5 * 2)

# Lógicos
print(True * True)
print(True * False)
print(False * False)

# Cadenas
print("Ntrrg" * 3)  # Solo con números

# Tuplas
print((1, 2) * 3)  # Solo con números

# Listas
print([3, 4] * 3)  # Solo con números
{% endhighlight %}

Resultado:

	10
	1
	0
	0
	NtrrgNtrrgNtrrg
	(1, 2, 1, 2, 1, 2)
	[1, 2, 1, 2, 1, 2]

- `/`:

	División.

{% highlight python %}
# Numéricos
print(12 / 3)
print(5 / 2)
print(5 / 2.0)  # Necesario para obtener el resultado decimal en Python 2
{% endhighlight %}

Resultado:

	4.0  # En Python 3 todas las divisiones retornan un decimal
	2.5
	2.5

- `%`:

	Módulo. Permite obtener el resto de una división.

{% highlight python %}
# Numéricos
print(12 % 3)
print(5 % 2)
{% endhighlight %}

Resultado:

	0
	1

- `//`:

	División entera.

{% highlight python %}
# Numéricos
print(12 // 3)
print(5 // 2)
print(5.3 // 2)
{% endhighlight %}

Resultado:

	6
	2  # No redondea
	2.0  # Si se usa algún decimal, retorna un decimal con ".0"

- `**`:

	Exponente, radical.

{% highlight python %}
# Numéricos
print(4 ** 2)
print(27 ** (1 / 3))  # Radical, retornan decimales
print(27 ** 0.3333333333333333)  # Equivalente al anterior
{% endhighlight %}

Resultado:

	16
	3.0
	3.0

&nbsp;

###Asignación:

- `=`:

		<identificador> = <expresión>

También existen los operadores de asignación aumentada, que nos permiten realizar operaciones en la misma variable si tener que hacer referencia a su identificador, con un ejemplo se entenderá mejor:

{% highlight python %}
numero = 1

numero = numero + 1

print(numero)

numero += 1  # Asignación aumentada

print(numero)
{% endhighlight %}

Resultado:

	2
	3

Aquí les dejo el resto de los operadores de asignación aumentada (de los operadores que conocemos hasta ahora):

{% highlight python %}
numero += 1
numero -= 2
numero *= 3
numero /= 4
numero %= 5
numero //= 6
numero **= 7
{% endhighlight %}

&nbsp;

##Notas

Aunque no es obligatorio, los estandares de Python recomiendan dejar un espacio antes y después de cada operador:

	15 - 4
	"Hola" + " " + "Mundo!"

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
	<a href="{% post_url 2015-11-08-curso-python-ii %}" data-force="true"><button title="Curso de Python II: Variables" class="button"><i class="fa fa-chevron-left fa-fw"></i></button></a>

	<a href="{% post_url 2015-11-11-curso-python-iv %}" data-force="true"><button title="Curso de Python IV: Entrada y salida" class="button"><i class="fa fa-chevron-right fa-fw"></i></button></a>
</div>