package list

import (
	"github.com/GAQF202/servidor-rest/Structs"
)

type NodoCola struct {
	//CADA DATO DE LA COLA ES UN PRODUCTO
	product   Structs.CodeProduct
	siguiente *NodoCola
}

func newNodoCola(producto Structs.CodeProduct) *NodoCola {
	return &NodoCola{producto, nil}
}

type Queue struct {
	first    *NodoCola
	last     *NodoCola
	contador int
}

func NewQueue() *Queue {
	return &Queue{nil, nil, 0}
}

func (q *Queue) isVoid() bool {
	return q.first == nil
}

func (q *Queue) Add(producto *Structs.CodeProduct) {
	nuevo := newNodoCola(*producto)
	if q.isVoid() {
		q.first = nuevo
	} else {
		q.last.siguiente = nuevo
	}
	q.last = nuevo
	q.contador++
}

func (q *Queue) Pop() *Structs.CodeProduct {
	aux := q.first.product

	q.first = q.first.siguiente
	q.contador--
	return &aux
}

var Productos Structs.CodigoDeProducto

func (q *Queue) Recorrer() {
	aux := q.first

	for aux != nil {
		Productos = append(Productos, aux.product.Codigo_producto)
		aux = aux.siguiente
	}
	//return productos
}

func (q *Queue) Start() *Structs.CodeProduct {
	return &q.first.product
}

func (q *Queue) Size() int {
	return q.contador
}
