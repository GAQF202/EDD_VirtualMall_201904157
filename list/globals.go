package list

import (
	"fmt"

	"github.com/GAQF202/servidor-rest/Structs"
)

//STRUCT PARA RETORNAR LOS INVENTARIOS
type InventoryType struct {
	Tienda       string
	Departamento string
	Calificacion int
	Products     []Structs.Product
}

// STRUCT QUE QUE RECIBE LAS TIENDAS
type Mytype struct {
	Datos []struct {
		Indice        string `json:"Indice"`
		Departamentos []struct {
			Nombre  string `json:"Nombre"`
			Tiendas []struct {
				Nombre       string `json:"Nombre"`
				Descripcion  string `json:"Descripcion"`
				Contacto     string `json:"Contacto"`
				Calificacion int    `json:"Calificacion"`
				Logo         string `json:"Logo"`
			}
		}
	}
}

var Dato Mytype
var GlobalVector []*Lista

//HALLAR POSICION ESPECIFICA DE UN NODO EN LA LISTA DE TIENDAS
func Get_position(Dep string, Name string, Cal int) int {
	first_dimention_size := len(Dato.Datos)
	second_dimention_size := len(Dato.Datos[0].Departamentos)
	Index := Name[:1]
	var position int
	var pos int

	for i := 0; i <= first_dimention_size-1; i++ {
		for j := 0; j <= second_dimention_size-1; j++ {
			position++
			if Dato.Datos[i].Indice == Index && Dato.Datos[i].Departamentos[j].Nombre == Dep && Cal <= 5 {
				pos = position
				pos = (((pos - 1) * 5) + Cal) - 1
			}
		}
	}
	return pos
}

func Imp() {
	fmt.Println(len(GlobalVector))
}
