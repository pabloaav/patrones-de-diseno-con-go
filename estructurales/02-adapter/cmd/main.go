package main

import (
	"fmt"
	"log"

	"github.com/Corrientes-Telecomunicaciones/patrones-de-diseno-con-go/estructurales/02-adapter/adapter"
)

func main() {
	var t string
	fmt.Print("ingrese tipo de transporte o medio de movilidad: ")
	_, err := fmt.Scanln(&t)
	if err != nil {
		log.Fatalf("no se pudo leer el medio de transporte: %v", err)
	}
	// en este punto el objeto es desconocido. Se crea a partir de lo ingresado
	// y se obtiene un objeto que cumple con la interface del Adapter
	// eso permite usarlo de acuerdo a sus caracteristicas
	transporteAdapter, err := adapter.MovilidadFactory(t)
	if err != nil {
		log.Fatalf("%v", err)
	}
	// el objeto se comporta al metodo Mover() de acuerdo o segun sea su naturaleza o tipo
	// se adapta al movimiento segun sus posibilidades
	transporteAdapter.Mover()
}
