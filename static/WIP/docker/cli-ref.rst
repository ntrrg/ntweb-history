.. site-description: Referencia de la interfaz de la línea de comandos de Docker.


=======================
Docker CLI (referencia)
=======================

-------
17.12.1
-------

.. contents:: Tabla de contenido

--config string  Location of client config files (default "~/.docker")
-D, --debug  Enable debug mode
-H, --host list  Daemon socket(s) to connect to
-l, --log-level string  Set the logging level ("debug"|"info"|"warn"|"error"|"fatal") (default "info")
--tls  Use TLS; implied by --tlsverify
--tlscacert string  Trust certs signed only by this CA (default "/root/.docker/ca.pem")
--tlscert string  Path to TLS certificate file (default "/root/.docker/cert.pem")
--tlskey string  Path to TLS key file (default "/root/.docker/key.pem")
--tlsverify  Use TLS and verify the remote

-v, --version
    .. code:: shell-session

       docker -v
       # Docker version 17.12.1-ce, build 7390fc6

    Muestra la versión y sale.

``image``
=========

Gestiona las imágenes que se han descargado.

``image ls``
------------

Muestra las imágenes que se han descargado.

``info``
========

.. code:: shell-session

    docker info
    # Containers: 2
    #  Running: 2
    #  Paused: 0
    #  Stopped: 0
    # Images: 3
    # Server Version: 17.12.1-ce
    # ...

Muestra información acerca de la instalación de Docker y el sistema.

.. code:: text

    docker info [OPCIONES]

-f PLANTILLA, --format PLANTILLA
    Altera la salida de acuerdo a la plantilla Go especificada.

``ps``
======

Muestra la lista de contenedores activos.

``run``
=======

.. code:: shell-session

    docker run hello-world
    # Hello from Docker!
    # This message shows that your installation appears to be working correctly.
    # ...

Inicia un contenedor y si no existe la imagen especificada, es descargada.

.. code:: text

    docker run [OPCIONES] IMAGEN [COMANDOS] [ARGUMENTOS...]

``version``
===========

.. code:: shell-session

    docker version
    # Client:
    #  Version:   17.12.1-ce
    # ...
    #
    # Server:
    #  Engine:
    #   Version:  17.12.1-ce
    # ...

Muestra de manera detallada la versión de Docker.

.. code:: text

    docker version [OPCIONES]

-f PLANTILLA, --format PLANTILLA
    Altera la salida de acuerdo a la plantilla Go especificada.
