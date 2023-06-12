package builder

// Director: Construye el objeto utilizando la interface.
type Sender struct {
	builder MessageBuilder // builder es el nombre del atributo del sender, que es de tipo MessageBuilder interface
}

// constructor
func (s *Sender) SetBuilder(b MessageBuilder) {
	s.builder = b
}

// BuildMessage a concrete message via MessageBuilder
func (s *Sender) BuildMessage(recipient, message string) (*MessageRepresented, error) {
	s.builder.SetRecipient(recipient).SetMessage(message)
	return s.builder.Build()
}
