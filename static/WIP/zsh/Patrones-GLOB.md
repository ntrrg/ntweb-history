* *: Cualquier cadena.
* ?: Cualquier caracter.
* \[\<clase\>\]: Cualquier caracter, rango o clase especificada.
	* \[:ascii:\]: Caracter ASCII.
	* \[:digit:\]: Numérico.
	* \[:xdigit:\]: Hexadecimal.
	* \[:space:\]: Espacio, tabulación, salto de línea, etc...
	* \[:blank:\]: Espacio en blanco o tabulación.
	* \[:alnum:\]: Alfanumérico.
	* \[:alpha:\]: Alfabético.
	* \[:cntrl:\]: Caracter de control.
	* \[:upper:\]: Mayúscula.
	* \[:lower:\]: Minúscula.
	* \[:print:\]: Caracter imprimible.
	* \[:graph:\]: Caracter imprimible que no sea un espacio en blanco.
	* \[:punct:\]: Caracter imprimible que no sea un espacio en blanco o un caracter alfanumérico.
	* \[\[\<clase\>\]\]: Evalua solo un caracter.
	* \[\[\<clase\>\]\<otraClase\>\]: Clases anidadas.
	* \[!\<clase\>\]: Negación.
* \<\<número\>-\<número\>\>: Rangos numéricos.
* (\<patrón\>): Agrupar.
* |: Operador lógico OR, es el operador con menor precedencia por lo que se recomienda usar con parentesis

		☮› ls -d /home/ntrrg/(D*|P*)  # Mostrará todos los archivos y carpetas que empiecen con "P" o "D"

* ^: Negación, tiene más precedencia que `/` por lo que solo negará un elemento.

		☮› ls /home/^ntrrg/Descargas/  # Mostrará el contenido de todas las carpetas de descargas de los usuarios que no sean "ntrrg"

* \<patrón\>~\<otroPatrón\>: Busca coincidencias con \<patrón\> y que no coincidan con \<otroPatrón\>, solo tiene mas precedencia que `|`, pueden especificarse mas patrones sepradandolos con `~`. 

		☮› ls -d /home/*/(*~Descargas~*.png)  # Mostrará todos los archivos y carpetas que no sean "Descargas" ni terminen en ".png"

&nbsp;

[Manual: Expansión de patrones GLOB](http://zsh.sourceforge.net/Doc/Release/Expansion.html#Glob-Operators)