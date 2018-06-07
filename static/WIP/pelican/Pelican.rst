#######
Pelican
#######

=====
3.7.0
=====

:Date: 2017-01-01T09:20:00-04:00
:Version: 0.1.0
:Author: Miguel Angel Rivera Notararigo (ntrrg) <ntrrgx@gmail.com>

.. Raw content

.. |br| raw:: html

    <br />

.. |bp| raw:: html

    <div style="page-break-after: always"></div>

.. Syntax highlight

.. role:: txt(code)
    :language: text

.. Custom roles

.. role:: emoji

.. Links

.. _Pelican: http://docs.getpelican.com/en/stable/
.. _reStructuredText: http://docutils.sourceforge.net/rst.html
.. _Markdown: https://es.wikipedia.org/wiki/Markdown
.. _HTML: https://es.wikipedia.org/wiki/HTML
.. _CSS: https://es.wikipedia.org/wiki/Hoja_de_estilos_en_cascada
.. _JavaScript: https://www.javascript.com/
.. _Jinja2: http://jinja.pocoo.org/
.. _Sublime Text 3: http://www.sublimetext.com/3
.. _Python: http://www.python.org/
.. _Virtualenv: https://virtualenv.pypa.io/en/stable/
.. _Chrome: https://www.google.com/chrome/

|br|

.. image:: images/pelican.png
    :alt: Pelican
    :align: center

|br|

Pelican_ es un generador de sitios web estáticos, este tipo de sitios es
llamado así porque una vez que se crean sus archivos, al acceder a ellos, el
contenido que se obtiene siempre será el mismo; no quiere decir que Pelican_
sea aburrido o algo así :emoji:`😅`, en realidad esta es una de sus mejores
cualidades pues:

* No hace falta procesar los archivos cuando alguien accede a ellos (más
  rápido! :emoji:`⚡`).

* No hay que preocuparse mucho por la seguridad (dudo que le hagan inyecciones
  a los :txt:`.rst` :emoji:`😂`).

* Puede alojarse con mucha facilidad.

Además de esto, también está la capacidad de escribir contenido facilmente con
un editor de texto usando reStructuredText_, Markdown_ o HTML_; soporte para
temas con HTML_, CSS_, JavaScript_ y Jinja2_; utilidades de desarrollo y
algunas cosas más que se conocerán usándolo. Fue desarrollado por Alexis
Métaireau y su primera versión estable fue liberada en diciembre del 2010.



|bp|

.. sectnum::
    :suffix: .

.. contents:: Tabla de contenido

|bp|



Herramientas necesarias
=======================

#. Un editor de texto (yo uso `Sublime Text 3`_)
#. Python_ >= 3.4 y virtualenv_ >= 15.1.0
#. Pelican_ 0.13.1 y opcionalmente Markdown_ 2.6.7
#. Un navegador web (yo uso Chrome_)



Pelican
-------

.. class:: os debian

    .. code:: sh

        mkdir ~/Entornos

    .. code:: sh

        virtualenv -p python3 ~/Entornos/pelican

    .. code:: sh

        source ~/Entornos/pelican/bin/activate

    .. code:: sh

        pip install pelican==3.7.0

.. class:: os windows

    .. code:: bat

        virtualenv %USERPROFILE%\Entornos\pelican

    .. code:: bat

        %USERPROFILE%\Entornos\pelican\Scripts\activate

    .. code:: bat

        pip install pelican==3.7.0

Si se escribirá contenido con Markdown_, debe instalarse el módulo de Python_
que se encarga de procesarlo

.. code:: sh
    :class: os debian

    pip install Markdown==2.6.7

.. code:: bat
    :class: os windows

    pip install Markdown==2.6.7



|bp|

Guía
====

__ pelican-quickstart_

