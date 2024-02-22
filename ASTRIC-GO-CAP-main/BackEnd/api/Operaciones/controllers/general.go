package controllers

import (
	"ASTRIC/BackEnd/api/operaciones/models"
	"ASTRIC/BackEnd/shared/ep"
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"time"
)

var historialOp = []models.OpRensponse{}

// Postoperacion Funcion que resuelve las operaciones y las guarda en el historial
func Postoperacion(w http.ResponseWriter, r *http.Request) {

	res := ep.NewResponse("Operacion", w)
	defer ep.ErrorControlResponse("/operaciones/resolver", w, r)

	var NuevaOperacion models.OpRequest
	response := models.OpRensponse{}

	date := time.Now()
	response.Date = date.Format("02/01/2006")

	err := json.NewDecoder(r.Body).Decode(&NuevaOperacion)
	if err != nil {
		res.ErrSend("Error al enviar el json")
		return
	}

	switch NuevaOperacion.Operator {
	case "+":
		response.Result = float64(NuevaOperacion.FirstNumber) + float64(NuevaOperacion.SecondNumber)
	case "-":
		response.Result = float64(NuevaOperacion.FirstNumber) - float64(NuevaOperacion.SecondNumber)
	case "*":
		response.Result = float64(NuevaOperacion.FirstNumber) * float64(NuevaOperacion.SecondNumber)
	case "/":
		response.Result = float64(NuevaOperacion.FirstNumber) / float64(NuevaOperacion.SecondNumber)
	case "%":
		response.Result = math.Mod(float64(NuevaOperacion.FirstNumber), float64(NuevaOperacion.SecondNumber))
	}

	response.Operation = strconv.Itoa(NuevaOperacion.FirstNumber) + NuevaOperacion.Operator + strconv.Itoa(NuevaOperacion.SecondNumber)
	historialOp = append(historialOp, response) // Guardar la operación en el historial

	w.Header().Set("Content-Type", "application/json")
	res.DatoSend(response)

}

// Gethistorial da el historial de operaciones realizadas
func Gethistorial(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("/operaciones/historial", w, r)

	res := ep.NewResponse("Operacion", w)
	res.DatoSend(historialOp)
}
