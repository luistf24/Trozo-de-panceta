package internal

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


type LightInfoFixture struct {
	temp LightInfo
}

func NewFixture() LightInfoFixture {
	var salida = new (LightInfoFixture)

	_, _ = salida.temp.generate("../data/test.json")

	return *salida
}

var fixture = NewFixture()


func TestErrorLightInfo(t *testing.T) {
	var test = new(LightInfo)

	_, err:= test.generate("../data/test.json")

	if err != nil {
		t.Fatal(err)
	}
}


func TestIsEmpty(t *testing.T) {
	assert.NotEqual(t, len(fixture.temp.brackets), 0, "No hay datos almacenados de la API")
}


func TestSize(t *testing.T) {
	assert.Equal(t, 24, len(fixture.temp.brackets), "El número de elementos almacenados de la API tiene que ser 24")
}


func TestGenerateTimeBracketTime(t * testing.T) {
	time, err := fixture.temp.generateTimeBracket(0)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, time.Hour, 0, "La hora no se ha almacenado correctamente en el tramo horario")
}


