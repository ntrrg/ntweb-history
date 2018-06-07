.. |br| raw:: html

    <br />

.. |bp| raw:: html

    <div style="page-break-after: always"></div>



.. role:: txt(code)
    :language: text

.. role:: py(code)
    :language: py



.. _Odoo: https://www.odoo.com
.. _Odoo Wikipedia: https://es.wikipedia.org/wiki/Odoo
.. _ERP: https://es.wikipedia.org/wiki/ERP
.. _Python: http://www.python.org/
.. _Debian: https://www.debian.org/
.. _PIL: https://pypi.python.org/pypi/PIL
.. _PostgreSQL: https://www.postgresql.org/
.. _NodeJS: https://nodejs.org/en/
.. _UNIX: https://es.wikipedia.org/wiki/Unix
.. _PROMPT: https://es.wikipedia.org/wiki/Prompt
.. _wkhtmltopdf: http://wkhtmltopdf.org/
.. _MVC: https://es.wikipedia.org/wiki/Modelo%E2%80%93vista%E2%80%93controlador
.. _Less: http://lesscss.org/



####
Odoo
####

===
9.0
===

:Date: 2016-08-27
:Version: 0.3.0
:Author: Miguel Angel Rivera Notararigo (ntrrg) <ntrrgx@gmail.com>
:Licence: MIT

|br|

.. image:: img/odoo.png
    :align: center

|br|

__ Odoo_
__ `Odoo Wikipedia`_
__ Debian_

Odoo_ es un ERP_, es decir, un software para dar soluciones empresariales,
aunque también ofrece otra gran cantidad de utilidades; si quieren saber más
sobre lo que es y lo que ofrece, pueden ir al `sitio oficial`__ o a su `página
en Wikipedia`__.

Yo uso `Debian 8 (Jessie)`__ por lo que este artículo está orientado a esta
distribución, pero es muy probable que sirva (exceptuando los comandos del
*Shell*) en cualquier sistema operativo con Python_ instalado; todos los
comandos que requieren privilegios de superusuario tienen :txt:`sudo` en
ellos, si el usuario que usan no pertenece al grupo :txt:`sudo` pueden
ejecutarlos con el usuario :txt:`root`, los demás comandos deben ser
ejecutados por un usuario común.

.. sectnum::
    :suffix: .

.. contents:: Tabla de contenido
    :backlinks: top

|bp|



Instalación
===========

Dependencias
------------

Verificar que todo esté actualizado e instalar lo que necesita Odoo_ del
sistema operativo:

.. code:: sh

    sudo apt-get update && sudo apt-get upgrade

.. code:: sh

    sudo apt-get install git libfreetype6-dev libjpeg-dev libldap2-dev \
    postgresql-server-dev-all python-dev python-passlib python-virtualenv \
    libpng12-dev libsasl2-dev libxslt1-dev npm postgresql zlib1g-dev

Para que Odoo_ pueda comunicarse con el motor de tipografías *FreeType* y el
interprete de NodeJS_ se deben crear algunos enlaces simbólicos:

.. code:: sh

    sudo ln -s /usr/include/freetype2 /usr/include/freetype

.. code:: sh

    sudo ln -s /usr/bin/nodejs /usr/bin/node

|br|

Instalar Less_ desde NodeJS_:

.. code:: sh

    sudo npm install -g less less-plugin-clean-css

|br|

Instalar wkhtmltopdf_:

.. admonition:: Nota

    __ http://download.gna.org/wkhtmltopdf/0.12/0.12.1/

    Si no usan Debian_, pueden descargar la versión que se adapte mejor a su
    distribución desde aquí__.

.. code:: sh

    wget -qO - http://goo.gl/H7xNAq > wkhtmltopdf.deb  # 32bits
    wget -qO - http://goo.gl/kYVtMI > wkhtmltopdf.deb  # 64bits

.. code:: sh

    sudo dpkg -i wkhtmltopdf.deb

.. code:: sh

    rm wkhtmltopdf.deb

|br|

Odoo_ no permite establecer conexiones a la base de datos con el usuario
predefinido de PostgreSQL_, por lo que debe crearse uno para el usuario que lo
ejecute:

