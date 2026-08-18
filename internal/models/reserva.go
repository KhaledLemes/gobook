package models

type Reserva struct {
	Propriedade Propriedade
	Quarto      Quarto

	Comodidades []Comodidades

	Reembolsavel bool

	Adultos  uint8
	Criancas uint8

	ValorNoite float64
	ValorTotal float64
	Taxas      float64
	Noites     uint8

	DataMarcado string
	Checkin     string
	Checkout    string
}
