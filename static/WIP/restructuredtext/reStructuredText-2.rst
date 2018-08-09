.. The MIT License (MIT)

.. Copyright (c) 2017 Miguel Angel Rivera Notararigo

.. Permission is hereby granted, free of charge, to any person obtaining a copy
.. of this software and associated documentation files (the "Software"), to deal
.. in the Software without restriction, including without limitation the rights
.. to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
.. copies of the Software, and to permit persons to whom the Software is
.. furnished to do so, subject to the following conditions:

.. The above copyright notice and this permission notice shall be included in all
.. copies or substantial portions of the Software.

.. THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
.. IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
.. FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
.. AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
.. LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
.. OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
.. SOFTWARE.

################
reStructuredText
################

======
0.13.1
======

:Author: Miguel Angel Rivera Notararigo
:Contact: ntrrgx@gmail.com
:Version: 0.5.2
:Date: 2016-12-30 22:00:00 -04:00
:License: `MIT <https://opensource.org/licenses/MIT>`_

.. Raw content

.. role:: raw-html(raw)
    :format: html

.. |br| raw:: html

    <br />

.. |bp| raw:: html

    <div style="page-break-after: always"></div>

.. Syntax highlight

.. role:: html(code)
    :language: html

.. role:: py(code)
    :language: python3

.. role:: rst(code)
    :language: rest

.. role:: sh(code)
    :language: sh

.. role:: txt(code)
    :language: text

.. Custom roles

.. role:: emoji

.. Links

.. _StructuredText: http://www.zope.org/DevHome/Members/jim/StructuredTextWiki/FrontPage
.. _Setext: http://docutils.sourceforge.net/mirror/setext.html
.. _Python Doc-SIG: http://www.python.org/sigs/doc-sig/
.. _Python: http://www.python.org/
.. _Pelican: http://docs.getpelican.com/en/stable/
.. _Sublime Text 3: http://www.sublimetext.com/3
.. _Docutils: http://docutils.sourceforge.net/
.. _Pygments: http://pygments.org/
.. _Debian: https://www.debian.org/
.. _HTML: https://es.wikipedia.org/wiki/HTML
.. _LaTeX: http://www.latex-project.org/
.. _manpages: https://www.kernel.org/doc/man-pages/
.. _ODT: https://es.wikipedia.org/wiki/OpenDocument
.. _XML: https://www.w3.org/XML/
.. _S5: http://meyerweb.com/eric/tools/s5/
.. _XeLaTeX: http://scripts.sil.org/cms/scripts/page.php?site_id=nrsi&id=xetex
.. _PyPI: https://pypi.python.org/pypi
.. _PEPs: https://www.python.org/dev/peps/
.. _RFCs: http://www.faqs.org/rfcs/
.. _Virtualenv: https://pypi.python.org/pypi/virtualenv
.. _Prompt: https://es.wikipedia.org/wiki/Prompt
.. _Chrome: https://www.google.com/chrome/
.. _CSV: https://es.wikipedia.org/wiki/CSV
.. _MS-DOS: https://es.wikipedia.org/wiki/MS-DOS
.. _Windows: https://www.microsoft.com/es-es/windows/
.. _Polyglot XHTML5: http://xmlplease.com/xhtml/xhtml5polyglot/
.. _Web scraping: https://es.wikipedia.org/wiki/Web_scraping

.. image:: images/rst.png
    :class: article-image
    :alt: reStructuredText

__ http://docutils.sourceforge.net/rst.html

Es un lenguaje de marcado, esto quiere decir que se encarga de interpretar
cierta sintaxis para estilizar texto común y corriente con el fin de mejorar
su presentación (o como yo lo diría, *«hacer que se vea mostro»* :emoji:`😄`).
La extensión usada para sus archivos es :txt:`.rst`; fue desarrollado (basado
en StructuredText_ y Setext_) por David Goodger y algunos miembros del `Python
Doc-SIG`_ para utilizarlo como lenguaje oficial de documentación en Python_,
pero a pesar de esto, no es necesario programar para poder usarlo pues
actualmente es usado por algunos generadores de sitios web estáticos (por
ejemplo Pelican_) para crear contenido. Para obtener más información acerca de
*reStructuredText*, es recomendable ir a la `página oficial`__ del proyecto.

|bp|

.. sectnum::
    :suffix: .

.. contents:: Tabla de contenido

|bp|



Herramientas necesarias
=======================

#. Un editor de texto (yo uso `Sublime Text 3`_)
#. Python_ >= 3.4 y virtualenv_ >= 15.1.0
#. Docutils_ 0.13.1 y opcionalmente Pygments_ 2.2.0 para resaltar sintaxis
#. Un navegador web (yo uso Chrome_)



Docutils
--------

.. class:: os os-debian

    .. code:: sh

        mkdir ~/Entornos

    .. code:: sh

        virtualenv -p python3 ~/Entornos/docutils

    .. code:: sh

        source ~/Entornos/docutils/bin/activate

    .. code:: sh

        pip install docutils==0.13.1 Pygments==2.2.0

.. class:: os os-windows

    .. code:: bat

        virtualenv %USERPROFILE%\Entornos\docutils

    .. code:: bat

        %USERPROFILE%\Entornos\docutils\Scripts\activate

    .. code:: bat

        pip install docutils==0.13.1 Pygments==2.2.0



|bp|

Guía
====

Para empezar a trabajar hay que crear un archivo :txt:`.rst` con el editor de
texto

:txt:`archivo.rst`

.. code:: rest
    :number-lines:

    ######
    Título
    ######

    **Hola** *mundo*!

.. admonition:: Nota

    En la sección Referencia_ se puede ver casi toda la sintaxis disponible.

Después se debe procesar el archivo usando Docutils_

.. admonition:: Advertencia
    :class: warning

    El entorno virtual debe estar activo, para verificar esto hay que buscar
    :txt:`(docutils)` en el prompt_ (o el nombre que se le haya puesto)

    .. figure:: images/virtualenv.png
        :alt: Entorno virtual activo
        :align: center

        Entorno virtual activo

    .. figure:: images/no-virtualenv.png
        :alt: Entorno virtual inactivo
        :align: center

        Entorno virtual inactivo

    Si no es así, se debe ejecutar el comando:

    .. code:: sh
        :class: os os-debian

        source ~/Entornos/docutils/bin/activate

    .. code:: bat
        :class: os os-windows

        %USERPROFILE%\Entornos\docutils\Scripts\activate

