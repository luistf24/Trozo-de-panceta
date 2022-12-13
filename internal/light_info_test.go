package internal

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"github.com/stretchr/testify/assert"
)


var temp LightInfo

func TestGetApi(t *testing.T) {

	var mensaje string
	var test bool

	test = true

	response, err := ioutil.ReadFile("../api/test.json")
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
	assert.NotEqual(t, len(temp.generate("../api/test.json").brackets), 0, "No hay datos almacenados de la API")
}


func TestSize(t *testing.T) {
	assert.Equal(t, 24, len(temp.brackets), "El número de elementos almacenados de la API tiene que ser 24")
}


func TestGenerateTimeBracketTime(t * testing.T) {
	var time TimeBracket
	time = temp.generateTimeBracket(0)

	assert.Equal(t, time.Hour, 0, "La hora no se ha almacenado correctamente en el tramo horario")
}


