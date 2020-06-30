package models

type kmeansResponse struct {
	Centroids []informe `json:"centroidsInforme"`
	Ncentroid []int     `json:"centroidsQuantity"`
}
