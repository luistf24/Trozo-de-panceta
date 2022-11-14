##############################
# MAKE FILE TROZO-DE-PANCETA #
##############################

# shell para el que está escrito este Makefile
SHELL = /bin/bash


#
# Paths y nombres de directorios
#

# Nombre del archivo binario
BINARY_NAME=trozo-de-panceta

# Ubicación donde irá main
MAIN = ./cmd/trozo-de-panceta/main.go

# Ubicación donde irá el binario
BIN = ./bin



#
# Instrucciones
#


# Construir el programa principal
build:
	@echo -e "Construyendo el proyecto ..."
	go build -o $(BIN)/$(BINARY_NAME) $(MAIN)


# Instalación de las dependencias
installdeps:
	@echo -e "Instalando las dependencias ..."
	go mod tidy


# Ejecutar el programa
run:
	./$(BIN)/$(BINARY_NAME)


# Limpiar el proyecto
clean:
	@echo -e "Limpiando los binarios ..."
	rm $(BIN)/$(BINARY_NAME)
	go clean ./...


# Comprobación de la sintaxis
check:
	@echo -e "Comprobando sintaxis del proyecto ..."
	go vet ./...


# Ejecución de todos los tests del proyecto
test:
	@echo -e "Ejecutando tests ..."
	go test ./...


# Imprime mensaje con las distintas opciones del Makefile
help:
	@echo -e "Las opciones disponibles son:"
	@echo -e "  - make build "
	@echo -e "  - make installdeps"
	@echo -e "  - make run"
	@echo -e "  - make clean"
	@echo -e "  - make check"
	@echo -e "  - make test"
	@echo -e "  - make help"


