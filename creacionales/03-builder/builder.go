package builder

// Builder o Interfaz Constructora: es una la interface con los métodos que deberán cumplir los constructores.
// Devolver la misma interface MessageBuilder en los metodos SetRecipient y SetMessage permite un patron en cadena
// Los metodos de estas interface son los "pasos de construccion"
type IBuilder interface {
	SetRecipient(recipient string) IBuilder // build step A
	SetMessage(message string) IBuilder     // build step B
	Build() (*ProductRepresented, error)    //  GetProduct, GetResult
}
