package list

import (
	"fmt"
	"strconv"

	"github.com/GAQF202/servidor-rest/Structs"
)

//STRUCT PARA TIENDA
type Store struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
	Logo         string
	//	Inventario   Estructures.AVL
}

type Nodo struct {
	siguiente, anterior *Nodo
	tienda              Store
	Inventario          AVL
}

type Lista struct {
	primero  *Nodo
	ultimo   *Nodo
	contador int
}

type GroupStores []Store

func NuevoNodo(tienda Store) *Nodo {
	//SIGUIENTE, ANTERIOR, TIENDA, INVENTARIO
	//return &Nodo{nil, nil, tienda}
	return &Nodo{siguiente: nil, anterior: nil, tienda: tienda}
}

func NewList() *Lista {
	return &Lista{nil, nil, 0}
}

func Imprimir(lista *Lista) {

	//fmt.Println(lista.primero.tienda)
	aux := lista.primero

	for aux != nil {

		fmt.Println("---------Tienda----------")

		fmt.Println(aux.tienda)

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

//ARREGLAR ESTA FUNCION
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

//BUSCA LA TIENDA Y GUARDA EN EL ARBOL LOS PRODUCTOS
func Get_store_node(name string, calification int, list *Lista, product Structs.Product) {
	aux := list.primero

	for aux != nil {
		if aux.tienda.Nombre == name && aux.tienda.Calificacion == calification {
			//fmt.Println(product.Nombre)
			aux.Inventario.Add(product)
		}
		aux = aux.siguiente
	}

}

func VerNodos(list *Lista) {
	aux := list.primero
	for aux != nil {
		if aux.Inventario.Raiz != nil {
			fmt.Println(aux.tienda.Nombre)
			aux.Inventario.Preorder(aux.Inventario.Raiz)
			fmt.Println("------------------")
		}
		aux = aux.siguiente
	}
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

func IsVoid(list *Lista) bool {
	return list.primero == nil
}

type Group []struct {
	Nombre       string `json:"Nombre"`
	Descripcion  string `json:"Descripcion"`
	Contacto     string `json:"Contacto"`
	Calificacion int    `json:"Calificacion"`
	Logo         string `json:"Logo"`
}

func Get_store(list *Lista) Group {
	var myGropu Group
	aux := list.primero

	for aux != nil {
		myGropu = append(myGropu, struct {
			Nombre       string "json:\"Nombre\""
			Descripcion  string "json:\"Descripcion\""
			Contacto     string "json:\"Contacto\""
			Calificacion int    "json:\"Calificacion\""
			Logo         string `json:"Logo"`
		}(aux.tienda))
		aux = aux.siguiente
	}
	return myGropu
}

func GetCalification(list *Lista) int {
	res := 0
	if list.primero != nil {
		res = list.primero.tienda.Calificacion
	}
	return res
}

//DEVUELVE EL JSON DEL INVENTARIO
func JsonInventory(name string, calification int, list *Lista, dep string) InventoryType {
	aux := list.primero
	var res InventoryType

	for aux != nil {
		if aux.tienda.Nombre == name && aux.tienda.Calificacion == calification {
			//inv := aux.Inventario.inorder(aux.Inventario.Raiz)
			res.Tienda = aux.tienda.Nombre
			res.Calificacion = aux.tienda.Calificacion
			res.Departamento = dep
			aux.Inventario.inorder(aux.Inventario.Raiz, &res)
		}
		aux = aux.siguiente
	}

	return res
}

//BUSCA LA TIENDA Y GUARDA EN EL ARBOL LOS PRODUCTOS
func Delete_product(name string, calification int, list *Lista, productos Structs.Product) {
	aux := list.primero

	for aux != nil {
		if aux.tienda.Nombre == name && aux.tienda.Calificacion == calification {
			aux.Inventario.searchAndDelete(productos, aux.Inventario.Raiz)
			//fmt.Println(product.Nombre)
			//aux.Inventario.buscar(aux.Inventario.)
		}
		aux = aux.siguiente
	}

}
