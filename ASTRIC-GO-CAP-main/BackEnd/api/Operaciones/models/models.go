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
	Error     string  `json:"error"`
	Resultado float64 `json:"resultado"`
	Fecha     string  `json:"fecha"`
}
type Cliente struct {
	Nombre   string `json:"cliente_nombre"`
	Apellido string `json:"cliente_apellido"`
	Edad     int    `json:"cliente_edad"`
}
