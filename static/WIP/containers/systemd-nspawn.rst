==============
systemd-nspawn
==============

----
v232
----

sudo apt install debootstrap systemd-container wpa_supplicant



sudo debootstrap --variant=minbase --include=dbus,locales,systemd-sysv stretch base http://httpredir.debian.org/debian
sudo systemd-nspawn --directory base/ passwd
sudo systemd-nspawn --register=no --boot --directory base/

    ln -sf /usr/share/zoneinfo/America/Caracas /etc/localtime
    dpkg-reconfigure locales
    reboot
    apt update
    apt upgrade -y

    <instalación de servicios>

    apt autoremove
    rm -rf /var/cache/apt/* /var/lib/apt/lists/*
    mv /usr/share/i18n/locales/en_US /tmp/
    rm -rf /usr/share/i18n/locales/??_* /usr/share/i18n/locales/???_* /usr/share/i18n/locales/eo /usr/share/i18n/locales/iso*
    mv /tmp/en_US /usr/share/i18n/locales/
    rm -rf /usr/share/locale/
    rm -rf /usr/share/man/
    rm -rf /usr/share/doc/
    rm -rf /usr/share/doc-base/

    rm -rf /usr/share/bug/
    rm -rf /usr/share/info/
    rm -rf /usr/share/X11/

    poweroff

sudo rm -rf base/etc/hostname
sudo rm -rf base/root/.bash_history base/var/log/*
mv base/ <nombre>

/var/lib/machines/<nombre>
/etc/systemd/nspawn/<nombre>.nspawn

Networking
==========

Por defecto el servicio de ``machinectl`` ejecuta los contenedores con la opción ``--network-veth``, por lo que el contenedor no compartirá la interfaz de red con el host, es decir que será inaccesible desde la red, para evitar esto se debe cambiar el comando de ejecución de ``machinectl``:

sudo mkdir /etc/systemd/system/systemd-nspawn@.service.d
sudo <editor> /etc/systemd/system/systemd-nspawn@.service.d/override.conf

    [Service]
        ExecStart=
        ExecStart=/usr/bin/systemd-nspawn --quiet --keep-unit --boot --link-journal=try-guest --settings=override --machine=%I


Antes de hacer este cambio era complicado (no se si imposible) hacer que un contenedor en específico corriera compartiendo la interfaz de red del host, ahora es mucho más fácil decidir el estilo de comunicación entre el host y el contenedor:

* Interfaz de ethernet virtual: simula una conexión punto a punto, no se puede acceder al contenedor desde la red (solo el host puede, creo que es obvio, pero por si acaso), para activarla se debe ejecutar ``systemd-nspawn --network-veth`` o crear un arhivo ``<nombre>.nspawn`` con:

      [Network]
          VirtualEthernet=yes

  Deben estar activados los servicios ``systemd-networkd`` y ``systemd-resolvd`` (solo si se usan configuraciones para DNS) en el host y el contenedor:

      wpa_passphrase <SSID> <contraseña> /etc/wpa_supplicant/wpa_supplicant-<interfaz inalámbrica>.conf
      sudo systemctl enable --now wpa_supplicant@<interfaz inalámbrica>.service

      sudo systemctl enable --now systemd-networkd.service
      sudo systemctl enable --now systemd-resolved.service
      sudo systemd-nspawn --directory base/

          systemctl enable --now systemd-networkd.service
          systemctl enable --now systemd-resolved.service
          ln -s /run/systemd/resolve/resolv.conf /etc/resolv.conf


* Switch virtual (bridge): simula una conexión entre un switch y el contenedor, se puede acceder desde la red al contenedor, para activarla se debe crear un bridge y ejecutar ``systemd-nspawn --network-bridge=<bridge>`` o crear un archivo ``<nombre>.nspawn`` con:

      [Network]
        Bridge=<bridge>

  /etc/systemd/network/25-wired.network:

      [Match]
          Name=enp*

      [Network]
          DHCP=yes
          Bridge=<bridge>

  /etc/systemd/network/25-wireless.network:

      [Match]
          Name=wlp*

      [Network]
          DHCP=yes
          Bridge=<bridge>

  /etc/systemd/network/25-bridge.netdev:

      [NetDev]
          Name=<bridge>
          Kind=bridge

  /etc/systemd/network/25-bridge.network:

      [Match]
          Name=<bridge>

      [Network]
          DHCP=yes

  Deben estar activados los servicios ``systemd-networkd`` y ``systemd-resolvd`` (solo si se usan configuraciones para DNS) en el host y el contenedor:

      wpa_passphrase <SSID> <contraseña> /etc/wpa_supplicant/wpa_supplicant-<interfaz inalámbrica>.conf
      sudo systemctl enable --now wpa_supplicant@<interfaz inalámbrica>.service

      sudo systemctl enable --now systemd-networkd.service
      sudo systemctl enable --now systemd-resolved.service
      sudo systemd-nspawn --directory base/

          systemctl enable --now systemd-networkd.service
          systemctl enable --now systemd-resolved.service
          ln -s /run/systemd/resolve/resolv.conf /etc/resolv.conf

  Si se quiere usar una IP estática se debe crear el siguiente archivo de configuración en el contenedor:

      sudo <editor> base/etc/systemd/network/25-<nombre>.network

          [Match]
              Name=host0

          [Network]
              Address=<IP>/<prefijo>
              Gateway=<IP de la puerta de enlace>
              DNS=<IP DNS> (se puede escribir múltiples veces)

          [Address]
              Broadcast=<IP del broadcast>

Resources and control
=====================

machinectl

systemd-cgtop

Para activar/desactivar el contenedor al arrancar el sistema:

    sudo systemd-nspawn {enable|disable} <nombre>
    sudo systemctl {enable|disable} systemd-nspawn@<nombre>

Para establecer límites de recursos:

    sudo systemctl set-property systemd-nspawn@<nombre> <propiedades>

    * CPUQuota=<porcentaje>%: puede usarse un porcentaje mayor a 100% para varios procesadores
    * MemoryHigh={<máximo uso de memoria>{K|M|G|T}|<porcentaje>%|infinty}
    * MemoryMax={<máximo uso de memoria>{K|M|G|T}|<porcentaje>%|infinty}: a diferencia de MemoryHigh, cuando se sobrepasa el límite mata el proceso, se debe usar MemoryHigh como principal control y MemoryMax como mecanismo de defensa
    * MemorySwapMax={<máximo uso de memoria>{K|M|G|T}|infinty}
    * TasksMax={<número>|<porcentaje>%}: Número de procesos

Referencias
===========

man debootstrap
man systemd-nspawn
man machinectl
man systemd.nspawn
man systemd.network
man systemd.resource-control

systemd-nspawn - ArchWiki. https://wiki.archlinux.org/index.php/Systemd-nspawn

Flockport - A Quick Look at Systemd Nspawn Containers. https://www.flockport.com/a-quick-look-at-systemd-nspawn-containers/

https://blog.selectel.com/systemd-containers-introduction-systemd-nspawn/

https://wiki.debian.org/ReduceDebian

https://wiki.archlinux.org/index.php/systemd-networkd

https://wiki.archlinux.org/index.php/Systemd-networkd_(Espa%C3%B1ol)