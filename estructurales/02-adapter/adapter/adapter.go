package adapter

import "github.com/Corrientes-Telecomunicaciones/patrones-de-diseno-con-go/estructurales/02-adapter/movilidad"

type Adapter interface {
	Mover()
}

// todos los adaptadores implementan la inteface Adapter
type AutomovilAdapter struct {
	Auto *movilidad.Automovil
}

// Para mover el auto hay que encenderlo y acelerar
func (a *AutomovilAdapter) Mover() {
	if !a.Auto.Encendido {
		a.Auto.Encender()
	}

	a.Auto.Acelerar()
}

type BicicletaAdapter struct {
	Bici *movilidad.Bicicleta
}

// Para mover la bicicleta hay que encenderlo y acelerar
func (b *BicicletaAdapter) Mover() {
	b.Bici.Montar()
	b.Bici.Avanzar()
}
