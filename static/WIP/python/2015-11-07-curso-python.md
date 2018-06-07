---
title: "Curso de Python"
description: "Python fue creado en 1991 por el holandés Guido van Rossum. Es un lenguaje interpretado, esto quiere decir que no es necesario compilarlo (convertirlo en lenguaje máquina) para ejecutarlo."
category: "informatica"
tags: "python curso"
type: "article"
postFolder: "Post/2"
image: "Python.png"
---

Antes de empezar, tenemos que conocer algunos conceptos que tal vez no nos ayuden mucho con el manejo del lenguaje, pero sí nos dejarán entender un poco qué es lo que realmente hace un programador:

- **Lógica**: Para evitar tecnicismos, se puede definir como la ciencia encargada de demostrar la coherencia de hechos o teorías.
- **Programación**: En la informática, consiste en crear una lista de instrucciones (algoritmo) para lograr que algún equipo realice tareas específicas.

Con base en estos dos conceptos, se crea la locución **Lógica de programación**, que decidí definir como la capacidad de crear algoritmos informáticos coherentes y optimizados, esta capacidad no está estrictamente ligada a algún lenguaje de programación, pero en este caso, usaremos Python para ponerla en práctica.

Python fue creado en 1991 por el holandés Guido van Rossum. Es un lenguaje interpretado, esto quiere decir que no es necesario compilarlo (convertirlo en lenguaje máquina) para ejecutarlo, aunque le resta un poco de rendimiento. Tiene varias implementaciones como Jython, que usa un intérprete (el programa que ejecuta el código) escrito en Java; PyPy, que esta escrito con el mismo Python y algunos más, pero nos enfocaremos en **CPython** que está escrito en C y es la implementación oficial del proyecto.

En mi opinión, es uno de los lenguajes de programación más apropiados para aprender a programar por las siguientes razones:

- Está preinstalado en muchas distribuciones de GNU/Linux.
- Al ser interpretado, su ejecución y corrección de errores serán mucho más fáciles.
- Es de alto nivel, por lo que permite que el código fuente sea legible y más sencillo de escribir.
- Obliga a sus programadores a respetar la identación (tabulaciones a la izquierda) y tiene estándares de sintaxis que enseñan buenas prácticas.
- Es uno de los lenguajes de programación más populares y con más demanda en el mundo.

&nbsp;

##Herramientas necesarias

1. Un editor de texto (yo uso [Sublime Text 3](http://www.sublimetext.com/3)).
2. El intérprete CPython (Aquí se explicará como instalarlo).

&nbsp;

##Primeros pasos

{% highlight python %}
print("Esto es Python con Ntrrg :B")
{% endhighlight %}

Actualmente Python está en la versión 3.5, pero debido a la gran cantidad de programas que se han hecho para las diferentes distribuciones de GNU/Linux con versiones anteriores, estas aún mantienen la 2.7, por lo que, aunque el curso esté orientado a la 3.5, intentaré mostrar como se escribirían las instrucciones en ambas.

&nbsp;

###Instalación:

	# apt install python python3

Python 2.7 ya viene preinstalado, por lo que no será necesario hacer nada más. Con esto se instalará el intérprete de Python (CPython) en la versión 3.4.2 y nuestra computadora ya podrá entender las instrucciones que le demos.

Tranquilos, no es un error lo de 3.4.2, como lo instalamos desde los repositorios de Debian, se instaló la última versión pero de sus repositorios; ya que la versión 3.5 es reciente, aún necesita ser probada por sus desarrolladores. Para nosotros no habrá problema, pero si de todas formas quieren instalar la versión 3.5 o están usando **Windows** hagan click [aquí]({% post_url 2015-11-20-instalar-python %}).

&nbsp;

###El intérprete:

	$ python3
	Python 3.4.2 (default, Oct  8 2014, 10:45:20) 
	[GCC 4.9.1] on linux
	Type "help", "copyright", "credits" or "license" for more information.
	>>> 

O con Python 2:

	$ python
	Python 2.7.9 (default, Mar  1 2015, 12:57:24) 
	[GCC 4.9.2] on linux2
	Type "help", "copyright", "credits" or "license" for more information.
	>>> 

Desde aquí podemos ejecutar las sentencias que queramos, como por ejemplo, la más famosa del mundo informático:

	>>> print("Hola mundo")
	Hola mundo

Claro, esto solo nos servirá para probar el comportamiento de algunas ordenes. Si queremos aprovechar al máximo a Python usaremos **scripts**, que simplemente son archivos con algoritmos muy específicos.

&nbsp;

##Scripts de Python

Igual que cualquier otro archivo, los scripts no necesitan una extensión específica, aunque por convenio se usa `.py`. Un ejemplo simple:

`prueba.py`:

{% highlight python %}
print("Hola mundo")
{% endhighlight %}

Se pueden ejecutar con el comando del intérprete:

	$ python3 prueba.py
	Hola mundo

En Python 2:

	$ python prueba.py
	Hola mundo

&nbsp;

O como un binario ejecutable:

	$ ./prueba.py

Pero antes de eso, tenemos que hacerle algunos arreglos, hay que decirle al sistema operativo con que debe ejecutar nuestro script y para eso usaremos un comentario preprocesado llamado **shebang**, que debe estar en primera línea:

{% highlight python %}
#!/usr/bin/env python3

print("Hola mundo")
{% endhighlight %}

En Python 2:

{% highlight python %}
#!/usr/bin/env python

print "Hola mundo"
{% endhighlight %}

Además del shebang, en Python 2 necesitaremos otro comentario preprocesado para poder usar caracteres con tildes:

{% highlight python %}
#!/usr/bin/env python
# -*- coding: utf-8 -*-

print "Hola mundo con Ñ"
{% endhighlight %}

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
	<a href="{% post_url 2015-11-08-curso-python-ii %}" data-force="true"><button title="Curso de Python II: Variables" class="button"><i class="fa fa-chevron-right fa-fw"></i></button></a>
</div>