.. code:: sh

    sudo passwd postgres

.. _comando-usuario-postgres:

.. code:: sh

    su postgres -c "createuser -sEP `whoami`"

|br|



Entorno virtual
---------------

No es obligatorio crear un entorno virtual, pero es recomendable hacerlo para
aislar los paquetes Python_ usados por Odoo_ y poder usar :txt:`pip` sin
privilegios de superusuario; para preparar un entorno virtual hay ejecutar los
siguientes comandos:

.. code:: sh

    cd

.. code:: sh

    virtualenv -p python2.7 odooEnv

.. code:: sh

    cd odooEnv

.. code:: sh

    source bin/activate  # Agregará (odooEnv) al PROMPT

.. code:: sh

    pip install -U Babel Jinja2 Mako MarkupSafe Pillow Python-Chart PyYAML \
    Werkzeug argparse decorator docutils feedparser gdata gevent greenlet \
    jcconv lxml mock ofxparse passlib psutil psycogreen psycopg2 pyPdf pydot \
    pyparsing pyserial python-dateutil python-ldap python-openid pytz pyusb \
    qrcode reportlab six suds-jurko vatnumber vobject wsgiref xlwt requests \
    watchdog

|br|



Descargar Odoo
--------------

.. code:: sh

    cd ~/odooEnv

.. code:: sh

    git clone git@github.com:odoo/odoo.git  # con SSH (recomendado)
    git clone https://github.com/odoo/odoo.git  # con HTTP

|br|



Ejecución
=========

Odoo_ posee un script que permite iniciar una instancia de él, podemos alterar
su comportamiento por medio de opciones al estilo comandos Unix_. Para
iniciarlo de la manera más sencilla, basta con entrar en el repositorio y
hacer uso del script:

.. admonition:: Nota

    En caso de que no esté activo el entorno virtual (el PROMPT_ no tiene
    :txt:`(odooEnv)`), ejecutar el comando:

    .. code:: sh

        source ~/odooEnv/bin/activate

.. code:: sh

    cd ~/odooEnv/odoo

.. code:: sh

    ./odoo.py

|br|



Configuración inicial
---------------------

|br|



Opciones
--------

.. admonition:: Nota

    Solo se muestran las opciones más relevantes, para obtener la lista
    completa se puede ejecutar el script con la opción :txt:`-h`.

--xmlrpc_port <puerto>  Puerto en el que se montará el servicio, su valor
                        predeterminado es :txt:`8069`.

--logfile <archivo>  Archivo donde se guardará el registro de actividades
                     hechas por Odoo_, su valor predefinido es
                     :txt:`/dev/stdout`.

--addons-path <carpetas>  Lista de carpetas separadas por coma en las que se
                          buscarán módulos.

--dev  Activa el modo desarrollador, útil para evitar reiniciar el servicio
       cuando se está trabajando sobre un módulo.

|br|

-d <base de datos>, --database <base de datos>  Base de datos a usar, es
                                                equivalente a seleccionar una
                                                base de datos en la interfaz
                                                gráfica.

--db_filter <expresión regular>  Solo muestra las bases de datos cuyos nombres
                                 coincidan con :txt:`<expresión regular>`.

-r <usuario>, --db_user <usuario>  Usuario PostgreSQL_ con el que se
                                   conectará, por defecto intenta conectarse
                                   con el nombre del usuario que ejecuta el
                                   servicio.

-w <contraseña>, --db_password <contraseña>  Contraseña del usuario
                                             PostgreSQL_.

--db_host <host>  Host en el que se encuentra el servidor PostgreSQL_.

--db_port <puerto>  Puerto en el corre el servidor PostgreSQL_.

--db_template <plantilla>  Plantilla con la que se crearán las bases de datos.

|br|

-i <módulos>, --init <módulos>  Lista de módulos separados por coma a
                                instalar o :txt:`all`, es equivalente a
                                instalar módulos desde la interfaz gráfica;
                                necesita la opción :txt:`-d`.

-u <módulos>, --update <módulos>  Lista de módulos separados por coma a
                                  actualizar o :txt:`all`, es equivalente a
                                  actualizar módulos desde la interfaz
                                  gráfica; necesita la opción :txt:`-d`.

