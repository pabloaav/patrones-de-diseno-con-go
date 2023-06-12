package builder

// Product: Objeto principal construido. Representa el objeto bajo construcción.
// El Producto tiene las caracteristicas tales que se describen en esta estructura
type Message struct {
	Recipient string `json:"recipient" xml:"recipient"`
	Text      string `json:"text" xml:"text"`
}

// Reresentación del objeto. El Producto Representado adquiere otros atributos segun la actual estructura
type MessageRepresented struct {
	Body   []byte
	Format string
}
