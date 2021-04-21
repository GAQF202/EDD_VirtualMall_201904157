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

	for i := 0; i < len(cuentas.Usuarios); i++ {
		cuentaActual := Structs.Usuario{cuentas.Usuarios[i].Dpi, cuentas.Usuarios[i].Nombre, cuentas.Usuarios[i].Correo, cuentas.Usuarios[i].Password, cuentas.Usuarios[i].Cuenta}
		//SE INSERTAN LOS USUARIOS EN ARBOL LOCAL
		ArbolDeUsuariosGlobal.Insert(cuentaActual)
	}

	//SE GRAFICA EL ARBOL A PARTIR DE LA RAIZ
	list.GraficaArbol = ""
	list.VerElementos(ArbolDeUsuariosGlobal.Root)
	fmt.Println(list.GraficaArbol)
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
	fmt.Println(list.GraficaArbol)
}

type getUser struct {
	Dpi      string `json:"Dpi"`
	Password string `json:"Password"`
}

var usuarioEncontrado getUser

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
		} else {
			json.NewEncoder(w).Encode("Contraseña incorrecta")
		}

	} else {
		json.NewEncoder(w).Encode("Usuario no encontrado")
	}

}
