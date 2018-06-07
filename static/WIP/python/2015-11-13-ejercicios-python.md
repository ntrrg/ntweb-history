---
title: "Ejercicios"
published: false
---

1. Cree un script que permita sumar dos números ingresados por la entrada estandar y muestre el resultado.
2. Cree un script que reciba dos cadenas por la entrada estandar y muestrelas en una línea separadas por una tabulación
3. Cree un script que sume los elementos de una lista, la cantidad de elementos y sus valores (numéricos) serán asignados por el programador. Debe mostrar el resutado final con este formato: `El resultado es: 10`.
4. Cree un script que borre los elementos repetidos de la tupla `(1, 2, 2, 2, 3, 3, 4, 4, 5, 5, 5)` y muestre el resultado.
5. Cree un script que reciba un número por la entrada estandar y muestre si es múltipo de 3.
6. Cree un script que reciba un número por la entrada estandar y muestre `Correcto!` solo si el número es par y menor que 20.

<div style="page-break-before: always"></div>

#Soluciones

1.

{% highlight python %}
n1 = int(input("Primer número: "))
n2 = int(input("Segundo número: "))

print(n1 + n2)
{% endhighlight %}

2.

{% highlight python %}
cadena = input("Primera cadena: ")
segundaCadena = input("Segunda cadena: ")

print(cadena + "\t" + segundaCadena)
{% endhighlight %}

3.

{% highlight python %}
lista = [1, 3, 5, 7]

print("El resultado es:", lista[0] + lista[1] + lista[2] + lista[3])
{% endhighlight %}

4.

{% highlight python %}
resultado = set((1, 2, 2, 2, 3, 3, 4, 4, 5, 5, 5))

print(resultado)
{% endhighlight %}

5.

{% highlight python %}
numero = int(input("Ingrese un número: "))

if numero % 3 == 0:
	print("Es múltiplo de 3")

else:  # Opcional
	print("No es múltiplo de 3")
{% endhighlight %}

6.

{% highlight python %}
numero = int(input("Ingrese un número: "))

if numero % 2 == 0 and numero < 20:
	print("Correcto!")

else:  # Opcional
	print("Incorrecto!")
{% endhighlight %}