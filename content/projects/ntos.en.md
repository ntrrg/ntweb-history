---
title: ntos
description: Debian live system with encrypted persistence.
metadata:
  source-code: https://github.com/ntrrg/ntos
  license: MIT
labels:
  - "[![Travis build btatus](https://travis-ci.com/ntrrg/ntos.svg?branch=master)](https://travis-ci.com/ntrrg/ntos)"
kinds:
  - operative-system
techs:
  - debian
  - shell-script
  - live-boot
  - travis-ci
  - make
  - shellcheck
---

**ntos** is a Debian live system with encrypted persistence by default. It is
attached at the Debian testing release cycle, so there are no version numbers
nor release names, just weekly builds.

**Requirements:**

* Debian GNU/Linux 10 (Buster)
* GNU Make

{{< toc >}}

```text
ROOTFS

IMAGE

INSTALL

  Usage: [DEV=/dev/sdX] [IPN=NUM] [NO_PERSISTENCE=true | IPP=NUM] make install

  Variables:

    * DEV: use the given USB device and don't ask for it.

    * IPN: use the given partition number as image target.

    * NO_PERSISTENCE: disable persistence.

    * PPN: use the given partition number as persistence target.

  Examples:

    Simple usage

    # make install

    No persistence

    # NO_PERSISTENCE=true make install

    Don't ask for device

    # DEV=/dev/sdb make install

    Don't ask for device or partition numbers

    # DEV=/dev/sdb IPN=1 PPN=2 make install
```

# Usage

1\. Get the image

```shell-session
$ wget -O /tmp/ntos-image.tar.gz \
  https://github.com/ntrrg/ntos/releases/download/w34/ntos-image-w34-x64.tar.gz
```

```shell-session
$ mkdir /tmp/image
```

```shell-session
# tar -xf /tmp/ntos-image.tar.gz -C /tmp/image
```

2\. Setup the parameters

```shell-session
$ EDITOR config.mk
```

3\. Install the image in a USB device

```shell-session
# make deps-install install
```

# Build

## Simple

1\. Setup the parameters

```shell-session
$ EDITOR config.mk
```

2\. Create the image

```shell-session
# make
```

3\. Install the image in a USB device

```shell-session
# make install
```

## Advanced

1\. Setup the parameters

```shell-session
$ EDITOR config.mk
```

2\. Install the build dependencies

```shell-session
# make deps
```

3\. Create the rootfs

```shell-session
# make rootfs
```

4\. Customize the rootfs, keep it as simple as possible.

```shell-session
# make login
```

5\. Create the image

```shell-session
# [NO_DEBIAN_INSTALLER=true] [ISO_URL=CUSTOM_ISO_URL] make image
```

6\. Install the image in a USB device

```shell-session
# [NO_PERSISTENCE=true] make install
```

# Acknowledgment

Working on this project I use/used:

* [Debian](https://www.debian.org/)

* [XFCE](https://xfce.org/)

* [st](https://st.suckless.org/)

* [Zsh](http://www.zsh.org/)

* [GNU Screen](https://www.gnu.org/software/screen)

* [Git](https://git-scm.com/)

* [EditorConfig](http://editorconfig.org/)

* [Vim](https://www.vim.org/)

* [GNU make](https://www.gnu.org/software/make/)

* [Chrome](https://www.google.com/chrome/browser/desktop/index.html)

* [Gogs](https://gogs.io/)

* [Github](https://github.com)

```shell-session
$ man live-boot
```

```shell-session
$ man live-config
```

```shell-session
$ man persistence.conf
```

<http://cosmolinux.no-ip.org/raconetlinux2/persistence.html>

<http://docs.kali.org/downloading/kali-linux-live-usb-persistence>

<http://willhaley.com/blog/create-a-custom-debian-live-environment/>

<https://wiki.debian.org/ReduceDebian>

<https://debian-live.alioth.debian.org/live-manual/stable/manual/html/live-manual.en.html>

<http://willhaley.com/blog/install-debian-usb/>

<https://wiki.archlinux.org/index.php/syslinux>

<http://www.syslinux.org/wiki/index.php>

<https://phenobarbital.wordpress.com/2011/05/13/linux-debiancanaima-en-soneview-n110-mini-laptop-classmate/>

<https://phenobarbital.wordpress.com/2011/07/13/debian-se-puede-tener-un-gnome-minimo/>
