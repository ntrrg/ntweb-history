.. role:: kbd

=============
Debian (NtOS)
=============

-----
v10.0
-----

Mounts
======

Manual
------

.. code:: sh

    su

.. code:: sh

    mkdir /media/ntrrg/Nt{Disk,Laptop,Server}

.. code:: sh

    chown -R ntrrg:ntrrg /media/ntrrg/*

.. code:: sh

    cryptsetup luksOpen /dev/NtDisk/Data NtDisk

.. code:: sh

    mount /dev/mapper/NtDisk /media/ntrrg/NtDisk

.. code:: sh

    mount /dev/NtLaptop/Data /media/ntrrg/NtLaptop

.. code:: sh

    # Remote (non root)
    exit
    sshfs nt.web.ve:/media/ntrrg/NtServer /media/ntrrg/NtServer

    # Local
    mount /dev/NtServer/Data /media/ntrrg/NtServer

Automatic
---------

.. code:: sh

    su

.. code:: sh

    vim /etc/crypttab

``/etc/crypttab``

.. code:: text
    :number-lines:

    NtDisk    /dev/NtDisk/Data    none    luks

.. code:: sh

    vim /etc/fstab

``/etc/fstab``

.. code:: text
    :number-lines:

    overlay    /    overlay    rw    0    0
    tmpfs    /tmp    tmpfs    nosuid,nodev    0    0

    /dev/NtFlash/Data     /media/ntrrg/NtFlash     btrfs    defaults    0    2
    # /dev/NtLaptop/Data    /media/ntrrg/NtLaptop    btrfs    defaults    0    0
    # /dev/mapper/NtDisk    /media/ntrrg/NtDisk      btrfs    defaults    0    0

    /media/ntrrg/NtLaptop/Server/home/ntrrg/.gnupg /home/ntrrg/.gnupg btrfs    bind    0    0
    /media/ntrrg/NtLaptop/Server/home/ntrrg/.ssh   /home/ntrrg/.ssh   btrfs    bind    0    0

    # Remote
    # ntrrg@nt.web.ve:/media/ntrrg/NtServer    /media/ntrrg/NtServer    fuse.sshfs    defaults,_netdev    0    0

    # Local
    # /dev/NtServer/Data    /media/ntrrg/NtServer    btrfs    defaults    0    0

    # /media/ntrrg/NtFlash/Server/var/lib/systemd-nspawn    /var/lib/machines    btrfs    bind    0    0
    /media/ntrrg/NtFlash/Server/var/lib/docker            /var/lib/docker      btrfs    bind    0    0
    # /media/ntrrg/NtFlash/Server/var/lib/rkt               /var/lib/rkt         btrfs    bind    0    0

Config file
===========

Live config
-----------

``${PERSISTENCE}/``:

.. code:: text
    :number-lines:

    /bin union
    /etc union
    /home union
    /media union
    /opt union
    /root union
    /sbin union
    /srv union
    /usr union
    /var union

Zsh
---

``~/.zshenv``:

.. code:: text
    :number-lines:

    export SERVER="/media/ntrrg/NtFlash/Server"
    export NTENVS="${SERVER}/var/lib/ntenvs"
    export GOENVS="${NTENVS}/go"
    export NODE_ENVS="${NTENVS}/node"

    export PATH="${SERVER}/bin:${PATH}"

    . go_activate
    . node_activate

Syncs
=====

``NtLaptop`` -> ``NtFlash``:

.. code:: sh

    rsync -uaHXh --delay-updates --delete-delay --progress \
      --exclude="/Server/" \
    /media/ntrrg/NtLaptop/ /media/ntrrg/NtFlash/

``NtFlash`` -> ``NtServer``:

.. code:: sh

    rsync -uaHXh --delay-updates --delete-delay --progress \
      --exclude="/Server/" \
    /media/ntrrg/NtFlash/ /media/ntrrg/NtDisk/srv/storage/data/ntrrg/

----

``NtServer`` -> ``NtDisk``:

.. code:: sh

    rsync -uaHXh --delay-updates --delete-delay --progress \
      --exclude="/var/lib/docker" --exclude="/var/lib/rkt" \
    /media/ntrrg/NtServer/ /media/ntrrg/NtDisk/

----

``NtServer`` -> ``NtFlash``:

.. code:: sh

    rsync -uaHXh --delay-updates --delete-delay --progress \
      --exclude="/debian/"
    /media/ntrrg/NtServer/srv/mirrors/ \
    /media/ntrrg/NtFlash/Server/srv/mirrors/

``NtFlash`` -> ``NtLaptop``:

.. code:: sh

    rsync -uaHXh --delay-updates --delete-delay --progress \
    /media/ntrrg/NtFlash/Server/srv/mirrors/ \
    /media/ntrrg/NtLaptop/Server/srv/mirrors/

----

``NtServer`` -> Dropbox:

.. code:: sh

    cd /media/ntrrg/NtServer/srv/storage/data/ntrrg/

    cp -rf --reflink \
    Backups Development Documents Pictures Ringtones Templates Work \
    /media/ntrrg/NtServer/srv/sync/data/ntrrg/Dropbox

    rsync --ignore-existing -ah --delete --progress \
      --exclude=".dropbox.cache/" --exclude=".dropbox" \
    Backups Development Documents Pictures Ringtones Templates Work \
    /media/ntrrg/NtServer/srv/sync/data/ntrrg/Dropbox/

----

``NtServer`` -> MEGA:

.. code:: sh

    cd /media/ntrrg/NtServer/srv/storage/data/ntrrg/

    cp -rf --reflink \
    Backups Development Documents Pictures Ringtones Templates Work \
    /media/ntrrg/NtServer/srv/sync/data/ntrrg/mega

    rsync --ignore-existing -ah --delete --progress \
      --exclude=".debris/" \
    Backups Development Documents Pictures Ringtones Templates Work \
    /media/ntrrg/NtServer/srv/sync/data/ntrmega/

Mirrors
-------

Alpine
++++++

.. code:: sh

    rsync -uaHXzh --delay-updates --delete-after --progress \
      --exclude="v2.*" --exclude="v3.[0-6]" \
      --exclude="/**/releases" --exclude="**/aarch64" --exclude="**/armhf" \
      --exclude="**/ppc64le" --exclude="**/s390x" --exclude="**/x86" \
    rsync://rsync.alpinelinux.org/alpine/ \
    /media/ntrrg/NtServer/srv/mirrors/alpine/

Debian
++++++

.. code:: sh

    cd

.. code:: sh

    bin/ftpsync sync:all

Installation
************

.. code:: sh

    TO="/media/ntrrg/NtServer/srv/mirrors/debian"
    RSYNC_HOST="ftp.us.debian.org"
    RSYNC_PATH="debian"
    ARCH_INCLUDE="amd64"

Keyboad shortcuts
=================

