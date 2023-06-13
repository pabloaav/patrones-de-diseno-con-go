package builder_test

import (
	"testing"

	builder "github.com/Corrientes-Telecomunicaciones/patrones-de-diseno-con-go/creacionales/03-builder"
)

/* la sintaxis del nombre de un test debe ser:
- La palabra Test seguida del nombre de la clase o estructura;
- seguida del nombre del metodo a testear
*/

func TestSender_BuildMessage(t *testing.T) {
	// variable tipo apuntador a producto concreto tipo JSON
	json := &builder.JSONMessageBuilder{}
	// variable tipo apuntador a producto concreto tipo XML
	xmlf := &builder.XMLMessageBuilder{}
	// La estructura directora ayuda a organizar el proceso de creación.
	director := &builder.Director{}

	// Paso 1: metodo constructor del Director.
	// Recibe un producto concreto que satisface la IBuilder, que es el contrato o interfaz
	// El cliente sólo necesita asociar un objeto constructor con una clase directora
	director.SetBuilder(json)                                      // asociar concrete product json con director
	jsonMsg, err := director.BuildMessage("Gopher", "Hola mundo!") // director.make()
	if err != nil {
		t.Fatalf("BuildMesage(): JSON: no se esperaba error, se recibió: %v", err)
	}

	t.Log(string(jsonMsg.Body))

	director.SetBuilder(xmlf)                                       // asociar concrete product xml con director
	xmlMsg, err := director.BuildMessage("XMLGopher", "Hola mundo") // director.make()
	if err != nil {
		t.Fatalf("BuildMesage(): XML: no se esperaba error, se recibió: %v", err)
	}

	t.Log(string(xmlMsg.Body))
}
