# Trozo-de-panceta
Repositorio creado para la asignatura IV    

## Problema   
El precio de la luz varía en distintos tramos horarios a lo largo del día haciendo que el precio por poner un electrodoméstico sea más alto o más bajo en función de la hora. 

## Solución
A partir de los electrodomésticos del usuario que vaya a poner ese día, generar una planificación de forma que se reduzca lo máximo posible el coste.

## Lógica de negocio:   
Los usuarios introducirán:    
- Su ubicación    
- Los electrodomésticos que tienen y su consumo   
- Cuales les interesa poner cada día, durante cuanto tiempo y cuantas veces    

Obtendrán un planning de como poner los distintos electrodomésticos a lo largo de la semana para ello se hará una regresión con los datos de los distintos tramos de las semanas anteriores, dado que los datos de los precios se publican el mismo día. La aplicación ajustará día a día el planning para corregir posibles fallos de la regresión obteniendo un planning óptimo diario y una predicción aproximada para los futuros días. Además se almacenarán los datos de los tramos diarios para mejorar las regresiones y los datos de los electrodomésticos de los usuarios para facilitar el registro de electrodomésticos a otros usuarios.            
Los principales clientes serán los usuarios de la aplicación los cuales pagaran una mensualidad por el uso de esta.   
