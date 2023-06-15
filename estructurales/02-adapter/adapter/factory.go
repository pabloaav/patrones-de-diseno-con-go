package adapter

import (
	"fmt"

	"github.com/Corrientes-Telecomunicaciones/patrones-de-diseno-con-go/estructurales/02-adapter/movilidad"
)

func MovilidadFactory(s string) (Adapter, error) {
	switch s {
	case "automovil":
		medio := movilidad.Automovil{}
		return &AutomovilAdapter{&medio}, nil
	case "bicicleta":
		medio := movilidad.Bicicleta{}
		return &BicicletaAdapter{&medio}, nil
	default:
		return nil, fmt.Errorf("%s", "error: lo ingresado no corresponde con ningun medio de movilidad")
	}
}
