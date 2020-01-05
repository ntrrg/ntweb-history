# Ntzsh, Miguel Angel Rivera Notararigo <ntrrgx@gmail.com>

#	1. Instalación:
#		# apt-get install zsh
#	2. Configuración:
#		· Usuario específico:
#			$ cp ZSH.txt ~/.zshrc
#			$ chsh -s /bin/zsh
#			$ zsh

#		· Todos los usuarios:
#			# cp ZSH.txt /etc/zsh/zshrc

#			Uno por uno:
#				# chsh -s /bin/zsh <usuario>

#			Desde el archivo de configuración
#				# <editor> /etc/passwd
#					root:x:0:0:root:/root:/bin/bash -> root:x:0:0:root:/root:/bin/zsh
#					ntrrg:x:1000:1000:Miguel Angel,,,:/home/ntrrg:/bin/bash -> ntrrg:x:1000:1000:Miguel Angel,,,:/home/ntrrg:/bin/zsh

# ╔═══════════════════════╗
# ║ Historial de comandos ║
# ╚═══════════════════════╝
	HISTFILE=~/.histfile
	HISTSIZE=1000
	SAVEHIST=1000



# ╔════════╗
# ║ PROMPT ║
# ╚════════╝
	autoload -U promptinit
	promptinit

#	setopt promptsubst  # Mejoras para el PROMPT como evaluación de expresiones aritmeticas, ejecución de funciones, etc..

#	┌─────┐
#	│ Git │
#	└─────┘
		autoload -U vcs_info
		zstyle ':vcs_info:*' enable git
		zstyle ':vcs_info:*' check-for-changes true  # Para habilitar %c y %u
		zstyle ':vcs_info:*' stagedstr ' ✔'
		zstyle ':vcs_info:*' unstagedstr ' ✘'
		zstyle ':vcs_info:*' formats "╞ %b%c%u ╡"

		function precmd() # Función predefinida de ZSH que ejecuta las sentencias antes de mostrar el PROMPT
		{
			vcs_info
			psvar=$vcs_info_msg_0_
		}

	PROMPT="%B%n %(!.%F{red}☢%f.%F{green}☮%f)›%b "
	RPROMPT="%B%v╞ %F{green}%1~%f ╡%b"



# ╔══════════════════╗
# ║ Autocompletación ║
# ╚══════════════════╝
	autoload -U compinit
	compinit

	zstyle ':completion:*' menu select  # Habilitar navegación entre las sugerencias, soporta el uso de las flechas
	zstyle ':completion:*:descriptions' format '%SCoincidencias:%s'
	zstyle ':completion:*:warnings' format '%SNo se encuentran coincidencias..%s'



#	┌─────────┐
#	│ Colores │
#	└─────────┘
		LS_COLORS='rs=0:di=01;34:ln=01;36:mh=00:pi=40;33:so=01;35:do=01;35:bd=40;33;01:cd=40;33;01:or=40;31;01:su=37;41:sg=30;43:ca=30;41:tw=30;42:ow=34;42:st=37;44:ex=01;32:*.tar=01;31:*.tgz=01;31:*.arj=01;31:*.taz=01;31:*.lzh=01;31:*.lzma=01;31:*.tlz=01;31:*.txz=01;31:*.zip=01;31:*.z=01;31:*.Z=01;31:*.dz=01;31:*.gz=01;31:*.lz=01;31:*.xz=01;31:*.bz2=01;31:*.bz=01;31:*.tbz=01;31:*.tbz2=01;31:*.tz=01;31:*.deb=01;31:*.rpm=01;31:*.jar=01;31:*.war=01;31:*.ear=01;31:*.sar=01;31:*.rar=01;31:*.ace=01;31:*.zoo=01;31:*.cpio=01;31:*.7z=01;31:*.rz=01;31:*.jpg=01;35:*.jpeg=01;35:*.gif=01;35:*.bmp=01;35:*.pbm=01;35:*.pgm=01;35:*.ppm=01;35:*.tga=01;35:*.xbm=01;35:*.xpm=01;35:*.tif=01;35:*.tiff=01;35:*.png=01;35:*.svg=01;35:*.svgz=01;35:*.mng=01;35:*.pcx=01;35:*.mov=01;35:*.mpg=01;35:*.mpeg=01;35:*.m2v=01;35:*.mkv=01;35:*.webm=01;35:*.ogm=01;35:*.mp4=01;35:*.m4v=01;35:*.mp4v=01;35:*.vob=01;35:*.qt=01;35:*.nuv=01;35:*.wmv=01;35:*.asf=01;35:*.rm=01;35:*.rmvb=01;35:*.flc=01;35:*.avi=01;35:*.fli=01;35:*.flv=01;35:*.gl=01;35:*.dl=01;35:*.xcf=01;35:*.xwd=01;35:*.yuv=01;35:*.cgm=01;35:*.emf=01;35:*.axv=01;35:*.anx=01;35:*.ogv=01;35:*.ogx=01;35:*.aac=00;36:*.au=00;36:*.flac=00;36:*.mid=00;36:*.midi=00;36:*.mka=00;36:*.mp3=00;36:*.mpc=00;36:*.ogg=00;36:*.ra=00;36:*.wav=00;36:*.axa=00;36:*.oga=00;36:*.spx=00;36:*.xspf=00;36:'
		export LS_COLORS

		zstyle ':completion:*:default' list-colors ${(s.:.)LS_COLORS}  # Color de las sugerencias



