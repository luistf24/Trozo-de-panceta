package internal

import "fmt"

// Error creado para comprobar la generación de LightInfo
type errorGenerateLightInfo struct {
	data string
}

func (e *errorGenerateLightInfo) Error() string {
	return fmt.Sprintf("Error en generateLightInfo debido a... %s", e.data)
}


//Error creado para comprobar la generación de TimeBracket a partir de LightInfo
type errorGenerateTimeBracket struct {
	data string
}

func (e *errorGenerateTimeBracket) Error() string {
	return fmt.Sprintf("Error en generateTimeBracket debido a... %s", e.data)
}
