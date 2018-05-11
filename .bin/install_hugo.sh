#!/bin/sh

cd /tmp &&
wget -c https://github.com/gohugoio/hugo/releases/download/v0.40.3/hugo_0.40.3_Linux-64bit.tar.gz &&
tar -xvf hugo_0.40.3_Linux-64bit.tar.gz &&
mv hugo "${OLDPWD}/.vendor/"