# ╔═══════════════════════════╗
# ║ Configuración del teclado ║
# ╚═══════════════════════════╝
	bindkey "^r" history-incremental-search-backward  # Buscar en el historial
	bindkey "${terminfo[kpp]}" up-line-or-search  # Busca hacia arriba en el historial según lo que esté escrito
	bindkey "${terminfo[knp]}" down-line-or-search  # Busca hacia abajo en el historial según lo que esté escrito



# ╔═══════╗
# ║ Alias ║
# ╚═══════╝
#	alias lsa=ls -A



# ╔════════╗
# ║ Extras ║
# ╚════════╝
#	┌───────────┐
#	│ Historial │
#	└───────────┘
		setopt histignorealldups  # Evita que se guarden lineas repetidas en el HISTFILE
		setopt histignorespace  # No guarda en el historial comandos que empiecen con un espacio en blanco
		setopt histreduceblanks  # Elimina los espacios innecesarios de los comandos antes de guardarlo en el HISTFILE

#	┌──────────────────┐
#	│ Autocompletación │
#	└──────────────────┘
		setopt correctall  # Muestra posibles correciones a errores de sintaxis en comandos y parametros
		setopt extendedglob  # Agrega los caracteres "^", "~" y "#" a los patrones GLOB
		setopt globdots  # Permite que los patrones GLOB reconozcan archivos ocultos sin tener que anteponer el "."
		setopt listrowsfirst  # Ordena de manera horizontal las sugerencias
#		setopt globcomplete  # En lugar de añadir todas las coincidencias con el patrón GLOB, muestra el menú de sugerencias. Ej: cd D*

#		setopt nocaseglob  # Evita la sensibilidad a mayusculas en patrones GLOB
		setopt nohashdirs  # Actualiza automáticamente la lista de binarios para la autocompletación
		setopt nohashcmds  # Cumple la misma función que la opción anterior

#	┌────────────────────────────┐
#	│ Comportamiento de comandos │
#	└────────────────────────────┘
		setopt autocd  # No es necesario escribir el comando "cd". Ej: cd Descargas -> Descargas
		setopt cdablevars  # Si se especifica una ruta relativa y no existe dentro de la carpeta actual, ZSH intentará buscar en ~/
		setopt interactivecomments  # Permite que ZSH ejecute comandos con comentarios. Ej: ls  # Mostrar contenido de una carpeta
		
		setopt noclobber  # Evita que ">" sobrescriba archivos y que ">>" los cree, se debe usar ">!" y ">>!" respectivamente para poder hacerlo

#	┌───────────────────────────┐
#	│ Configuración del teclado │
#	└───────────────────────────┘
#		setopt ignoreeof  # Evita la salida con "Ctrl + D"



