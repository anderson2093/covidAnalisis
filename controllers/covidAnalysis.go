package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func covidAnalysis(r http.ResponseWriter, request *http.Request) {
	//var bdy bodyCovid
	var bdy informe
	err := json.NewDecoder(request.Body).Decode(&bdy)
	if err != nil {
		http.Error(r, err.Error(), http.StatusBadRequest)
		fmt.Println("ERROR GAA")
		return
	}
	test := bdy
	//test := bdy.Informe
	fmt.Println(test)
	K := 1000
	Threads := 8

	xTest := []float32{test.AgeGroup, test.Sex, test.CardiovascularDisease, test.Diabetes, test.RespiratoryDisease,
		test.Hypertension, test.Cancer}
	fmt.Println("Age group: ", xTest[0])
	fmt.Println("Sex: ", xTest[1])
	fmt.Println("Cardiovascular Disease: ", xTest[2])
	fmt.Println("Diabetes: ", xTest[3])
	fmt.Println("Respiratory Disease: ", xTest[4])
	fmt.Println("Hypertension: ", xTest[5])
	fmt.Println("Cancer: ", xTest[6])

	dftemp := make([][]float32, len(xTrainCovid))
	for i := range xTrainCovid {
		dftemp[i] = make([]float32, len(xTrainCovid[i]))
		copy(dftemp[i], xTrainCovid[i])
	}
	dftemp = append(dftemp, xTest)
	//dftemp = normalize(dftemp)
	dftemp, _, _ = standarization(dftemp)

	newdftemp := make([][]float32, len(yTrainCovid))
	for i := 0; i < len(newdftemp); i++ {
		newdftemp[i] = make([]float32, len(dftemp[0]))
		copy(newdftemp[i], dftemp[i])
	}
	copy(xTest, dftemp[len(dftemp)-1])

	clase, ocurs = multiKNN(newdftemp, yTrainCovid, xTest, K, Threads)
	fmt.Println("La clase resultante es: ", clase)
}
