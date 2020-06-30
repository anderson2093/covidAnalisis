package models

type pacient struct {
	AgeGroup              float32 `json:"age_group"`
	Sex                   float32 `json:"sex"`
	CardiovascularDisease float32 `json:"cardiovascular_disease"`
	Diabetes              float32 `json:"diabetes"`
	RespiratoryDisease    float32 `json:"respiratory_disease"`
	Hypertension          float32 `json:"hypertension"`
	Cancer                float32 `json:"cancer"`
}
