---
title: "Curso de Python II: Variables"
description: "Las variables son espacios de la RAM en los que podemos almacenar datos, están compuestas por un identificador, que apunta a la dirección del espacio de memoria que tiene asignado y un valor, que contiene los datos como tal."
category: "informatica"
tags: "python curso"
type: "article"
postFolder: "Post/2"
image: "Python.png"
---

{{ page.description }} Haciéndole honor a su nombre, tienen la capacidad de modificar sus valores en cada ejecución de los scripts y para hacerlo más interesante, hasta en una ejecución pueden variar múltiples veces.

Si estudiamos una de las operaciones aritmética básica es posible entender mucho más fácil esta definición, intentemos con la suma. Desde pequeños aprendimos que `2 + 2 = 4`, pero también aprendimos que la suma no está limitada a solo esos números pues existe una cantidad infinita de números y expresiones que podríamos probar; entonces, estaría mal definir la suma como `2 + 2` pues nadie entendería que los `2` de esa ecuación pueden modificarse, por lo que tenemos que usar variables para arreglarlo, si decimos que la suma es `número + otroNúmero` se convertiría en una definición mucho más generalizada y fácil de entender, además de cumplirse la estructura que les dije al principio, un identificador (`número` y `otroNúmero`) y su valor (los datos que usemos para probar la sentencia).

{% highlight python %}
numero = 2
otroNumero = 3

print(numero + otroNumero)

numero = 5
otroNumero = 1

print(numero + otroNumero)
{% endhighlight %}

Que resultaría algo como:

	5
	6

&nbsp;

##Declaración

Python es un lenguaje de tipado dinámico, esto quiere decir que no es necesario decirle de que tipo son las variables que usamos pues él lo determinará por nosotros, lo único que quedaría por hacer es asignarles valores:

	<identificador> = <expresión>

Los identificadores no pueden escribirse con caracteres especiales (solo los guiones bajos están permitidos), tildes o empezar con números; además es recomendable que permitan determinar a que se refieren con solo leerlos. Ejemplos:

- Usos correctos:
	
		numero
		numero2
		otro_numero
		otroNumero

- Usos incorrectos:

		número
		2numero
		otro-numero

Una expresión puede ser un valor literal (valores escritos directamente), el resultado de una operación y hasta invocaciones de funciones o métodos (esto lo veremos más adelante, intenten acordarse o anótenlo :B). Ejemplos:

{% highlight python %}
# Valores literales

numero = 1
otroNumero = 2
frase = "Hola mundo"

# Resultado de operaciones

suma = numero + otroNumero
resta = otroNumero - numero
{% endhighlight %}

&nbsp;

##Eliminación

Además de crear variables, también es posible eliminarlas cuando ya no son necesarias para liberar memoria. Es bastante sencillo, solo hay que usar la palabra reservada (palabras que el lenguaje interpreta como ordenes) `del`, ejemplo:

{% highlight python %}
numero = 1

del numero
{% endhighlight %}

&nbsp;

##Tipos

Aunque antes dije que Python determinaba el tipo de variable por nosotros, no me refería a que lo hiciera mágicamente (aunque parece :D), realmente lo hace porque estudia el dato que le estamos asignando y determina cual se adapta más, por esto necesitamos conocer con cuales cuenta el lenguaje.

###Simples:

- **Numéricos**:

		1  # Entero (int)
		2147483647  # Máximo valor de un entero en 32 bits
		9223372036854775807  # Máximo valor de un entero en 64 bits

		1L  # Entero largo (long)
		2147483648  # Primer valor de un entero largo en 32 bits
		9223372036854775808  # Primer valor de un entero largo en 64 bits

		1.0  # Decimal (float)
		1j  # Complejo (complex)

		0b1011  # Binario (bin)
		0o13  # Octal (oct)
		0xb  # Hexadecimal (hex)

	A partir de Python 3 se eliminaron los enteros largos, por lo que los enteros perdieron su límite máximo.

- **Lógicos**:

		True  # Verdadero
		False  # Falso
		None  # Vacío

	Los datos lógicos, la mayoría de las veces, sirven para comprobar condiciones. Es probable que no le vean mucha utilidad por ahora, pero más adelante podremos crear algoritmos mucho más flexibles con ellos.

###Compuestos

Pueden contener varios datos, a diferencia de los simples que solo pueden almacenar un valor literal, por esta razón, no podemos acceder a uno de sus valores simplemente indicando su identificador sino que debemos usar una sintaxis especial que varía en algunos casos.

- **Cadenas de caracteres**:

		'Hola mundo'
		"Hola mundo"
		"""
			Hola
			Mundo
		"""

	Es el único dato compuesto que no necesita de un delimitador para sus elementos, pues cada caracter representa una posición. Dentro de las cadenas se pueden usar tecnicas de escape anteponiéndoles `\`, esto causa un comportamiento diferente al habitual, algunos de los posibles escapes son:

	- \\\<caracter>: Evita que \<caracter> tenga un comportamiento especial.

			'Apóstrofo (\')'
			"Comillas (\")"
			"Backslash \\"

		Muestra:

			Apóstrofo (')
			Comillas (")
			Backslash \

		Para escribir `"` en las cadenas multilineales no hace falta escaparlas.

	- \b: Borrar un caracter.

			"Hola\b"

		Muestra:

			Hol

	- \n: Salto de línea.

		\r: Retorno de carro.

			"Hola\nMundo"
			"Hola\r\nMundo"  # Para Windows

		Muestra:

			Hola
			Mundo

	- \t: Tabulador.

			"\tHola"

		Muestra:

			    Hola

	- \v: Tabulador vertical.

			"Hola\vMundo"

		Muestra:

			Hola
			    Mundo

	Para acceder a valores específicos se usa la siguiente sintaxis

		cadena = "Hola"

		print(cadena[0])

	Y el resultado sería:

		H  # Las posiciones (índices) empiezan en 0

