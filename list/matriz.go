package list

import "fmt"

type nodo struct {
	//Estos atributos son especificos para la matriz
	x/*, y */ int //Saber en que cabecera estoy
	y             string
	//producto                          *Product //tipo de objeto
	product                           *Queue
	izquierda, derecha, arriba, abajo *nodo //nodos con los que nos desplazamos dentro de la matriz
	//Estos atributos son especificos para la lista
	header              int //tipo interno de la cabecera
	headerVertical      string
	siguiente, anterior *nodo // nodos con los que nos vamos a desplazar dentro de las listas
}

type lista struct {
	first, last *nodo
}

type matriz struct {
	lst_h, lst_v *lista
}

func nodoMatriz(x int, y string, producto *Queue) *nodo {
	return &nodo{x, y, producto, nil, nil, nil, nil, 0, "", nil, nil}
}

//INSERTA LOS NODOS DE DIAS
func nodoLista(header int) *nodo {
	return &nodo{0, "", nil, nil, nil, nil, nil, header, "", nil, nil}
}

//INSERTA LOS NODOS DE CATEGORIA
func nodoListaVertical(header string) *nodo {
	return &nodo{0, "", nil, nil, nil, nil, nil, 0, header, nil, nil}
}

func newLista() *lista {
	return &lista{nil, nil}
}

//Se cambio a primer letra mayuscula para poder acceder
func NewMatriz() *matriz {
	return &matriz{newLista(), newLista()}
}

func (n *nodo) headerX() int    { return n.x }
func (n *nodo) headerY() string { return n.y }

/*func (n *nodo) toString() string {
	return "Nombre: " + n.producto.Nombre + "\nDescripcion: " + n.producto.Descripcion
}*/

//METODO PARA ORDENAR DIAS
func (l *lista) ordenarHorizontal(nuevo *nodo) {
	aux := l.first
	for aux != nil {
		if nuevo.header > aux.header {
			aux = aux.siguiente
		} else {
			if aux == l.first {
				nuevo.siguiente = aux
				aux.anterior = nuevo
				l.first = nuevo
			} else {
				nuevo.anterior = aux.anterior
				aux.anterior.siguiente = nuevo
				nuevo.siguiente = aux
				aux.anterior = nuevo
			}
			return
		}
	}
	l.last.siguiente = nuevo
	nuevo.anterior = l.last
	l.last = nuevo
}

//METODO PARA ORDENAR LISTA DE CATEGORIAS
func (l *lista) ordenarVertical(nuevo *nodo) {
	aux := l.first
	for aux != nil {
		if nuevo.header != aux.header {
			aux = aux.siguiente
		} else {
			if aux == l.first {
				nuevo.siguiente = aux
				aux.anterior = nuevo
				l.first = nuevo
			} else {
				nuevo.anterior = aux.anterior
				aux.anterior.siguiente = nuevo
				nuevo.siguiente = aux
				aux.anterior = nuevo
			}
			return
		}
	}
	l.last.siguiente = nuevo
	nuevo.anterior = l.last
	l.last = nuevo
}

//METODO PARA INSERTAR CABECERA DE HORIZONTAL DE DIAS
func (l *lista) insert(header int) {
	nuevo := nodoLista(header)
	if l.first == nil {
		l.first = nuevo
		l.last = nuevo
	} else {
		l.ordenarHorizontal(nuevo)
	}
}

//METODO PARA INSERTAR CABECERA DE VERTICAL DE CATEGORIAS
func (l *lista) insertY(header string) {
	nuevo := nodoListaVertical(header)
	if l.first == nil {
		l.first = nuevo
		l.last = nuevo
	} else {
		l.ordenarVertical(nuevo)
	}
}

//BUSQUEDA DE DIAS
func (l *lista) search(header int) *nodo {
	temp := l.first
	for temp != nil {
		if temp.header == header {
			return temp
		}
		temp = temp.siguiente
	}
	return nil
}

//BUSQUEDA DE CATEGORIAS
func (l *lista) searchVertical(header string) *nodo {
	temp := l.first
	for temp != nil {
		if temp.headerVertical == header {
			return temp
		}
		temp = temp.siguiente
	}
	return nil
}

func (l *lista) print() {
	temp := l.first
	for temp != nil {
		fmt.Println("Cabecera:", temp.header)
		temp = temp.siguiente
	}
}

func (m *matriz) Insert(producto *Queue, x int, y string) {
	h := m.lst_h.search(x)
	v := m.lst_v.searchVertical(y)

	if h == nil && v == nil {
		m.noExisten(producto, x, y)
	} else if h == nil && v != nil {
		m.existeVertical(producto, x, y)
	} else if h != nil && v == nil {
		m.existeHorizontal(producto, x, y)
	} else {
		m.existen(producto, x, y)
	}
}

func (n *nodo) print() {
	fmt.Println("x:", n.x, "y:", n.y)
}

//INSERTA CADA ELEMENTO DE LA MATRIZ DE MES REPETIDO EN LA MATRIZ DE MES YA EXISTENTE
func (m *matriz) ColumnMajor(Matriz *matriz) {
	cabecera := m.lst_v.first

	for cabecera != nil {
		aux := cabecera.derecha
		for aux != nil {
			//aux.print(aux.abajo)
			/*fmt.Println(aux.product.first.product.Codigo_producto)*/
			Matriz.Insert(aux.product, aux.x, aux.y)
			aux = aux.derecha
		}
		cabecera = cabecera.siguiente
	}
}

