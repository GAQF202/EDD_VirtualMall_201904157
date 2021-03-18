package Products

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/GAQF202/servidor-rest/Structs"
	"github.com/GAQF202/servidor-rest/list"
)

//STRUCT PARA RECIBIR LOS INVENTARIOS
type InventoryType struct {
	Inventarios []struct {
		Tienda       string `json:"Tienda"`
		Departamento string `json:"Departamento"`
		Calificacion int    `json:"Calificacion"`
		Productos    []struct {
			Nombre      string  `json:"Nombre"`
			Codigo      int     `json:"Codigo"`
			Descripcion string  `json:"Descripcion"`
			Precio      float64 `json:"Precio"`
			Cantidad    int     `json:"Cantidad"`
			Imagen      string  `json:"Imagen"`
		}
	}
}

//STRUCT PARA GUARDAR PRODUCTOS
type Product struct {
	Nombre      string
	Codigo      int
	Descripcion string
	Precio      float64
	cantidad    int
	Imagen      string
}

func LoadInv(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	var Inventory InventoryType

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}

	json.Unmarshal([]byte(reqBody), &Inventory)
	fmt.Println(Inventory)
	//SE LLAMA A LA FUNCION PARA CREAR LOS INVENTARIOS DENTRO DE LAS TIENDAS
	//add_inventory(Inventory)
}

//FUNCION PARA INSERTAR LOS INVENTARIOS DENTRO DE LAS TIENDAS
func add_inventory(inventory InventoryType) {

	for i := 0; i < len(inventory.Inventarios); i++ {
		Position := list.Get_position(inventory.Inventarios[i].Departamento, inventory.Inventarios[i].Tienda, inventory.Inventarios[i].Calificacion)
		for j := 0; j < len(inventory.Inventarios[i].Productos); j++ {
			tmp := inventory.Inventarios[i].Productos[j]
			product := Structs.Product{tmp.Nombre, tmp.Codigo, tmp.Descripcion, tmp.Precio, tmp.Cantidad, tmp.Imagen}
			list.Get_store_node(inventory.Inventarios[i].Tienda, inventory.Inventarios[i].Calificacion, list.GlobalVector[Position], product)
		}
		list.VerNodos(list.GlobalVector[Position])
	}
}