.. code:: sh
    :class: os os-debian

    rst2html.py archivo.rst archivo.html

.. code:: bat
    :class: os os-windows

    rst2html.py archivo.rst archivo.html

Y para ver los resultados se abre :txt:`archivo.html` con el navegador

.. code:: sh
    :class: os os-debian

    google-chrome archivo.html

.. code:: bat
    :class: os os-windows

    archivo.html



|bp|

Referencia
==========

Títulos
-------

Deben estar al inicio del documento, solo pueden crearse un título y un
subtítulo.

.. admonition:: Sintaxis
    :class: syntax

    Para definirlos se usan los caracteres: :txt:`#`, :txt:`=`, :txt:`-`,
    :txt:`+`, :txt:`*`, :txt:`"`, :txt:`'`, :txt:`~`, :txt:`^`, :txt:`_`,
    :txt:`:`, :txt:`<` o :txt:`>`.

    .. code:: text

        #######
        <texto>
        #######

.. code:: rest

    ######
    Título
    ######

    =========
    Subtítulo
    =========

**Resultado:**

.. image:: images/titulos.png
    :width: 90%
    :align: center



Secciones
---------

.. admonition:: Sintaxis
    :class: syntax

    Para definirlas se usan los caracteres: :txt:`#`, :txt:`=`, :txt:`-`,
    :txt:`+`, :txt:`*`, :txt:`"`, :txt:`'`, :txt:`~`, :txt:`^`, :txt:`_`,
    :txt:`:`, :txt:`<` o :txt:`>`.

    .. code:: text

        <texto>
        =======

.. code:: rest

    ######
    Título
    ######

    =========
    Subtítulo
    =========

    Sección 1:
    ==========

    Sección 1.1:
    ------------

    Sección 2:
    ==========

**Resultado:**

.. image:: images/secciones.png
    :width: 90%
    :align: center



Tablas de contenido
-------------------

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        .. contents:: [<title>]
            [:backlinks: top | entry | none]
            [:depth: <profundidad>]
            [:name: <id>]
            [:class: <clase[ ...]>]

    Parámetros:
      :title: título de la tabla de contenido. Su valor predefinido es
              :txt:`Contents`.

    Opciones:
      :backlinks: permite agregar enlaces a los títulos de las secciones que
                  apuntan a la tabla de contenido. Soporta los valores
                  :txt:`top` que dirige al inicio, :txt:`entry` que dirige al
                  elemento (predeterminado) y :txt:`none` que deshabilita los
                  enlaces.

      :depth: profundidad de la tabla de contenido. Por defecto muestra todas
              las secciones.

      :name: identificador HTML_.
      :class: clases HTML_ separadas por espacios.

.. code:: rest

    .. contents:: Tabla de contenido
        :depth: 1

**Resultado:**

.. image:: images/tabla-contenido.png

Si se quieren enumerar las secciones, se puede usar la directiva
:txt:`sectnum`

.. _sectnum-directive:

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        .. sectnum::
            [:start: <número inicial>]
            [:depth: <profundidad>]
            [:prefix: <prefijo>]
            [:suffix: <sufijo>]

    Opciones:
      :start: número en el que inicia. Su valor predefinido es :txt:`1`.

      :depth: profundidad de secciones a enumerar. Por defecto enumera todas
              las secciones.

      :prefix: texto a agregar antes de la numeración.
      :suffix: texto a agregar después de la numeración.

.. code:: rest

    .. sectnum::
        :suffix: .

    .. contents:: Tabla de contenido
        :depth: 1

**Resultado:**

.. image:: images/tabla-contenido-num.png



Párrafos
--------

.. code:: rest

    Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
    tempor incididunt ut labore et dolore magna aliqua.

    Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
    aliquip ex ea commodo consequat.

**Resultado:**

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua.

Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat.



Transiciones
------------

.. code:: rest

    Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
    tempor incididunt ut labore et dolore magna aliqua.

    ----

    Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
    aliquip ex ea commodo consequat.

**Resultado:**

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua.

----

Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat.

.. admonition:: Advertencia
    :class: warning

    No pueden usarse justo después de las secciones.



Cursiva
-------

.. code:: rest

    *Ntrrg*

**Resultado:**

*Ntrrg*



Negrita
-------

.. code:: rest

    **Ntrrg**

**Resultado:**

**Ntrrg**



Enlaces
-------

Directos
++++++++

.. code:: rest

    http://www.ntrrg.com.ve

    ntrrgx@gmail.com

**Resultado:**

http://www.ntrrg.com.ve

ntrrgx@gmail.com



Referencias
+++++++++++

.. admonition:: Sintaxis (definición)
    :class: syntax

    .. code:: text

        .. _<identificador>: <URL>

    Parámetros:
      :URL: ruta.

.. admonition:: Sintaxis (uso)
    :class: syntax

    .. code:: text

        <identificador>_

    En caso de que el identificador esté compuesto por varias palabras se debe
    encerrar entre comillas invertidas.

.. code:: rest

    .. _página: http://www.ntrrg.com.ve
    .. _correo electrónico: mailto:ntrrgx@gmail.com

    Página_

    `Correo electrónico`_

    `Referencia directa <http://www.ntrrg.com.ve>`_

**Resultado:**

.. _página: http://www.ntrrg.com.ve
.. _correo electrónico: mailto:ntrrgx@gmail.com

Página_

`Correo electrónico`_

`Referencia directa <http://www.ntrrg.com.ve>`_



Referencias internas
++++++++++++++++++++

Permiten referenciar elementos del documento.

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        <identificador de un elemento>_

    En caso de que el identificador esté compuesto por varias palabras se debe
    encerrar entre comillas invertidas.

.. code:: rest

    `Referencias internas`_

    `referencias internas`_

    `REFERENCIAS INTERNAS`_

**Resultado:**

`Referencias internas`_

`referencias internas`_

`REFERENCIAS INTERNAS`_



Referencias internas vacías
+++++++++++++++++++++++++++

Permiten crear anclajes en el documento.

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        <identificador>_

        .. _<identificador>:

    En caso de que el identificador esté compuesto por varias palabras se debe
    encerrar entre comillas invertidas.

.. code:: rest

    `Referencia vacía`_

    Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
    tempor incididunt ut labore et dolore magna aliqua.

    .. _Referencia vacía:

    Objetivo de referencia vacía

    Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
    aliquip ex ea commodo consequat.

**Resultado:**

`Referencia vacía`_

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua.

.. _Referencia vacía:

