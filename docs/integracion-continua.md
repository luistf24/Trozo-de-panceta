# Integración Continua
Para la búsqueda me he centrado en aquellos sistemas de integración continua que se puedan usar con el leguaje del proyecto que es go y que estén integrandos en la nube, y para la elección me he basado en que sean gratuitos, los permisos que pidan para integrarse en GitHub o los datos para crearse una cuenta en estos.

## **Búsqueda de CI:**
En la búsqueda de CI he encontrado: [Travis](https://circleci.com/), [AppVeyor](https://www.appveyor.com/), [Cirrus](https://github.com/apps/cirrus-ci) y las propias [Github Actions](https://docs.github.com/es/actions)    

- **Travis:** Es una plataforma de integración, que como las anteriores, automatiza proceoss de creación, prueba e implementación, aunque además añade otras cosas nuevas como una sintaxis más sencilla, soporte para pull request, test en sistemas operativos como Mac, Linux y iOS además de arquitecturas únicas como AMD Graviton 64, IBM PowerPC o IMB Z. Además al igual que CircleCI te permite trabajar con ella en la nube. Se ha **descartado** ya que pide un número de tarjeta a pesar de ser gratuita y el vat.    

- **AppVeyor:** Es una plataforma de integración, que ofrece lo que las anteriores, pero lo que la distingue es es que utiliza Microsoft Build(MSBuild) Engine para crear paquetes e incluir parámetros adicionales el los archivos de respuesta. Al igual que CircleCI y Travis se puede trabajar con ella en la nube pero es más dificil de configurar. Se ha **descartado** ya que pide acceso a las organizaciones en las que estés.    

- **Cirrus:** Es otra plataforma de integración con la novedad de que es tratada como una aplicación de github por lo que es muy fácil integrarla.   

- **GitHub Actions:** Sistema de acciones que se ejecutan cada vez que sucede algún evento en el sistema, son muy amplias y van integradas en el propio repositorio por lo que es muy fácil de integrar.    

Siguiendo los criterios de elección me he quedado con: **Cirrus** y **GitHub Actions**. Dado que hay que implementar varios test, me he decantado por implementar:    
- Un test de versiones del lenguage con **Cirrus**.     
- Un test que ejecuta los tests unitarios desarrollados anteriormente con **GitHub Actions**.    
