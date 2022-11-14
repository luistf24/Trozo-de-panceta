# Trozo-de-panceta
Repositorio creado para la asignatura IV    

## Estado del proyecto:   
En ejecución    

## Documentación:  
- [Personas](https://github.com/luistf24/Trozo-de-panceta/blob/objetivo-1/docs/personas.md)   
- [Milestones](https://github.com/luistf24/Trozo-de-panceta/blob/objetivo-1/docs/milestones.md)   
- [Historias de usuario](https://github.com/luistf24/Trozo-de-panceta/blob/objetivo-1/docs/user-stories.md)    

## Problema   
El precio de la luz varía en distintos tramos horarios a lo largo del día haciendo que el precio por poner un electrodoméstico sea más alto o más bajo en función de la hora. 

## Solución
A partir de los electrodomésticos del usuario que vaya a poner ese día, generar una planificación de forma que se reduzca lo máximo posible el coste.

## Lógica de negocio:   
Los usuarios introducirán:    
- Su ubicación (península, Ceuta y Melilla, Baleares, Canarias)    
- Los electrodomésticos que tienen y su consumo   
- Cuales les interesa poner cada día, durante cuanto tiempo (en función del tiempo se buscarán tramos de distinta duración) y cuantas veces    

Obtendrán un planning de como poner los distintos electrodomésticos a lo largo de la semana para ello se hará una regresión con los datos de los distintos tramos de las semanas anteriores, dado que los datos de los precios se publican el mismo día.    
La aplicación ajustará día a día el planning para corregir posibles fallos de la regresión obteniendo un planning óptimo diario y una predicción aproximada para los futuros días.    
Además se almacenarán los datos de los tramos diarios para mejorar las regresiones y los datos de los electrodomésticos de los usuarios para facilitar el registro de electrodomésticos a otros usuarios.            
Los principales clientes serán los usuarios de la aplicación los cuales pagaran una mensualidad por el uso de esta.   

## Herramientas utilizadas:     
- **Lenguaje de programación:** Go      
- **Task manager:** Make     
	- Instrucciones:     
		- ``opcion build`` -> Construye el proyecto    
		- ``opcion installdeps`` -> Instala las dependencias necesarias para el proyecto y limpia las que no     
		- ``opcion run`` -> Ejecuta el programa     
		- ``opcion clean`` -> Limpia el proyecto     
		- ``opcion check`` -> Analiza la sintaxis de todo el código usado en el proyecto     
		- ``opcion test`` -> Ejecuta todos los test del fichero, tienen que ser de la forma: nombre_test.go     
		- ``opcion help`` -> Imprime un mensaje con todas las opciones disponibles     
