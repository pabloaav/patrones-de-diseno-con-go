package main

import (
	"time"

	"github.com/Corrientes-Telecomunicaciones/patrones-de-diseno-con-go/comportamiento/state"
)

func main() {
	initialized := state.Initialized{}
	po := state.PurchaseOrder{
		Client: "Alexys Lozada",
		Date:   time.Now(),
		State:  initialized,
	}

	for po.State.Run(&po) {
	}
}
