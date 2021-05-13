package Products

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/GAQF202/servidor-rest/Structs"
	"github.com/GAQF202/servidor-rest/dijkstra"
	"github.com/GAQF202/servidor-rest/dot"
	"github.com/GAQF202/servidor-rest/list"
)

//VARIABLE GLOBAR PARA ALMACENAR EL ARBOL DE ANIOS
var Pedidos list.AVLYear

//STRUCT PARA RECIBIR LOS INVENTARIOS
type InventoryType struct {
	Inventarios []struct {
		Tienda       string `json:"Tienda"`
		Departamento string `json:"Departamento"`
		Calificacion int    `json:"Calificacion"`
		Productos    []struct {
			Nombre         string  `json:"Nombre"`
			Codigo         int     `json:"Codigo"`
			Descripcion    string  `json:"Descripcion"`
			Precio         float64 `json:"Precio"`
			Cantidad       int     `json:"Cantidad"`
			Imagen         string  `json:"Imagen"`
			Almacenamiento string  `json:"Almacenamiento"`
		}
	}
}

type CarritoType struct {
	Tienda       string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int    `json:"Calificacion"`
	Productos    []struct {
		Nombre         string  `json:"Nombre"`
		Codigo         int     `json:"Codigo"`
		Descripcion    string  `json:"Descripcion"`
		Precio         float64 `json:"Precio"`
		Cantidad       int     `json:"Cantidad"`
		Imagen         string  `json:"Imagen"`
		Almacenamiento string  `json:"Almacenamiento"`
	}
}

//STRUCT PARA RECIBIR LOS PEDIDOS
type OrderType struct {
	Pedidos []struct {
		Fecha        string `json:"Fecha"`
		Tienda       string `json:"Tienda"`
		Departamento string `json:"Departamento"`
		Calificacion int    `json:"Calificacion"`
		Productos    []struct {
			Codigo int `json:"Codigo"`
		}
	}
}

//STRUCT PARA GUARDAR PRODUCTOS
type Product struct {
	Nombre         string
	Codigo         int
	Descripcion    string
	Precio         float64
	cantidad       int
	Imagen         string
	Almacenamiento string
}

var Inventory InventoryType

//FUNCION PARA INGRESAR LOS INVENTARIOS EN EL STRUCT
func LoadInv(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}

	json.Unmarshal([]byte(reqBody), &Inventory)
	//SE LLAMA A LA FUNCION PARA CREAR LOS INVENTARIOS DENTRO DE LAS TIENDAS
	add_inventory(Inventory)

}

//FUNCION PARA INSERTAR LOS INVENTARIOS DENTRO DE LAS TIENDAS
func add_inventory(inventory InventoryType) {
	var hashProductos []dijkstra.Hashable
	for i := 0; i < len(inventory.Inventarios); i++ {
		Position := list.Get_position(inventory.Inventarios[i].Departamento, inventory.Inventarios[i].Tienda, inventory.Inventarios[i].Calificacion)
		for j := 0; j < len(inventory.Inventarios[i].Productos); j++ {
			tmp := inventory.Inventarios[i].Productos[j]
			product := Structs.Product{tmp.Nombre, tmp.Codigo, tmp.Descripcion, tmp.Precio, tmp.Cantidad, tmp.Imagen, tmp.Almacenamiento}
			hashProductos = append(hashProductos, dijkstra.Block(product.Nombre+"\\n"+product.Almacenamiento+"\\n"+strconv.FormatFloat(product.Precio, 'E', -1, 32)+"\\n"+product.Descripcion))
			list.Get_store_node(inventory.Inventarios[i].Tienda, inventory.Inventarios[i].Calificacion, list.GlobalVector[Position], product)
		}
		list.VerNodos(list.GlobalVector[Position])
	}
	dijkstra.PrintTree(dijkstra.BuildTree(hashProductos)[0].(dijkstra.Node))
	dot.CrearArchivoEvery(dijkstra.DotMerkleTree+"}", "txt", "DotAnios")
	dot.GraphEvery("MerkleProductos", "jpg", "DotAnios")
	dijkstra.DotMerkleTree = "digraph { node [shape=box, style=\"filled\", fillcolor=\"#61e665\"];"
}

