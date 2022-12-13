package internal

import (
	"encoding/json"
	"fmt"
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


func (l *LightInfo) generate(file string) LightInfo{

	info, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print("Error al abrir el JSON")
	}

	var responseObject Bracket

	json.Unmarshal(info, &responseObject)

	var temporal [24]float32
	for index, element := range responseObject.Prueba {
		temporal[index] = element.Precio
	}

	l.brackets = temporal
	return *l
}


func (l *LightInfo) generateTimeBracket(i int) TimeBracket {

	var out TimeBracket
	if(i<24 && i>=0) {
		out.Hour	= i
		out.Price	= l.brackets[i]
		return out
	}

	//Hago hour 24 porque como nunca va a valer eso me sirve para indicar que hay un error 
	out.Hour = 24
	return out
}

