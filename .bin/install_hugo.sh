#!/bin/sh

DEST="$OLDPWD/.vendor/hugo"
RELEASE="0.42.1"
PACKAGE="hugo_${RELEASE}_Linux-64bit.tar.gz"

cd /tmp || exit 1

wget -c "https://github.com/gohugoio/hugo/releases/download/v$RELEASE/$PACKAGE"
tar -xf "$PACKAGE"
cp hugo "$DEST"
chmod +x "$DEST"

