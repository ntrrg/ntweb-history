	bindkey "<tecla>" <ZLEWidget>

#### Lista de teclas:
* ^\<tecla\>: `Ctrl` + \<tecla\>
* ^[\<tecla\>: `Alt` + \<tecla\>
* ${terminfo[kcuu1]}: `↑`
* ${terminfo[kcuf1]}: `→`
* ${terminfo[kcud1]}: `↓`
* ${terminfo[kcub1]}: `←`
* ;5A: `Ctrl` + `↑`
* ;5C: `Ctrl` + `→`
* ;5B: `Ctrl` + `↓`
* ;5D: `Ctrl` + `←`
* ${terminfo[kich1]}: `Insert`
* ${terminfo[kdch1]}: `Supr`
* ${terminfo[khome]}: `Inicio`
* ${terminfo[kend]}: `Fin`
* ${terminfo[kpp]}: `Re Pág`
* ${terminfo[knp]}: `Av Pág`

#### Widgets del usuario:
	<función> ()
	{
		<sentencias>
	}

	zle -N <función>

&nbsp;

[Manual: Configuración del teclado y resaltado sintaxis de Zsh](http://zsh.sourceforge.net/Doc/Release/Zsh-Line-Editor.html)

[Manual: Definición de funciones en Zsh](http://zsh.sourceforge.net/Doc/Release/Functions.html)