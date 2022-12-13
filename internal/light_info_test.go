package internal

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var temp LightInfo

func TestErrorLightInfo(t *testing.T) {
	_, err := temp.generate("../api/test.json")

	if err != nil {
		t.Fatal(err)
	}
}


func TestIsEmpty(t *testing.T) {
	assert.NotEqual(t, len(temp.brackets), 0, "No hay datos almacenados de la API")
}


func TestSize(t *testing.T) {
	assert.Equal(t, 24, len(temp.brackets), "El número de elementos almacenados de la API tiene que ser 24")
}


func TestGenerateTimeBracketTime(t * testing.T) {
	var time TimeBracket
	time = temp.generateTimeBracket(0)

	assert.Equal(t, time.Hour, 0, "La hora no se ha almacenado correctamente en el tramo horario")
}


