---
title: Instalar Go 1.11
date: 2018-09-09T18:35:47-04:00
description: Instalar Go es bastante sencillo, con solo seguir unas pocas instrucciones cualquiera puede hacerlo.
image: /uploads/gopher.png
categories:
  - tecnología
tags:
  - guías
  - instalaciones
  - lenguajes-de-programación
  - go
---

La forma más rápida de instalar es descargando la versión binaria, que es la
que explico aquí, si lo quieren compilar desde el código fuente, en
[esta sección](#desde-el-código-fuente) explico el procedimiento para
hacerlo.

1\. Descargar el paquete con los binarios

```shell-session
$ wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz
```

2\. Verificar que se haya descargado correctamente

```shell-session
$ sha256sum go1.11.linux-amd64.tar.gz
b3fcf280ff86558e0559e185b601c9eade0fd24c900b4c63cd14d1d38613e499  go1.11.linux-amd64.tar.gz
```

3\. Descomprimirlo en `/usr/local`

```shell-session
# tar -xvf go1.11.linux-amd64.tar.gz -C /usr/local
```

4\. Agregar los binarios a la lista de comandos del sistema

```shell-session
# ln -s /usr/local/go/bin/* /usr/bin/
```

5\. Verificar que se haya instalado correctamente

```shell-session
$ go version
go version go1.11 linux/amd64
```

6\. Eliminar los archivos necesarios para la instalación

```shell-session
$ rm go1.11.linux-amd64.tar.gz
```

{{% note %}}

Es posible instalar Go en una ruta personalizada e incluso sin permisos de
super usuario, los pasos serían muy parecidos a los anteriores, solo que hay
que cambiar las rutas y opcionalmente (si se quieren usar algunas utilidades
como ver la documentación de los paquetes) definir la variable de entorno
`GOROOT`, que le dice al sistema donde está instalado Go.

```shell-session
$ export GOROOT="$HOME/.local/go"
```

Para que se mantenga depués de cerrar el terminal se debe agregar a los
archivos de configuración del Shell.

**Zsh:**

```shell-session
$ echo "export GOROOT=\"$GOROOT\"" >> ~/.zshenv
```

**Bash:**

```shell-session
$ echo "export GOROOT=\"$GOROOT\"" >> ~/.profile
```

{{% /note %}}

# Desde el código fuente

<!--lint disable no-undefined-references no-shortcut-reference-link-->

[Cómo contribuir a Go]: {{< relref "/blog/contribute-to-go.es.md" >}}

Para este método también es necesario tener los binarios, pues desde la
versión 1.5, el compilador de Go está escrito en Go 😅, por lo que solo
tiene sentido usar este procedimiento si se tiene pensado [modificar el código
fuente][Cómo contribuir a Go].

<!--lint enable no-undefined-references no-shortcut-reference-link-->

**Nota:** puede que se necesite [Git](https://git-scm.com/) en algunas
ocasiones, depende de como se realice la instalación.

1\. Descargar el paquete con los binarios

```shell-session
$ wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz
```

2\. Verificar que se haya descargado correctamente

```shell-session
$ sha256sum go1.11.linux-amd64.tar.gz
b3fcf280ff86558e0559e185b601c9eade0fd24c900b4c63cd14d1d38613e499  go1.11.linux-amd64.tar.gz
```

3\. Descomprimirlo

```shell-session
$ tar -xvf go1.11.linux-amd64.tar.gz
```

4\. Renombrar la carpeta de Go (para evitar cualquier conflicto)

```shell-session
$ mv go toolchain
```

5\. Establecer la variable de entorno `GOROOT_BOOTSTRAP`, que determina donde
   buscar el compilador

```shell-session
$ export GOROOT_BOOTSTRAP="$PWD/toolchain"
```

6\. Obtener el código fuente

**Paquete:**

```shell-session
$ wget https://dl.google.com/go/go1.11.src.tar.gz
```

```shell-session
$ sha256sum go1.11.src.tar.gz
afc1e12f5fe49a471e3aae7d906c73e9d5b1fdd36d52d72652dde8f6250152fb  go1.11.src.tar.gz
```

```shell-session
$ tar -xvf go1.11.src.tar.gz
```

**Git:**

```shell-session
$ git clone -b go1.11 --depth 1 https://go.googlesource.com/go
```

7\. ¡Compilar!

```shell-session
$ cd go/src
```

```shell-session
$ ./all.bash
```

**Nota:** el script `all.bash` también ejecuta todas las pruebas (que es
recomendable hacerlo), para saltarse las pruebas y solo compilar, se debe usar
el script `make.bash`.

Al terminar, deberían existir nuevos recursos (entre esos, los binarios) en la
carpeta del código fuente.

8\. Mover a `/usr/local`

```shell-session
$ cd ../../
```

```shell-session
# mv go /usr/local/
```

9\. Agregar los binarios a la lista de comandos del sistema

```shell-session
# ln -s /usr/local/go/bin/* /usr/bin/
```

10\. Verificar que se haya instalado correctamente

```shell-session
$ go version
go version go1.11 linux/amd64
```

11\. Eliminar los archivos necesarios para la instalación

```shell-session
$ rm -r go1.11.linux-amd64.tar.gz go1.11.src.tar.gz toolchain
```

Para instalar algunas utilidades más (como **godoc**, que permite visualizar la
documentación de paquetes) se debe usar el comando:

```shell-session
$ go get -v golang.org/x/tools/cmd/...
```

Estos paquetes deberían instalarse en la carpeta a la que esté apuntando la
variable de entorno `$GOPATH` (que por defecto es `~/go`).

# Atribuciones

**Go Team.** *Getting Started.* <https://golang.org/doc/install>

**Go Team.** *Installing Go from source.* <https://golang.org/doc/install/source>