# ╔═════════════╗
# ║ Referencias ║
# ╚═════════════╝
#	┌────────┐
#	│ PROMPT │
#	└────────┘
#		PROMPT=""  # Prompt
#		RPROMPT=""  # Prompt izquierdo

#		Componentes:
#			%l -> TTY activa
#			%M -> Nombre del equipo
#			%n -> Nombre de usuario
#			%# -> Privilegios
#			%? -> Estado del último comando ejecutado
#			%h -> Número de elementos en el HISTFILE
#			%j -> Cantidad de procesos activos en segundo plano
#			%v -> Muestra el valor del arreglo "psvar"
#			%<número>~ -> Ruta actual
#				· 0 o sin especificar mostrará la ruta completa
#				· Positivo mostrará la cantidad de carpetas indicada a partir de la actual
#				· Negativo mostrará la cantidad de carpetas indicada a partir de la raiz

#		Estilos:
#			%F{<color>}<cadena>%f -> Color de fuente
#			%K{<color>}<cadena>%k -> Color de fondo
#			%B<cadena>%b -> Negrita
#			%U<cadena>%u -> Subrayado
#			%S<cadena>%s -> Resaltar (Invertir colores)
#			%<longitud>{<cadena>%} -> Segmentos

#			Colores:
#				0, black -> Negro
#				1, red -> Rojo
#				2, green -> Verde
#				3, yellow -> Amarillo
#				4, blue -> Azul
#				5, magenta -> Morado
#				6, cyan -> Turquesa
#				7, white -> Blanco

#		Condiciones:
#			%<número>(<operador>.<salidaVerdadero>.<salidaFalso>)

#			Operadores:
#				! -> Retorna verdadero si el shell se ejecuta con privilegios de superusuario
#				? -> Retorna verdadero si el estado del ultimo comando es <número>
#				~ -> Retorma verdadero si la ruta actual tiene al menos <número> de ancestros sin contar "/"
#				j -> Retorna verdadero si la cantidad de trabajos actuales es al menos <número>
#				v -> Retorna verdadero si el arreglo "psvar" tiene al menos <número> elementos
#				V -> Retorna verdadero si el elemento <número> del arreglo "psvar" no esta vacío
#				D -> Retorna verdadero si el mes actual es <número>, teniendo Enero como 0
#				d -> Retorna verdadero si el dia del mes es <número>
#				w -> Retorna verdadero si el dia de la semana es <número>, teniendo Domingo como 0
#				T -> Retorna verdadero si la hora es <número>
#				t -> Retorna verdadero si los minutos son <número>

#		Fecha y hora:
#			%D{<formato>} -> Formato personalizado
#				%c -> Formato recomendado por localidad
#				%x -> Formato recomendado por localidad sin hora
#				%X -> Formato recomendado por localidad sin fecha

#				Escape:
#					%t -> Tabulación
#					%n -> Salto de línea
#					%% -> Caracter %

#				Dia:
#					%u -> Número del dia de la semana
#					%w -> Número del dia de la semana, 0 - Domingo
#					%A -> Nombre del dia de la semana
#					%a -> Nombre del dia de la semana abreviado

#					%f -> Número del dia del mes
#					%d -> Número del dia del mes con 0
#					%e -> Número del dia del mes con " " en lugar de 0

#					%j -> Número del dia del año con 0

#				Semana:
#					%U -> Número de la semana del año

#				Mes:
#					%B -> Nombre del mes
#					%b -> Nombre del mes abreviado
#					%m -> Número del mes con 0

#				Año:
#					%Y -> Año
#					%y -> Ultimos dos números del año

#				Hora:
#					%K -> Hora en formato 24h
#					%H -> Hora en formato 24h con 0

#					%L -> Hora en formato 12h
#					%I -> Hora en formato 12h con 0

#				Minuto:
#					%M -> Minuto con 0

#				Segundo:
#					%s -> Segundo UNIX
#					%S -> Segundo con 0

#				Zona horaria:
#					%z -> Zona horaria
#					%Z -> Nombre de la zona horaria

#		Manual: http://zsh.sourceforge.net/Doc/Release/Prompt-Expansion.html

