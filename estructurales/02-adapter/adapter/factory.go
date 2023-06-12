package adapter

import (
	"github.com/Corrientes-Telecomunicaciones/patrones-de-diseno-con-go/estructurales/02-adapter/auto"
	"github.com/Corrientes-Telecomunicaciones/patrones-de-diseno-con-go/estructurales/02-adapter/bici"
)

func Factory(s string) Adapter {
	switch s {
	case "automovil":
		d := auto.Automovil{}
		return &AutomovilAdapter{&d}
	case "bicicleta":
		d := bici.Bicicleta{}
		return &BicicletaAdapter{&d}
	}

	return nil
}
