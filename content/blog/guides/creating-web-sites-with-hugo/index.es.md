---
title: Crear sitios web con Hugo
date: 2018-05-05T01:19:25-04:00
categories:
  - tecnología
  - frontend
tags:
  - generador-de-sitios-estáticos
  - hugo
  - sitios-web
draft: true
---

Hugo es un generador de sitios web estáticos, este tipo de sitios es llamado
así porque una vez que se crean sus archivos, al acceder a ellos, el contenido
que se obtiene siempre será el mismo.

> El contenido estático es aquél que permanece invariable desde el momento en
> que su autor lo crea. Es decir, no depende de quién lo visualice ni en que
> momento lo haga.
>
> \- Según [este artículo][Contenido dinámico] de Wikipedia.

[Contenido dinámico]: https://es.wikipedia.org/wiki/Contenido_din%C3%A1mico

No quiere decir que sean aburridos o algo así 😅, en realidad esta es una de
sus mejores cualidades pues:

* No hace falta procesar los archivos cuando alguien accede a ellos (más
  rápido! ⚡).

* No hay que preocuparse mucho por la seguridad (dudo que le hagan
  inyecciones a los `.yaml` 😂).

* Pueden alojarse con mucha facilidad.

Fue desarrollado inicialmente por Steve Francia y su primera versión estable
fue liberada en 2012. A la fecha (mientras escribía esto 😂) su última
versión estable es la número [0.40.2][latest-release] y algunas de sus
características son:

* Es el generador de sitios estáticos más rápido en la actualidad.

* Solo hace falta un archivo ejecutable para utilizarlo.

* Soporte para temas, que pueden crearse con cualquier tecnología web (HTML,
  CSS, JavaScript, etc...) y [Go HTML Templates][].
  
* Utilidades de desarrollo que hacen más sencillo el trabajo de edición.

* Gracias a su popularidad, se han ido creando aplicaciones como
  [Netlify CMS][] y [Forestry][] que le facilitan la vida a los editores.

* Un montón más que pueden verse [aquí][Hugo features].

[latest-release]: https://github.com/gohugoio/hugo/releases/tag/v0.40.2
[Go HTML Templates]: http://golang.org/pkg/html/template/
[Netlify CMS]: https://www.netlifycms.org/
[Forestry]: https://forestry.io
[Hugo features]: https://gohugo.io/about/features/

Es un [CMS](https://es.wikipedia.org/wiki/Sistema_de_gesti%C3%B3n_de_contenidos), esto quiere decir que él se encarga de facilitarnos la vida al momento de generar contenido para nuestro sitio; pero, ¿por qué nos la facilita? Bueno porque después de configurarlo, solo bastará con crear un archivo y él armará las entradas al articulo automáticamente. A diferencia de la mayoría de los CMS, el contenido que se obtiene de Jekyll es **"estático"**.

Pero antes de verlo como una desventaja, es bueno saber que:

* Da acceso directo al documento (más rápido! <i class="fa fa-bolt"></i>).
* No hay que preocuparse mucho por la seguridad (dudo que le hagan inyecciones a los YAML :B).
* Puede alojarse en los servidores de [GitHub](https://pages.github.com/).
* Existe AJAX.

En fin.. para realizar sus tareas usa Ruby que es un lenguaje de programación muuy conocido y se encargará de procesar todos los archivos y conventirlos en el sitio, Liquid que es el lenguaje de manejo de plantillas y nos permitirá tratar la mayoría de elementos (páginas, posts, etc...) del sitio como variables comunes y Markdown que es un lenguaje de marcado con el cual se escribirán las páginas. De estos tres lenguajes, si queremos crear el blog será necesario conocer (no ser un experto) los dos últimos y algo de desarrollo web, en el caso de que se use alguna [plantilla](http://jekyllthemes.org/), solo habrá que investigar sobre Markdown. Pero por ahora con seguir los pasos de este post, tal vez, no haga falta un nuevo ciclo de estudio.
