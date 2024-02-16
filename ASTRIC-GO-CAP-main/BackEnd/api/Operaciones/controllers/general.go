package controllers

import (
	"ASTRIC/BackEnd/api/operaciones/models"
	"ASTRIC/BackEnd/shared/ep"
	"encoding/json"
	"math"
	"net/http"
	"strconv"
)

var historialOp = []models.OpRensponse{}

// Postoperacion Funcion que resuelve las operaciones y las guarda en el historial
func Postoperacion(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("/operaciones/resolver", w, r)
	var NuevaOperacion models.OpRequest
	response := models.OpRensponse{}
	err := json.NewDecoder(r.Body).Decode(&NuevaOperacion)
	if err != nil {
		response.Error = err.Error()
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
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
		if NuevaOperacion.SegundoNumero == 0 {
			response.Error = "No se puede dividir por cero"
		} else {
			response.Resultado = float64(NuevaOperacion.PrimerNumero) / float64(NuevaOperacion.SegundoNumero)
		}
	case "%":
		response.Resultado = math.Mod(float64(NuevaOperacion.PrimerNumero), float64(NuevaOperacion.SegundoNumero))
	default:
		response.Error = "Operador no soportado"
	}

	response.Operacion = strconv.Itoa(NuevaOperacion.PrimerNumero) + NuevaOperacion.Operador + strconv.Itoa(NuevaOperacion.SegundoNumero)
	historialOp = append(historialOp, response) // Guardar la operación en el historial

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Gethistorial da el historial de operaciones realizadas
func Gethistorial(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("/operaciones/historial", w, r)
	json.NewEncoder(w).Encode(historialOp)
}
