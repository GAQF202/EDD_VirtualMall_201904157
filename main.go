package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/GAQF202/servidor-rest/dot"
	"github.com/GAQF202/servidor-rest/list"

	"github.com/gorilla/mux"
)

var vector []*list.Lista

func Linear(doc Mytype) {
	var temp_vector []*list.Lista

	first_dimention_size := len(doc.Datos)
	second_dimention_size := len(doc.Datos[0].Departamentos)
	var position int

	for i := 0; i <= first_dimention_size-1; i++ {
		for j := 0; j <= second_dimention_size-1; j++ {
			position++
			//CREACION DE LAS CALIFICACIONES
			for d := 1; d <= 5; d++ {
				var MyList *list.Lista = list.NewList()
				temp_vector = append(temp_vector, MyList)
			}
			//INSERSION AL VECTOR
			for k := 0; k <= len(doc.Datos[i].Departamentos[j].Tiendas)-1; k++ {
				if doc.Datos[i].Departamentos[j].Tiendas[k].Calificacion < 6 {
					MyStore := list.Store(doc.Datos[i].Departamentos[j].Tiendas[k])
					pos := ((position - 1) * 5) + doc.Datos[i].Departamentos[j].Tiendas[k].Calificacion
					list.Insertar(&MyStore, temp_vector[pos-1])
				}
			}
		}
	}
	vector = temp_vector
}
func Grafi(w http.ResponseWriter, r *http.Request) {

	dot_inst := "digraph G{ \n node[style=\"filled\",fillcolor=\"#8df7ef\",shape=\"record\"]  VectorNode[label=\""
	var lists string
	count := -1

	for g := 0; g < len(vector); g++ {
		count++
		if g == (len(vector) - 1) {
			dot_inst += "<" + strconv.Itoa(g) + ">" + strconv.Itoa(count) + "\"]"
		} else {
			dot_inst += "<" + strconv.Itoa(g) + ">" + strconv.Itoa(count) + "|"
		}

		lists += list.GetDotList(vector[g], g)
	}

	dot_inst += lists + "\n }"
	dot.CrearArchivo(dot_inst)
	dot.Graph()
}

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
			}
		}
	}
}

type ReturnedType struct {
	Datos []struct {
		Indice        string
		Departamentos []struct {
			Nombre  string
			Tiendas []list.Store
		}
	}
}

type E_pos struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
}

type D_pos struct {
	Categoria    string `json:"Categoria"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
}

func Get_position(Dep string, Name string, Cal int) int {
	first_dimention_size := len(dato.Datos)
	second_dimention_size := len(dato.Datos[0].Departamentos)
	Index := Name[:1]
	var position int
	var pos int

	for i := 0; i <= first_dimention_size-1; i++ {
		for j := 0; j <= second_dimention_size-1; j++ {
			position++
			if dato.Datos[i].Indice == Index && dato.Datos[i].Departamentos[j].Nombre == Dep && Cal <= 5 {
				pos = position
				pos = (((pos - 1) * 5) + Cal) - 1
			}
		}
	}
	return pos
}

func Delete_Store(w http.ResponseWriter, r *http.Request) {
	var info D_pos

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}
	json.Unmarshal([]byte(reqBody), &info)
	pos := Get_position(info.Categoria, info.Nombre, info.Calificacion)
	list.Delete_Node(vector[pos], info.Nombre, info.Calificacion)

}

func Browser(w http.ResponseWriter, r *http.Request) {
	var info E_pos
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}
	json.Unmarshal([]byte(reqBody), &info)

	first_dimention_size := len(dato.Datos)
	second_dimention_size := len(dato.Datos[0].Departamentos)
	Dep := info.Departamento
	Index := info.Nombre[:1]
	Cal := info.Calificacion
	var position int
	var pos int

	fmt.Println(Dep, Index)
	for i := 0; i <= first_dimention_size-1; i++ {
		for j := 0; j <= second_dimention_size-1; j++ {
			position++
			if dato.Datos[i].Indice == Index && dato.Datos[i].Departamentos[j].Nombre == Dep && Cal <= 5 {
				pos = position
				pos = (((pos - 1) * 5) + Cal) - 1
			}
		}
	}

	fmt.Println(list.Store_Browser(info.Nombre, info.Calificacion, vector[pos]))
	res := list.Store_Browser(info.Nombre, info.Calificacion, vector[pos])
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if res.Calificacion != 0 {
		json.NewEncoder(w).Encode(res)
	} else {
		json.NewEncoder(w).Encode("No existe dicha tienda")
	}

}

var dato Mytype

func CreateData(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}

	json.Unmarshal([]byte(reqBody), &dato)
	Linear(dato)
}

func Linear_Browser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id := parametros["numero"]
	number, err := strconv.Atoi(id)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err == nil && list.Get_Group(vector[number]) != nil {
		json.NewEncoder(w).Encode(list.Get_Group(vector[number]))
	} else {
		json.NewEncoder(w).Encode("No hay registro de tiendas en el indice: " + strconv.Itoa(number))
	}

}

func GetData(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("stores")
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/getArreglo", Grafi).Methods(("GET"))
	router.HandleFunc("/", CreateData).Methods("POST")
	router.HandleFunc("/id/{numero}", Linear_Browser).Methods("GET")
	router.HandleFunc("/Eliminar", Delete_Store).Methods(("POST"))
	router.HandleFunc("/TiendaEspecifica", Browser).Methods(("POST"))
	log.Fatal(http.ListenAndServe(":3000", router))

}