Objetivo de referencia vacía

Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat.



Referencias anónimas
++++++++++++++++++++

Permiten crear apuntadores a enlaces.

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        __ <enlace>

        <texto>__

    * Deben existir tantas definiciones como referencias, es decir, por cada
      :rst:`__ <enlace>` debe existir un :rst:`<texto>__`.

    * El orden de las definiciones determinan el orden en que se usarán.

.. code:: rest

    .. _Python: http://www.python.org/
    __ Python_
    __ http://docutils.sourceforge.net/rst.html

    `Python 3`__ y reStructuredText__ hacen muy buen equipo!

**Resultado:**

__ Python_
__ http://docutils.sourceforge.net/rst.html

`Python 3`__ y reStructuredText__ hacen muy buen equipo!



Imágenes
--------

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        .. image:: <URL>
            [:alt: <texto>]
            [:height: <altura>]
            [:width: <anchura>]
            [:align: left | center | right]
            [:target: <enlace>]
            [:name: <id>]
            [:class: <clase[ ...]>]

    Parámetros:
      :URL: ruta de la imagen.

    Opciones:
      :alt: texto a mostrar en caso de que no se pueda cargar la imagen.
      :height: altura.
      :width: anchura.
      :align: alineación horizontal.
      :target: objeto de enlace que se activa al hacer clic sobre la imagen.
      :name: identificador HTML_.
      :class: clases HTML_ separadas por espacios.

.. code:: rest

    .. image:: images/luffy.jpg
        :alt: Monkey D. Luffy
        :height: 200px
        :width: 200px
        :align: center
        :target: Luffy_
        :name: Luffy
        :class: chibi

**Resultado:**

.. image:: images/luffy.jpg
    :alt: Monkey D. Luffy
    :height: 200px
    :width: 200px
    :align: center
    :target: Luffy_
    :name: Luffy
    :class: chibi



Figuras
-------

