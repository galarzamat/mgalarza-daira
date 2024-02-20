package routers

import (
	"ASTRIC/BackEnd/ASTRIC/helper/routes"
	rOperaciones "ASTRIC/BackEnd/api/operaciones/routes"

	"github.com/gorilla/mux"
)

// RutasPrincipales Manejador de rutas principales, donde se declaran los prefijos
func RutasPrincipales(rout *mux.Router) {
	// Crear un nuevo enrutador para las rutas de operaciones
	routOpereaciones := routes.NewPrefix(rout, "/operaciones", "Modulo de operaciones")
	routersInformes := routes.NewRouter(routOpereaciones)
	rOperaciones.RutasModulo(routersInformes)

}
