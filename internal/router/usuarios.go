package router

import (
	controller "gobook/internal/controllers"
)

var usersRoutes = []Route{
	{
		URI:        "/usuarios",
		Metodo:     "POST",
		Func:       controller.CriaUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios",
		Metodo:     "GET",
		Func:       controller.BuscaPorID,
		RequerAuth: true,
	},
}
