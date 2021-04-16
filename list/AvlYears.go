package list

import (
	"fmt"
	"strconv"

	"github.com/GAQF202/servidor-rest/Structs"
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
		avl.buscar(nuevo.Anio.Year, subAr).Anio.Meses.Insertar(&nuevo.Anio.Meses.primero.Mes)
		//fmt.Println(subAr.Anio.Meses.ultimo.Mes.Month)
		//fmt.Println("Es AQUI", avl.buscar(subAr.Anio.Year, subAr).Anio.Meses)

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

//FUNCION PARA GUARDAR LOS ANIOS
var Anios []Structs.Anio

//FUNCION PARA GUARDAR LOS ANIOS Y MESES EN EL STRUCT
func (avl AVLYear) Inorder(tmp *nodoAnio) []Structs.Anio {

	if tmp != nil {
		avl.Inorder(tmp.hizq)
		//SE CREA EL STRUCT DE MESES
		meses := tmp.Anio.Meses.Recorrer_insertar()
		//SE CREA EL UN STRUCT PARA CADA ANIO
		anioActual := Structs.Anio{tmp.Anio.Year, meses}
		//SE INSERTA CADA ANIO NUEVO EN EL STRUCT
		Anios = append(Anios, anioActual)
		//SE SIGUE RECORRIENDO EL ARBOL
		avl.Inorder(tmp.hder)
	}
	Res = []Structs.Mes{}
	return Anios
}

//FUNCION PARA RECORRER EL ARBOL INORDER
var Ver = ""

func (avl AVLYear) RecorrerInOrder(tmp *nodoAnio) string {

	if tmp != nil {
		avl.RecorrerInOrder(tmp.hizq)
		Ver += strconv.Itoa(tmp.Anio.Year) + "->" + tmp.Anio.Meses.GetCodigoInterno(strconv.Itoa(tmp.Anio.Year))
		tmp.Anio.Meses.GetCodigoInterno(strconv.Itoa(tmp.Anio.Year))
		avl.RecorrerInOrder(tmp.hder)
	}
	return Ver
}

//FUNCION PARA GUARDAR LOS ANIOS Y MESES EN EL STRUCT
func (avl AVLYear) BuscarAnio(tmp *nodoAnio, mes string, anio int) {

	if tmp != nil {
		avl.BuscarAnio(tmp.hizq, mes, anio)
		if tmp.Anio.Year == anio {
			tmp.Anio.Meses.buscarMes(mes)
			fmt.Println(tmp.Anio.Meses.primero.Mes)
		}
		avl.BuscarAnio(tmp.hder, mes, anio)
	}
}

//METODO PARA OBTENER EL CODIGO DOT DEL ARBOl
var Arbol string

func (avl *nodoAnio) GetCodigoInterno(r *nodoAnio) string {

	if r.hizq != nil {
		Arbol += strconv.Itoa(r.Anio.Year) + "->" + strconv.Itoa(r.hizq.Anio.Year) + "\n"
		r.hizq.GetCodigoInterno(r.hizq)
	}
	if r.hder != nil {
		Arbol += strconv.Itoa(r.Anio.Year) + "->" + strconv.Itoa(r.hder.Anio.Year) + "\n"
		r.hder.GetCodigoInterno(r.hder)
	}
	/*if r.hizq != nil {
		res += r.hizq.GetCodigoInterno(r.hizq) + "->" + strconv.Itoa(r.Anio.Year) + "\n"
	}
	if r.hder != nil {
		res += r.hder.GetCodigoInterno(r.hder) + "->" + strconv.Itoa(r.Anio.Year) + "\n"
	}*/

	return Arbol
}
