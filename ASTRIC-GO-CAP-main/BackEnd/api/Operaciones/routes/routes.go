package routes

import (
	"ASTRIC/BackEnd/api/operaciones/controllers"
	"ASTRIC/BackEnd/shared/routes"
)

func rutas(ruta routes.TipoRuta) {
	ruta("/historial", controllers.Gethistorial, "GET", "Ruta de historial")
	ruta("/resolver", controllers.Postoperacion, "POST", "Ruta de Resolver operaciones")
}
