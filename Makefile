##############################
# MAKE FILE TROZO-DE-PANCETA #
##############################


#
# Paths y nombres de directorios
#

# Nombre del archivo binario
BINARY_NAME=trozo-de-panceta


# Ubicación donde irá el binario
BIN = ./bin


# Ubicación donde irá el código
CODE_FOLDERS = internal/*

#
# Instrucciones
#

help: # Imprime mensaje con las distintas opciones del Makefile
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done


build: # Construir el programa principal
	@echo -e "Construyendo el proyecto ..."
	go build ./...


installdeps: # Instalación de las dependencias
	@echo -e "Instalando las dependencias ..."
	go mod tidy


run: # Ejecutar el programa
	./$(BIN)/$(BINARY_NAME)


clean: # Limpiar el proyecto
	@echo -e "Limpiando los binarios ..."
	rm $(BIN)/$(BINARY_NAME)
	go clean ./...


check: # Comprobación de la sintaxis
	@echo -e "Comprobando sintaxis del proyecto ..."
	 gofmt -l $(CODE_FOLDERS)

test: # Ejecución de todos los tests del proyecto
	@echo -e "Ejecutando tests ..."
	go test ./... -v

docker: # Ejecución desde un docker
	@echo -e "Ejecutando desde docker ..."
	docker run -u 1001 -t -v `pwd`:/app/test luistf24/trozo-de-panceta

