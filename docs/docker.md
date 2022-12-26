# Contenedores
En la búsqueda me he centrado en contenedores oficiales de docker con golang con la versión que he trabajado y basados en linux, en caso de tener instalado demasiadas cosas por defecto también se contempló la idea de modificar los docker files de la búsqueda. En cuanto al criterio de selección me he centrado en el tamaño de los contenedores ya que había una diferencia muy significativa entre unos y otros

## **Búsqueda de contenedores:**
Los contenedores los he buscado en Dockerhub, donde había diversos contenedores para diversas versiones de go en función de las distribuciones, las que más destacaban para la versión que buscaba de go son: [alpine](https://github.com/docker-library/golang/blob/ee5d5d94498b7bd72a5047a2980cf762b0d7236f/1.19/alpine3.16/Dockerfile) y [debian](https://github.com/docker-library/golang/blob/ee5d5d94498b7bd72a5047a2980cf762b0d7236f/1.19/bullseye/Dockerfile)

## **Elección del contenedor:**
Como se ha comentado anteriormente el criterio de elección ha sido el tamaño del contendor y en este caso ha sido fácil decantarse por el contendor de **alpine** ya que en cuanto al tamaño, el de alpine es algo menos que la mitad del tamaño del contenedor de debian.   
Además este como se puede ver en el Dockerfile, tiene instalado lo justo para funcionar.
