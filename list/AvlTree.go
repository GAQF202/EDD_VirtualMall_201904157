package list

import (
	"fmt"

	"github.com/GAQF202/servidor-rest/Structs"
)

//STRUCT PARA GUARDAR PRODUCTOS
type Product struct {
	Nombre      string
	Codigo      int
	Descripcion string
	Precio      float64
	cantidad    int
	Imagen      string
}

func hola() {
	//c := Products.InventoryType{}
}

type NodoArbol struct {
	dato         Structs.Product
	equilibrador int
	hizq, hder   *NodoArbol
}

// BST is a set of sorted Nodes
type AVL struct {
	Raiz *NodoArbol
}

func (avl AVL) buscar(value int, r *NodoArbol) *NodoArbol {

	if avl.Raiz == nil {
		return nil
	} else if r.dato.Codigo == value {
		return r
	} else if r.dato.Codigo < value {
		return avl.buscar(value, r.hder)
	}
	return avl.buscar(value, r.hizq)
}

func (avl AVL) obtenerEquilibrio(x *NodoArbol) int {
	if x == nil {
		return -1
	} else {
		return x.equilibrador
	}
}

//ROTACION SIMPLE IZQUIERDA
func (avl AVL) rotacionIzq(x *NodoArbol) *NodoArbol {
	aux := x.hizq
	x.hizq = aux.hder
	aux.hder = x

	x.equilibrador = max(avl.obtenerEquilibrio(x.hizq), avl.obtenerEquilibrio(x.hder)) + 1
	aux.equilibrador = max(avl.obtenerEquilibrio(aux.hizq), avl.obtenerEquilibrio(aux.hder)) + 1
	return aux
}

//ROTACION SIMPLE DERECHA
func (avl AVL) rotacionDer(x *NodoArbol) *NodoArbol {
	aux := x.hder
	x.hder = aux.hizq
	aux.hizq = x
	x.equilibrador = max(avl.obtenerEquilibrio(x.hizq), avl.obtenerEquilibrio(x.hder)) + 1
	aux.equilibrador = max(avl.obtenerEquilibrio(aux.hizq), avl.obtenerEquilibrio(aux.hder)) + 1
	return aux
}

//ROTACION DOBLE IZQUIERDA
func (avl AVL) rotacionDobleIzq(x *NodoArbol) *NodoArbol {
	var aux *NodoArbol

	x.hizq = avl.rotacionDer(x.hizq)
	aux = avl.rotacionIzq(x)
	return aux
}

//ROTACION DOBLE DERECHA
func (avl AVL) rotacionDobleDer(x *NodoArbol) *NodoArbol {
	var aux *NodoArbol

	x.hder = avl.rotacionIzq(x.hder)
	aux = avl.rotacionDer(x)
	return aux
}

//METODO PARA OBTENER Y ACTUALIZAR ALTURA

func (avl AVL) _add(nuevo *NodoArbol, subAr *NodoArbol) *NodoArbol {
	padre := subAr

	if nuevo.dato.Codigo < subAr.dato.Codigo {
		if subAr.hizq == nil {
			subAr.hizq = nuevo
		} else {
			subAr.hizq = avl._add(nuevo, subAr.hizq)
			if ((avl.obtenerEquilibrio(subAr.hizq)) - (avl.obtenerEquilibrio(subAr.hder))) == 2 {
				if nuevo.dato.Codigo < subAr.hizq.dato.Codigo {
					padre = avl.rotacionIzq(subAr)
				} else {
					padre = avl.rotacionDobleIzq(subAr)
				}
			}
		}
	} else if nuevo.dato.Codigo > subAr.dato.Codigo {
		if subAr.hder == nil {
			subAr.hder = nuevo
		} else {
			subAr.hder = avl._add(nuevo, subAr.hder)
			if ((avl.obtenerEquilibrio(subAr.hder)) - (avl.obtenerEquilibrio(subAr.hizq))) == 2 {
				if nuevo.dato.Codigo > subAr.hder.dato.Codigo {
					padre = avl.rotacionDer(subAr)
				} else {
					padre = avl.rotacionDobleDer(subAr)
				}
			}
		}
	} else {
		fmt.Println("Nodo duplicado")
	}
	if subAr.hizq == nil && subAr.hder != nil {
		subAr.equilibrador = subAr.hder.equilibrador + 1
	} else if subAr.hder == nil && subAr.hizq != nil {
		subAr.equilibrador = subAr.hizq.equilibrador + 1
	} else {
		subAr.equilibrador = max(avl.obtenerEquilibrio(subAr.hizq), avl.obtenerEquilibrio(subAr.hder)) + 1
	}
	return padre
}

//METODO PARA INSERTAR
func (avl *AVL) Add(value Structs.Product) {
	nuevo := &NodoArbol{dato: value}
	if avl.Raiz == nil {
		avl.Raiz = nuevo
	} else {
		avl.Raiz = avl._add(nuevo, avl.Raiz)
	}
}

func max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func (bst AVL) inorder(tmp *NodoArbol) {
	if tmp != nil {
		bst.inorder(tmp.hizq)
		fmt.Print(tmp.dato.Nombre, " ")
		bst.inorder(tmp.hder)
	}
}

func (bst AVL) Preorder(tmp *NodoArbol) {

	if tmp != nil {
		fmt.Print(tmp.dato.Nombre, " ")
		bst.Preorder(tmp.hizq)
		bst.Preorder(tmp.hder)
	}
}

func main() {
	t := AVL{}
	t.Preorder(t.Raiz)
}
