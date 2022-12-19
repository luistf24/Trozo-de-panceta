package inferencia

import (
	"Trozo-de-panceta/internal"
	"testing"
	"github.com/stretchr/testify/assert"
)


var prueba [5]float32

func TestCargaInternal(t *testing.T) {
	var test = internal.Appliance{Id:"Prueba", Consumption:1, UsageTime:1}
	assert.Equal(t, test.Id, "Prueba", "No se ha cargado correctamente el paquete internal del proyecto")
}


func TestCalcReq(t *testing.T) {
	x := []float32{1.1, 1.3, 1.4}
	y := []float32{2.1, 2.1, 2.5}

	temp, err := calcReq(x,y)
	prueba = temp
	if err != nil {
		t.Fatal(err)
	}
}


func TestMediaX(t *testing.T) {
	assert.LessOrEqual(t, float32(0), prueba[0], "La media de números positivos tiene que ser mayor que 0")
}


func TestMediaY(t *testing.T) {
	assert.LessOrEqual(t, float32(0), prueba[1], "La media de números positivos tiene que ser mayor que 0")
}


func TestVarianzaX(t *testing.T) {
	assert.LessOrEqual(t, float32(0), prueba[2], "La varianza es siempre positiva")
}


func TestVarianzaY(t *testing.T){
	assert.LessOrEqual(t, float32(0), prueba[3], "La varianza es siempre postiva")
}


func TestCalcRR(t *testing.T) {
	x := []float32{1.1, 1.3, 1.4}
	y := []float32{2.1, 2.1, 2.5}

	recta, err := calcRR(x,y)
	if err != nil {
		t.Fatal(err)
	}

	assert.LessOrEqual(t, float32(0), recta(1), "Por los datos introducidos la recta no debería tocar al 0")
}
