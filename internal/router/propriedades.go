package router

import controller "gobook/internal/controllers"

var propriedadesRouter = []Route{
	{
		URI:        "/propriedades",
		Metodo:     "GET",
		Func:       controller.MostraTodasPropriedas,
		RequerAuth: false,
	},
	{
		URI:        "/propriedades/{id}",
		Metodo:     "GET",
		Func:       controller.BuscaPropriedadePorID,
		RequerAuth: false,
	},
	{
		URI:        "/propriedades",
		Metodo:     "POST",
		Func:       controller.CriarPropriedade,
		RequerAuth: true,
	},
	{
		URI:        "/propriedades",
		Metodo:     "UPDATE",
		Func:       controller.EditarPropriedade,
		RequerAuth: true,
	},
	{
		URI:        "/propriedades/{id}",
		Metodo:     "DELETE",
		Func:       controller.DeletaPropriedadePorID,
		RequerAuth: true,
	},
}
