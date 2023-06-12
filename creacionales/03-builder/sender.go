package builder

// Director: Construye el objeto utilizando la interface.
type Sender struct {
	builder MessageBuilder // builder es el nombre del atributo del sender, que es de tipo MessageBuilder interface
}

// constructor. Podria llamarse NewBuilder
func (s *Sender) SetBuilder(b MessageBuilder) {
	s.builder = b
}

// BuildMessage es el metodo que contruye un producto representado
// Ya que el Director esta compuesto por un producto concreto, este metodo setea los atributos del producto modelo.
// Devuelve un producto representado, segun sea el producto concreto que recibio en su constructor
func (s *Sender) BuildMessage(recipient, message string) (*MessageRepresented, error) {
	s.builder.SetRecipient(recipient).SetMessage(message)
	return s.builder.Build()
}
