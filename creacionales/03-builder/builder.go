package builder

// Builder: define la interface con los métodos que deberán cumplir los constructores.
// En este caso la abstraccion es una clase que se llama Message
type MessageBuilder interface {
	SetRecipient(recipient string) MessageBuilder
	SetMessage(message string) MessageBuilder
	Build() (*MessageRepresented, error)
}
