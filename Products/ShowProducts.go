package Products

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

//RESTA LOS PEDIDOS SELECCIONADOS EN EL CARRITO DE COMPRAS
func removeProd(inventory []CarritoType) {

	//SE ALMACENAN LOS PRODUCTOS OBTENIDOS
	for i := 0; i < len(inventory); i++ {
		//BUSCA LA POSICION DE LA TIENDA
		Position := list.Get_position(inventory[i].Departamento, inventory[i].Tienda, inventory[i].Calificacion)
		for j := 0; j < len(inventory[i].Productos); j++ {
			p := inventory[i].Productos[j]
			producto := Structs.Product{p.Nombre, p.Codigo, p.Descripcion, p.Precio, p.Cantidad, p.Imagen}

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
