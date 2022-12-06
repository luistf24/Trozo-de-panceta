package internal

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type LightInfo struct {
	brackets map[int]float32
}

type Bracket struct {
	Tramo0 Price `json:"00-01"`
	Tramo1 Price `json:"01-02"`
	Tramo2 Price `json:"02-03"`
	Tramo3 Price `json:"03-04"`
	Tramo4 Price `json:"04-05"`
	Tramo5 Price `json:"05-06"`
	Tramo6 Price `json:"06-07"`
	Tramo7 Price `json:"07-08"`
	Tramo8 Price `json:"08-09"`
	Tramo9 Price `json:"09-10"`
	Tramo10 Price `json:"10-11"`
	Tramo11 Price `json:"11-12"`
	Tramo12 Price `json:"12-13"`
	Tramo13 Price `json:"13-14"`
	Tramo14 Price `json:"14-15"`
	Tramo15 Price `json:"15-16"`
	Tramo16 Price `json:"16-17"`
	Tramo17 Price `json:"17-18"`
	Tramo18 Price `json:"18-19"`
	Tramo19 Price `json:"19-20"`
	Tramo20 Price `json:"20-21"`
	Tramo21 Price `json:"21-22"`
	Tramo22 Price `json:"22-23"`
	Tramo23 Price `json:"23-24"`
}

type Price struct {
	Precio float32 `json:"price"`
}


func (l *LightInfo) generate() LightInfo{

	l.brackets = make(map[int]float32)

	response, err := http.Get("https://api.preciodelaluz.org/v1/prices/all?zone=PCB")
	if err != nil {
		fmt.Print("Un error ha ocurrido")
	}


	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print("Un error ha ocurrido 2")
	}

	defer response.Body.Close()

	var responseObject Bracket

	json.Unmarshal(responseData, &responseObject)

	l.brackets[0] = responseObject.Tramo0.Precio
	l.brackets[1] = responseObject.Tramo1.Precio
	l.brackets[2] = responseObject.Tramo2.Precio
	l.brackets[3] = responseObject.Tramo3.Precio
	l.brackets[4] = responseObject.Tramo4.Precio
	l.brackets[5] = responseObject.Tramo5.Precio
	l.brackets[6] = responseObject.Tramo6.Precio
	l.brackets[7] = responseObject.Tramo7.Precio
	l.brackets[8] = responseObject.Tramo8.Precio
	l.brackets[9] = responseObject.Tramo9.Precio
	l.brackets[10] = responseObject.Tramo10.Precio
	l.brackets[11] = responseObject.Tramo11.Precio
	l.brackets[12] = responseObject.Tramo12.Precio
	l.brackets[13] = responseObject.Tramo13.Precio
	l.brackets[14] = responseObject.Tramo4.Precio
	l.brackets[15] = responseObject.Tramo15.Precio
	l.brackets[16] = responseObject.Tramo16.Precio
	l.brackets[17] = responseObject.Tramo17.Precio
	l.brackets[18] = responseObject.Tramo18.Precio
	l.brackets[19] = responseObject.Tramo19.Precio
	l.brackets[20] = responseObject.Tramo20.Precio
	l.brackets[21] = responseObject.Tramo21.Precio
	l.brackets[22] = responseObject.Tramo22.Precio
	l.brackets[23] = responseObject.Tramo23.Precio

	return *l
}
