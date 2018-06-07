---
title: "Curso de Python VIII: POO"
description: "La programación orientada a objetos es uno de los paradigmas más populares, permite representar cualquier cosa por medio de sus características y capacidades."
category: "informatica"
tags: "python curso"
type: "article"
postFolder: "Post/2"
image: "Python.png"
---

{{ page.description }} Para lograr esta representación, los lenguajes de programación cuentan con algo llamado **Clases** que describen de forma generalizada alguna especie o tipo de elemento y **Objetos**, que son los elementos como tal (una instancia); para entenderlo mejor, la raza humana puede llamarse una clase y cada uno de nosotros es una instancia de ella.

Dentro de las clases se declaran los **atributos** (características) y **métodos** (capacidades) que poseerán los objetos, siguiendo con el ejemplo, los atributos serían la altura, el peso y cualquier otro dato descriptivo; los métodos serían respirar, caminar y cualquier otra cosa que pueda hacer un humano. Para definir una clase en Python hay que seguir la sintaxis:

	class <Identificador>:
		[<atributo> = <expresión>]

		[def <nombreMétodo>(self[, <parámetro>[, ...]]):
			<sentencias>
			[return <expresión>]]

Con el ejemplo:

{% highlight python %}
class Humano:
	altura = 1.87
	peso = 68

	def caminar(self):
		print("Caminando! :B")
{% endhighlight %}

Como ven, la primera letra del identificador de la clase está en mayúscula, no es obligatorio pero es parte del estándar. Los atributos se comportan como variables comunes y los métodos como funciones. El parámetro `self` es obligatorio (aunque no su nombre) y hará referencia al objeto que está invocando el método (después de mostrarles como instanciar una clase les explico mejor esto).

&nbsp;

Para instanciar una clase (crear un objeto):

	<identificador> = <Clase>()

	<identificador>.<atributo>  # Obtener valor de un atributo
	<identificador>.<método>()  # Ejecutar un método

Con el ejemplo:

{% highlight python %}
persona = Humano()

print(persona.altura)
persona.caminar()
{% endhighlight %}

Resultado:

	1.87
	Caminando! :B

Para acceder a los elementos del objeto se usa el operador `.`. Cuando se invocan métodos, se pasa al mismo objeto como parámetro (por eso es obligatorio tener al menos un parámetro y el estándar recomienda llamarlo `self`) y en el caso de que algún método necesite elementos del objeto bastará con definirse de la siguiente manera:

{% highlight python %}
class Humano:
	altura = 1.87
	peso = 68

	def caminar(self):
		print("Caminando! :B")

	def imc(self):
		return self.peso / self.altura ** 2

persona = Humano()

print(persona.imc())
{% endhighlight %}

Resultado:

	19.445794846864363

&nbsp;

Aunque ya aprendimos a crear e instanciar clases, hasta ahora lo hemos hecho de manera estática pues no todos los humanos tienen la misma altura y peso, para definir valores específicos por objeto hace falta usar un método especial llamado `__init__`, que se invoca automáticamente durante la instanciación. Les muestro el cambio a la sintaxis y después el ejemplo:

	class <Identificador>:
		[<atributo> = <expresión>]

		[def __init__(self[, <parámetro>[, ...]]):
			self.<atributo> = <expresión>
			<sentencias>]

		[def <nombreMétodo>(self[, <parámetro>[, ...]]):
			<sentencias>
			[return <expresión>]]

	<identificador> = <Clase>([<parámetro>[, ...]])

{% highlight python %}
class Humano:
	def __init__(self, altura, peso):
		self.altura = altura
		self.peso = peso

	def caminar(self):
		print("Caminando! :B")

	def imc(self):
		return self.peso / self.altura ** 2

persona = Humano(1.87, 68)
persona2 = Humano(1.67, 56)
{% endhighlight %}

&nbsp;

##Herencia

Consiste en crear nuevas clases con base en otras (superclases), para lograrlo simplemente hay que agregarlas entre paréntesis al lado del identificador de la nueva clase:

	class <Identificador>[(<superclase>[, ...])]:

{% highlight python %}
class Humano:
	def __init__(self, altura, peso):
		self.altura = altura
		self.peso = peso

	def caminar(self):
		print("Caminando! :B")

	def imc(self):
		return self.peso / self.altura ** 2

class Mujer(Humano):
	sexo = "Femenino ♀"

class Hombre(Humano):
	sexo = "Masculino ♂"

persona = Hombre(1.87, 68)
persona2 = Mujer(1.67, 56)

print(persona.sexo, persona.imc())
print(persona2.sexo, persona2.imc())
{% endhighlight %}

Resultado:

	Masculino ♂ 19.445794846864363
	Femenino ♀ 20.07960127648894

Los atributos y métodos de la nueva clase sobrescribirán los de las superclases si existen, en el caso de las superclases, los elementos de las más cercanas al identificador se mantendrán (de izquierda a derecha). Para reescribir el `__init__` solo es necesario agregarlo en la nueva clase pero también hará falta invocar el de la superclase dentro de el:

	class <Identificador>[(<superclase>[, ...])]:
		def __init__(self[, <parámetro>[, ...]]):
			<superClase>.__init__(self[, <parámetro>[, ...]])

			[self.<atributo> = <expresión>]

