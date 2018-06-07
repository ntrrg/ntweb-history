---
title: "Curso de Python IV: Entrada y salida"
description: "La mayoría de los algoritmos tienen el patrón \"Entrada, proceso y salida\" que consiste en recibir un dato, procesarlo y generar/mostrar información."
category: "informatica"
tags: "python curso"
type: "article"
postFolder: "Post/2"
image: "Python.png"
---

{{ page.description }} Para cumplir con esto, Python tiene varios métodos que cada programador escoje según necesite; dentro de las posibilidades están los archivos, las bases de datos, los dispositivos especiales (pantallas Braille, controles, etc...), varios más y los que usaremos nosotros por ahora, el teclado y la pantalla, que harán de entrada y salida respectivamente.

&nbsp;

##Salida

Hasta ahora hemos usado varias veces este método, nos permite mostrar por la **salida estandar** (normalmente el monitor) lo que necesitamos. Desde la versión 3 de Python, `print` dejó de ser una palabra reservada para convertirse en una función (esto lo veremos pronto y por eso no entraré en tanto detalle, pero es bueno tenerlo en cuenta). Para usarla solo necesitamos cumplir con la sintaxis:

	print([<expresión>[, ...]])

Si no se entiende mucho, aquí van unos ejemplos:

{% highlight python %}
print("Ntrrg")
print()  # Línea vacía
print(5 + 6)
print("Número " + str(11))  # Concatenar
print("Número", 11)  # Separador automático
print("Hola", "Mundo", "!")  # Se pueden especificar múltiples expresiones
print("Hola", "Mundo" + "!")  # Evitar el separador automático
{% endhighlight %}

Resultado:

	Ntrrg

	11
	Número 11
	Número 11
	Hola Mundo !
	Hola Mundo!

Para que esto sirva en Python 2, bastaría con borrar los parentesis:

{% highlight python %}
print "Ntrrg"
print
print 5 + 6
print "Número " + str(11)
print "Número", 11
print "Hola", "Mundo", "!"
print "Hola", "Mundo" + "!"
{% endhighlight %}

&nbsp;

##Entrada

Para darle más flexibilidad y usabilidad a nuestro script, necesitamos que pueda recibir datos directamente del usuario, para eso usaremos la **entrada estandar** (normalmente el teclado). Su sintaxis es muy parecida a la de `print`, solo que no permite varias expresiones separados por comas pero si concatenadas:

	input([<expresión>])

Ejemplos:

{% highlight python %}
input()  # Frena la ejecución hasta que se ingrese algún dato

nombre = input("Indique su nombre: ")

numero = int(input("Ingrese un número: "))  # Los datos recibidos son siempre serán cadenas

texto = "Ingrese otro número: "

otroNumero = int(input("Por favor, " + texto))  # Uso de variables y concatenación

print(numero + otroNumero)
{% endhighlight %}

Para que esto sirva en Python 2, tenemos que cambiar `input` por `raw_input`:

{% highlight python %}
raw_input()  # Frena la ejecución hasta que se ingrese algún dato

nombre = raw_input("Indique su nombre: ")

numero = int(raw_input("Ingrese un número: "))  # Los datos recibidos son siempre serán cadenas

texto = "Ingrese otro número: "

otroNumero = int(raw_input("Por favor, " + texto))  # Uso de variables y concatenación

print(numero + otroNumero)
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
	<a href="{% post_url 2015-11-11-curso-python-iii %}" data-force="true"><button title="Curso de Python III: Operadores" class="button"><i class="fa fa-chevron-left fa-fw"></i></button></a>

	<a href="{% post_url 2015-11-13-curso-python-v %}" data-force="true"><button title="Curso de Python III: Estructuras de decición" class="button"><i class="fa fa-chevron-right fa-fw"></i></button></a>
</div>