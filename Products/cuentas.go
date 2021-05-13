package Products

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/GAQF202/servidor-rest/Structs"
	"github.com/GAQF202/servidor-rest/dijkstra"
	"github.com/GAQF202/servidor-rest/dot"
	"github.com/GAQF202/servidor-rest/hashes"
	"github.com/GAQF202/servidor-rest/list"
)

type acount struct {
	Usuarios []struct {
		Dpi      int    `json:"Dpi"`
		Nombre   string `json:"Nombre"`
		Correo   string `json:"Correo"`
		Password string `json:"Password"`
		Cuenta   string `json:"Cuenta"`
	}
}

type user struct {
	Dpi      string `json:"Dpi"`
	Nombre   string `json:"Nombre"`
	Correo   string `json:"Correo"`
	Password string `json:"Password"`
	Cuenta   string `json:"Cuenta"`
}

type UsuarioActual struct {
	Usuario string
	DPI     int
}

var UsuarioGlobalActual UsuarioActual

//VARIABLE GLOBAL PARA GUARDAR TODOS LOS USUARIOS
var usuarios acount
var usuario user

func LoadAcounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}

	json.Unmarshal([]byte(reqBody), &usuarios)
	InsertInBTree(usuarios)
}

var ArbolDeUsuariosGlobal = list.NewBTree(5)

func InsertInBTree(cuentas acount) {
	var hashUsuarios []dijkstra.Hashable
	for i := 0; i < len(cuentas.Usuarios); i++ {
		cuentaActual := Structs.Usuario{cuentas.Usuarios[i].Dpi, cuentas.Usuarios[i].Nombre, cuentas.Usuarios[i].Correo, cuentas.Usuarios[i].Password, cuentas.Usuarios[i].Cuenta}
		hashUsuarios = append(hashUsuarios, dijkstra.Block(cuentaActual.Nombre+"\\n"+cuentaActual.Correo+"\\n"+cuentaActual.Cuenta+"\\n"+strconv.Itoa(cuentaActual.Dpi)))
		//SE INSERTAN LOS USUARIOS EN ARBOL LOCAL
		ArbolDeUsuariosGlobal.Insert(cuentaActual)
	}
	dijkstra.PrintTree(dijkstra.BuildTree(hashUsuarios)[0].(dijkstra.Node))
	dot.CrearArchivoEvery(dijkstra.DotMerkleTree+"}", "txt", "DotAnios")
	dot.GraphEvery("MerkleUsuarios", "jpg", "DotAnios")
	dijkstra.DotMerkleTree = "digraph { node [shape=box, style=\"filled\", fillcolor=\"#61e665\"];"

	//SE GRAFICA EL ARBOL A PARTIR DE LA RAIZ
	list.GraficaArbol = ""
	list.GraficaArbolDatosSensibles = ""
	list.VerElementos(ArbolDeUsuariosGlobal.Root)
	//fmt.Println(list.GraficaArbol)
	ReporteUsuarios(list.GraficaArbol)
	//fmt.Println(list.GraficaArbolDatosSensibles)
	//fmt.Println(ArbolDeUsuariosGlobal.Root)
	//SE PASA EL ARBOL LOCAL A GLOBAL
	//ArbolDeUsuariosGlobal = *ArbolDeUsuarios
}

//GUARDA UN SOLO USUARIO EN EL ARBOL DE CUENTAS
func LoadAcount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}

	json.Unmarshal([]byte(reqBody), &usuario)
	dpi, err := strconv.Atoi(usuario.Dpi)
	cuentaActual := Structs.Usuario{dpi, usuario.Nombre, usuario.Correo, usuario.Password, usuario.Cuenta}
	ArbolDeUsuariosGlobal.Insert(cuentaActual)

	//SE IMPRIME EL ARBOL A PARTIR DE LA RAIZ
	list.GraficaArbol = ""
	list.VerElementos(ArbolDeUsuariosGlobal.Root)
	ReporteUsuarios(list.GraficaArbol)
	//fmt.Println(list.GraficaArbol)
}

type getUser struct {
	Dpi      string `json:"Dpi"`
	Password string `json:"Password"`
}

type GetComent struct {
	Usuario    string `json:"Usuario"`
	Dpi        string `json:"Dpi"`
	Comentario string `json:"Comentario"`
}

var usuarioEncontrado getUser

//FUNCION PARA OBTENER EL USUARIO LOGUEADO ACTUAL
func GetUsuarioActual(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//SE MANDA EL JSON AL BODY
	json.NewEncoder(w).Encode(UsuarioGlobalActual)
	//fmt.Println(UsuarioGlobalActual)
}

//FUNCION PARA OBTENER USUARIOS
func GetUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)

	json.Unmarshal([]byte(reqBody), &usuarioEncontrado)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}

	dpi, err := strconv.Atoi(usuarioEncontrado.Dpi)
	encontrado := ArbolDeUsuariosGlobal.Buscar(dpi, ArbolDeUsuariosGlobal.Root)

	//SE MANDA EL JSON AL BODY
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(encontrado)
	if encontrado.Dpi != 0 {
		if usuarioEncontrado.Password == encontrado.Password {
			json.NewEncoder(w).Encode(encontrado)

			UsuarioGlobalActual = UsuarioActual{encontrado.Nombre, encontrado.Dpi}
		} else {
			json.NewEncoder(w).Encode("Contraseña incorrecta")
		}

	} else {
		json.NewEncoder(w).Encode("Usuario no encontrado")
	}

}

//VARIABLE PARA CUARDAR EL COMENTARIO
var comentario Structs.GetComent

//GUARDA LA POSICION DE LA TIENDA ACTUAL
var PosicionTiendaActual list.InventoryType
var PosicionVectorActual int

//FUNCION PARA ENVIO DE COMENTARIOS
func SendComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, err := ioutil.ReadAll(r.Body)

	json.Unmarshal([]byte(reqBody), &comentario)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(comentario)
	//fmt.Println("Posicion", PosicionTiendaActual)
	comentarioActual := hashes.Comentario{Usuario: comentario.Usuario, Comentario: comentario.Comentario}
	list.GuardarComentarios(PosicionTiendaActual.Tienda, PosicionTiendaActual.Calificacion, list.GlobalVector[PosicionVectorActual], comentario, comentarioActual)
}

//FUNCION PARA OBTENCION DE COMENTARIOS
func GetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	v := list.ObtenerComentarios(PosicionTiendaActual.Tienda, PosicionTiendaActual.Calificacion, list.GlobalVector[PosicionVectorActual])
	res := v.GetElements()
	json.NewEncoder(w).Encode(res)
	//fmt.Println(res)
}