Existen dos maneras de crear un proyecto (bueno, las que conozco :emoji:`😅`),
la primera es hacerlo desde cero o como le dicen en inglés *from scratch*; y
la otra es usar el comando :txt:`pelican-quickstart` que genera
automáticamente una estructura básica. Yo tengo la mala costumbre de ser un
fiel amante del *from scratch*, así que esta guía muestra como hacerlo de
dicha manera, pero para los que no tengan esta costumbre, aquí__ se ve como
hacerlo con el comando.

.. code:: text

    .
    ├─ content/
    ├─ pelicanconf.py
    └─ publishconf.py

.. admonition:: Creador de proyectos
    :name: pelican-quickstart

    Hasta este punto la estructura del proyecto es equivalente a la que se
    obtendría con el comando :txt:`pelican-quickstart`, para hacerlo de esta
    manera se ejectuan los siguientes comandos:

    .. class:: os debian

        .. code:: sh

            mkdir ~/NtBlog

        .. code:: sh

            cd ~/NtBlog

        .. code:: sh

            pelican-quickstart

        A partir de aquí se deben responder algunas preguntas:

        #. ¿Dónde se creará el proyecto? :txt:`.`
        #. ¿Cuál será el nombre del sitio? :txt:`NtBlog`
        #. ¿Quién será el autor? :txt:`Miguel Angel Rivera Notararigo`
        #. ¿Cuál será el lenguaje principal? :txt:`es`
        #. ¿Se usará un dominio? :txt:`y`
        #. ¿Cuál es el dominio? :txt:`https://www.ntrrg.com.ve`
        #. ¿Quiere activar la paginación? :txt:`y`
        #. ¿Cuántos artículos se mostrarán por página? :txt:`5`
        #. ¿Cuál zona horaria aplicará para el sitio? :txt:`America/Caracas`

        #. ¿Quiere que se creen archivos para automatizar la generación del
           sitio? :txt:`n`

        #. ¿Quiere agregar algunas utilidades para el desarrollo? :txt:`n`

        Si se quiere conocer más sobre este comando y sus opciones, en la
        sección `Creador de proyectos`_ se explican la mayoría de sus
        utilidades y opciones.

La carpeta :txt:`content/` es la que se encarga de almacenar todos los archivos que serán procesados durante la generación del sitio, si esta carpeta está vacía, se obtendrá un mensaje como :txt:`WARNING: No valid files found in content.`, aunque igual se generará el sitio, solo que no habrá nada que mostrar :emoji:`😅`; para evitar que esto pase, hay que agregar contenido y es suficiente con crear un archivo con uno de los formatos aceptados por Pelican_ (reStructuredText_, Markdown_ o HTML_), por ejemplo:

:txt:`content/ejemplo.rst` (**reStructuredText**):



:txt:`content/ejemplo.md` (**Markdown**):



:txt:`content/ejemplo.html` (**HTML**):



|bp|

Referencia
==========

Línea de comandos
-----------------

Creador de proyectos
++++++++++++++++++++

.. code:: text

    pelican-quickstart [<Opciones>]

-h, --help  Muestra la ayuda.

-p <Ruta>, --path=<Ruta>  Ruta en la que se creará el proyecto. Su valor
                          predefinido es :txt:`.`, que equivale a la carpeta
                          actual.

-t <Título>, --title=<Título>  Título del proyecto. Por defecto :txt:`None`.
-a <Autor>, --author=<Autor>  Autor del proyecto. Por defecto :txt:`None`.
-l <Idioma>, --lang=<Idioma>  Idioma del proyecto. Por defecto :txt:`None`.



Generador de sitios
+++++++++++++++++++

.. code:: text

    pelican [<Opciones>] [<Ruta>]

.. admonition:: Nota

    Si no se especifica :txt:`<Ruta>` se usará por defecto :txt:`content/`
    para obtener el contenido a procesar.

-h, --help  Muestra la ayuda.

-o <Ruta>, --output <Ruta>  Ruta en la que se generará el sitio. Por defecto
                            :txt:`output/`.



|bp|

Atribuciones
============

**Pelican team.** *Pelican 3.7.0 — Pelican 3.7.0 documentation.* http://docs.getpelican.com/en/stable/index.html