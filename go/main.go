package main

import (
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OpRequest struct {
	Operador      string `json:"operador"`
	PrimerNumero  int    `json:"primernumero"`
	SegundoNumero int    `json:"segundonumero"`
}
type OpRensponse struct {
	Operacion string  `json:"operacion"`
	Error     string  `json:"error"`
	Resultado float64 `json:"resultado"`
	Fecha     string  `json:"fecha"`
}

var historialOp = []OpRensponse{}

func postOperacion(c *gin.Context) {
	var NuevaOperacion OpRequest
	c.BindJSON(&NuevaOperacion)

	response := OpRensponse{}
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

func getHistorial(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, historialOp)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/historial", getHistorial)
	router.POST("/resolver", postOperacion)
	router.Run("localhost:8080")

}