{% highlight python %}
class Humano:
	def __init__(self, altura, peso):
		self.altura = altura
		self.peso = peso

	def caminar(self):
		print("Caminando! :B")

	def imc(self):
		return self.peso / self.altura ** 2

class Mujer(Humano):
	sexo = "Femenino ♀"

	def __init__(self, altura, peso, baile):
		Humano.__init__(self, altura, peso)  # Instanciación de Humano

		self.baile = baile

class Hombre(Humano):
	sexo = "Masculino ♂"

persona = Hombre(1.87, 68)
persona2 = Mujer(1.67, 56, "Salsa")

print(persona.sexo, persona.imc())
print(persona2.sexo, persona2.imc(), "Baile favorito:", persona2.baile)
{% endhighlight %}

Resultado:

	Masculino ♂ 19.445794846864363
	Femenino ♀ 20.07960127648894 Baile favorito: Salsa

&nbsp;

##Polimorfismo

Es un técnica que nos permite alterar el comportamiento de los objetos y así usarlos donde se espera un dato de un tipo diferente, para aprovechar esta cualidad del lenguaje tenemos que usar métodos especiales (parecidos a `__init__`) y el método tradicional de cambio de tipos.

- **Enteros**: se activa usando `int(<objeto>)`.

		def __int__(self):
			<sentencias>
			return <entero>

- **Decimales**: se activa usando `float(<objeto>)`.

		def __float__(self):
			<sentencias>
			return <decimal>

- **Complejos**: se activa usando `complex(<objeto>)`.

		def __complex__(self):
			<sentencias>
			return <númeroComplejo>

- **Lógicos**: se activa usando `bool(<objeto>)`.

		def __bool__(self):
			<sentencias>
			return True | False

- **Cadenas**: se activa usando `str(<objeto>)` o `print(<objeto>)`.

		def __str__(self):
			<sentencias>
			return <cadena>

&nbsp;

##Encapsulamiento

Permite ocultar atributos y métodos para que sean accesibles solo desde métodos propios del objeto, su utilidad principalmente es asegurar que ciertos datos sean leídos o escritos cumpliendo con ciertas reglas, por ejemplo: el estado de un usuario dependerá de si inició sesión y al intentarlo ingresó datos válidos:

{% highlight python %}
class Usuario:
	def __init__(self, nombre, contrasena):
		self.nombre = nombre
		self.__estado = self.__iniciar(contrasena)

	def __iniciar(self, contrasena):
		# Datos correctos
		if contrasena == "1234":
			return True

		# Datos incorrectos
		else:
			return False

	def estado(self):
		return self.__estado

	def cerrar(self):
		self.__estado = False
{% endhighlight %}

El atributo `estado` contiene un valor que solo podrá determinarse después de realizar algunas comprobaciones en bases de datos, archivos o algún otro método (en el ejemplo no se hacen para evitar complicar las cosas :B); ya que debe realizarse estrictamente ese procedimiento, no puede modificarse con libertad el valor de este atributo y para establecer la limitación se debe declarar como un atributo **privado**, la forma de lograrlo es agregar dos guiones bajos antes de su nombre `__estado` y de esta manera, aunque intentemos leer o escribir este atributo fuera de los métodos del objeto, Python no nos dejará. Para poder leer el valor de `__estado` se define un método **público** (`estado()`) y otros dos (`__iniciar` **privado** y `cerrar()` **público**) para escribirlo.

{% highlight python %}
usuario = Usuario("ntrrg", "1234")
print(usuario.nombre, "Estado:", "Conectado" if usuario.estado() else "Desconectado")
usuario.cerrar()
print(usuario.nombre, "Estado:", "Conectado" if usuario.estado() else "Desconectado")
{% endhighlight %}

Resultado:

	ntrrg Estado: Conectado
	ntrrg Estado: Desconectado

&nbsp;

Existe un pequeño truco para poder acceder a los atributos y métodos privados:

	<objeto>._<clase>__<atributoPrivado>
	<objeto>._<clase>__<métodoPrivado>()

{% highlight python %}
usuario = Usuario("ntrrg", "1234")
print(usuario._Usuario__estado)
usuario.cerrar()
print(usuario._Usuario__estado)
usuario._Usuario__estado = usuario._Usuario__iniciar("1234")
print(usuario._Usuario__estado)
{% endhighlight %}

Resultado:

	True
	False
	True

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
	<a href="{% post_url 2015-12-02-curso-python-vii %}" data-force="true"><button title="Curso de Python VII: Funciones" class="button"><i class="fa fa-chevron-left fa-fw"></i></button></a>

	<a href="{% post_url 2015-12-11-curso-python-ix %}" data-force="true"><button title="Curso de Python VII: Programación modular" class="button"><i class="fa fa-chevron-right fa-fw"></i></button></a>
</div>