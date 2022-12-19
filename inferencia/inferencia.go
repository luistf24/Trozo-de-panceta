package inferencia

import (
	"math"
)

/**
 * @return: Array con 5 elementos donde se almacena:
 *	salida[0] -> Media de X
 *	salida[1] -> Media de Y
 *	salida[2] -> Varianza de X
 *	salida[3] -> Varianza de Y
 *	salida[4] -> Covarianza de X y de Y
 */
func calcReq(x []float32, y[]float32)([5]float32, error) {
	var salida [5]float32
	tamX := len(x)
	tamY := len(y)

	if(tamX != tamY) {
		return salida, &errorCalcReq{"los tamaños de x e y tienen que ser iguales"}
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
	salida[0] = tempx/float32(tamX)
	salida[1] = tempy/float32(tamX)

	// Cálculamos las varianzas de X e Y
	var powx float64
	var powy float64
	powx = 0
	powy = 0
	for i := 0; i < tamX; i++ {
		powx += math.Pow(float64(x[i]-salida[0]), 2)
		powy += math.Pow(float64(y[i]-salida[1]), 2)
	}
	salida[2] = float32(powx)/float32(tamX)
	salida[3] = float32(powy)/float32(tamX)

	// Cálculamos la covarianza
	tempx = 0
	for i := 0; i < tamX; i++ {
		tempx +=(x[i]-salida[0])*(y[i]-salida[1])
	}
	salida[4] = tempx/float32(tamX)

	return salida, nil
}



func calcRR(x []float32, y[]float32)(func(int)float32, error) {
	datos, err := calcReq(x,y)

	if err != nil {
		return nil, &errorCalcRR{" erro calculando los datos necesarios usando la función calcReq"}
	}

	return func(t int) float32 {
        return datos[4]*(float32(t)-datos[0])/datos[2]+datos[1]
    }, nil
}