|br|



Archivo de configuración
++++++++++++++++++++++++

Es posible definir las opciones anteriores por medio de un archivo especial,
que será leido por el script y de esta manera se evitará escribir todas las
opciones nuevamente en cada inicio del servicio. El script permite usar las
siguientes opciones para trabajar con este archivo:

-c <archivo>, --config <archivo>  Archivo de configuración, su valor
                                  predefinido es :txt:`~/.openerp_serverrc`.

-s, --save  Guarda la configuración actual en el archivo especificado en la
            opción :txt:`-c`.

Las opciones nombradas en Opciones_ se definirían de la siguiente manera:

.. code:: ini

    [options]
    xmlrpc_port = <puerto>
    logfile = <archivo>
    addons_path = <carpetas>
    dev_mode = True | False

    db_name = <base de datos>
    dbfilter = <expresión regular>
    db_user = <usuario>
    db_password = <contraseña>
    db_host = <host>
    db_port = <puerto>
    db_template = <plantilla>

|br|



Módulos
=======

Los módulos son elementos (aplicaciones, temas, etc...) que pueden agregar
funcionalidades o alterar el comportamiento/aspecto de Odoo_. Su estructura, a
pesar de seguir el patrón MVC_, puede variar según la finalidad con la que se
cree, pero generalmente debería ser algo como:

.. code:: text

    .
    ├─ controllers/
    │  ├─ __init__.py
    │  └─ main.py
    ├─ data/
    │  ├─ <modelo>_data.xml
    │  └─ <modelo>_demo.xml
    ├─ models/
    │  ├─ __init__.py
    │  └─ <modelo>.py
    ├─ static/
    │  ├─ css/
    │  ├─ img/
    │  └─ js/
    ├─ views/
    │  ├─ <modelo>_templates.xml
    │  └─ <modelo>_views.xml
    ├─ __init__.py
    └─ __openerp__.py

|br|

Modelos
-------

.. code:: text

    ├─ models/
    │  ├─ __init__.py
    │  └─ <modelo>.py

Son la representación de los componentes de un sistema por medio de clases
Python_ que heredan de la clase :py:`openerp.models.Model`, ejemplo:

.. code:: python

    from openerp.models import fields, Model


    class Persona(Model):
        u"""Representación de persona."""

        _name = "persona"
        _description = u"Representación de persona."

        nombre = fields.Char(string="Nombre")
        apellido = fields.Char(string="Apellido")

|bp|



Solución de errores
===================

Instalación
-----------

Instalando wkhtmltopdf
++++++++++++++++++++++

Si se muestra algún error de dependencias, hay que ejecutar el comando:

.. code:: sh

    sudo apt-get install -f




Creando el usuario PostgreSQL
+++++++++++++++++++++++++++++

.. image:: img/error-usuario-postgres.png
    :align: center
    :height: 300px

Este error se genera porque el servicio de PostgreSQL_ no está activo, para
arreglarlo basta con iniciar el servicio:

.. code:: sh

    sudo /etc/init.d/postgres restart

__ `comando-usuario-postgres`_

[Regresar__]

|br|



Ejecución
---------

Procesando imágenes
+++++++++++++++++++

Para solucionar este error se debe reinstalar PIL_ con los siguientes
comandos:

.. code:: sh

    rm -r OdooEnv/lib/python2.7/site-packages/PIL*

.. code:: sh

    wget http://effbot.org/downloads/Imaging-1.1.7.tar.gz  # 32bits
    wget http://effbot.org/downloads/Imaging-1.1.7-x64.tar.gz  # 64bits

.. code:: sh

    tar -xvf Imaging-1.1.7*

.. code:: sh

    cd Imaging-1.1.7

.. code:: sh

    python setup.py install

.. code:: sh

    cd ..

.. code:: sh

    rm -r Imaging-1.1.7*

|bp|



Referencias
===========

*Installing Odoo*. https://www.odoo.com/documentation/9.0/setup/install.html#source-install

*Command-line interface: odoo.py*. https://www.odoo.com/documentation/9.0/reference/cmdline.html