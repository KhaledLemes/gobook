package models

type Comodidades struct {
	CafeDaManha    bool `json:"cafe_da_manha"`
	Estacionamento bool `json:"estacionamento"`
	Wifi           bool `json:"wifi"`
	VistaParaOMar  bool `json:"vista_para_omar"`
	FrenteAoMar    bool `json:"frente_ao_mar"`
}
