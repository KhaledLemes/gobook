package models

type Reserva struct {
	Propriedade Propriedade `json:"propriedad"`
	Quarto      Quarto      `json:"quarto"`

	Comodidades []Comodidades `json:"comodidades"`

	Reembolsavel bool `json:"reembolsavel"`

	Adultos  uint8 `json:"adultos"`
	Criancas uint8 `json:"criancas"`

	ValorNoite float64 `json:"valor_noite"`
	ValorTotal float64 `json:"valor_total"`
	Taxas      float64 `json:"taxas"`
	Noites     uint8   `json:"noites"`

	DataMarcado string `json:"data_marcado"`
	Checkin     string `json:"checkin"`
	Checkout    string `json:"checkout"`
}
