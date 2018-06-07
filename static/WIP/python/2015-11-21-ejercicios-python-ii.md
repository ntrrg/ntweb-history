---
title: "Ejercicios II"
published: false
---

1. Cree un script que pregunte cual es la raiz cúbica de 27 hasta que reciba la respuesta correcta.

2. Cree un script que muestre solo los números divisibles entre 5 del 0 al 50 (incluido el 50).

3. Cree un script que muestre los números primos (2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97) hasta el 100 preguntado en cada repetición si desea continuar mostrandolos, si al final se escriben 25 números debe mostrar `Estos son los números primos menores a 100`.

<div style="page-break-before: always"></div>

#Soluciones

I.

{% highlight python %}
raiz = 0;

while raiz != 3:
	raiz = int(input("Cual es la raiz cúbica de 27? "))
{% endhighlight %}

II.

{% highlight python %}
for numero in range(51):
	if numero % 5 != 0:
		continue

	print(numero)
{% endhighlight %}

III.

{% highlight python %}
primos = (2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97)

for numero in primos:
	print(numero)

	if input("Desea continuar? (s) ") != "s":
		break

else:
	print("Estos son los números primos menores a 100")
{% endhighlight %}