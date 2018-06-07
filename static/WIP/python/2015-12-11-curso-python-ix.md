---
title: "Curso de Python IX: Programación modular"
description: "La programación modular consiste en subdividir programas de manera organizada para mejorar la legibilidad y aumentar significativamente la reutilización de código."
category: "informatica"
tags: "python curso"
type: "article"
postFolder: "Post/2"
image: "Python.png"
---

{{ page.description }} Les mostraré su utilidad con este ejemplo: tenemos un script que se encarga de crear figuras geométricas:

{% highlight python %}
class Triangulo:
	nombre = "Triángulo"

	def __init__(self, lados):
		self.lados = lados

		print("Se creó un triángulo!")

class Circulo:
	nombre = "Círculo"

	def __init__(self, radio):
		self.radio = radio

		print("Se creó un círculo!")

class Cuadrado:
	nombre = "Cuadrado"

	def __init__(self, lado):
		self.lado = lado

		print("Se creó un cuadrado!")

a = Triangulo((3, 2, 2))
b = Circulo(1)
c = Cuadrado(2)
{% endhighlight %}

Como se definen todas las clases en el mismo script, se vuelve un poco largo el código y si alguien quiere leerlo probablemente le parezca enredado, así que la mejor forma de solucionar este problema es separarlo en varios archivos:

`figuras.py`:

{% highlight python %}
class Triangulo:
	nombre = "Triángulo"

	def __init__(self, lados):
		self.lados = lados

		print("Se creó un triángulo!")

class Circulo:
	nombre = "Círculo"

	def __init__(self, radio):
		self.radio = radio

		print("Se creó un círculo!")

class Cuadrado:
	nombre = "Cuadrado"

	def __init__(self, lado):
		self.lado = lado

		print("Se creó un cuadrado!")
{% endhighlight %}

`script.py`:

{% highlight python %}
a = Triangulo((3, 2, 2))
b = Circulo(1)
c = Cuadrado(2)
{% endhighlight %}

Claro, así todavía no funciona pues solo lo separamos, si queremos utilizar las clases tenemos que importar el archivo que las contiene (`figuras.py`) y para esto existen dos formas:

- Crear un objeto que contenga todos los elementos (variables, funciones, clases, etc...) del archivo:

{% highlight python %}
import figuras

a = figuras.Triangulo((3, 2, 2))
b = figuras.Circulo(1)
c = figuras.Cuadrado(2)
{% endhighlight %}

Se puede utilizar la palabra reservada `as` para renombrar el objeto `import figuras as fig` en caso de que ya exista algún elemento del script con el mismo identificador o simplemente para reducir su tamaño.

- Agregar los elementos del archivo importado al actual:

{% highlight python %}
from figuras import *

a = Triangulo((3, 2, 2))
b = Circulo(1)
c = Cuadrado(2)
{% endhighlight %}

El `*` simboliza todos los elementos, si se quisiera importar alguno en específico solo haría falta cambiar el `*` por su nombre `from figuras import Triangulo`. Al igual que el método anterior puede usarse la palabra reservada `as` pero solo cuando se importa un elemento `from figuras import Triangulo as Tri` o varios separados por comas `from figuras import Triangulo as Tri, Circulo as Cir`.

&nbsp;

En ambos casos, el archivo `figuras.py` debe estar en la misma carpeta que `script.py`. Los archivos que se importan son llamados **Módulos** y su función principal es contener las herramientas necesarias para ejecutar ciertos procesos. Python también usa **Paquetes** que son simplemente carpetas que contienen módulos, en el mundo de la programación se les conoce como bibliotecas o repositorios de código; para configurar paquetes solo hace falta crear en el interior de una carpeta un archivo vacío llamado `__init__.py`, bastante sencillo no? para usar los módulos dentro de un paquete tenemos que anteponer el nombre del paquete cuando lo importamos, adaptemos el ejemplo para entenderlo mejor aún:

	.
	├─ geometria/
	│  ├── __init__.py
	│  ├── circulo.py
	│  ├── cuadrado.py
	│  └── triangulo.py
	└─ script.py

`circulo.py`:

{% highlight python %}
class Circulo:
	nombre = "Círculo"

	def __init__(self, radio):
		self.radio = radio

		print("Se creó un círculo!")
{% endhighlight %}

`cuadrado.py`:

{% highlight python %}
class Cuadrado:
	nombre = "Cuadrado"

	def __init__(self, lado):
		self.lado = lado

		print("Se creó un cuadrado!")
{% endhighlight %}

`triangulo.py`:

{% highlight python %}
class Triangulo:
	nombre = "Triángulo"

	def __init__(self, lados):
		self.lados = lados

		print("Se creó un triángulo!")
{% endhighlight %}

`script.py`:

{% highlight python %}
import geometria.circulo, geometria.cuadrado, geometria.triangulo

