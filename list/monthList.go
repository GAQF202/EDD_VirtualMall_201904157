package list

import (
	"github.com/GAQF202/servidor-rest/Structs"
)

type Month struct {
	Month  string
	Matriz *matriz
}

type NodoMes struct {
	siguiente, anterior *NodoMes
	Mes                 Month
}

type ListaMes struct {
	primero  *NodoMes
	ultimo   *NodoMes
	contador int
}

func NuevoNodoMes(mes Month) *NodoMes {
	//SIGUIENTE, ANTERIOR, TIENDA, INVENTARIO
	return &NodoMes{siguiente: nil, anterior: nil, Mes: mes}
}

func NewListMes() *ListaMes {
	return &ListaMes{nil, nil, 0}
}

func (lista *ListaMes) Insertar(mes *Month) {
	var nuevo *NodoMes = NuevoNodoMes(*mes)

	if !lista.existe(mes) {
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
	} else {
		lista.insertarExistente(mes)
	}

}

//SI EL MES EXISTEMANDA A INSERTAR LOS ELEMENTOS DE LA MATRIZ DENTRO DE LA EXISTENTE
func (lista *ListaMes) insertarExistente(mes *Month) {
	//fmt.Println(lista.primero.Mes.Month)

	var actual *NodoMes
	actual = lista.primero

	for actual != nil {
		if actual.Mes.Month == mes.Month {

			//actual.Mes = *mes
			//mes.Matriz.ColumnMajor(actual.Mes.Matriz)
			actual.Mes.Matriz.ColumnMajor(mes.Matriz)
			//mes.Matriz.ColumnMajor(mes.Matriz)
			//actual.Mes.Matriz.Insert()
		}
		actual = actual.siguiente
	}

}

//INSERTA LOS DATOS DE LA MATRIZ DEL MES REPETIDO EN EL EXISTENTE
func (lista *ListaMes) existe(mes *Month) bool {
	actual := lista.primero
	for actual != nil {
		if actual.Mes.Month == mes.Month {
			return true
		}
		actual = actual.siguiente
	}
	return false
}

func (list *ListaMes) IsVoid() bool {
	return list.primero == nil
}

//VARIABLE PARA GUARDAR LOS MESES
var Res []Structs.Mes

func (list *ListaMes) Recorrer_insertar() []Structs.Mes {
	aux := list.primero

	for aux != nil {
		//CREO EL MES Y LO INSERTO
		mes := Structs.Mes{aux.Mes.Month}
		Res = append(Res, mes)
		//SIGO RECORRIENDO LA LISTA DE MESES
		aux = aux.siguiente
	}
	return Res
}

var Meses Structs.MesSeleccionado

//FUNCION PARA BUSCAR EL MES SELECCIONADO
func (list *ListaMes) buscarMes(mes string) {
	aux := list.primero

	for aux != nil {
		if aux.Mes.Month == mes {
			aux.Mes.Matriz.ColMa()
			mes := Structs.MesSeleccionado{aux.Mes.Month, Cola}
			Meses = mes
		}
		aux = aux.siguiente
	}
}

func (list *ListaMes) GetCodigoInterno(anio string) string {
	aux := list.primero
	res := ""

	for aux != nil {

		if aux.siguiente != nil {
			res += aux.Mes.Month + anio + "->" + aux.siguiente.Mes.Month + anio + "\n"
			res += aux.Mes.Month + anio + "[label=\"" + aux.Mes.Month + "\" shape=box]\n"
		}
		if aux.anterior != nil {
			res += aux.Mes.Month + anio + "->" + aux.anterior.Mes.Month + anio + "\n"
			res += aux.Mes.Month + anio + "[label=\"" + aux.Mes.Month + "\" shape=box]\n"
		}
		if aux.siguiente == nil && aux.anterior == nil {
			res += aux.Mes.Month + anio + "\n"
			res += aux.Mes.Month + anio + "[label=\"" + aux.Mes.Month + "\" shape=box]\n"
		}

		aux = aux.siguiente
	}
	return res
}

/*func Delete_Node(lista *Lista, Name string, Cal int) {
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
}*/
