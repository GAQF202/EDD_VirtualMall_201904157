package hashes

import (
	"fmt"
	"math"
)

type Comentario struct {
	Usuario    string
	Comentario string
}

type Tupla struct {
	clave int
	valor Comentario
}

func NewTupla(clave int, valor Comentario) *Tupla {
	return &Tupla{clave: clave, valor: valor}
}

type Nodo struct {
	indice int
	lista  []Tupla
}

func NewNodo(indice int) *Nodo {
	return &Nodo{indice: indice}
}

type HashTable struct {
	vector      []*Nodo
	elementos   int
	factorCarga float64
	size        int
}

func (ht *HashTable) GetAtributos() {
	fmt.Println("Tamaño:", ht.size, "Elementos:", ht.elementos, "Factor Carga:", ht.factorCarga)
}

func NewHashTable(size int) *HashTable {
	ht := &HashTable{elementos: 0, factorCarga: 0, size: size}

	for i := 0; i < size; i++ {
		ht.vector = append(ht.vector, nil)
	}
	return ht
}

func (ht *HashTable) funcionHash(id int) int {

	//METODO DE MULTIPLICACION
	mod := math.Mod(0.6556, 1)
	posicion := (ht.size * id * int(mod))
	if posicion > ht.size {
		return posicion - ht.size
	}
	return posicion
}

func (ht *HashTable) rehashing() {
	siguente := ht.size
	factor := 0.0

	for factor < 0.3 {
		factor = float64(ht.elementos) / float64(siguente)
		siguente++
	}

	ht_temporal := NewHashTable(siguente)

	for _, nodo := range ht.vector {
		for _, tupla := range nodo.lista {
			ht_temporal.Insertar(int(tupla.clave), tupla.clave, tupla.valor)
		}
	}

	ht.vector = ht_temporal.vector
	ht.elementos = ht_temporal.elementos
	ht.size = ht_temporal.size
	ht.factorCarga = ht_temporal.factorCarga
}

func (ht *HashTable) Insertar(id int, clave int, valor Comentario) {
	posicion := ht.funcionHash(id)

	if ht.vector[posicion] != nil {
		nuevo := NewTupla(clave, valor)
		ht.vector[posicion].lista = append(ht.vector[posicion].lista, *nuevo)
	} else {
		nuevo := NewNodo(posicion)
		nuevo.lista = append(nuevo.lista, *NewTupla(clave, valor))
		ht.vector[posicion] = nuevo
		ht.elementos++
		ht.factorCarga = float64(ht.elementos) / float64(ht.size)
	}
	if ht.factorCarga > 0.6 {
		//REHASHING
		ht.rehashing()
	}
}

func (nodo *Nodo) print() {
	for i := 0; i < len(nodo.lista); i++ {
		fmt.Println("indice:", i, "Valor:", nodo.lista[i].valor)
	}
}

//FUNCION PARA PRINTIAR LOS ELEMENTOS DE LA TABLA HASH
func (ht *HashTable) Print() {
	for i := 0; i < ht.size; i++ {
		if ht.vector[i] == nil {
			fmt.Println("Posicion:", i, "vacia")
		} else {
			fmt.Println("Posicion:", i)
			ht.vector[i].print()
		}
	}
}

//OBTENER LOSCOMENTARIOS DE LOS NODOS INTERNOS
func (nodo *Nodo) elInternos() []Comentario {
	var elementos []Comentario
	for i := 0; i < len(nodo.lista); i++ {
		elementos = append(elementos, nodo.lista[i].valor)
		fmt.Println(nodo.lista[i].valor)
	}
	return elementos
}

func (ht *HashTable) GetElements() []Comentario {
	var elementos []Comentario
	for i := 0; i < ht.size; i++ {
		if ht.vector[i] != nil {
			elementos = ht.vector[i].elInternos()
		}
	}
	return elementos
}

func stringtoascii(entrada string) int32 {
	cod := []rune(entrada)
	var temp int32
	temp = 0
	for i := 0; i < len(cod); i++ {
		temp = cod[i] + temp
	}
	return temp
}