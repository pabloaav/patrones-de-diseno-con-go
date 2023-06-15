package movilidad

import "fmt"

/* Bicicleta tiene un solo metodo propio por su caracteristica */
type Bicicleta struct {
	Marca string
	Color string
}

func (b *Bicicleta) Montar() {
	fmt.Println("Montando!")
}

func (b *Bicicleta) Avanzar() {
	fmt.Println("Avanzando!")
}
