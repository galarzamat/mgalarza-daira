package models

//Oprequest estructura para realizar operaciones
type OpRequest struct {
	Operador      string `json:"operador"`
	PrimerNumero  int    `json:"primernumero"`
	SegundoNumero int    `json:"segundonumero"`
}

//OpRensponse estructura para mostrar las operaciones
type OpRensponse struct {
	Operacion string  `json:"operacion"`
	Resultado float64 `json:"resultado"`
	Fecha     string  `json:"fecha"`
}
