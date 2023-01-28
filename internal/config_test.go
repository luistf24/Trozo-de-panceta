package internal

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {

	test, err := NewData()

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, test.LogLevel, "No se ha obtenido correctamente el secreto")
}
