package local

import (
	"github.com/Corrientes-Telecomunicaciones/patrones-de-diseno-con-go/estructurales/01-proxy/book"
)

type Proxy interface {
	GetByID(ID uint) book.Book
	GetAll() book.Books
}
