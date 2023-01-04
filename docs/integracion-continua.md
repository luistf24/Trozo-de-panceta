# Integración Continua     
Para la búsqueda me he centrado en aquellos sistemas de integración continua que se puedan usar con el leguaje del proyecto y que estén integrandos en la nube. Dentro de estos he seleccionado para usar en el proyecto aquellos que sean gratuitos, que no se excedan en los permisos para vincularlos con GitHub y que no demanden ciertos datos personales como puede ser el número de tarjeta.   

## **Búsqueda de CI:**    
En la búsqueda de CI he encontrado: [Travis](https://www.travis-ci.com/), [AppVeyor](https://www.appveyor.com/), [Cirrus](https://github.com/apps/cirrus-ci) y las propias [Github Actions](https://docs.github.com/es/actions)    

- **Travis:** Es una plataforma de integración que automatiza procesos de creación, prueba e implementación, aunque además añade otras cosas nuevas como una sintaxis más sencilla, soporte para pull request, test en sistemas operativos como Mac, Linux y iOS además de arquitecturas únicas como AMD Graviton 64, IBM PowerPC o IMB Z. Además al igual que CircleCI te permite trabajar con ella en la nube. No se ha seleciconado ya que pide un número de tarjeta a pesar de ser gratuita y el vat.    

- **AppVeyor:** Es una plataforma de integración, que ofrece lo que las anteriores, pero lo que la distingue es es que utiliza Microsoft Build(MSBuild) Engine para crear paquetes e incluir parámetros adicionales el los archivos de respuesta. Al igual que CircleCI y Travis se puede trabajar con ella en la nube pero es más dificil de configurar. No se ha seleccionado ya que pide acceso a las organizaciones en las que estés.    

- **Cirrus:** Es otra plataforma de integración con la novedad de que es tratada como una aplicación de github por lo que es muy fácil integrarla.   

- **GitHub Actions:** Sistema de acciones que se ejecutan cada vez que sucede algún evento en el sistema, son muy amplias y van integradas en el propio repositorio por lo que es muy fácil de integrar.    

Siguiendo los criterios de selección me he quedado con: **Cirrus** y **GitHub Actions**. Dado que hay que implementar varios test, me he decantado por implementar:    
- Un test de versiones del lenguage con **Cirrus** en el se testearan las versiones 1.18 y 1.19 ya que la 1.20 es la que se usa en el docker actual.     
- Un test que ejecuta los tests unitarios desarrollados anteriormente con **GitHub Actions**.    
