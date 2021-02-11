package list

import (
	"fmt"
	"strconv"
)

type Store struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
}

type Nodo struct {
	siguiente, anterior *Nodo
	tienda              Store
}

type Lista struct {
	primero  *Nodo
	ultimo   *Nodo
	contador int
}

func NuevoNodo(tienda Store) *Nodo {
	return &Nodo{nil, nil, tienda}
}

func NewList() *Lista {
	return &Lista{nil, nil, 0}
}

func Imprimir(lista *Lista) {
	aux := lista.primero

	for aux != nil {

		fmt.Println("---------Tienda----------")

		fmt.Println("nodo anterior:", aux.tienda)

		aux = aux.siguiente
	}
}

func Insertar(tienda *Store, lista *Lista) {
	var nuevo *Nodo = NuevoNodo(*tienda)

	if lista.primero == nil {
		lista.primero = nuevo
		lista.ultimo = nuevo
		lista.primero.siguiente, lista.primero.anterior = nil, nil
		lista.contador++
	} else {
		aux := lista.ultimo
		lista.ultimo.siguiente = nuevo
		lista.ultimo = lista.ultimo.siguiente
		lista.ultimo.anterior = aux

		lista.contador++
	}

}

func CreateNode(index int, label string, color string) string {

	return "node" + strconv.Itoa(index) + "[label=\"" + label + "\",fillcolor=\"white\"," + "shape=\"record\"" + ",color=\"" + color + "\"]"
}

func GetDotList(lista *Lista, corr int) string {
	count := corr
	aux := lista.primero
	var dot_inst string

	if aux != nil {
		for aux != nil {
			count++
			actual, anterior, siguiente_anterior := "node"+strconv.Itoa(count), "node"+strconv.Itoa(count+1), "node"+strconv.Itoa(count)
			dot_inst += " " + CreateNode(count, aux.tienda.Nombre, "pink")

			if aux.siguiente != nil {
				dot_inst += " " + actual + "->" + anterior + "->" + siguiente_anterior
			}

			if aux.anterior == nil {
				dot_inst += " VectorNode:\"" + strconv.Itoa(corr) + "\" -> " + actual
			}

			aux = aux.siguiente
		}
	}

	return dot_inst
}
