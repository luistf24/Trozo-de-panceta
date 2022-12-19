package inferencia

import (
	"Trozo-de-panceta/internal"
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestCargaInternal(t *testing.T) {
	var test = internal.Appliance{Id:"Prueba", Consumption:1, UsageTime:1}
	assert.Equal(t, test.Id, "Prueba", "No se ha cargado correctamente el paquete internal del proyecto")
}


func TestCalcReq(t *testing.T) {
	var x []float32
	var y []float32

	_, err := calcReq(x,y)

	if err != nil {
		t.Fatal(err)
	}
}