Las figuras son imágenes que poseen un título y una leyenda (descripción).

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        .. figure:: <URL>
            [:alt: <texto>]
            [:height: <altura>]
            [:width: <anchura>]
            [:figwidth: <anchura>]
            [:align: left | center | right]
            [:target: <enlace>]
            [:name: <id>]
            [:class: <clase[ ...]>]
            [:figclass: <clase[ ...>]

            [<título>]

            [<leyenda>]

    Parámetros:
      :URL: ruta de la imagen.

    Opciones:
      :alt: texto a mostrar en caso de que no se pueda cargar la imagen.
      :height: altura.
      :width: anchura.
      :figwidth: anchura del contenedor.
      :align: alineación horizontal.
      :target: objeto de enlace que se activa al hacer clic sobre la imagen.
      :name: identificador HTML_.
      :class: clases HTML_ separadas por espacios.
      :figclass: clases HTML_ del contenedor separadas por espacios.

.. code:: rest

    .. figure:: images/luffy.jpg
        :alt: Monkey D. Luffy
        :height: 200px
        :width: 200px
        :figwidth: 400px
        :align: center
        :target: `Figura Luffy`_
        :name: Figura Luffy
        :class: chibi
        :figclass: anime

        Monkey D. Luffy

        Integrante de la tripulación de los **Mugiwara**.

        +-------------+
        | Recompensas |
        +=============+
        | 500.000.000 |
        +-------------+
        | 400.000.000 |
        +-------------+
        | 300.000.000 |
        +-------------+
        | 100.000.000 |
        +-------------+
        |  30.000.000 |
        +-------------+

**Resultado:**

.. figure:: images/luffy.jpg
    :alt: Monkey D. Luffy
    :height: 200px
    :width: 200px
    :figwidth: 400px
    :align: center
    :target: `Figura Luffy`_
    :name: Figura Luffy
    :class: chibi
    :figclass: anime

    Monkey D. Luffy

    Integrante de la tripulación de los **Mugiwara**.

    +-------------+
    | Recompensas |
    +=============+
    | 500.000.000 |
    +-------------+
    | 400.000.000 |
    +-------------+
    | 300.000.000 |
    +-------------+
    | 100.000.000 |
    +-------------+
    |  30.000.000 |
    +-------------+



Listas
------

Desordenadas
++++++++++++

.. admonition:: Sintaxis
    :class: syntax

    Para definirlas se usan los caracteres: :txt:`*`, :txt:`+` y :txt:`-`.

    .. code:: text

        * <elemento>

.. code:: rest

    * Primero
    * Segundo
      con más texto

      Y otro párrafo
    * Tercero

**Resultado:**

* Primero
* Segundo
  con más texto

  Y otro párrafo
* Tercero



Ordenadas
+++++++++

.. admonition:: Sintaxis
    :class: syntax

    Pueden definirse listas enumeradas simples, con números romanos
    (minúsculas y mayúsculas) y letras (minúsculas y mayúsculas).

    .. code:: text

        #. <elemento>

.. code:: rest

    #. Primero
       con más texto
    #. Segundo

    1. Primero
       con más texto
    #. Segundo

    i. Primero
       con más texto
    #. Segundo

    I. Primero
       con más texto
    #. Segundo

    a. Primero
       con más texto
    #. Segundo

    A. Primero
       con más texto
    #. Segundo

**Resultado:**

#. Primero
   con más texto
#. Segundo

1. Primero
   con más texto
#. Segundo

i. Primero
   con más texto
#. Segundo

I. Primero
   con más texto
#. Segundo

a. Primero
   con más texto
#. Segundo

A. Primero
   con más texto
#. Segundo



Definición
++++++++++

Permiten crear pares de *concepto-definición* como un diccionario.

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        <concepto>
          <definición>

.. code:: rest

    Concepto
      Definición

**Resultado:**

Concepto
  Definición



Tablas
------

Simples
+++++++

.. code:: rest

    +----------+---------------+
    | Cabecera | Otra cabecera |
    +==========+===============+
    |          |    Celda 2    |
    | Celda 1  +---------------+
    |          |    Celda 3    |
    +----------+---------------+
    |         Celda 4          |
    +--------------------------+

**Resultado:**

+----------+---------------+
| Cabecera | Otra cabecera |
+==========+===============+
|          |    Celda 2    |
| Celda 1  +---------------+
|          |    Celda 3    |
+----------+---------------+
|         Celda 4          |
+--------------------------+



Listas
++++++

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        .. list-table:: [<title>]
            [:widths: <anchura>[ ...]]
            [:header-rows: <cantidad>]
            [:stub-columns: <cantidad>]
            [:name: <id HTML>]
            [:class: <clases HTML>]

            * - <celda 1-1>
              - <celda 1-2>

    Parámetros;
      :title: título de la tabla.

    Opciones:
      :widths: anchura de las celdas, por cada celda debe haber una anchura.
      :header-rows: cantidad de filas que forman parte de la cabecera.
      :stub-columns: cantidad de columnas a resaltar.
      :name: identificador HTML_.
      :class: clases HTML_ separadas por espacios.

.. code:: rest

    .. list-table:: OVAs Hellsing
        :widths: 15 10
        :header-rows: 1
        :stub-columns: 1

        * - Nombre
          - Duración

        * - OVA 1
          - 50:25

        * - OVA 2
          - 51:48

        * - OVA 3
          - 49:17

**Resultado:**

.. list-table:: OVAs Hellsing
    :widths: 15 10
    :header-rows: 1
    :stub-columns: 1

    * - Nombre
      - Duración

    * - OVA 1
      - 50:25

    * - OVA 2
      - 51:48

    * - OVA 3
      - 49:17



CSV
+++

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        .. csv-table:: [<title>]
            [:widths: <anchura>[ ...]]
            [:header: <cabecera>]
            [:header-rows: <cantidad>]
            [:stub-columns: <cantidad>]
            [:delim: <caracter>]
            [:quote: <caracter>]
            [:escape: <caracter>]
            [:file: <ruta archivo local>]
            [:url: <ruta archivo remoto>]
            [:name: <id HTML>]
            [:class: <clases HTML>]

            [<datos en formato CSV>]

    Parámetros;
      :title: título de la tabla.

    Opciones:
      :widths: anchura de las celdas, por cada celda debe haber una anchura.
      :header: cabecera en formato CSV_.

      :header-rows: cantidad de filas que forman parte de la cabecera
                    después de :rst:`:header:`.

      :stub-columns: cantidad de columnas a resaltar.
      :delim: caracter delimitador de campos. Por defecto es :txt:`,`.

      :quote: caracter para definir cadenas de caracteres. Por defecto es
              :txt:`"`.

      :escape: caracter para escapar. Por defecto es ``\``.
      :file: archivo local con datos CSV_.
      :url: archivo remoto con datos CSV_.
      :name: identificador HTML_.
      :class: clases HTML_ separadas por espacios.

.. code:: rest

    .. csv-table:: Título
        :header: "Cabecera 1","Cabecera 2"

        "Celda 1-1","Celda 1-2"
        "Celda 2-1","Celda 2-2"
        "Celda 3-1","Celda 3-2"
        "Celda 4-1","Celda 4-2"
        "Celda 5-1","Celda 5-2"

**Resultado:**

.. csv-table:: Título
    :header: "Cabecera 1","Cabecera 2"

    "Celda 1-1","Celda 1-2"
    "Celda 2-1","Celda 2-2"
    "Celda 3-1","Celda 3-2"
    "Celda 4-1","Celda 4-2"
    "Celda 5-1","Celda 5-2"



Campos
++++++

Field name "Author": author element.
"Authors": authors.
"Organization": organization.
"Contact": contact.
"Address": address.
"Version": version.
"Status": status.
"Date": date.
"Copyright": copyright.
"Dedication": topic.
"Abstract": topic.

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        :<campo>: <descripción>

    __ reStructuredText_

    Si se escriben justo después del título o el subtítulo, establecerán
    metadatos del documento (como en el inicio__ de éste artículo)

.. code:: rest

    :Author: Miguel Angel Rivera Notararigo (ntrrg) <ntrrgx@gmail.com>
    :Licence: MIT
    :Version: 0.5.2
    :Date: 2016-12-30

**Resultado:**

:Author: Miguel Angel Rivera Notararigo (ntrrg) <ntrrgx@gmail.com>
:Licence: MIT
:Version: 0.5.2
:Date: 2016-12-30



Opciones
++++++++

.. code:: rest

    -a  Corta
    -b <arg>  Con un argumento
    -c <arg[,...]>  Con varios argumentos
    --a-larga  Larga
    --b-larga=<arg>  Con un argumento
    --b-larga=<arg[,...]>  Con varios argumentos
    -o, --opcion-doble  Doble
    /o  Estilo MS-DOS

**Resultado:**

-a  Corta
-b <arg>  Con un argumento
-c <arg[,...]>  Con varios argumentos
--a-larga  Larga
--b-larga=<arg>  Con un argumento
--b-larga=<arg[,...]>  Con varios argumentos
-o, --opcion-doble  Doble
/o  Estilo MS-DOS_



Citas
-----

Bloque
++++++

.. admonition:: Sintaxis
    :class: syntax

    Deben estar cuatro espacios delante del elemento anterior.

    .. code:: text

            <cuerpo>

            [-- <autor>]

        ..

    Los últimos dos puntos (:rst:`..`) no son necesarios, aquí se usan para
    que **reStructuredText** reconozca los espacios en blanco a la izquierda
    sin contenido extra en el bloque de código, si no, se vería así:

    .. code:: text

            <cuerpo>

            [-- <autor>]

.. code:: rest

        Cuerpo de la cita

        -- Autor

    ..

**Resultado:**

    Cuerpo de la cita

    -- Autor

|br|

Es posible crear citas dentro de citas:

.. code:: rest

        Cita con autor

            Cita dentro de cita

        -- Autor

    ..

**Resultado:**

    Cita con autor

        Cita dentro de cita

    -- Autor



Notas de pie
++++++++++++

.. admonition:: Sintaxis
    :class: syntax

    En este caso, los corchetes (:txt:`[]`) no representan condicionalidad.

    .. code:: text

        [# | *]_

        .. [# | *] <descripción>

.. code:: rest

    Python [#]_ y reStructuredText[#]_.

    .. [#] Lenguaje de programación
    .. [#] Lenguaje de documentanción

    Notas con símbolos [*]_

    .. [*] Se pueden usar símbolos para las notas de pie

**Resultado:**

Python [#]_ y reStructuredText [#]_.

.. [#] Lenguaje de programación
.. [#] Lenguaje de documentanción

Notas con símbolos [*]_

.. [*] Se pueden usar símbolos para las notas de pie



Citas de pie
++++++++++++

.. admonition:: Sintaxis
    :class: syntax

    En este caso, los corchetes (:txt:`[]`) no representan condicionalidad.

    .. code:: text

        [<identificador>]_

        .. [<identificador>] <descripción>

.. code:: rest

    [Py]_ y [reST]_.

    .. [Py] Python
    .. [reST] reStructuredText

**Resultado:**

[Py]_ y [reST]_.

.. [Py] Python
.. [reST] reStructuredText



Fórmulas
--------

:math:`\times 10e^\infty`

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        .. math::
            [:name: <id>]
            [:class: <clase[ ...]>]

            <fórmula en sintaxis LaTeX>

.. code:: rest

    .. math::

        E=mc^2

**Resultado:**

.. math::

    E=mc^2



Bloques lineales
----------------

.. code:: rest

    | Lo bloques lineales son útiles para
    |     definir direcciones y versos pues
    |         mantienen los saltos de líneas e indentaciones.
    |
    | Para mantener las líneas largas solo
      hay que iniciar al nivel del bloque
      y sin agregar el caracter :txt:`|`.

**Resultado:**

| Lo bloques lineales son útiles para
|     definir direcciones y versos pues
|         mantienen los saltos de líneas e indentaciones.
|
| Para mantener las líneas largas solo
  hay que iniciar al nivel del bloque
  y sin agregar el caracter :txt:`|`.



Avisos
------

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        .. admonition:: <title>
            [:name: <id>]
            [:class: <clase[ ...]>]

            <cuerpo del aviso>

    Parámetros;
      :title: título del aviso.

    Opciones:
      :name: identificador HTML_.
      :class: clases HTML_ separadas por espacios.

    Existen clases predefinidas que personalizan los avisos, algunas de ellas
    son:

    * :txt:`attention`
    * :txt:`caution`
    * :txt:`danger`
    * :txt:`error`
    * :txt:`hint`
    * :txt:`important`
    * :txt:`note`
    * :txt:`tip`
    * :txt:`warning`

.. code:: rest

    .. admonition:: Título

        Cuerpo del aviso

**Resultado:**

.. admonition:: Título

    Cuerpo del aviso



Barras laterales
----------------

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        .. sidebar:: <title>
            [:subtitle: <texto>]
            [:name: <id>]
            [:class: <clase[ ...]>]

            <cuerpo>

    Parámetros;
      :title: título de la barra lateral.

    Opciones:
      :subtitle: subtítulo de la barra lateral.
      :name: identificador HTML_.
      :class: clases HTML_ separadas por espacios.

.. code:: rest

    .. sidebar:: Título

        Cuerpo de la barra lateral

**Resultado:**

.. sidebar:: Título

    Cuerpo de la barra lateral

|br|
|br|
|br|



Literales
---------

Lineales
++++++++

.. code:: rest

    Así se escribe en la salida estándar con Python: ``print("Hola mundo!")``.

**Resultado:**

Así se escribe en la salida estándar con Python: ``print("Hola mundo!")``.



Bloque
++++++

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        [<texto>]::

            <código>

    Si no se especifica :txt:`<texto>` no se mostrarán los :txt:`:`.

.. code:: rest

    Código::

        def cadena():
            return "Python en reStructuredText"

        print(cadena())

**Resultado:**

Código::

    def cadena():
        return "Python en reStructuredText"

    print(cadena())



Resaltar sintaxis
-----------------

Lineal
++++++

.. admonition:: Advertencia
    :class: warning

    Pygments_ debe estar instalado.

.. admonition:: Sintaxis (deficinición)
    :class: syntax

    .. code:: text

        .. role:: <name>(code)
            :language: <lenguaje>

    Parámetros:
      :name: identificador del rol.

    Opciones:
      __ http://pygments.org/languages/

      :language: lenguaje usado, debe estar en la lista de `lenguajes
                 soportados`__ por Pygments_; también se puede usar el comando
                 :sh:`pygmentize -L lexer` para obtener la lista completa o
                 :sh:`pygmentize -L lexer | grep -iA 1 <lenguaje>` para uno en
                 específico (este último no sirve en Windows_).

.. admonition:: Sintaxis (uso)
    :class: syntax

    .. code:: text

        :<name>:`<código>`

.. code:: rest

    .. role:: py(code)
        :language: python3

    Así se escribe en la salida estándar con Python:
    :py:`print("Hola mundo!")`.

**Resultado:**

Así se escribe en la salida estándar con Python:
:py:`print("Hola mundo!")`.



Bloque
++++++

.. admonition:: Advertencia
    :class: warning

    Pygments_ debe estar instalado.

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        .. code:: <language>
            [:number-lines: [<número inicial>]]

            <código>

    Parámetros:
      __ http://pygments.org/languages/

      :language: lenguaje usado, debe estar en la lista de `lenguajes
                 soportados`__ por Pygments_; también se puede usar el comando
                 :sh:`pygmentize -L lexer` para obtener la lista completa o
                 :sh:`pygmentize -L lexer | grep -iA 1 <lenguaje>` para uno en
                 específico (este último no sirve en Windows_).

    Opciones:
      :number-lines: enumera las líneas, puede recibir el número en que
                     iniciará.

.. code:: rest

    .. code:: python3
        :number-lines:

        def cadena():
            return "Código Python con sintaxis resaltada"

        print(cadena())

**Resultado:**

.. code:: python3
    :number-lines:

    def cadena():
        return "Código Python con sintaxis resaltada"

    print(cadena())



Archivos
++++++++

.. admonition:: Advertencia
    :class: warning

    Pygments_ debe estar instalado.

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        .. include:: <URL>
            :code: <lenguaje>
            [:start-line: <número de línea>]
            [:end-line: <número de línea>]
            [:start-after: <texto>]
            [:end-before: <texto>]
            [:number-lines: <número inicial>]
            [:tab-width: <cantidad de caracteres>]

    Opciones:
      __ http://pygments.org/languages/

      :code: lenguaje usado, debe estar en la lista de `lenguajes
             soportados`__ por Pygments_; también se puede usar el comando
             :sh:`pygmentize -L lexer` para obtener la lista completa o
             :sh:`pygmentize -L lexer | grep -iA 1 <lenguaje>` para uno en
             específico (este último no sirve en Windows_).

      :start-line: línea desde la que mostrar contenido.

      :end-line: línea hasta la que mostrar contenido (no incluye la línea
                 especificada).

      :start-after: texto desde el que mostrar contenido.

      :end-before: texto hasta el que mostrar contenido.

      :number-lines: enumera las líneas, puede recibir el número en que
                     iniciará.

      :tab-width: cantidad de caracteres para las tabulaciones.

.. code:: rest

    .. include:: prueba.py
        :code: python3
        :number-lines:

**Resultado:**

.. include:: prueba.py
    :code: python3
    :number-lines:



Contenido no procesado
----------------------

Lineal
++++++

.. admonition:: Sintaxis (definición)
    :class: syntax

    .. code:: text

        .. role:: raw-<name>(raw)
            :format: <formato>

    Opciones:
      :format: formato original del contenido.

.. admonition:: Sintaxis (uso)
    :class: syntax

    .. code:: text

        :raw-<name>:`<contenido>`

.. code:: rest

    .. role:: raw-html(raw)
        :format: html

    Hola :raw-html:`<strong>mundo</strong>!`

**Resultado:**

Hola :raw-html:`<strong>mundo</strong>!`



Bloque
++++++

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        .. raw:: <format>
            [:file: <ruta archivo local>]
            [:url: <ruta archivo remoto>]

            [<contenido>]

    Parámetros:
      :format: formato original del contenido.

    Opciones:
      :file: lee el contenido desde un archivo local.
      :url: lee el contenido desde un archivo remoto.

.. code:: rest

    .. raw:: html

        <h1 class="align-center">Título no procesado</h1>

**Resultado:**

.. raw:: html

    <h1 class="align-center">Título no procesado</h1>



Clases
------

.. admonition:: Sintaxis
    :class: syntax

    .. code:: text

        .. class:: <clase[ ...]>

        <objetivo de la clase>

    Para aplicar a múltiples elementos se debe indentar:

    .. code:: text

        .. class:: <clase[ ...]>

            <objetivo de la clase>

            <otro objetivo de la clase>

.. code:: rest

    .. class:: align-center

    Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
    tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
    quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
    consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
    cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
    proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

**Resultado:**

.. class:: align-center

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.



Roles
-----

.. admonition:: Sintaxis (definición)
    :class: syntax

    .. code:: text

        .. role:: <name>[(<rol padre>)]
            [:class: <clase[ ...]>]

    * El nombre del rol (:txt:`<name>`) será aplicado como una clase al
      elemento; si se especifica la opción :rst:`:class:`, las clases
      definidas allí serán aplicadas al elemento en lugar del nombre del rol.

    * Los roles predefinidos son:

      * :rst:`:emphasis:`: Es equivalente a :rst:`*texto*`.
      * :rst:`:strong:`: Es equivalente a :rst:`**texto**`.

      * :rst:`:title:`: Se usa para citar nombres de libros o marcas
        registradas, es equivalente a ```texto```.

      * :rst:`:sub:`: Permite escribir subíndices.
      * :rst:`:sup:`: Permite escribir superíndices.
      * :rst:`:literal:`: Es equivalente a ````texto````.
      * :rst:`:math:`: Genera fórmulas a partir de sintaxis LaTeX_.

      * :rst:`:PEP:`: Permite hacer referencia a las PEPs_ especificando su
        número.

      * :rst:`:RFC:`: Permite hacer referencia a las RFCs_ especificando su
        número.

      * :rst:`:code:`: Resalta la sintaxis de código fuente, ver la sección
        `Resaltar Sintaxis`_.

      * :rst:`:raw:`: Evita que el contenido sea procesado por Docutils_, ver
        la sección `Contenido no procesado`_.

.. admonition:: Sintaxis (uso)
    :class: syntax

    .. code:: text

        :<rol>:`<name>`



Substituciones
--------------

.. admonition:: Sintaxis (definición)
    :class: syntax

    .. code:: text

        .. |<identificador>| <directiva>

.. admonition:: Sintaxis (uso)
    :class: syntax

    .. code:: text

        |<identificador|

.. code:: rest

    .. |reST| replace:: reStructuredText

    |reST|

    |reST|

    .. |img| image:: images/luffy.jpg
        :alt: Monkey D. Luffy
        :height: 200px
        :width: 200px

    |img|

    .. |nick| raw:: html

        <h1>Ntrrg</h1>

    |nick|

**Resultado:**

.. |reST| replace:: reStructuredText

|reST|

|reST|

.. |img| image:: images/luffy.jpg
    :alt: Monkey D. Luffy
    :height: 200px
    :width: 200px

|img|

.. |nick| raw:: html

    <h1>Ntrrg</h1>

|nick|



Línea de comandos
-----------------

.. code:: text

    rst2html5.py [<opciones>] [<origen>.rst [<destino>.html]]

Los valores predefinidos para :txt:`<origen>.rst` y :txt:`<destino>.html` son
la entrada estándar (*stdin*) y la salida estándar (*stdout*) respectivamente,
por esta razón es opcional especificarlos.

.. admonition:: Nota

    Aunque este artículo está enfocado a :txt:`rst2html5.py`, Docutils_
    permite obtener varios tipos de archivos a partir de un :txt:`.rst`, para
    los que sufran de curiosidad académica, esta es la lista:

    * HTML_ (:txt:`rst2html5.py`)
    * LaTeX_ (:txt:`rst2latex.py`)
    * manpages_ (:txt:`rst2man.py`)
    * ODT_ (:txt:`rst2odt.py`)
    * pseudo-XML (:txt:`rst2pseudoxml.py`)
    * XML_ (:txt:`rst2xml.py`)
    * S5_ (:txt:`rst2s5.py`)
    * XeLaTeX_ (:txt:`rst2xetex.py`)

Opciones
++++++++

--title=<título>
    Configura el título del documento como metadata.

.. class:: exclusive

    --generator, -g
        Agrega información del generador al final del documento.

    --no-generator
        No muestra información del generador (predeterminado).

.. class:: exclusive

    --date, -d
        Agrega la fecha (UTC) al final del documento.

    --time, -t
        Agrega la fecha y la hora (UTC) al final del documento, implica
        :txt:`--date, -d`.

    --no-datestamp
        No muestra la fecha y la hora (predeterminado).

.. class:: exclusive

    --source-link, -s
        Agrega un enlace al código fuente (:txt:`.rst`) al final del
        documento.

    --source-url=<URL>
        URL del código fuente, implica :txt:`--source-link, -s`.

    --no-source-link
        No muestra el enlace del código fuente (predeterminado).

.. class:: exclusive

    __ `Tablas de contenido`_
    __ `Tablas de contenido`_
    __ `Tablas de contenido`_

    --toc-entry-backlinks
        Los enlaces de los títulos de las secciones apuntan a su elemento en
        la tabla de contenido, también puede configurarse en su definición__
        (predeterminado).

    --toc-top-backlinks
        Los enlaces de los títulos de las secciones apuntan al inicio de la
        tabla de contenido, también puede configurarse en su definición__.

    --no-toc-backlinks
        Elimina los enlaces de los títulos de las secciones, también puede
        configurarse en su definición__.

.. class:: exclusive

    --footnote-backlinks
        Agrega enlaces entre las notas de pie (también las citas de pie) y su
        referencia (predeterminado).

    --no-footnote-backlinks
        Elimina los enlaces entre las notas de pie (también las citas de pie)
        y su referencia.

.. class:: exclusive

    --section-numbering
        Enumera las secciones (predeterminado).

    --no-section-numbering
        No enumera las secciones.

.. admonition:: Bug
    :class: bug

    __ sectnum-directive_

    La opción :txt:`--section-numbering` no funciona actualmente, pero puede
    usarse la directiva :txt:`sectnum` (`ver sintaxis`__).

.. class:: exclusive

    --strip-comments
        No procesa los comentarios de :txt:`<origen>.rst`.

    --leave-comments
        Procesa los comentarios de :txt:`<origen>.rst` (predeterminado).

--strip-elements-with-class=<clase[,...]>
    Elimina los elementos que tengan las clases :txt:`<clase[,...]>`.

--strip-class=<clase[,...]>
    Elimina las clases :txt:`<clase[,...]>` de todos los elementos.

--warnings=<archivo>
    Guarda los mensaje obtenidos durante la ejecución del write en el archivo
    :txt:`<archivo>`.

.. class:: exclusive

    --report=<nivel>, -r <nivel>
        Muestra los mensajes generados por el writer a partir del nivel
        :txt:`<nivel>`, admite los valores :txt:`info`, :txt:`warning`
        (predeterminado), :txt:`error`, :txt:`severe` o :txt:`none`.

    --verbose, -v
        Equivale a :txt:`--report=info, -r info`.

    --quiet, -q
        Equivale a :txt:`--report=none, -r none`.

.. class:: exclusive

    --halt=<nivel>
        Detiene la ejecución del writer si se produce algún mensaje de nivel
        :txt:`<nivel>`.

    --strict
        Detiene la ejecución del writer si se produce cualquier mensaje,
        equivale a :txt:`--halt=info`.

.. class:: exclusive

    --debug
        Muestra mensajes de depuración por la salida de error estandar
        (*stderr*).

    --no-debug
        No muestra mensaje de depuración (predeterminado).

.. class:: exclusive

    --traceback
        Muestra las excepciones de Python_ que se levanten durante la
        ejecución del writer.

    --no-traceback
        No muestra las excepciones de Python_ que se levanten durante la
        ejecución del writer (predeterminado).

--input-encoding=<codificación[:manejador]>, -i <codificación[:manejador]>
    Codificación de :txt:`<origen>.rst` o la entrada estándar (*stdin*) en su
    defecto; :txt:`<codificación>` admite los valores :txt:`ASCII`,
    :txt:`UTF-8`, :txt:`UTF-16` o :txt:`UTF-32`; :txt:`<manejador>` admite los
    valores :txt:`strict` (predeterminado), :txt:`ignore` o :txt:`replace`.

--input-encoding-error-handler=<manejador>
    Manejador que se usará si se encuentran errores de codificación procesando
    :txt:`<origen>.rst` o la entrada estándar (*stdin*) en su defecto, admite
    los valores :txt:`strict` (predeterminado), :txt:`ignore` o
    :txt:`replace`.

--output-encoding=<codificación[:manejador]>, -o <codificación[:manejador]>
    Codificación de :txt:`<destino>.html` o la salida estándar (*stdout*) en
    su defecto; :txt:`<codificación>` admite los valores :txt:`ASCII`,
    :txt:`UTF-8` (predeterminado), :txt:`UTF-16` o :txt:`UTF-32`;
    :txt:`<manejador>` admite los valores :txt:`strict` (predeterminado),
    :txt:`ignore`, :txt:`replace`, :txt:`xmlcharrefreplace` o
    :txt:`backslashreplace`.

--output-encoding-error-handler=<manejador>
    Manejador se usará si se encuentran errores de codificación procesando
    :txt:`<destino>.html` o la salida estándar (*stdout*) en su defecto,
    admite los valores :txt:`strict` (predeterminado), :txt:`ignore`,
    :txt:`replace`, :txt:`xmlcharrefreplace` o :txt:`backslashreplace`.

--error-encoding=<codificación[:manejador]>, -e <codificación[:manejador]>
    Codificación de la salida de error estándar (*stderr*);
    :txt:`<codificación>` admite los valores :txt:`ASCII`, :txt:`UTF-8`
    (predeterminado), :txt:`UTF-16` o :txt:`UTF-32`; :txt:`<manejador>` admite
    los valores :txt:`strict`, :txt:`ignore`, :txt:`replace`,
    :txt:`xmlcharrefreplace` o :txt:`backslashreplace` (predeterminado).

--error-encoding-error-handler=<manejador>
    Manejador se usará si se encuentran errores de codificación procesando la
    salida de error estándar (*stderr*), admite los valores :txt:`strict`,
    :txt:`ignore`, :txt:`replace`, :txt:`xmlcharrefreplace` o
    :txt:`backslashreplace` (predeterminado).

__ https://tools.ietf.org/html/bcp47

--language=<idioma>, -l <idioma>
    Idioma del documento (según la `BCP 47`__), su valor predeterminado es
    :txt:`en`.

--record-dependencies=<archivo>
    Escribe las dependencias para procesar :txt:`<origen>.rst` en
    :txt:`<archivo>`.

__ http://docutils.sourceforge.net/docs/user/config.html#configuration-file-syntax

--config=<archivo>
    Lee la configuración desde :txt:`<archivo>`, debe seguir esta__ sintaxis.

--version, -V
    Muestra la versión de Docutils_.

--help, -h
    Muestra la ayuda del writer.

|br|

--pep-references
    Busca y procesa las PEP en el documento, por ejemplo: :txt:`PEP 258`.

--pep-base-url=<URL>
    Ruta en la que encontrar las PEP, su valor predeterminado es
    :txt:`http://www.python.org/dev/peps/`.

--pep-file-url-template=<plantilla>
    Plantilla para construir el nombre de las PEP, su valor predeterminado es
    :txt:`pep-%04d`.

--rfc-references
    Busca y procesa las RFC en el documento, por ejemplo: :txt:`RFC 822`.

--rfc-base-url=<URL>
    Ruta en la que encontrar las RFC, su valor predeterminado es
    :txt:`http://tools.ietf.org/html/`.

--tab-width=<ancho>
    Cantidad de caracteres para las hard tabs (tabulaciones, no espacios), su
    valor predeterminado es :txt:`8`.

.. class:: exclusive

    --trim-footnote-reference-space
        Elimina los espacios antes de las referencias de las notas de pie.

    --leave-footnote-reference-space
        Mantiene los espacios antes de las referencias de las notas de pie
        (predeterminado).

.. class:: exclusive

    --no-file-insertion
        No permite incluir archivos con :txt:`include` o :txt:`raw`.

    --file-insertion-enabled
        Permite incluir archivos con :txt:`include` o :txt:`raw`
        (predeterminado).

.. class:: exclusive

    --no-raw
        No permite usar la directiva :txt:`raw`.

    --raw-enabled
        Permite usar la directiva :txt:`raw` (predeterminado).

--syntax-highlight=<formato>
    Formato de las clases HTML_ que se usarán para resaltar la sintaxis,
    admite los valores :txt:`long` (predeterminado), :txt:`short` o
    :txt:`none`.

    .. code:: rest

        .. code:: python3

            print("Hola mundo!")

    .. code:: python3

        print("Hola mundo!")

    :txt:`--syntax-highlight=long`:

    .. code:: html
        :number-lines:

        <pre class="code python3 literal-block">
          <code>
            <span class="name builtin">print</span>
            <span class="punctuation">(</span>
            <span class="literal string double">
              "Hola mundo!"
            </span>
            <span class="punctuation">)</span>
          </code>
        </pre>

    :txt:`--syntax-highlight=short`:

    .. code:: html
        :number-lines:

        <pre class="code python3 literal-block">
          <code>
            <span class="nb">print</span>
            <span class="p">(</span>
            <span class="s2">
              "Hola mundo!"
            </span>
            <span class="p">)</span>
          </code>
        </pre>

|br|

--no-doc-title
    Transforma el título y el subtítulo del documento en secciones.

--no-doc-info
    No aplica un formato especial a los campo bibliográficos.

.. class:: exclusive

    --section-subtitles
        Permite agregar subtítulos a las secciones.

    --no-section-subtitles
        No permite agregar subtítulos a las secciones (predeterminado).

|br|

--template=<plantilla>
    Plantilla con la estructura para construir documentos, su valor
    predeterminado es :txt:`~/Entornos/docutils/lib/python3.4/site-packages/
    docutils/writers/html5_polyglot/template.txt`

--stylesheet=<URL[,...]>
    URLs de las hojas de estilo alojadas en algún servidor.

--stylesheet-path=<ruta[,...]>
    Rutas de las hojas de estilo, se debe escribir después de la opción
    :txt:`--stylesheet` si se usa, su valor predeterminado es
    :txt:`minimal.css, plain.css`.

--stylesheet-dirs=<carpeta[,...]>
    Rutas de las carpetas en las que :txt:`--stylesheet-path` busca hojas de
    estilo, su valor predeterminado es :txt:`".", "~/Entornos/docutils/lib/
    python3.4/site-packages/docutils/writers/html5_polyglot"`.

.. class:: exclusive

    --embed-stylesheet
        Agrega las hojas de estilo en :txt:`<destino>.html` (predeterminado).

    --link-stylesheet
        Solo referencia las hojas de estilo en :txt:`<destino>.html`.

--initial-header-level=<nivel>
    Nivel de título inicial para las secciones (:html:`<h1>`, :html:`<h2>`,
    etc...), no afecta a los títulos de documento, su valor predeterminado es
    :txt:`1`.

--footnote-references=<formato>
    Formato para las notas de pie, recibe los valores :txt:`brackets`
    (predeterminado) o :txt:`superscript`.

--attribution=<formato>
    Formato en que se escribe el autor de las citas, recibe los valores
    :txt:`dash` (predeterminado), :txt:`parens` y :txt:`none`.

.. class:: exclusive

    --compact-lists
        Elimina el interlineado entre las listas simples (predeterminado).

    --no-compact-lists
        No elimina el interlineado entre las listas simples.

.. class:: exclusive

    --compact-field-lists
        Elimina el interlineado entre las listas de campos simples
        (predeterminado).

    --no-compact-field-lists
        No elimina el interlineado entre las listas de campos simples.

--table-style=<clase[,...]>
    Agrega clases a todas las tablas del documento, algunas clases
    predefinidas son: :txt:`borderless`, :txt:`booktabs`, :txt:`align-left`,
    :txt:`align-center`, :txt:`align-right`, :txt:`colwidths-auto`.

--math-output=<formato [hoja_de_estilo.css]>
    Formato para procesar las fómulas matemáticas, admite los valores
    :txt:`MathML`, :txt:`HTML math.css` (predeterminado),
    :txt:`MathJax` o :txt:`LaTeX`. La hoja de estilo solo aplica para
    :txt:`HTML`, será buscada en :txt:`--stylesheet-dirs` y solo se usará si
    se procesa alguna fórmula.

.. class:: exclusive

    --xml-declaration
        Agrega la etiqueta de declaración de XML_.

    --no-xml-declaration
        Elimina la etiqueta de declaración de XML_ (predeterminado).

--cloak-email-addresses
    Ofusca las direcciones de correo para evitar el `Web scraping`_.



|bp|

Historial de cambios
====================

0.6.0
-----

Cambiado
++++++++

* Se actualizó el writer de :txt:`rst2html.py` a :txt:`rst2html5.py`, la
  principal diferencia es que el HTML_ generado es `Polyglot XHTML5`_.

* Se agregó la lista completa de opciones para el writer :txt:`rst2html5.py`.

* Se actualizó Pygments_ 2.1.3 a la versión 2.2.0.

* Se movió la referencia de la línea de comandos al final de la sección
  **Referencia**.

Borrado
+++++++

* Las opciones :txt:`--field-name-limit` y :txt:`--option-limit` ya que el
  nuevo writer usa listas de deficinión en lugar de tablas.



|bp|

Referencias
===========

**Docutils Team.** *reStructuredText.* http://docutils.sourceforge.net/rst.html

**Steve George.** *Writing and highlighting source code in reStructured Text (RST).* http://www.futurile.net/2015/08/07/writing-highlighting-code-restructured-text/