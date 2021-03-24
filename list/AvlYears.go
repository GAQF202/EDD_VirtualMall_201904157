package list

import (
	"fmt"
)

type Year struct {
	Year  int
	Meses *ListaMes
}

type nodoAnio struct {
	//CADA DATO ES UNA MATRIZ
	Anio         *Year
	equilibrador int
	hizq, hder   *nodoAnio
}

// BST is a set of sorted Nodes
type AVLYear struct {
	Raiz *nodoAnio
}

func (avl AVLYear) buscar(value int, r *nodoAnio) *nodoAnio {

	if avl.Raiz == nil {
		return nil
	} else if r.Anio.Year == value {
		return r
	} else if r.Anio.Year < value {
		return avl.buscar(value, r.hder)
	}
	return avl.buscar(value, r.hizq)
}

func (avl AVLYear) obtenerEquilibrio(x *nodoAnio) int {
	if x == nil {
		return -1
	} else {
		return x.equilibrador
	}
}

//ROTACION SIMPLE IZQUIERDA
func (avl AVLYear) rotacionIzq(x *nodoAnio) *nodoAnio {
	aux := x.hizq
	x.hizq = aux.hder
	aux.hder = x

	x.equilibrador = max(avl.obtenerEquilibrio(x.hizq), avl.obtenerEquilibrio(x.hder)) + 1
	aux.equilibrador = max(avl.obtenerEquilibrio(aux.hizq), avl.obtenerEquilibrio(aux.hder)) + 1
	return aux
}

//ROTACION SIMPLE DERECHA
func (avl AVLYear) rotacionDer(x *nodoAnio) *nodoAnio {
	aux := x.hder
	x.hder = aux.hizq
	aux.hizq = x
	x.equilibrador = max(avl.obtenerEquilibrio(x.hizq), avl.obtenerEquilibrio(x.hder)) + 1
	aux.equilibrador = max(avl.obtenerEquilibrio(aux.hizq), avl.obtenerEquilibrio(aux.hder)) + 1
	return aux
}

//ROTACION DOBLE IZQUIERDA
func (avl AVLYear) rotacionDobleIzq(x *nodoAnio) *nodoAnio {
	var aux *nodoAnio

	x.hizq = avl.rotacionDer(x.hizq)
	aux = avl.rotacionIzq(x)
	return aux
}

//ROTACION DOBLE DERECHA
func (avl AVLYear) rotacionDobleDer(x *nodoAnio) *nodoAnio {
	var aux *nodoAnio

	x.hder = avl.rotacionIzq(x.hder)
	aux = avl.rotacionDer(x)
	return aux
}

//METODO PARA OBTENER Y ACTUALIZAR ALTURA

func (avl AVLYear) _add(nuevo *nodoAnio, subAr *nodoAnio) *nodoAnio {
	padre := subAr

	if nuevo.Anio.Year < subAr.Anio.Year {
		if subAr.hizq == nil {
			subAr.hizq = nuevo
		} else {
			subAr.hizq = avl._add(nuevo, subAr.hizq)
			if ((avl.obtenerEquilibrio(subAr.hizq)) - (avl.obtenerEquilibrio(subAr.hder))) == 2 {
				if nuevo.Anio.Year < subAr.hizq.Anio.Year {
					padre = avl.rotacionIzq(subAr)
				} else {
					padre = avl.rotacionDobleIzq(subAr)
				}
			}
		}
	} else if nuevo.Anio.Year > subAr.Anio.Year {
		if subAr.hder == nil {
			subAr.hder = nuevo
		} else {
			subAr.hder = avl._add(nuevo, subAr.hder)
			if ((avl.obtenerEquilibrio(subAr.hder)) - (avl.obtenerEquilibrio(subAr.hizq))) == 2 {
				if nuevo.Anio.Year > subAr.hder.Anio.Year {
					padre = avl.rotacionDer(subAr)
				} else {
					padre = avl.rotacionDobleDer(subAr)
				}
			}
		}
	} else {
		fmt.Println("Nodo duplicado")
		//SE INSERTA EN EL NODO EXISTENTE
		avl.buscar(subAr.Anio.Year, subAr).Anio.Meses.Insertar(&subAr.Anio.Meses.primero.Mes)
		//mes := ListaMes{subAr.Anio.Meses.primero, subAr.Anio.Meses.ultimo, 5}
		//avl.buscar(subAr.Anio.Year, subAr).Anio.Meses.Insertar(&mes.primero.Mes)
		//fmt.Println(subAr.Anio.Meses.primero.Mes)
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
func (avl *AVLYear) Add(year *Year) {
	nuevo := &nodoAnio{Anio: year}
	if avl.Raiz == nil {
		avl.Raiz = nuevo
	} else {
		avl.Raiz = avl._add(nuevo, avl.Raiz)
	}
}

func (avl AVLYear) Inorder(tmp *nodoAnio) {
	if tmp != nil {
		avl.Inorder(tmp.hizq)
		fmt.Print(tmp.Anio.Year, " ")
		avl.Inorder(tmp.hder)
	}
}

func (avl AVLYear) Preorder(tmp *nodoAnio) {

	if tmp != nil {
		fmt.Print(tmp.Anio.Year, " ")
		avl.Preorder(tmp.hizq)
		avl.Preorder(tmp.hder)
	}
}
