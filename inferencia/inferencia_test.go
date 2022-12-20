package inferencia

import (
	"Trozo-de-panceta/internal"
	"testing"
	"github.com/stretchr/testify/assert"
)


type PruebaFixture struct {
	prueba [5]float32
	err error
}

func NewPrueba() PruebaFixture {
	var salida = new(PruebaFixture)

	x := []float32{1.1, 1.3, 1.4}
	y := []float32{2.1, 2.1, 2.5}

	salida.prueba, salida.err = calculateRequisites(x,y)
	return *salida
}

var fixture = NewPrueba()


func TestCargaInternal(t *testing.T) {
	var test = internal.Appliance{Id:"Prueba", Consumption:1, UsageTime:1}
	assert.Equal(t, test.Id, "Prueba", "No se ha cargado correctamente el paquete internal del proyecto")
}


func TestCalcReq(t *testing.T) {
	if fixture.err != nil {
		t.Fatal(fixture.err)
	}
}


func TestMediaX(t *testing.T) {
	assert.LessOrEqual(t, float32(0), fixture.prueba[0], "La media de números positivos tiene que ser mayor que 0")
}


func TestMediaY(t *testing.T) {
	assert.LessOrEqual(t, float32(0), fixture.prueba[1], "La media de números positivos tiene que ser mayor que 0")
}


func TestVarianzaX(t *testing.T) {
	assert.LessOrEqual(t, float32(0), fixture.prueba[2], "La varianza es siempre positiva")
}


func TestVarianzaY(t *testing.T){
	assert.LessOrEqual(t, float32(0), fixture.prueba[3], "La varianza es siempre postiva")
}


func TestCalcRR(t *testing.T) {
	dias			:= []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	precioTramo13	:= []float32{122.1, 211.1, 123.5, 111.4, 167.5, 300.5, 299.1, 301.5, 302.1, 202.1, 222.1, 300.4}

	recta, err := calculateRegressionLine(dias, precioTramo13)
	if err != nil {
		t.Fatal(err)
	}

	assert.LessOrEqual(t, float32(0), recta(13), "Por los datos introducidos la recta no debería tocar al 0")
}
