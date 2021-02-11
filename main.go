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

	first_dimention_size := len(doc.Datos)
	second_dimention_size := len(doc.Datos[0].Departamentos)
	var position int

	for i := 0; i <= first_dimention_size-1; i++ {
		for j := 0; j <= second_dimention_size-1; j++ {
			position++
			for d := 1; d <= 5; d++ {
				var MyList *list.Lista = list.NewList()
				vector = append(vector, MyList)
			}
			for k := 0; k <= len(doc.Datos[i].Departamentos[j].Tiendas)-1; k++ {
				MyStore := list.Store(doc.Datos[i].Departamentos[j].Tiendas[k])
				pos := ((position - 1) * 5) + doc.Datos[i].Departamentos[j].Tiendas[k].Calificacion
				list.Insertar(&MyStore, vector[pos-1])
			}
		}
	}

	/*dot_inst := "digraph G{ node[style=\"filled\",fillcolor=\"#8df7ef\",shape=\"record\"]  VectorNode[label=\""
	var lists string
	count := 0

	for g := 0; g < len(vector); g++ {
		count++
		if g == (len(vector) - 1) {
			dot_inst += "<" + strconv.Itoa(g) + ">" + strconv.Itoa(count) + "\"]"
		} else {
			dot_inst += "<" + strconv.Itoa(g) + ">" + strconv.Itoa(count) + "|"
		}
		if count == 5 {
			count = 0
		}
		lists += list.GetDotList(vector[g], g)
	}

	dot_inst += lists + " }"
	fmt.Println(dot_inst)*/
	Grafi()

}
func Grafi() {

	dot_inst := "digraph G{ \n node[style=\"filled\",fillcolor=\"#8df7ef\",shape=\"record\"]  VectorNode[label=\""
	var lists string
	count := 0

	for g := 0; g < len(vector); g++ {
		count++
		if g == (len(vector) - 1) {
			dot_inst += "<" + strconv.Itoa(g) + ">" + strconv.Itoa(count) + "\"]"
		} else {
			dot_inst += "<" + strconv.Itoa(g) + ">" + strconv.Itoa(count) + "|"
		}
		if count == 5 {
			count = 0
		}
		lists += list.GetDotList(vector[g], g)
	}

	dot_inst += lists + "\n }"
	//fmt.Println(dot_inst)
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

func CreateData(w http.ResponseWriter, r *http.Request) {
	//var dato Dato
	var dato Mytype

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}

	json.Unmarshal([]byte(reqBody), &dato)
	Linear(dato)
	//fmt.Println(dato.Datos[0].Departamentos[0].Tiendas[0])
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dato)
}

func GetData(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("stores")
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute).Methods(("GET"))
	router.HandleFunc("/", CreateData).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))

}
