package builder

// Director: Construye el objeto utilizando la interface.
// el Director recibe cualquier producto concreto que cumpla con la interface IBuilder
// la clase directora puede ser un buen lugar donde colocar distintas rutinas de construcción para poder reutilizarlas a lo largo del programa.
type Director struct {
	builder IBuilder // builder es el nombre del atributo del director, que es de tipo IBuilder interface
}

// constructor. Podria llamarse NewBuilder
func (s *Director) SetBuilder(b IBuilder) {
	s.builder = b
}

// BuildMessage es el metodo que contruye un producto representado
// Ya que el Director esta compuesto por un producto concreto, este metodo setea los atributos del producto modelo.
// Devuelve un producto representado, segun sea el producto concreto que recibio en su constructor
// En terminos del ejemplo, el cuerpo de este metodo depende de la logica de construccion del metodo Build() de cada constructor concreto
func (s *Director) BuildMessage(recipient, message string) (*ProductRepresented, error) {
	// sea cual fuere el Producto Concreto, sabemos que podemos decirle que setee sus atributos
	// puesto que implementa una interface IBuilder
	s.builder.SetRecipient(recipient).SetMessage(message) // rutina de construccion
	// sea cual fuere el Producto Concreto, el metodo Build() hara una logica particular que adaptara los atributos recibidos, para crear un Producto Representado acorde al tipo de producto concreto actual
	return s.builder.Build()
}
