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

type GroupStores []Store

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

func Store_Browser(name string, calification int, list *Lista) Store {
	aux := list.primero
	var result Store

	for aux != nil {
		if aux.tienda.Nombre == name && aux.tienda.Calificacion == calification {
			result = aux.tienda
		}
		aux = aux.siguiente
	}
	return result
}

func CreateNode(index int, label string, color string) string {

	return "node" + strconv.Itoa(index) + "[label=\"" + label + "\",fillcolor=\"white\"," + "shape=\"record\"" + ",color=\"" + color + "\"]"
}

var count int

func GetDotList(lista *Lista, corr int) string {

	aux := lista.primero
	var dot_inst string

	if aux != nil {
		for aux != nil {
			count++
			actual, anterior, siguiente_anterior := "node"+strconv.Itoa(count), "node"+strconv.Itoa(count+1), "node"+strconv.Itoa(count)
			dot_inst += " " + CreateNode(count, aux.tienda.Nombre, "pink")
			if aux.siguiente != nil {
				dot_inst += "\n " + actual + "->" + anterior + "->" + siguiente_anterior + "\n"
			}

			if aux.anterior == nil {
				dot_inst += "\n VectorNode:\"" + strconv.Itoa(corr) + "\" -> " + actual + "\n"
			}

			aux = aux.siguiente
		}
	}

	return dot_inst
}

func Get_Group(lista *Lista) GroupStores {
	aux := lista.primero
	var result GroupStores

	for aux != nil {
		result = append(result, aux.tienda)
		aux = aux.siguiente
	}
	return result
}

func Delete_Node(lista *Lista, Name string, Cal int) {
	//SI LA LISTA ESTA VACIA
	if lista != nil {
		if lista.primero == lista.ultimo && lista.primero.tienda.Nombre == Name && lista.primero.tienda.Calificacion == Cal {
			lista.primero, lista.ultimo = nil, nil
		} else if lista.primero.tienda.Nombre == Name && lista.primero.tienda.Calificacion == Cal {
			lista.primero = lista.primero.siguiente
			lista.primero.anterior = nil
		} else {
			anterior := lista.primero
			temp := lista.primero.siguiente

			for temp != nil && temp.tienda.Nombre != Name && temp.tienda.Calificacion != Cal {
				anterior = anterior.siguiente
				temp = temp.siguiente
			}
			if temp != nil {
				anterior.siguiente = temp.siguiente
				if temp == lista.ultimo {
					lista.ultimo = anterior
				}
			}
		}
	}
}
