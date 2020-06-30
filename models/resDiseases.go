package models

type resDiseases struct {
	Centroids []pacient `json:"centroids"`
	Ncentroid []int     `json:"ncentroid"`
}
