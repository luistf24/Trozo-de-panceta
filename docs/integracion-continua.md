# Integración Continua     
Para la búsqueda me he centrado en aquellos sistemas de integración continua que se puedan usar con el lenguaje del proyecto, que estén integrados en la nube y que se puedan usar en GitHub. Dentro de estos he seleccionado para usar en el proyecto aquellos que sean gratuitos, que no se excedan en los permisos para vincularlos con GitHub y que no demanden ciertos datos personales como puede ser el número de tarjeta.   

## **Búsqueda de CI:**    
En la búsqueda de CI he encontrado: [AppVeyor](https://www.appveyor.com/), [Cirrus](https://github.com/apps/cirrus-ci) y las propias [Github Actions](https://docs.github.com/es/actions)    

- **AppVeyor:** Es una plataforma de integración usada para construir y testear proyectos en GitHub y otros sitios como GitLab y BitBucket. No se ha seleccionado ya que pide acceso a las organizaciones en las que estés.    

- **Cirrus:** Es otra plataforma de integración que es tratada como una aplicación de GitHub, que es la forma oficial y recomendada de integrarse con GitHub ya que ofrecen permisos mucho más granulares para acceder a los datos y además, no actuan como usuarios autentificados como ocurre con las OAuth App.   

- **GitHub Actions:** Sistema de acciones que se ejecutan cada vez que sucede algún evento en el sistema, son muy amplias y van integradas en el propio repositorio por lo que es muy fácil de integrar.    

Siguiendo los criterios de selección me he quedado con: **Cirrus** y **GitHub Actions**. Dado que hay que implementar varios test, me he decantado por implementar:    
- Un test de versiones del lenguage con **Cirrus** en el se testearan las versiones 1.18 y 1.19 ya que la 1.20 es la que se usa en la imagen del contenedor de pruebas.     
- Un test que ejecuta los tests unitarios desarrollados anteriormente con **GitHub Actions**.    
