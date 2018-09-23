.. role:: kbd

=============
Debian (NtOS)
=============

-----
v10.0
-----

Config files
============

Mounts
======

``/etc/crypttab``

.. code:: text
    :number-lines:

    NtDisk    /dev/NtDisk/Data    none    luks

.. code:: shell-session

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

Syncs
=====

``NtFlash`` -> Backup:

.. code:: shell-session

    # rsync -uaXHh --delay-updates --delete-delay --progress \
      --exclude="/var/" \
      /media/ntrrg/NtFlash/ /path/to/backup/

``NtFlash`` -> Dropbox:

.. code:: shell-session

    $ rsync -uah --delay-updates --delete-delay --progress \
      --exclude=".dropbox.cache/" --exclude=".dropbox" \
      /media/ntrrg/NtFlash/srv/storage/data/ntrrg/ \
      /home/ntrrg/Dropbox/

``NtFlash`` -> MEGA:

.. code:: shell-session

    $ cp -af --reflink \
      /media/ntrrg/NtFlash/srv/storage/data/* \
      /media/ntrrg/NtFlash/var/mega/

.. code:: shell-session

    $ rsync --ignore-existing -ah --delete --progress \
      --exclude=".debris/" \
      /media/ntrrg/NtFlash/srv/storage/data/ \
      /media/ntrrg/NtFlash/var/mega/

``NtFlash`` -> ``NtServer``:

.. code:: shell-session

    $ rsync -e "ssh -p 8022" -uaXHh --delay-updates --delete-delay --progress \
      --exclude="_/games" --exclude="_/videos" \
      /media/ntrrg/NtFlash/srv/storage/data/_ \
      /media/ntrrg/NtFlash/srv/storage/data/ntrrg \
      ntrrg@home.nt.web.ve:/media/ntrrg/NtServer/srv/storage/data/

----

``NtServer`` -> Backup:

.. code:: shell-session

    # rsync -uaXHh --delay-updates --delete-delay --progress \
      --exclude="/var/" \
      /media/ntrrg/NtServer/ /path/to/backup/

Mirrors
-------

Alpine
++++++

.. code:: shell-session

    $ rsync -uaHXzh --delay-updates --delete-after --progress \
      --exclude="/v2.*/" --exclude="/v3.[0-6]/" --exclude="/edge/" \
      --exclude="/**/releases" --exclude="**/aarch64" --exclude="**/armhf" \
      --exclude="**/ppc64le" --exclude="**/s390x" --exclude="**/x86" \
      rsync://rsync.alpinelinux.org/alpine/ \
      /media/ntrrg/NtServer/srv/mirrors/alpine/

Debian
++++++

.. code:: shell-session

    cd

.. code:: shell-session

    bin/ftpsync sync:all

Installation
************

.. code:: shell-session

    TO="/media/ntrrg/NtServer/srv/mirrors/debian"
    RSYNC_HOST="ftp.us.debian.org"
    RSYNC_PATH="debian"
    ARCH_INCLUDE="amd64"

