package models

//Oprequest estructura para realizar operaciones
type OpRequest struct {
	Operator     string `json:"operator"`
	FirstNumber  int    `json:"firstnumber"`
	SecondNumber int    `json:"secondnumber"`
}

//OpRensponse estructura para mostrar las operaciones
type OpRensponse struct {
	Operation string  `json:"operation"`
	Result    float64 `json:"result"`
	Date      string  `json:"date"`
}
