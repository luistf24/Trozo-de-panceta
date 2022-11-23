# **Gestor de dependencias**

La gestión de dependencias en go es semidescentralizada, permitiendo tratar a cada repositorio de git como un módulo más. No existe un gestor de dependencias como tal, y lo que se usa es **Go Module** donde las dependencias van en los archivos ``go.mod`` donde se define los module path de un módulo y los requisitos de las dependencias y los archivos ``go.sum`` que se encargan de gestionar nuestras versiones e indicar el hash de integridad de cada una de ellas.   

