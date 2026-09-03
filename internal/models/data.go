package models

var EstadosBrasil = []string{
	"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Distrito Federal",
	"Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul",
	"Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí",
	"Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia",
	"Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins",
}

var Categorias = []Categoria{
	Chale,
	Apartamento,
	Hotel,
	Pousada,
}

var Roles = []Role{
	RoleAdmin,
	RoleOwner,
	RoleGuest,
}