func (m *matriz) noExisten(producto *Queue, x int, y string) {
	m.lst_h.insert(x)  //insertamos en la lista que emula la cabecera horizontal
	m.lst_v.insertY(y) //insertamos en la lista que emula la cabecera vertical

	h := m.lst_h.search(x)         //vamos a buscar el nodo que acabamos de insertar para poder enlazarlo
	v := m.lst_v.searchVertical(y) //vamos a buscar el nodo que acabamos de insertar para poder enlazarlo

	nuevo := nodoMatriz(x, y, producto) //creamos nuevo nodo tipo matriz

	h.abajo = nuevo  //enlazamos el nodo horizontal hacia abajo
	nuevo.arriba = h //enlazmos el nuevo nodo hacia arriba

	v.derecha = nuevo   //enlazamos el nodo vertical hacia la derecha
	nuevo.izquierda = v //enlazamos el nuevo nodo hacia la izquierda
	nuevo.derecha = nil
	nuevo.abajo = nil
}

//METODO PARA CUANDO EXISTE LA CABECERA VERTICAL
func (m *matriz) existeVertical(producto *Queue, x int, y string) {
	m.lst_h.insert(x)                   //insertamos en la lista que emula la cabecera horizontal
	h := m.lst_h.search(x)              //BUSCO EL NODO EN LA LISTA HORIZONTAL
	v := m.lst_v.searchVertical(y)      //BUSCO EL NODO EN LA LISTA VERTICAL
	nuevo := nodoMatriz(x, y, producto) //creamos nuevo nodo tipo matriz
	agregado := false
	aux := v.derecha

	var cabecera int

	for aux != nil {
		cabecera = aux.headerX()
		if cabecera < x {
			aux = aux.derecha
		} else {
			nuevo.derecha = aux
			nuevo.izquierda = aux.izquierda
			aux.izquierda.derecha = nuevo
			aux.izquierda = nuevo
			agregado = true
			break
		}
	}

	if agregado == false {
		aux = v.derecha
		for aux.derecha != nil {
			aux = aux.derecha
		}
		nuevo.izquierda = aux
		aux.derecha = nuevo
	}
	nuevo.arriba = h
	h.abajo = nuevo
}

//METODO PARA CUANDO EXISTE LA CABECERA HORIZONTAL
func (m *matriz) existeHorizontal(producto *Queue, x int, y string) {
	m.lst_v.insertY(y)                  //insertamos en la lista que emula la cabecera horizontal
	h := m.lst_h.search(x)              //BUSCO EL NODO EN LA LISTA HORIZONTAL
	v := m.lst_v.searchVertical(y)      //BUSCO EL NODO EN LA LISTA VERTICAL
	nuevo := nodoMatriz(x, y, producto) //creamos nuevo nodo tipo matriz
	agregado := false
	aux := h.abajo

	var cabeceraVertical string

	for aux != nil {
		cabeceraVertical = aux.headerY()
		if cabeceraVertical != y {
			aux = aux.abajo
		} else {
			nuevo.abajo = aux
			nuevo.arriba = aux.arriba
			aux.arriba.abajo = nuevo
			aux.arriba = nuevo
			agregado = true
			break
		}
	}

	if agregado == false {
		aux = h.abajo
		for aux.abajo != nil {
			aux = aux.abajo
		}
		nuevo.arriba = aux
		aux.abajo = nuevo
	}
	nuevo.izquierda = v
	v.derecha = nuevo
}

func (m *matriz) existen(producto *Queue, x int, y string) {

	h := m.lst_h.search(x)              //BUSCO EL NODO EN LA LISTA HORIZONTAL
	v := m.lst_v.searchVertical(y)      //BUSCO EL NODO EN LA LISTA VERTICAL
	nuevo := nodoMatriz(x, y, producto) //creamos nuevo nodo tipo matriz
	agregado := false

	aux := v.derecha
	var cabecera int
	var cabeceraVertical string

	for aux != nil {
		cabecera = aux.headerX()
		if cabecera < x {
			aux = aux.derecha
		} else {
			nuevo.derecha = aux
			nuevo.izquierda = aux.izquierda
			aux.izquierda.derecha = nuevo
			aux.izquierda = nuevo
			agregado = true
			break
		}
	}

	//agregado = false
	//aux = v.derecha

	if agregado == false {
		aux = v.derecha
		for aux.derecha != nil {
			aux = aux.derecha
		}
		nuevo.izquierda = aux
		aux.derecha = nuevo
	}

	agregado = false
	aux = h.abajo

	for aux != nil {
		cabeceraVertical = aux.headerY()
		if cabeceraVertical != y {
			aux = aux.abajo
		} else {
			nuevo.abajo = aux
			nuevo.arriba = aux.arriba
			aux.arriba.abajo = nuevo
			aux.arriba = nuevo
			agregado = true
			break
		}
	}

	if agregado == false {
		aux = h.abajo
		for aux.abajo != nil {
			aux = aux.abajo
		}
		nuevo.arriba = aux
		aux.abajo = nuevo
	}

}
