package router

import controller "gobook/internal/handlers"

var loginRoute = []Route{
	{
		URI:        "/login",
		Metodo:     "GET",
		Func:       controller.Login,
		RequerAuth: false,
	},
}
