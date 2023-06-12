package builder_test

import (
	"testing"

	builder "github.com/Corrientes-Telecomunicaciones/patrones-de-diseno-con-go/creacionales/03-builder"
)

/* la sintaxis del nombre de un test debe ser:
- La palabbra Test seguida del nombre de la clase o estructura;
- seguida del nombre del metodo a testear
*/

func TestSender_BuildMessage(t *testing.T) {
	// variable tipo apuntador a producto concreto tipo JSON
	json := &builder.JSONMessageBuilder{}
	// variable tipo apuntador a producto concreto tipo XML
	xml := &builder.XMLMessageBuilder{}
	// variable tipo apuntador a struct Sender
	sender := &builder.Sender{}

	// Paso 1: metodo constructor del Director.
	// Recibe un producto concreto que satisface el builder, que es el contrato o interfaz
	sender.SetBuilder(json)
	jsonMsg, err := sender.BuildMessage("Gopher", "Hola mundo!")
	if err != nil {
		t.Fatalf("BuildMesage(): JSON: no se esperaba error, se recibió: %v", err)
	}

	t.Log(string(jsonMsg.Body))

	sender.SetBuilder(xml)
	xmlMsg, err := sender.BuildMessage("Gopher", "Hola mundo")
	if err != nil {
		t.Fatalf("BuildMesage(): XML: no se esperaba error, se recibió: %v", err)
	}

	t.Log(string(xmlMsg.Body))
}
