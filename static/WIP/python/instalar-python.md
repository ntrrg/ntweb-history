---
title: "Instalar Python 3.5"
description: "Aquí verán como poder instalar Python 3.5 en Debian 8 y en Windows, además de como utilizar el intérprete interactivo que se instala con él."
category: "informatica"
tags: "python instalacion"
type: "article"
postFolder: "Post/3"
---

{{ page.description }}

&nbsp;

###Debian 8:

Con este comando se realiza todo el proceso, pero si quieren pueden hacerlo paso a paso para ver como se hace la instalación.

	# wget https://www.python.org/ftp/python/3.5.0/Python-3.5.0.tar.xz && tar xvf Python-3.5.0.tar.xz && cd Python-3.5.0 && ./configure; make; make test; make install; cd .. && rm -rf Python-3.5.0 Python-3.5.0.tar.xz

&nbsp;

Hay que descargar el código fuente desde la página oficial con el navegador haciendo click [aquí](https://www.python.org/ftp/python/3.5.0/Python-3.5.0.tar.xz) o desde el terminal:

	$ wget https://www.python.org/ftp/python/3.5.0/Python-3.5.0.tar.xz

Cuando termine descomprimimos el paquete:

	$ tar xvf Python-3.5.0.tar.xz

Entramos a la carpeta y empezamos a compilar e instalar (tarda un poquito):

	$ cd Python-3.5.0
	$ ./configure
	$ make
	$ make test
	# make install

Al terminar, pueden borrar los archivos:

	$ cd ..
	$ rm -rf Python-3.5.0 Python-3.5.0.tar.xz

Listo! Para poder usar el intérprete solo deben escribir el comando:

	$ python3.5
	Python 3.5.0 (default, Nov 20 2015, 13:31:09) 
	[GCC 4.9.2] on linux
	Type "help", "copyright", "credits" or "license" for more information.
	>>> 

Para ejecutar los scripts sería igual:

	$ python3.5 script.py

Y para el shebang:

{% highlight python %}
#!/usr/bin/env python3.5
{% endhighlight %}

&nbsp;

###Windows:

Descargamos los ejecutables de [x86](https://www.python.org/ftp/python/3.5.0/python-3.5.0.exe) o [x64](https://www.python.org/ftp/python/3.5.0/python-3.5.0-amd64.exe) según el sistema operativo. Cuando termine la descarga lo ejecutamos, marcamos la opción **Add Python 3.5 to PATH** y hacemos click en **Install Now**:

![Instalación]({{ site.gDriveFolder }}/{{ page.postFolder }}/1.jpg)

Despues de que cargue, hacemos click en **Close** y para abrir el intérprete buscamos la aplicación **Python 3.5**:

![Intérprete]({{ site.gDriveFolder }}/{{ page.postFolder }}/2.jpg)

Para ejecutar los scripts hay dos opciones:

- Desde el SHELL:

	Para poder usar esta opción necesitamos agregar Python a PATH (durante la instalación) o tendremos que escribir la ruta completa al intérprete.

		python script.py

- Desde la aplicación **IDLE**:

	Buscamos la aplicación **IDLE**, al ejecutarla abrirá una ventana en la que al pulsar `Ctrl + O` podemos seleccionar el script y para ejecutarlo pulsamos `F5`.

	![IDLE]({{ site.gDriveFolder }}/{{ page.postFolder }}/3.jpg)