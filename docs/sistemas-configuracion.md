# Servicios de configuración    
La configuración no puede encontrarse en el código ya que esto supondría un problema de seguridad, por ello tenemos que incluir diferentes posibles niveles de la misma, estos serán:    
- Depósitos clave-valor distribuidos    
- Configuración del entorno con archivos de configuración    
- Configuración desde la línea de ordenes    

Para la primera opción tenemos que buscar depósitos clave-valor distribuidos que permitan, a diferentes copias de un mismo servicio, compartir todos los elementos de la configuración. Para esta labor se podría utilizar una base de datos como [MongoDB](https://www.mongodb.com/) o [Redis](https://redis.io/), pero me he centrado en aquellas que su uso específico sea este.
Los más usados son [etcd](https://etcd.io/) y [consul](https://www.consul.io/). Me he decantado por **Redis** ya que su instalación, despliegue y configuración es menos compleja que la de Consul.

Para la configuración del entorno con archivos de configuración, he buscado aquellos que se puedan usar con *Go* y he elegido elegido aquel que pueda ser usado en el mayor número de niveles sin usar otras bibliotecas. He encontrado [cleanenv](https://github.com/ilyakaznacheev/cleanenv), [viper](https://github.com/spf13/viper) y [godotenv](https://github.com/joho/godotenv).
Me he decantado por **Viper** ya que como dicen en su repositorio es una solución de la configuración completa y permite leer desde la línea de órdenes, variables de entorno y depósitos clave-valor distribuidos, mientras que ni **cleanenv** ni **godotenv** leen de depósitos clave-valor.
