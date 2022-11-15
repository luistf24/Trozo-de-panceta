# Gestor de tareas:
Para automatizar las tereas he mirado distintos gestores de tareas:
- **Mask:**  Mask es un gestor de tareas que está en markdown y este es uno de sus puntos fuertes, además al estar escrito en Rust su instalación es muy sencilla usando cargo, además se le puede inyectar código en diversos lenguajes.     

- **Make:** Make es un gestor de tareas que vienen por defecto en muchas distribuciones linux, es fácil de usar contiene reglas de dependencia, macros y las suffix rules. Además comprueba las fechas de modificación compilando solo los archivos que han sido actualizados reduciendo el tiempo de compilación.    

- **Task:** Task es un task runner diseñado para ser más simple y más fácil de usar que Make. Está escrito en go y entre sus ventajas están que es fácil de instalar, multiplataforma y al igual que make busca reducir el tiempo de compilación, compilando aquellos archivos que han sido modificados.     

## **Elección del gestor de tareas:**   
Una vez introducido los gestores más destacables que he encontrado, procedo a explicar por cual me he decantado y porqué.     
Los trest gestores de tareas son más que suficientes para mi proyecto, ya que el usar go facilita mucho las cosas.    
Mask fué una opción que sopesé, es fácil de instalar y usa Markdown, lo cual facilita mucho el aprendizaje de este, pero a pesar de ello ofrece demasiado para lo que yo busco y además la  notación y las reglas por mucho que estén en markdown me dejaba un archivo de configuración bastante sucio y poco legible.     
El siguiente fue Task y este me convenció mucho más, es muy similar a Make en lo bueno pero usa yml, si bien el archivo de configuración es muy legible y fácil de entender no acaba de tener todo lo que ofrece Make, como el uso de variables, macros, o por lo menos yo no he encontrado como hacerlo.     
Al final me decanté por Make, es más complejo que Task pero quizá porque ya venía de haberlo usado antes, me ha resultado mucho más comodo usarlo, y si bien tiene esa complejidad extra en muchos aspectos, para lo exigido en mi proyecto no aparece.
