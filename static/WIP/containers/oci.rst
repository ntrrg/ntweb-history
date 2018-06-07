# OCI Runtime Specification v1.0.0

---

**Contenedor:** es un entorno aislado y previamente configurado en el que se ejecutará determinado software sin tener que preocuparse por cumplir sus dependencias. Su funcionalidad es parecida a la de una máquina virtual, solo que en este caso no se virtualiza el hardware, sino solo el software.

---

Es una lista de reglas que pretende estandarizar el ciclo de vida de un contenedor, desde su creación hasta su ejecución, por medio de imágenes (plantillas) que podrán ser implementadas (crear contenedores basados en ellas) y trasladadas fácilmente.

## Principios

Cada imagen cumple con cinco principios que le otorgan la etiqueta de *Standart Container* (Contenedor Estándar) a los contenedores que sean creados a partir de ella, estos principios son:

1. **Standart operations** (operaciones estándar): cada contenedor puede ser creado, iniciado y detenido usando herramientas de manejo de contenedores que respeten esta especificación; copiado y capturado usando utilidades tradicionales del sistema de archivos; y descargado o cargado usando herramientas de red comunes.

2. **Content-agnostic** (independiente de contenido): todas las operaciones del principio uno pueden ser ejecutadas sin importar el contenido del contenedor.

3. **Infrastructure-agnostic** (independiente de infraestructura): cada contenedor puede ser implementado en cualquier tipo de infraestructura (computadores personales, IoT, servidores caseros, SaaS, PaaS, IaaS, etc).

4. **Designed for automation** (diseñado para automatización): la mayoría de las tareas necesarias para la implementación de un contenedor pueden ser automatizadas (con servicios como CI/CD).

5. **Industrial-grade delivery** (implementación a nivel industrial): cumpliendo con todos los principios anteriores, un *Standart Container* tiene la capacidad tanto de realizar tareas simples y de baja demanda, como de asegurar la ejecución de tareas que requieran de alta disponibilidad y gran consumo por medio de la creación de múltiples instancias sobre la marcha, evitando el desperdicio de recursos y disminuyendo los costos.

## Manejo de contenedores

Las operaciones que pueden realizarse con cualquier herramienta de manejo de contenedores que siga esta especificación son:

* `create`: crea un contenedor a partir de una imagen.

* `state`: solicita el estado de un contenedor, la información que se obtiene puede variar según la herramienta, pero obligatoriamente debe mostrar los siguientes datos:

    * `ociVersion`: número de versión de la especificación por la que se rige.

    * `id`: identificador que se le asigna al momento de usar una imagen.

    * `status`: estado en el que se encuentra, pueden variar según la herramienta que se use, pero todas deben respetar los siguientes estados:

        * `creating`: se está creando a partir de una imagen.

        * `created`: ya ha sido creado pero no se encuentra activo.

        * `running`: está ejecutando sus procesos.

        * `stopped`: ya ha terminado con todas sus tareas u ocurrió otro evento que forzara su salida.

    * `pid`: número de proceso que tiene asignado, solo aplica para los estados `created` o `running`.

    * `bundle`: ruta absoluta a la imagen con la que fue creado.

    * `annotations`: información extra que permite conocer más sobre su origen, sus funciones, su infraestructura, etc.

* `start`: inicia un contenedor y ejecuta sus procesos, debe estar en el estado `created`.

* `kill`: detiene un contenedor que se encuentre activo, debe estar en alguno de los estados `created` o `running`.

* `delete`: elimina un contenedor que se haya detenido, debe estar en el estado `stopped`.

## Ciclo de vida

1. Se ejecuta el comando `create` e inicia la creación de un contenedor a partir de la imagen especificada, se le asigna el estado `creating`.

2. Al finalizar las tareas de creación, el contenedor pasa al estado `created` y espera para ser iniciado.

3. Se ejecuta el comando `start` para iniciar el contenedor.

