package client_one

import "github.com/Corrientes-Telecomunicaciones/patrones-de-diseno-con-go/creacionales/01-singleton/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
