# Alpine

```shell-session
$ rsync -uaHXzh --delete-after --progress \
    --exclude="/v2.*/" --exclude="/v3.[0-8]/" --exclude="/edge/" \
    --exclude="/**/releases" --exclude="**/aarch64" --exclude="**/armhf" \
    --exclude="**/armv7" --exclude="**/ppc64le" --exclude="**/s390x" \
    --exclude="**/x86" \
    rsync://rsync.alpinelinux.org/alpine/ \
    /media/ntServer/srv/mirrors/alpine/
```

# Debian

```shell-session
$ cd
```

```shell-session
$ git clone https://ftp-master.debian.org/git/archvsync.git
```

```shell-session
$ cp -r archvsync/bin ./
```

```shell-session
$ mkdir etc/
```

```shell-session
cat <<EOF > etc/ftpsync.conf
TO="/media/ntServer/srv/mirrors/debian"
RSYNC_HOST="ftp.us.debian.org"
RSYNC_PATH="debian"
ARCH_INCLUDE="amd64"
EOF
```

```shell-session
$ bin/ftpsync sync:all
```

