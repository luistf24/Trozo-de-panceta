# Integración Continua
Para la búsqueda me he centrado en aquellos sistemas de integración continua que se puedan usar con el leguaje del proyecto que es go, y para la elección me he basado en que se puedan utilizar desde la nube, permitiendome así ahorrarme la instalación de este y además de que no dependa de mi ordenador, lo extendido que esté su uso, ya que en la búsqueda de muchos aparecen sus clientes más relevantes, que se puedan integrar fácilmente con github, que no sean de pago y los permisos que piden al integrarse en gihub.

## **Búsqueda de CI:**
En la búsqueda de CI he encontrado: [Jekins](https://www.jenkins.io/), [CircleCI](https://circleci.com/), [Travis](https://circleci.com/), [AppVeyor](https://www.appveyor.com/), [Cirrus](https://github.com/apps/cirrus-ci) y las propias [GithubActions](https://docs.github.com/es/actions)    
- **Jekins:** Jekins es un servidor de automatización open source que puede ser usado para automatizar todo tipo de tareas relacionadas con la construción, testing y despliegue de software, tiene que ser instalado tanto en docker, distintos sistemas operativos y también usando JRE.    

- **CircleCI:** Es una plataforma de integración continua para automatizar los procesos de creación, prueba e implementación, se trabaja con ella en la nube, es muy fácil de configurar y también es muy fácil de integrar con github.    

- **Travis:** Es una plataforma de integración, que como las anteriores, automatiza proceoss de creación, prueba e implementación, aunque además añade otras cosas nuevas como una sintaxis más sencilla, soporte para pull request, test en sistemas operativos como Mac, Linux y iOS además de arquitecturas únicas como AMD Graviton 64, IBM PowerPC o IMB Z. Además al igual que CircleCI te permite trabajar con ella en la nube.    

- **AppVeyor:** Es una plataforma de integración, que ofrece lo que las anteriores, pero lo que la distingue es es que utiliza Microsoft Build(MSBuild) Engine para crear paquetes e incluir parámetros adicionales el los archivos de respuesta. Al igual que CircleCI y Travis se puede trabajar con ella en la nube pero es más dificil de configurar.    

- **Cirrus:** Es otra plataforma de integración con la novedad de que es tratada como una aplicación de github por lo que es muy fácil integrarla.   

- **GitHub Actions:** Sistema de acciones que se ejecutan cada vez que sucede algún evento en el sistema, son muy amplias y van integradas en el propio repositorio por lo que es muy fácil de integrar.    

## **Elección de los CI para el proyecto:**    

Siguiendo los criterios de elección:   

- Travis se ha descartado porque pide meter la tarjeta y también el vat.   

- AppVeyor se ha descartado porque pide acceso a las organizaciones en las que estés y sin el permiso de los miembros de esta no me parece correcto usarla además su configuración es peor que la de las demás.    

- Jekins se ha descartado porque hay que instalarlo ya que no se puede ejecutar en la nube.

Por lo que me he quedado con: Cirrus, CircleCI y GitHub Actions. Dado que hay que implementar varios test, me he decantado por implementar:    
- Un test de versiones del lenguage con Cirrus ya que se pide evitar CircleCI en los hitos.     
- Un test que ejecuta los test desarrollados en un docker, esto lo implementé en el objtivo 5 con GitHub Actions.    
