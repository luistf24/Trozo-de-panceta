package internal

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetApi(t *testing.T) {
	response, err := ioutil.ReadFile("./test.json")

	var mensaje string
	var test bool

	test = true

	if err != nil {
		mensaje = "Error al abrir el json"
		test = false
	}

	var bracket Bracket

	err = json.Unmarshal(response, &bracket)
	if err != nil {
		mensaje = "Error durante Unmarshal"
		test = false
	}

	assert.Equal(t, true, test, mensaje)
}

func TestIsEmpty(t *testing.T) {
	assert.NotEqual(t, len(generate().brackets), 0, "No hay datos almacenados de la API")
}

func TestSize(t *testing.T) {

	assert.Equal(t, 24, len(generate().brackets), "El número de elementos almacenados de la API tiene que ser 24")
}

