package builder

import "encoding/json"

/* Concrete Builder: Implementa la interface builder para entregar el objeto concreto.
Un Constructor Concreto implementa la interfaz del Builder, con todos los metodos propios  referidos a la construccion
y necesidad de un Product. En este caso un destinatario y un texto para el mensaje.
Devuelve el Producto Representado, es decir, un producto concreto que adopta la forma de la struct MessageRepresented,
en este ejemplo.
La funcion Build() no tiene parametros. Es la encargada de tomar los atributos de un producto concreto y adaptarlos,
a la struct de un Producto Representado.
Cada struct del producto concreto usa composicion con la strcuct del producto modelo u objeto bajo construcción.
*/

// JSONMessageBuilder is concrete builder
type JSONMessageBuilder struct {
	message Product
}

// SetRecipient asigna el destinatario del mensaje
func (b *JSONMessageBuilder) SetRecipient(recipient string) IBuilder {
	b.message.Recipient = recipient
	return b
}

// SetMessage asigna el mensaje a enviar
func (b *JSONMessageBuilder) SetMessage(text string) IBuilder {
	b.message.Text = text
	return b
}

// Build construye el objeto y lo representa en JSON
// Cada Producto Concreto construye su propia representacion al modo que mejor le convenga
func (b *JSONMessageBuilder) Build() (*ProductRepresented, error) {
	data, err := json.Marshal(b.message)
	if err != nil {
		return nil, err
	}

	return &ProductRepresented{Body: data, Format: "JSON"}, nil
}
