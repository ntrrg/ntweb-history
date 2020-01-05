# ntLaptop

```shell-session
$ export NTLAPTOP="/media/host/srv/baul/data"
```

## Local backup

```shell-session
$ rsync -uaXHh --delete-delay --progress \
    "$NTLAPTOP/" /path/to/backup/
```

## Non-Linux FS local backup

```shell-session
$ rsync -urth --delete-delay --progress \
    "$NTLAPTOP/" /path/to/backup/
```

## ntServer

### Rsync

```shell-session
$ rsync -uaXHh --delay-updates --delete-delay --progress \
    "$NTLAPTOP/ntrrg/" ntrrg@s6.nt.web.ve::ntrrg/
```

```shell-session
$ rsync -uaXHh --delay-updates --delete-delay --progress \
    "$NTLAPTOP/software/" ntrrg@s6.nt.web.ve::software/
```

### SFTP

```shell-session
$ rsync -uaXHh --delay-updates --delete-delay --progress \
    "$NTLAPTOP/ntrrg/" ntrrg@s6.nt.web.ve:
```

## MEGA:

```shell-session
$ cd "$NTLAPTOP"
```

```shell-session
$ cp -af --reflink "$NTLAPTOP/ntrrg" /media/host/home/ntrrg/MEGA/
```

```shell-session
$ rsync -r --ignore-existing --delete --progress \
    --exclude=".debris/" \
    "$NTLAPTOP/ntrrg/" /media/host/home/ntrrg/MEGA/ntrrg/
```

