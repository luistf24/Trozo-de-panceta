package inferencia
import "fmt"

// Error creado para comprobar la generación de medias, covarianza y varianzas
type errorCalcReq struct {
	data string
}

func (e *errorCalcReq) Error() string {
	return fmt.Sprintf("Error en CalcReq debido a... %s", e.data)
}


// Error creado para comprobar la generación de rectas de regresión
type errorCalcRR struct {
	data string
}

func (e *errorCalcRR) Error() string {
	return fmt.Sprintf("Error en CalcRR debido a... %s", e.data)
}