//FUNCION PARA INGRESAR LOS PEDIDOS EN EL STRUCT
func LoadOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)

	var Order OrderType

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}

	json.Unmarshal([]byte(reqBody), &Order)
	//SE LLAMA A LA FUNCION PARA CREAR LOS INVENTARIOS DENTRO DE LAS TIENDAS
	add_orders(Order)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Order)
	reporteAnios()
	reporteMeses()
	list.Ver = ""
}

//CREA LA ESTRUCTURA PARA CADA NODO DE ANIO
func add_orders(orders OrderType) {
	//fmt.Println(orders)
	for i := 0; i < len(orders.Pedidos); i++ {
		//SE SEPARA LA FECHA POR DIA, MES Y ANIO
		fecha := strings.Split(orders.Pedidos[i].Fecha, "-")
		dia, err := strconv.Atoi(fecha[0])
		m, err := strconv.Atoi(fecha[1])
		mes := Structs.Get_month(m)
		anio, err := strconv.Atoi(fecha[2])
		if err != nil {
			fmt.Println("Error al ingresar fecha de pedido")
		}

		categoria := orders.Pedidos[i].Departamento

		matriz := list.NewMatriz()
		soloMes := &list.Month{mes, matriz}
		mesActual := list.NewListMes()
		mesActual.Insertar(soloMes)
		anioActual := list.Year{anio, mesActual}

		for j := 0; j < len(orders.Pedidos[i].Productos); j++ {
			cola := list.NewQueue()
			cola.Add(&Structs.CodeProduct{orders.Pedidos[i].Productos[j].Codigo})
			matriz.Insert(cola, dia, categoria)
		}
		Pedidos.Add(&anioActual)
	}

	//Pedidos.Preorder(Pedidos.Raiz)
}

//STRUCT PARA ENVIAR EL JSON DE AÑOS Y MESES
var JsonArbolAnios []Structs.Anio

func JsonMatriz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//SE RECORRE EL ARBOL PARA INGRESAR LOS AÑOS EN EL STRUCT
	JsonArbolAnios = Pedidos.Inorder(Pedidos.Raiz)

	//SE MANDA EL JSON AL BODY
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(JsonArbolAnios)
	//SE REINICIAN LOS STRUCTS DE LAS ESTRUCTURAS INTERNAS
	list.Anios = []Structs.Anio{}
	list.Res = []Structs.Mes{}
}

type MesAnio struct {
	Anio int    `json:"Anio"`
	Mes  string `json:"Mes"`
}

//STRUCT PARA GUARDAR LA MATRIZ SELECCIONADA
var Matrix []Structs.MesSeleccionado

var MyA MesAnio

//FUNCION PARA ENVIAR EL DATO DE UNA MATRIZ ESPECIFICA
func Matriz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)

	json.Unmarshal([]byte(reqBody), &MyA)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}
	//SE MANDA EL JSON AL BODY
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	getJsonMatrix(MyA)
	json.NewEncoder(w).Encode(Matrix)
	Matrix = nil
}

//FUNCION PARA RETORNAR EL JSON DE RESPUES SEGUN LA SELECCION DE MES Y ANIO
func getJsonMatrix(YearAndMonth MesAnio) {
	Pedidos.BuscarAnio(Pedidos.Raiz, YearAndMonth.Mes, YearAndMonth.Anio)
	//Matrix = list.Meses
	Matrix = append(Matrix, list.Meses)
	fmt.Println(Matrix)

	list.Meses = Structs.MesSeleccionado{}
	list.Cola = []Structs.Cola{}
	list.Productos = Structs.CodigoDeProducto{}
}

func GetMatriz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)

	json.Unmarshal([]byte(reqBody), &MyA)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}
	//SE MANDA EL JSON AL BODY
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(Matrix)
}
