package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/GAQF202/servidor-rest/Products"
	"github.com/GAQF202/servidor-rest/Structs"
	"github.com/GAQF202/servidor-rest/dijkstra"
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
	var hashTiendas []dijkstra.Hashable

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
					hashTiendas = append(hashTiendas, dijkstra.Block(MyStore.Nombre+"\\n"+MyStore.Descripcion+"\\n"+MyStore.Contacto))
				}
			}
		}
	}
	dijkstra.PrintTree(dijkstra.BuildTree(hashTiendas)[0].(dijkstra.Node))
	//fmt.Println(dijkstra.DotMerkleTree)
	dot.CrearArchivoEvery(dijkstra.DotMerkleTree+"}", "txt", "DotAnios")
	dot.GraphEvery("MerkleTiendas", "jpg", "DotAnios")
	dijkstra.DotMerkleTree = "digraph { node [shape=box, style=\"filled\", fillcolor=\"#61e665\"];"
	vector = temp_vector
	list.GlobalVector = vector
}
func Grafi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
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
	dot.CrearArchivo(dot_inst, "txt")
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
				Logo         string `json:"Logo"`
			}
		}
	}
}

//STRUCT PARA BUSQUEDA
type E_pos struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
}

//STRUCT PARA ELIMINACION
type D_pos struct {
	Categoria    string `json:"Categoria"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
}

//HALLAR POSICION ESPECIFICA DE UN NODO EN LA LISTA DE TIENDAS
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

//FUNCION PARA ELIMINAR TIENDA
func Delete_Store(w http.ResponseWriter, r *http.Request) {
	var info D_pos

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte una tienda existente")
	}
	json.Unmarshal([]byte(reqBody), &info)
	pos := Get_position(info.Categoria, info.Nombre, info.Calificacion)
	list.Delete_Node(vector[pos], info.Nombre, info.Calificacion)
}

//BUSQUEDA ESPECIFICA DE TIENDAS
func Browser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var info E_pos
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte una tienda existente")
	}
	json.Unmarshal([]byte(reqBody), &info)

	first_dimention_size := len(dato.Datos)
	second_dimention_size := len(dato.Datos[0].Departamentos)
	Dep := info.Departamento
	Index := info.Nombre[:1]
	Cal := info.Calificacion
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.Unmarshal([]byte(reqBody), &dato)
	Linear(dato)

	//JSON DE RESPUESTA
	json.NewEncoder(w).Encode(dato)
	//ENVIA LOS DATOS PARA MANEJAR LA VARIABLE EN TODO EL PROGRAMA
	list.Dato = list.Mytype(dato)
}

func Linear_Browser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id := parametros["numero"]
	number, err := strconv.Atoi(id)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err == nil && list.Get_Group(vector[number]) != nil {
		json.NewEncoder(w).Encode(list.Get_Group(vector[number]))
	} else {
		json.NewEncoder(w).Encode("No hay registro de tiendas en el indice: " + strconv.Itoa(number))
	}

}

func Delete_all() {
	first_dimention_size := len(dato.Datos)
	second_dimention_size := len(dato.Datos[0].Departamentos)
	for i := 0; i <= first_dimention_size-1; i++ {
		for j := 0; j <= second_dimention_size-1; j++ {
			dato.Datos[i].Departamentos[j].Tiendas = nil
		}
	}
}

//FUNCION PARA GUARDAR ARCHIVO JSON DE SALIDA
func Json_Returned(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := 1
	Delete_all()
	for k := 0; k < len(vector); k++ {
		count++
		if !list.IsVoid(vector[k]) {
			VecPos := ((count - list.GetCalification(vector[k])) / 5)
			Insert_in_myType(vector[k], k, VecPos)
		}
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dato)
	file, _ := json.MarshalIndent(dato, "", " ")
	_ = ioutil.WriteFile("salida.json", file, 0644)

}

func Insert_in_myType(lista *list.Lista, position int, VecPos int) {
	first_dimention_size := len(dato.Datos)
	second_dimention_size := len(dato.Datos[0].Departamentos)
	count := -1
	for i := 0; i <= first_dimention_size-1; i++ {
		for j := 0; j <= second_dimention_size-1; j++ {
			count++
			if count == VecPos {
				dato.Datos[i].Departamentos[j].Tiendas = append(dato.Datos[i].Departamentos[j].Tiendas, list.Get_store(vector[position])...)
			}
		}
	}
}

//var cartProducts []Products.CarritoType

func Carrito(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)
	var carrito Products.CarritoType
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.Unmarshal([]byte(reqBody), &carrito)
	//GUARDA LOS PRODUCTOS
	//fmt.Println(carrito)

	Products.CartProducts = append(Products.CartProducts, carrito)

	//JSON DE RESPUESTA
	json.NewEncoder(w).Encode(carrito)
	//SE VAN AL MACENANDO LAS RUTAS DE PEDIDOS
	for i := 0; i < len(carrito.Productos); i++ {
		if !Structs.ExisteEstacion(Structs.EstacionesDePedidos, carrito.Productos[i].Almacenamiento) {
			Structs.EstacionesDePedidos = append(Structs.EstacionesDePedidos, carrito.Productos[i].Almacenamiento)
		}
	}
	//DESPUES DE IR AL DESPACHO SE PASA AL INICIO NUEVAMENTE
	//SE CALCULA EL RECORRIDO
	Products.CalcularRecorrido(Graph, Structs.EstacionesDePedidos, d.PosicionInicialRobot, d.PosicionInicialRobot, d.Entrega)
	fmt.Println(Products.Ruta)
	Products.GraficarRuta()
}

func getCarrito(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Products.CartProducts)
}

func ResetearRuta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var nuevo []string
	Products.CartProducts = nil
	Products.Ruta = nuevo
	Structs.EstacionesDePedidos = nil
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getArreglo", Grafi).Methods(("GET"))
	router.HandleFunc("/guardar", Json_Returned).Methods(("GET"))
	router.HandleFunc("/cargartienda", CreateData).Methods("POST")
	router.HandleFunc("/id/{numero}", Linear_Browser).Methods("GET")
	router.HandleFunc("/Eliminar", Delete_Store).Methods(("DELETE"))
	router.HandleFunc("/TiendaEspecifica", Browser).Methods(("POST"))

	//ENPOINT PARA OBTENER LA EL JSON DE ARBOL DE AÑOS Y LISTA DE MESES
	router.HandleFunc("/Pedidos", Products.JsonMatriz).Methods(("POST"))
	//ENDPOINT PARA OBTENER LA MATRIZ Y LAS COLAS
	router.HandleFunc("/Matriz", Products.Matriz).Methods(("POST"))
	router.HandleFunc("/Matriz", Products.GetMatriz).Methods(("GET"))

	//ENVIO DE PRODUCTOS CONFIRMADOS EN EL CARRITO DE COMPRAS
	router.HandleFunc("/comprar", Products.Cobrar).Methods(("POST"))
	router.HandleFunc("/comprar", Products.Comprados).Methods(("GET"))

	//ENVIO DE CARRITO DE COMPRAS
	router.HandleFunc("/EnviarCarrito", Carrito).Methods(("POST"))
	router.HandleFunc("/EnviarCarrito", getCarrito).Methods(("GET"))
	//router.HandleFunc("/elSel", D_Select).Methods(("POST"))
	router.HandleFunc("/elSel", Products.Delete_Select).Methods(("POST"))

	//BUSQUEDA DE TIENDA PARA RETORNAR PRODUCTOS
	router.HandleFunc("/getTienda", Products.Tienda).Methods(("POST"))

	//RUTAS PARA CARGA DE PRODUCTOS
	router.HandleFunc("/CargarInventarios", Products.LoadInv).Methods(("POST"))

	//RUTA PARA CARGA DE PEDIDOS
	router.HandleFunc("/CargarPedidos", Products.LoadOrders).Methods(("POST"))

	//ENPOINTS PARA GRAFICAR REPORTES DE ESTRUCTURAS
	router.HandleFunc("/getMatriz", getCarrito).Methods(("GET"))
	router.HandleFunc("/getAnios", Products.GetAnios).Methods(("GET"))
	router.HandleFunc("/getMeses", getCarrito).Methods(("GET"))

	//RUTA PARA OBTENER LOS CAMINOS MAS CORTOS
	router.HandleFunc("/getRecorrido", Products.RealizarRecorrido).Methods(("GET"))
	//SE BORRAN LOS CAMINOS DEL PRODUCTO ACTUAL
	router.HandleFunc("/resetearRuta", ResetearRuta).Methods(("GET"))

	//FASE 3
	//RUTA DE CARGA DE USUARIOS
	router.HandleFunc("/usuarios", Products.LoadAcounts).Methods(("POST"))
	router.HandleFunc("/usuario", Products.LoadAcount).Methods(("POST"))
	//RUTA PARA BUSQUEDA DE USUARIO
	router.HandleFunc("/getUsuario", Products.GetUsuario).Methods(("POST"))
	//RUTA PARA BUSQUEDA DEL USUARIO LOGEADO ACTUAL
	router.HandleFunc("/getUsuarioActual", Products.GetUsuarioActual).Methods(("GET"))
	//RUTA PARA ENVIO DE COMENTARIO
	router.HandleFunc("/sendComent", Products.SendComment).Methods(("POST"))
	//RUTA PARA OBTENER LOS COMENTARIOS ALMACENADOS EN LA TABLA HASH
	router.HandleFunc("/sendComent", Products.GetComments).Methods(("GET"))

	router.HandleFunc("/recorrido", recorrido).Methods(("POST"))

	log.Fatal(http.ListenAndServe(":3000", router))

}

//STRUCT PARA RECIBIR EL GRAFO
type rec struct {
	Nodos []struct {
		Nombre  string `json:"Nombre"`
		Enlaces []struct {
			Nombre    string `json:"Nombre"`
			Distancia int    `json:"Distancia"`
		}
	}
	PosicionInicialRobot string `json:"PosicionInicialRobot"`
	Entrega              string `json:"Entrega"`
}

//FUNCION PARA ENVIAR EL MAPA DE ALMACENES
var d Structs.Rec

func recorrido(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.Unmarshal([]byte(reqBody), &d)
	//JSON DE RESPUESTA
	json.NewEncoder(w).Encode(d)
	RecogerPedidos(Structs.EstacionesDePedidos)
	//EL MAPA QUE SE OBTIENE SE MANDA A UNA VARIABLE GLOBAL PARA PODER GRAFICAR
	Products.MapaGlobal = d
	//Products.GraficarRuta()
}

//EN EL GRAPH SE GUARDAN TODAS LAS RUTAS CON LOS PESOS
var Graph Products.Graph

//SE GUARDAN LOS CAMINOS MAS CORTOS ENTRE CADA NODO
func RecogerPedidos(estaciones []string) {

	adyacencia := make([][]float64, len(d.Nodos))
	for i := range adyacencia {
		adyacencia[i] = make([]float64, len(d.Nodos))
	}

	//SE LLENA LA MATRIZ DE ADYACENCIA CON NUMEROS GRANDES
	for i := 0; i < len(d.Nodos); i++ {
		for j := 0; j < len(d.Nodos); j++ {
			adyacencia[i][j] = 1000000
		}
	}

	graph := Products.NewGraph()
	for i := 0; i < len(d.Nodos); i++ {
		for j := 0; j < len(d.Nodos[i].Enlaces); j++ {

			for k := 0; k < len(d.Nodos); k++ {
				if d.Nodos[k].Nombre == d.Nodos[i].Enlaces[j].Nombre {
					//SE INCERTAN LOS VALORES EN LA MATRIZ DE ADYACENCIA PARA LOS ENLACES
					//adyacencia[j][k] = float64(d.Nodos[i].Enlaces[j].Distancia)
					graph.AddEdge(d.Nodos[i].Nombre, d.Nodos[i].Enlaces[j].Nombre, d.Nodos[i].Enlaces[j].Distancia)
				}
			}
		}
	}

	Graph = *graph

}
