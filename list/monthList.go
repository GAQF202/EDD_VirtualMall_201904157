package list

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

func (lista ListaMes) insertarExistente(mes *Month) {
	//fmt.Println(lista.primero.Mes.Month)

	actual := lista.primero

	for actual != nil {
		if actual.Mes.Month == mes.Month {
			actual.Mes = *mes
			mes.Matriz.ColumnMajor(actual.Mes.Matriz)
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
