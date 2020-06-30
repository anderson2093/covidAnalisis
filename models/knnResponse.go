package models

type knnResponse struct {
	ClassResult         int `json:"classResult"`
	DontHaveCoronavirus int `json:"dontHaveCoronavirus"`
	HaveCoronavirus     int `json:"haveCoronavirus"`
}
