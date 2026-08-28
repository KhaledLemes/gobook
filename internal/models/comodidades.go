package models

type Comodidades struct {
	CafeDaManha    bool `json:"cafe_da_manha"`
	Estacionamento bool `json:"estacionamento"`
	Wifi           bool `json:"wifi"`
	VistaMar       bool `json:"vista_mar"`
	FrenteAoMar    bool `json:"frente_ao_mar"`
}
