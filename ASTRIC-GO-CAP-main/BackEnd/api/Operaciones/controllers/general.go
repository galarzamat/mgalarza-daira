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
	response.Fecha = date.Format("02/01/2006")

	err := json.NewDecoder(r.Body).Decode(&NuevaOperacion)
	if err != nil {
		res.ErrSend("Error al enviar el json")
		return
	}

	switch NuevaOperacion.Operador {
	case "+":
		response.Resultado = float64(NuevaOperacion.PrimerNumero) + float64(NuevaOperacion.SegundoNumero)
	case "-":
		response.Resultado = float64(NuevaOperacion.PrimerNumero) - float64(NuevaOperacion.SegundoNumero)
	case "*":
		response.Resultado = float64(NuevaOperacion.PrimerNumero) * float64(NuevaOperacion.SegundoNumero)
	case "/":
		response.Resultado = float64(NuevaOperacion.PrimerNumero) / float64(NuevaOperacion.SegundoNumero)
	case "%":
		response.Resultado = math.Mod(float64(NuevaOperacion.PrimerNumero), float64(NuevaOperacion.SegundoNumero))
	}

	response.Operacion = strconv.Itoa(NuevaOperacion.PrimerNumero) + NuevaOperacion.Operador + strconv.Itoa(NuevaOperacion.SegundoNumero)
	historialOp = append(historialOp, response) // Guardar la operación en el historial

	w.Header().Set("Content-Type", "application/json")
	res.DatoSend(response)

}

// Gethistorial da el historial de operaciones realizadas
func Gethistorial(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("/operaciones/historial", w, r)
	json.NewEncoder(w).Encode(historialOp)
}
