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


# Ejecutar el programa
run:
	./$(BIN)/$(BINARY_NAME)


# Limpiar el proyecto
clean:
	@echo -e "Limpiando los binarios ..."
	rm $(BIN)/$(BINARY_NAME)
	go clean ./...


