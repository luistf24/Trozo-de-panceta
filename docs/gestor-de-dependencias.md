# **Gestor de dependencias**

Los gestores de dependencias más usados en go son: **Vendoring**, **Dep** y **Go Modules**   
Me he decantado por **Go Modules** ya que con **Vendoring** hacer upgrade y downgrade de un paquete puede ser muy complejo, puede no funcionar para las versiones de algunos paquetes e incrementa el tamaño del código y con **Dep** tenemos que no es nativo, puede tener problemas con las interdependencias y no se describe muy bien cuales dependencias son directas e indirectas.     
**Go Modules** no tiene ninguno de estos problemas, además es el más utilizado.
