package Products

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/GAQF202/servidor-rest/Structs"
	"github.com/GAQF202/servidor-rest/list"
)

//VARIABLE PARA ALMACENAR LOS PRODUCTOS DEL CARRITO
var CartProducts []CarritoType

//STRUCT PARA BUSQUEDA DE TIENDA
type Buscar_tienda struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
}

var tienda Buscar_tienda

//BUSQUEDA DE TIENDA PARA MOSTRAR LOS PRODUCTOS
func Tienda(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Inserte una tienda existente")
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.Unmarshal([]byte(reqBody), &tienda)

	res := Get_Producto(tienda)
	//JSON DE RESPUESTA
	json.NewEncoder(w).Encode(res)
}

type TiendaEcontrada struct {
	Productos []struct{}
}

func Get_Producto(tienda Buscar_tienda) list.InventoryType {
	Position := list.Get_position(tienda.Departamento, tienda.Nombre, tienda.Calificacion)
	return list.JsonInventory(tienda.Nombre, tienda.Calificacion, list.GlobalVector[Position], tienda.Departamento)
}

//FUNCION PARA RECIBIR LAS COMPRAS DEL CARRITO
var carrito CarritoType

func Cobrar(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Inserte una tienda existente")
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.Unmarshal([]byte(reqBody), &carrito)

	//fmt.Println(carrito)
	removeProd(CartProducts)
	//AL TERMINAR LA COMPRA EL CARRITO SE QUEDA VACIO
	CartProducts = nil
}

func Comprados(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Inserte una tienda existente")
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.Unmarshal([]byte(reqBody), &carrito)
	json.NewEncoder(w).Encode(CartProducts)
}

//RESTA LOS PEDIDOS SELECCIONADOS EN EL CARRITO DE COMPRAS
func removeProd(inventory []CarritoType) {

	//SE ALMACENAN LOS PRODUCTOS OBTENIDOS
	for i := 0; i < len(inventory); i++ {
		//BUSCA LA POSICION DE LA TIENDA
		Position := list.Get_position(inventory[i].Departamento, inventory[i].Tienda, inventory[i].Calificacion)
		for j := 0; j < len(inventory[i].Productos); j++ {
			p := inventory[i].Productos[j]
			producto := Structs.Product{p.Nombre, p.Codigo, p.Descripcion, p.Precio, p.Cantidad, p.Imagen, p.Almacenamiento}

			list.Delete_product(inventory[i].Tienda, inventory[i].Calificacion, list.GlobalVector[Position], producto)
		}
	}

}

//FUNCION PARA RECIBIR LAS COMPRAS DEL CARRITO
var elim CarritoType

func Delete_Select(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Inserte una tienda existente")
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.Unmarshal([]byte(reqBody), &elim)

	json.NewEncoder(w).Encode(elim)
	/*for i := 0; i < len(CartProducts); i++ {
		for j := 0; j < len(elim.Productos); j++ {
			//BUSQUEDA DE TIENDA Y PRODUCTO
			if CartProducts[i].Tienda == elim.Tienda && CartProducts[i].Calificacion == elim.Calificacion && CartProducts[i].Departamento == elim.Departamento && CartProducts[i].Productos[j].Codigo == elim.Productos[j].Codigo {
				fmt.Println(CartProducts[i])
				CartProducts[i].Productos[j].Cantidad = 0
			}

		}
	}*/

}

var Ruta []string

//FUNCION PARA CALCULAR EL CAMINO MAS CORTO
func CalcularRecorrido(grafo Graph, paradas []string, inicio string, inicioAbs string, fin string) {
	var paradasTemporal []ByWay
	//SE GUARDAN TODAS LAS POSIBLES RUTAS DESDE EL DESTINO PARAMETRO
	for _, est := range paradas {
		ParadaActual := grafo.GetPath(inicio, est)
		paradasTemporal = append(paradasTemporal, ParadaActual)
	}

	var paradaAInsertar ByWay
	//SE BUSCA QUE RUTA ES MAS CORTA DESDE EL DESTINO PARAMETRO
	for _, parada := range paradasTemporal {
		if paradaAInsertar.PesoTotal < parada.PesoTotal {
			paradaAInsertar = parada
		}
	}
	//SE INSERTA EN LA VARIABLE GLOBAL DE RUTAS LA RUTA MENOR
	for i := 0; i < len(paradaAInsertar.Estaciones); i++ {
		Ruta = append(Ruta, paradaAInsertar.Estaciones[i])
	}
	//SI AUN HAY RUTAS POR RECORRER SE REPITE EL PROCESO
	if len(paradas) != 0 {
		NuevoSlice := remove(paradas, 0)
		CalcularRecorrido(grafo, NuevoSlice, paradaAInsertar.Estaciones[1], inicioAbs, fin)
	} else {
		ParadaActual := grafo.GetPath(Ruta[len(Ruta)-1], fin)
		for i := 0; i < len(ParadaActual.Estaciones); i++ {
			Ruta = append(Ruta, ParadaActual.Estaciones[i])
		}
		Regreso := grafo.GetPath(fin, inicioAbs)
		for i := 0; i < len(Regreso.Estaciones); i++ {
			Ruta = append(Ruta, Regreso.Estaciones[i])
		}
	}

}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

var MapaGlobal Structs.Rec

func RealizarRecorrido(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	//JSON DE RESPUESTA
	json.NewEncoder(w).Encode(Ruta)
	GraficarRuta()
}

func GraficarRuta() {
	dot := ""
	for i := 0; i < len(MapaGlobal.Nodos); i++ {
		dot += MapaGlobal.Nodos[i].Nombre + " [label=\"" + MapaGlobal.Nodos[i].Nombre + "\"];\n"
		for j := 0; j < len(MapaGlobal.Nodos[i].Enlaces); j++ {
			dot += MapaGlobal.Nodos[i].Nombre + " -- " + MapaGlobal.Nodos[i].Enlaces[j].Nombre + "[label = " + strconv.Itoa(MapaGlobal.Nodos[i].Enlaces[j].Distancia) + "];\n"
		}
	}

	for i := 0; i < len(Ruta); i++ {
		if i < len(Ruta)-1 {
			if Ruta[i] != Ruta[i+1] {
				dot += Ruta[i] + "--" + Ruta[i+1] + "[color=\"#f02c2c\"];"
			}
		}
	}
	ReporteRecorrido(dot)
}
