---
title: "Curso de Python X: Lectura/Escritura de archivos"
description: "Una de las principales maneras de almacenar y controlar datos es mediante la lectura y escritura de archivos."
category: "informatica"
tags: "python curso"
type: "article"
postFolder: "Post/2"
image: "Python.png"
---

{{ page.description }} Para lograr esto, Python tiene una instrucción que activa estas tareas por medio de objetos que representan a los archivos. Les explicaré todo a partir de ejemplos y al final les mostraré la sintaxis.

###Lectura:

`archivo.txt`:

	Python con Ntrrg

	Este es un curso básico para aprender
	a programar con Python sin sufrir en 
	el intento :B

Para crear el objeto que haga referencia al archivo basta con escribir la siguiente instrucción:

{% highlight python %}
archivo = open("archivo.txt")
{% endhighlight %}

Si queremos obtener el contenido del archivo, existen tres métodos:

- `archivo.read()`: Lee todo el contenido desde la posición del puntero, este último es quien recorre caracter a caracter el archivo y nos permite obtener su contenido.

{% highlight python %}
archivo = open("archivo.txt")  # El puntero inicia en el byte 0 (P)

print(archivo.read())  # Lleva el puntero hasta el final
{% endhighlight %}

Resultado:

	Python con Ntrrg

	Este es un curso básico para aprender
	a programar con Python sin sufrir en 
	el intento :B

También es posible indicar cuantos caracteres (bytes) queremos leer después del puntero:

{% highlight python %}
archivo = open("archivo.txt")

print(archivo.read(16))  # Leerá 16 bytes
{% endhighlight %}

Resultado

	Python con Ntrrg

Como el puntero se mueve durante la lectura, si quisiéramos leer varias veces el contenido tendríamos que regresarlo y para esto se utiliza el método `.seek()`:

{% highlight python %}
archivo = open("archivo.txt")

print(archivo.read(16))

archivo.seek(0)  # Mueve el puntero al byte 0

print(archivo.read())
{% endhighlight %}

Resultado:

	Python con Ntrrg
	Python con Ntrrg

	Este es un curso básico para aprender
	a programar con Python sin sufrir en 
	el intento :B

Para saber en que byte se encuentra el puntero se usa el método `archivo.tell()`.

- `archivo.readline()`: Lee desde la ubicación del puntero hasta el siguiente final de línea.

{% highlight python %}
archivo = open("archivo.txt")

print(archivo.read())  # Lee hasta el final de la línea e incluye el salto
{% endhighlight %}

Resultado:

	Python con Ntrrg
	# Línea vacía

Si se indica una cantidad de bytes, al igual que en el método anterior, se leerán los bytes especificados después del puntero, pero con la diferencia de que si se encuentra con un salto de línea finalizará la lectura:

{% highlight python %}
archivo = open("archivo.txt")

print(archivo.read(30))
{% endhighlight %}

Resultado:

	Python con Ntrrg
	# Línea vacía

- `archivo.readlines()`: Lee todas las líneas como elementos de una lista desde la posición del puntero.

{% highlight python %}
archivo = open("archivo.txt")

print(archivo.readlines())
{% endhighlight %}

Resultado:

	['Python con Ntrrg\n', '\n', 'Este es un curso básico para aprender\n', 'a programar con Python sin sufrir en \n', 'el intento :B']

&nbsp;

###Escritura:

Se usa la misma instrucción, pero con un parámetro diferente:

{% highlight python %}
archivo = open("archivo.txt", "w")  # Escritura
archivo = open("archivo.txt", "a")  # Adición
{% endhighlight %}

Existen dos tipos de escrituras: la escritura tradicional que borra el contenido existente en el archivo especificado, para agregar los nuevos datos y la adición, que mantiene el contenido y agrega al final todo lo que se intente escribir. A diferencia de la lectura, poseen dos métodos:

- `archivo.write(<expresión>)`: Escribe la expresión en donde esté el puntero y en caso de que sea adición siempre escribirá al final del contenido sin importar donde esté el puntero.

{% highlight python %}
archivo = open("archivo.txt", "w")

archivo.write("Hola")
{% endhighlight %}

`archivo.txt`:

	Hola

&nbsp;

{% highlight python %}
archivo = open("archivo.txt", "w")

archivo.write("HOLA")

archivo.seek(1)

archivo.write("ola")  # Escribe donde está el puntero
{% endhighlight %}

`archivo.txt`:

	Hola

&nbsp;

{% highlight python %}
archivo = open("archivo.txt", "a")

archivo.write(" esto es Python")

archivo.seek(0)

archivo.write(" con Ntrrg")  # Escribe al final, sin importar el puntero
{% endhighlight %}

`archivo.txt`:

	Hola esto es Python con Ntrrg

- `archivo.writelines(<datoCompuesto>)`: Permite escribir los elementos de un dato compuesto en el que todos son cadenas, no agrega los saltos de línea automáticamente y al igual que el método anterior, su comportamiento dependerá del tipo de escritura.

{% highlight python %}
archivo = open("archivo.txt", "w")

tupla = ("Hola :B\n", "Soy Bata..")

archivo.writelines(tupla)
{% endhighlight %}

`archivo.txt`:

	Hola :B
	Soy Bata..

&nbsp;

###Combinación de métodos:

- `open("<archivo>", "r+")`: Lectura y escritura. Mantiene el contenido del archivo. Posiciona el puntero en el byte 0 y mientras no se use el método ´.seek()´ escribirá al final del contenido. Genera un error si no existe el archivo.

- `open("<archivo>", "w+")`: Escritura y lectura. Se comporta igual que la escritura tradicional, la única diferencia es que permite usar métodos de lectura.

- `open("<archivo>", "a+")`: Adición y lectura. Posiciona el puntero en el último byte. Se comporta igual que la adición, la única diferencia es que permite usar métodos de lectura.

&nbsp;

###Cierre:

Después de haber realizado todas las actividades que se esperaban con el archivo, debe cerrarse y para eso existe el método `.close()`, luego de esto ya no podrá realizarse ninguna actividad sobre el archivo con el objeto a menos que vuelva a definirse.

{% highlight python %}
archivo = open("archivo.txt")

print(archivo.read())

archivo.seek(0)
archivo.close()

print(archivo.read())
{% endhighlight %}

Resultado:

	Hola :B
	Soy Bata..
	Traceback (most recent call last):
	  File "script.py", line 8, in <module>
	    print(archivo.read())
	ValueError: I/O operation on closed file.

También existe la estructura `with` que se encarga de cerrar el archivo después que finalicen las sentencias dentro de el:

{% highlight python %}
with open("archivo.txt") as archivo:
	print(archivo.read())

print(archivo.read())
{% endhighlight %}

Resultado:

	Hola :B
	Soy Bata..
	Traceback (most recent call last):
	  File "script.py", line 4, in <module>
	    print(archivo.read())
	ValueError: I/O operation on closed file.

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
	<a href="{% post_url 2015-12-11-curso-python-ix %}" data-force="true"><button title="Curso de Python IX: Programación modular" class="button"><i class="fa fa-chevron-left fa-fw"></i></button></a>
</div>