a = geometria.triangulo.Triangulo((3, 2, 2))
b = geometria.circula.Circulo(1)
c = geometria.cuadrado.Cuadrado(2)
{% endhighlight %}

&nbsp;

Si se crean más subpaquetes, el proceso sería el mismo:

	.
	└─ paquete/
	   ├── __init__.py
	   └─ subpaquete/
	      ├── __init__.py
	      └── modulo.py

Y para importar los módulos:

{% highlight python %}
import paquete.subpaquete.modulo
import paquete.subpaquete.modulo as modulo  # Se puede usar "as"
{% endhighlight %}

&nbsp;

Hasta ahora, los módulos y paquetes que creemos deben estar en la misma carpeta del script que los importará, si queremos alterar este comportamiento tenemos que decirle a Python donde queremos que los busque y para esto existen dos formas:

- Temporal (mientras dure la ejecución del script):

Hay una variable de Python que contiene una lista de ubicaciones donde buscar módulos, solo haría falta agregarle la ruta a la carpeta que contiene los módulos y ya podríamos usarlos aunque con una diferencia, no será necesario crear el archivo `__init__.py` ni anteponer el nombre del paquete (solo en la carpeta principal, sus subcarpetas si necesitan que se hagan las dos). Les explico con un ejemplo:

	.
	├─ geometria/
	│  ├── circulo.py
	│  ├── cuadrado.py
	│  └── triangulo.py
	└─ Python/
	   └── script.py

`script.py`:

{% highlight python %}
import sys

sys.path += ["../geometria"]

import triangulo

a = triangulo.Triangulo((3, 2, 2))
{% endhighlight %}

Resultado:

	Se creó un triángulo!

Como ven, `sys.path` soporta rutas relativas, pero esto es un poco peligroso pues esta ruta será relativa a la carpeta desde la que se está ejecutando el script y no en la que está almacenado, es decir, si lo ejecutamos desde `/home/ntrrg/Descargas/Python/` (que es donde yo lo guardé) entonces `../geometria` equivaldrá a `/home/ntrrg/Descargas/geometria/` y  el resultado sería el esperado; si se ejecuta desde `/home/ntrrg/Descargas/` entonces equivaldrá a `/home/ntrrg/geometria/` y apuntará a una ruta que no existe. Con una ruta absoluta quedaría así:

`script.py`:

{% highlight python %}
import sys

sys.path += ["/home/ntrrg/Descargas/geometria"]

import triangulo

a = triangulo.Triangulo((3, 2, 2))
{% endhighlight %}

- Permanente:

Se debe agregar la siguiente línea en el archivo `/etc/profile` y reiniciar la sesión de usuario:

	export PYTHONPATH=$PYTHONPATH:/home/ntrrg/Descargas/geometria

	# Para agregar varias ubicaciones:
	export PYTHONPATH=$PYTHONPATH:<ruta>:<otraRuta>

Me disculpo con los que estén usando Windows, pero realmente no conozco el equivalente a este método.

&nbsp;

Todas las sentencias escritas en el cuerpo de los módulos serán ejecutadas cuando se importen, aunque es posible crear un bloque que se ejecute solo en caso de que el módulo sea llamado como un script, que utilidad tiene esto? bueno, imaginen que quieren obtener el área de un triángulo usando el módulo `triangulo.py`, sería muy fastidioso tener que crear un script solo para importarlo, así que les muestro como usarlo directamente:

	# El primer parámetro será la altura y el segundo la base
	python3 geometria/triangulo.py 3 3

Y los cambios en el módulo:

`triangulo.py`:

{% highlight python %}
class Triangulo:
	nombre = "Triángulo"

	def __init__(self, lados):
		self.lados = lados

		print("Se creó un triángulo!")

# Bloque que se ejecuta solo al llamarse como script

if __name__ == "__main__":
	import sys  # Esta biblioteca nos deja leer los parámetros

	# sys.argv[0] contiene la ruta con la que fue ejecutado el módulo
	base = int(sys.argv[1])
	altura = int(sys.argv[2])

	print(base * altura / 2)
{% endhighlight %}

&nbsp;

Por último, el estándar de Python tiene ciertas recomendaciones al momento de importar módulos:

- Escribir las instrucciones de importación en las primeras líneas del script.
- Ordenar los paquetes/módulos en tres secciones:

		<paquetes/módulos de Python>
		# Una línea vacía
		<paquetes/módulos de herramientas externas>
		# Una línea vacía
		<paquetes/módulos propios>

- Organizar alfabéticamente los paquetes/módulos.

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
	<a href="{% post_url 2015-12-06-curso-python-viii %}" data-force="true"><button title="Curso de Python VIII: POO" class="button"><i class="fa fa-chevron-left fa-fw"></i></button></a>
	
	<a href="{% post_url 2015-12-14-curso-python-x %}" data-force="true"><button title="Curso de Python X: Lectura/Escritura de archivos" class="button"><i class="fa fa-chevron-right fa-fw"></i></button></a>
</div>