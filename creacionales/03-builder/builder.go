package builder

// Builder: es una la interface con los métodos que deberán cumplir los constructores.
// En este caso la abstraccion es una clase que se llama Message.
// Devolver la misma interface MessageBuilder en los metodos SetRecipient y SetMessage permite un patron en cadena
type IBuilder interface {
	SetRecipient(recipient string) IBuilder
	SetMessage(message string) IBuilder
	Build() (*ProductRepresented, error)
}
