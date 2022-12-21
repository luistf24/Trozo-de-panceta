package inferencia

import (
	"math"
)

type DatosRegresion struct {
	mediaX		float32
	mediaY		float32
	varianzaX	float32
	varianzaY	float32
	covarianza	float32
}

// La función calculateRequisites generará un struct con todos los datos necesarios
// para calcular una recta de regresión.
// Los parámetros de entrada son x e y dos arrays float32 que serán los datos obtenidos
// en el caso del proyecto x es los días e y son los precios de un tramo en esos días
func calculateRequisites(x []float32, y[]float32)(DatosRegresion, error) {
	var salida = new(DatosRegresion)
	tamX := len(x)
	tamY := len(y)

	if(tamX != tamY) {
		return *salida, &errorCalcReq{"los tamaños de x e y tienen que ser iguales"}
	}

	// Cálculamos las medias de X e Y
	var tempx float32
	var tempy float32
	tempx = 0
	tempy = 0
	for i := 0; i < tamX; i++ {
		tempx += x[i]
		tempy += y[i]
	}
	salida.mediaX = tempx/float32(tamX)
	salida.mediaY = tempy/float32(tamX)

	// Cálculamos las varianzas de X e Y
	var powx float64
	var powy float64
	powx = 0
	powy = 0
	for i := 0; i < tamX; i++ {
		powx += math.Pow(float64(x[i]-salida.mediaX), 2)
		powy += math.Pow(float64(y[i]-salida.mediaY), 2)
	}
	salida.varianzaX = float32(powx)/float32(tamX)
	salida.varianzaY = float32(powy)/float32(tamX)

	// Cálculamos la covarianza
	tempx = 0
	for i := 0; i < tamX; i++ {
		tempx +=(x[i]-salida.mediaX)*(y[i]-salida.mediaY)
	}
	salida.covarianza = tempx/float32(tamX)

	return *salida, nil
}


// Esta función genera a partir de dos array de datos, en el caso del proyecto serán días
// y precios de un tramo y devolverá una función que recibe como entrada un día y devuelve 
// una aproximación del tramo
func calculatePredictFunction(x []float32, y[]float32)(func(int)float32, error) {
	datos, err := calculateRequisites(x,y)

	if err != nil {
		return nil, &errorCalcRR{" error calculando los datos necesarios usando la función calcReq"}
	}

	if datos.varianzaX == 0 {
		return nil, &errorCalcRR{" error, la varianza de X no puede ser 0"}
	}

	return func(t int) float32 {
        return datos.covarianza*(float32(t)-datos.mediaX)/datos.varianzaX+datos.mediaY
	}, nil
}
