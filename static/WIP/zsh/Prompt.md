	PROMPT=""  # Prompt
	RPROMPT=""  # Prompt izquierdo

### Componentes:

* %l: TTY activa.
* %M: Nombre del equipo.
* %n: Nombre del usuario.
* %#: Privilegios.
* %?: Estado del último comando ejecutado.
* %h: Número de elementos en el HISTFILE.
* %j: Cantidad de trabajos activos.
* %\<índice\>v: Muestra el valor del arreglo `psvar`.
* %\<número\>~: Ruta actual, 
	* 0 o sin especificar mostrará la ruta completa.
	* Positivo mostrará la cantidad de carpetas indicada a partir de la actual.
	* Negativo mostrará la cantidad de carpetas indicada a partir de la raiz.

### Estilos:

* %B\<cadena\>%b: Negrita.
* %U\<cadena\>%u: Subrayado.
* %S\<cadena\>%s: Resaltar (Invertir colores).
* %\<longitud\>{\<cadena\>%}: Segmentos.
* %F{\<color\>}\<cadena\>%f: Color de fuente.
* %K{\<color\>}\<cadena\>%k: Color de fondo.
  * 0, black: Negro.
  * 1, red: Rojo.
  * 2, green: Verde.
  * 3, yellow: Amarillo.
  * 4, blue: Azul.
  * 5, magenta: Morado.
  * 6, cyan: Turquesa.
  * 7, white: Blanco.

### Condiciones:
	PROMPT="%\<número\>(\<operador\>.\<salidaVerdadero\>.\<salidaFalso\>)"

* !: Retorna verdadero si el shell se ejecuta con privilegios de superusuario.
* ?: Retorna verdadero si el estado del ultimo comando es \<número\>.
* ~: Retorma verdadero si la ruta actual tiene al menos \<número\> de ancestros sin contar `/`.
* j: Retorna verdadero si la cantidad de trabajos actuales es al menos \<número\>.
* v: Retorna verdadero si el arreglo `psvar` tiene al menos \<número\> elementos.
* V: Retorna verdadero si el elemento \<número\> del arreglo `psvar` no esta vacío.
* D: Retorna verdadero si el mes actual es \<número\>, teniendo Enero como 0.
* d: Retorna verdadero si el dia del mes es \<número\>.
* w: Retorna verdadero si el dia de la semana es \<número\>, teniendo Domingo como 0.
* T: Retorna verdadero si la hora es \<número\>.
* t: Retorna verdadero si los minutos son \<número\>.

### Fecha y hora:
	PROMPT="%D{\<formato\>}"

* %c: Formato recomendado por localidad.
* %x: Formato recomendado por localidad sin hora.
* %X: Formato recomendado por localidad sin fecha.
* Escape:
	* %t: Tabulación.
	* %n: Salto de línea.
	* %%: Caracter %.
* Dia:
	* %u: Número del dia de la semana.
	* %w: Número del dia de la semana, teniendo Domingo como 0.
	* %A: Nombre del dia de la semana.
	* %a: Nombre del dia de la semana abreviado.
	* %f: Número del dia del mes.
	* %d: Número del dia del mes con 0.
	* %e: Número del dia del mes con " " en lugar de 0.
	* %j: Número del dia del año con 0.
* Semana:
	* %U: Número de la semana del año.
* Mes:
	* %B: Nombre del mes.
	* %b: Nombre del mes abreviado.
	* %m: Número del mes con 0.
* Año:
	* %Y: Año.
	* %y: Ultimos dos números del año.
* Hora:
	* %K: Hora en formato 24h.
	* %H: Hora en formato 24h con 0.
	* %L: Hora en formato 12h.
	* %I: Hora en formato 12h con 0.
* Minuto:
	* %M: Minuto con 0.
* Segundo:
	* %s: Segundo UNIX.
	* %S: Segundo con 0.
* Zona horaria:
	* %z: Zona horaria.
	* %Z: Nombre de la zona horaria.

### Git:
* %s: Nombre del vcs.
* %b: Rama.
* %c: Cambios en el áre de preparación.
* %u: Cambios fuera del área de preparación.
* %r: Nombre de la carpeta raiz del repositorio.
* %S: Ruta actual relativa al repositorio.

&nbsp;

[Manual: Expansión del Prompt](http://zsh.sourceforge.net/Doc/Release/Prompt-Expansion.html)