#		Git:
#			Componentes:
#				%s -> Nombre del vcs
#				%b -> Rama
#				%c -> Cambios en el áre de preparación
#				%u -> Cambios fuera del área de preparación
#				%r -> Nombre de la carpeta raiz del repositorio
#				%S -> Ruta actual relativa al repositorio

#	┌───────────────────────────┐
#	│ Configuración del teclado │
#	└───────────────────────────┘
#		Para definir una nueva acción:
#			bindkey "<tecla>" <ZLEWidget>

#		Lista de teclas:
#			^<tecla> -> Ctrl + <tecla>
#			^[<tecla> -> Alt + <tecla>
#			${terminfo[kich1]} -> Insert
#			${terminfo[kdch1]} -> Supr
#			${terminfo[khome]} -> Inicio
#			${terminfo[kend]} -> Fin
#			${terminfo[kpp]} -> Re Pág
#			${terminfo[knp]} -> Av Pág

#		Manual: http://zsh.sourceforge.net/Doc/Release/Zsh-Line-Editor.html

#		Para crear nuevos Widgets:
#		<función> ()
#		{
#			<sentencias>
#		}
#	
#		zle -N <función>

#		Funciones: http://zsh.sourceforge.net/Doc/Release/Functions.html

#	┌───────────────┐
#	│ Patrones GLOB │
#	└───────────────┘
#		* -> Cualquier cadena
#		? -> Cualquier caracter
#		[<clase>] -> Cualquier caracter, rango o clase especificada
#			[:ascii:] -> Caracter ASCII
#			[:digit:] -> Numérico
#			[:xdigit:] -> Hexadecimal
#			[:space:] -> Espacio, tabulación, salto de línea, etc...
#			[:blank:] -> Espacio en blanco o una tabulación
#			[:alnum:] -> Alfanumérico
#			[:alpha:] -> Alfabético
#			[:cntrl:] -> Caracter de control
#			[:upper:] -> Mayúscula
#			[:lower:] -> Minúscula
#			[:print:] -> Caracter imprimible
#			[:graph:] -> Caracter imprimible que no sea un espacio en blanco
#			[:punct:] -> Caracter imprimible que no sea un espacio en blanco o un caracter alfanumérico
#
#			[[<clase>]] -> Evalua solo un caracter
#			[[<clase>]<otraClase>] -> Clases anidadas
#			[!<clase] -> Negación
#		<<número>-<número>> -> Rangos numéricos
#		(<patrón>) -> Agrupar
#		| -> Operador lógico OR, se debe usar con "(<patrón>)|(<otroPatrón>)" pues es el operador con menor precedencia
#		^ -> Negación, tiene más precedencia que "/" por lo que solo negará un elemento. Ej: "ls /home/^usuario/Descargas/", mostrará el contenido de todas las carpetas de descargas de los usuarios que no sean "usuario"
#		<patrón>~<otroPatrón> -> Busca coincidencias con <patrón> y que no coincidan con <otroPatrón>, solo tienes mas precedencia que "|", pueden especificarse mas patrones sepradandolos con "~". Ej: "ls -d /home/*/D*~Dropbox", mostrará todas las carpetas y archivos que empiecen con "D" y no sean "Dropbox"
#		<patrón># -> El patrón puede aparecer 0 o mas veces de manera recursiva, tiene gran precedencia por lo que se recomienda usar parentesis. Ej: (<patrón>)#
#		<patrón>## -> El patrón puede aparecer 1 o mas veces de manera recursiva, tiene gran precedencia por lo que se recomienda usar parentesis. Ej: (<patrón>)##

#		Manual: http://zsh.sourceforge.net/Doc/Release/Expansion.html#Glob-Operators

#	Manual ZSH: http://zsh.sourceforge.net/Doc/Release/zsh.html#Top
#	Lista de opciones: http://linux.die.net/man/1/zshoptions

#	https://wiki.archlinux.org/index.php/Zsh
#	https://wiki.gentoo.org/wiki/Zsh/Guide
#	https://github.com/robbyrussell/oh-my-zsh