- **Tuplas**:

		(10, "Valor 2", 30.0, 0b100, 0x5)
		(1,)  # Definición de una tupla de un elemento

	Para acceder a valores específicos se usa la siguiente sintaxis

		tupla = (10, "Valor 2", 30.0, 0b100, 0x5)

		print(tupla[1])

	Y el resultado sería:

		Valor 2

- **Listas**:

		[10, "Valor 2", 30.0, 0b100, 0x5]

	Para acceder a valores específicos se usa la siguiente sintaxis

		lista = (10, "Valor 2", 30.0, 0b100, 0x5)

		print(lista[-1])  # Busca de manera inversa

	Y el resultado sería:

		5

- **Conjuntos**:

		{10, "Valor 2", 30.0, 0b100, 0x5}

	Los conjuntos eliminan automáticamente elementos repetidos y permiten realizar operaciones algebraicas de conjuntos como unión, intersección, diferencia y diferencia simétrica (cuando veamos los métodos estudiaremos esto).

- **Diccionarios**:

		{"Clave 1": "Valor 1", "Clave 2": 2, "Clave 3": 3.0}

	Los diccionarios tienen la capacidad de asignar un identificador (clave) a cada uno de sus elementos. Para acceder a valores específicos se usa la siguiente sintaxis

		diccionario =  {"Clave 1": "Valor 1", "Clave 2": 2, "Clave 3": 3.0}

		print(diccionario["Clave 3"])

	Y el resultado sería:

		3.0

&nbsp;

##Cambio de tipo

###Simples:

- **Numéricos**:

	**Enteros**:

		int(10.6)  # Retorna 10, no redondea
		int(0b110)  # Retorna 6
		int(0o14)  # Retorna 12
		int(0x1e)  # Retorna 30

	**Decimales**:

		float(5)  # Retorna 5.0
		float(0b110)  # Retorna 6.0
		float(0o14)  # Retorna 12.0
		float(0x1e)  # Retorna 30.0

	**Complejos**:

		complex(5)  # Retorna (5+0j)
		complex(10.6)  # Retorna (10.6+0j)
		complex(0b110)  # Retorna (6+0j)
		complex(0o14)  # Retorna (12+0j)
		complex(0x1e)  # Retorna (30+0j)

	**Binarios**:

		bin(5)  # Retorna '0b101'
		bin(0o14)  # Retorna '0b1100'
		bin(0x1e)  # Retorna '0b11110'

	Retorna una cadena con el número binario pues Python lo interpretaría.

	**Octales**:

		oct(5)  # Retorna '0o5'
		oct(0b110)  # Retorna '0o6'
		oct(0x1e)  # Retorna '0o36'

	Retorna una cadena con el número octal pues Python lo interpretaría.

	**Hexadecimales**:

		hex(5)  # Retorna '0x5'
		hex(0b110)  # Retorna '0x6'
		hex(0o14)  # Retorna '0xc'

	Retorna una cadena con el número hexadecimal pues Python lo interpretaría.

- **Lógicos**:

		bool(0)  # Retorna False
		bool(1)  # Retorna True
		bool("")  # Retorna False
		bool("M")  # Retorna True
		bool([])  # Retorna False
		bool([1, 2])  # Retorna True

###Compuestos:

- **Cadenas**:

		str(10)  # Retorna '10'
		str(0b110)  # Retorna '6'
		str(True)  # Retorna 'True'
		str([1, 2, 3, 4, 5])  # Retorna '[1, 2, 3, 4, 5]'
		str({3, 1, 2})  # Retorna '{1, 2, 3}', los conjuntos se ordenan al cambiar (solo los valores numéricos, los demás estarán desordenados y pueden aparecer entre los números ordenados)

- **Tuplas**:

		tuple("Ntrrg")  # Retorna ('N', 't', 'r', 'r', 'g')
		tuple([1, 2, 3, 4, 5])  # Retorna (1, 2, 3, 4, 5)
		tuple({3, 1, 2})  # Retorna (1, 2, 3)
		tuple({"Clave 1" : 1, "Clave 2" : 2})  # Retorna ('Clave 2', 'Clave 1'), el orden es irrelevante para los diccionarios

- **Listas**:

		list("Ntrrg")  # Retorna ['N', 't', 'r', 'r', 'g']
		list((1, 2, 3, 4, 5))  # Retorna [1, 2, 3, 4, 5]
		list({3, 1, 2})  # Retorna [1, 2, 3]
		list({"Clave 1" : 1, "Clave 2" : 2})  # Retorna ['Clave 2', 'Clave 1']

- **Conjuntos**:

		set("Ntrrg")  # Retorna {'r', 'g', 'N', 't'}
		set((1, 2, 3, 4, 5))  # Retorna {1, 2, 3, 4, 5}

		# Versión inmutable

		frozenset("Ntrrg")  # Retorna frozenset({'r', 'g', 'N', 't'})
		frozenset((1, 2, 3, 4, 5))  # Retorna frozenset({1, 2, 3, 4, 5})

- **Diccionarios**:

		dict([("Nombre", "MA"), ("Apellido", "RN")])  # Retorna {'Nombre': 'MA', 'Apellido': 'RN'}

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
	<a href="{% post_url 2015-11-07-curso-python %}" data-force="true"><button title="Curso de Python" class="button"><i class="fa fa-chevron-left fa-fw"></i></button></a>

	<a href="{% post_url 2015-11-11-curso-python-iii %}" data-force="true"><button title="Curso de Python III: Operadores" class="button"><i class="fa fa-chevron-right fa-fw"></i></button></a>
</div>