4. Se lanzan los *prestart hooks* (ganchos de pre-inicio), que son ordenes que se deben ejecutar antes de iniciar el contenedor. En caso de que alguno falle, se salta hasta el paso 9 del ciclo de vida.

5. El contenedor es iniciado y cambia de estado a `running`.

6. Se lanzan los *poststart hooks* (ganchos de post-inicio), que son ordenes que se deben ejecutar después de iniciar el contenedor.

7. El contenedor termina de ejecutar sus procesos, ocurre algún error o se ejecuta el comando `kill`; pasa al estado `stopped`.

8. Se ejecuta el comando `delete`.

9. Se elimina el contenedor.

10. Se lanzan los *poststop hooks* (ganchos de post-fin), que son ordenes que se deben ejecutar después de detener el contenedor.

## Creación de imágenes

Una imagen es un *bundle* (conjunto de elementos) que definirá el comportamiento y el contenido de un contenedor creado a partir de ella. Los elementos que la componen son:

1. [Configuración](#1-configuracion-configjson)

2. [Estructura del sistema de archivos](#2-estructura-del-sistema-de-archivos-rootfs)

### 1. Configuración (`config.json`)

Es un archivo JSON donde se establecen las configuraciones y procesos que deberán tener todos los contenedores que usen la imagen. Las propiedades que se pueden encontrar en este archivo son:

* `ociVersion` (cadena): número de versión de la especificación por la que se rige.

* `root` (objeto): estructura del sistema de archivos.

    * `path` (cadena): ruta relativa o absoluta a la estructura del sistema de archivos.

    * `readonly` (booleano, opcional, `false`): determina si el sistema de archivos será de solo lectura dentro del contenedor.

* `mounts` (arreglo, opcional): lista de montajes en orden de prioridad, cada elemento debe ser un objeto con las siguientes propiedades:

    * `destination` (cadena): ruta absoluta (tomando como raíz al contenedor) al punto de montaje.

    * `type` (cadena): tipo del sistema de archivos a montar.

    * `source` (cadena): ruta absoluta o relativa a la estructura del sistema de archivos del elemento a montar.

    * `options` (arreglo, opcional): opciones de montaje, todas las opciones soportadas pueden encontrarse [aquí](http://man7.org/linux/man-pages/man8/mount.8.html#FILESYSTEM-INDEPENDENT_MOUNT_OPTIONS) y [aquí](http://man7.org/linux/man-pages/man8/mount.8.html#FILESYSTEM-SPECIFIC_MOUNT_OPTIONS).

### 2. Estructura del sistema de archivos (`rootfs/`)

Es una carpeta que contiene la estructura del sistema operativo (o solo la aplicación) que correrá el contenedor.

```
$ mkdir miContenedor
$ cd miContenedor
$ wget -c 'http://dl-cdn.alpinelinux.org/alpine/v3.7/releases/x86_64/alpine-minirootfs-3.7.0-x86_64.tar.gz'
...
$ mkdir rootfs
# tar -xvf alpine-minirootfs-3.7.0-x86_64.tar.gz -C rootfs
$ rm alpine-minirootfs-3.7.0-x86_64.tar.gz
$ echo "{ ... }" > config.json
```

```
$ mkdir miContenedor
$ cd miContenedor
$ wget -c 'https://example.com/miContenedor.tar.gz'
...
# tar -xf miContenedor.tar.gz
$ ls
config.json  rootfs
```

<br/>

---

Si el proceso parece muy complejo o intimidante (para mí lo fue, pero no dejé de leer 😂), no es estrictamente necesario crear las imágenes de forma manual, ya que existen algunas herramientas que permiten realizar esta tarea de una manera más amigable, algunas de ellas son:

* [buildah](https://github.com/projectatomic/buildah)
* [umoci](https://umo.ci)
* [acbuild](https://github.com/containers/build)

---

## Referencias y atribuciones

[runtime-spec v1.0.0](https://github.com/opencontainers/runtime-spec/tree/v1.0.0)
