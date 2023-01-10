# Servicios de configuración    
La configuración no puede encontrarse en el código ya que esto supondría un problema de seguridad, por ello tenemos que incluir diferentes posibles niveles de la misma, estos serán:    
- Depósitos clave-valor distribuidos    
- Configuración del entorno con archivos de configuración    
- Configuración desde la línea de ordenes    

Para la primera opción tenemos que buscar depósitos clave-valor distribuidos que permitan, a diferentes copias de un mismo servicio, compartir todos los elementos de la configuración. Para esta labor se podría utilizar una base de datos como [MongoDB](https://www.mongodb.com/) o [Redis](https://redis.io/), pero me he centrado en aquellas que su uso específico sea este.
Los más usados son [etcd](https://etcd.io/) y [consul](https://www.consul.io/). Me he decantado por **Redis** ya que su instalación, despliegue y configuración es menos compleja que la de Consul.

