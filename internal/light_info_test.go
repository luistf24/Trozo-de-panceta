package internal

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestGetApi(t *testing.T) {
	response, err := http.Get("https://api.preciodelaluz.org/v1/prices/all?zone=PCB")

	var mensaje string
	var test bool

	test = true

	if err != nil {
		mensaje = "Error al establecer la conexión con la API"
		test = false
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		mensaje = "Error al obtener los datos de la API"
		test = false
	}

	assert.Equal(t, true, test, mensaje)
}

func TestIsEmpty(t *testing.T) {
	prueba := new(LightInfo);

	assert.NotEqual(t, len(prueba.generate()), 0, "No hay datos almacenados de la API")
}

func TestSize(t *testing.T) {
	prueba := new(LightInfo);

	assert.Equal(t, 24, prueba.generate(), "El número de elementos almacenados de la API tiene que ser 24")
}

