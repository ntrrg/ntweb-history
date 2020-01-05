.. _BusyBox: https://play.google.com/store/apps/details?id=stericson.busybox

################
Comandos Android
################

Terminal
========

Forzar modo Fastboot
--------------------

.. code:: sh

    sudo echo -n boot-fastboot | busybox dd of=/dev/block/misc count=1 conv=sync; sync; reboot

.. admonition:: Nota

    Necesita BusyBox_

Instalación de aplicaciones
---------------------------

.. code:: sh

    sudo pm install <paquete>.apk

.. admonition:: Nota

    También es posible instalar aplicaciones copiando los .apk a la ruta "/data/app/" o "/system/app/" para hacerla una app del sistema