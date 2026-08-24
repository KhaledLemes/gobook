package router

import controller "gobook/internal/controllers"

var loginRoute = Route{
	URI:        "/login",
	Metodo:     "GET",
	Func:       controller.Login,
	RequerAuth: false,
}
