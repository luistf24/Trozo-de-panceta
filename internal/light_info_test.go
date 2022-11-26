package internal

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	prueba := new(LightInfo);

	assert.Empty(t, prueba.generate(), "La obtención de datos de la API ha fallado")
}

func TestSize(t *testing.T) {
	prueba := new(LightInfo);

	assert.Equal(t, 24, prueba.generate(), "El número de elementos almacenados de la API tiene que ser 24")
}
