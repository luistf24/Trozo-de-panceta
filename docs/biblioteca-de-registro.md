# Biblioteca de registro   
En la búsqueda me he centrado en encontrar las bibliotecas de logs que permitan un registro de niveles estructurado, ya que esto hace que tengan un formato consistente que puede ser fácilmente procesado. También nos tiene que permitir elegir la salida por consola o por un archivo, que sirva para el mayor número de versiones posibles de go y que tenga soporte para futuras versiones de este.    
Además será importante que esté bien documentada.    
## Bibliotecas que he encontrado:   
- [log:](https://pkg.go.dev/log) Es la biblioteca estándar de golang para los logs, es muy sencilla de usar pero no cuenta con la posibilidad de estructurar los logs por lo que queda descartada.    
- [Zap:](https://pkg.go.dev/go.uber.org/zap#section-readme) Es una biblioteca de logging con niveles estructurados lo que la hace muy relevante es lo fácil que es de usar y el excelente rendimiento que tiene, pero tiene el inconveniente de que solo sirve para las últimas versiones de go por lo que queda descartada.   
- [Zerolog:](https://github.com/rs/zerolog) Como dice en su repositorio Zerolog está diseñado para garantizar una gran experiencia de desarrollo y un excelente rendimiento al igual que Zap, pero dado que no he visto ninguna limitación respecto a la versión de go y un historial de actualización muy reciente la hace una gran candidata para usar en el proyecto.   
- [Logrus:](https://github.com/sirupsen/logrus) Al igual que las anteriores es una biblioteca de loggin estructurada con el inconveniente de que está en *maintenance-mode* por lo que no añadirá nuevas mejoras y se limitarán solo a dar soporte para seguridad, por lo que queda descartada.   
- [Apex/log:](https://github.com/apex/log) Es una biblioteca de loggin estructurada pero al contrario que las demás está muy desactualizada y apenas hay documentación acerca de esta, por lo que queda descartada.   

Al final la biblioteca por la que me he decantado es **Zerolog** ya que es la única que cumple todos los requisitos.   
