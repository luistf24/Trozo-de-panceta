package internal

import (
	"encoding/json"
	"io/ioutil"
)

type LightInfo struct {
	brackets [24]float32
}

type Bracket struct {
	Prueba []Price `json:"brackets"`
}

type Price struct {
	Precio float32 `json:"price"`
}


func (l *LightInfo) generate(file string) (LightInfo, error){

	info, err := ioutil.ReadFile(file)
	if err != nil {
		return *l, &errorGenerateLightInfo{"error leyendo el archivo json"}
	}

	var responseObject Bracket

	err = json.Unmarshal(info, &responseObject)
	if err != nil {
		return *l, &errorGenerateLightInfo{"error haciendo Unmarshal del json"}
	}

	var temporal [24]float32
	for index, element := range responseObject.Prueba {
		temporal[index] = element.Precio
	}

	l.brackets = temporal
	return *l, nil
}


func (l *LightInfo) generateTimeBracket(i int) TimeBracket {

	var out TimeBracket
	if(i<24 && i>=0) {
		out.Hour	= i
		out.Price	= l.brackets[i]
		return out
	}

	out.Hour = 24
	return out
}

