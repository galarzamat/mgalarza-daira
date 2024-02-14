package controllers

import (
	"ASTRIC/BackEnd/api/operaciones/models"
	"ASTRIC/BackEnd/shared/ep"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var historialOp = []models.OpRensponse{}

// Prueba esta es una funcion de prueba
func Prueba(w http.ResponseWriter, r *http.Request) {

	defer ep.ErrorControlResponse("/operaciones/prueba", w, r)

	res := ep.NewResponse("Prueba de funcionamiento", w)

	var cliene models.Cliente

	cliene.Nombre = "Nombre"
	cliene.Edad = 45
	cliene.Apellido = "Apellido"

	res.DatoSend(cliene)

}

// Postoperacion Funcion que resuelve las operaciones y las guarda en el historial
func Postoperacion(c *gin.Context, w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("/operaciones/resolver", w, r)
	var NuevaOperacion models.OpRequest
	c.BindJSON(&NuevaOperacion)

	response := models.OpRensponse{}
	response.Fecha = time.Now().Format("2/1/2006")
	switch NuevaOperacion.Operador {
	case "+":
		response.Resultado = float64(NuevaOperacion.PrimerNumero) + float64(NuevaOperacion.SegundoNumero)
	case "-":
		response.Resultado = float64(NuevaOperacion.PrimerNumero) - float64(NuevaOperacion.SegundoNumero)
	case "*":
		response.Resultado = float64(NuevaOperacion.PrimerNumero) * float64(NuevaOperacion.SegundoNumero)
	case "/":
		if NuevaOperacion.SegundoNumero == 0 {
			response.Error = "No se puede dividir un numero por 0"
		} else {
			response.Resultado = float64(NuevaOperacion.PrimerNumero) / float64(NuevaOperacion.SegundoNumero)
		}
	case "%":
		response.Resultado = math.Mod(float64(NuevaOperacion.PrimerNumero), float64(NuevaOperacion.SegundoNumero))
	}
	response.Operacion = strconv.Itoa(NuevaOperacion.PrimerNumero) + NuevaOperacion.Operador + strconv.Itoa(NuevaOperacion.SegundoNumero)
	historialOp = append(historialOp, response)
	c.IndentedJSON(http.StatusCreated, response)
}

// Gethistorial da el historial de operaciones realizadas
func Gethistorial(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, historialOp)

}
