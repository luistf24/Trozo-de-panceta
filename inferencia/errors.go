package inferencia
import "fmt"

// Error creado para comprobar la generación de medias, covarianza y varianzas
type errorCalcReq struct {
	data string
}

func (e *errorCalcReq) Error() string {
	return fmt.Sprintf("Error en CalcReq debido a... %s", e.data)
}
