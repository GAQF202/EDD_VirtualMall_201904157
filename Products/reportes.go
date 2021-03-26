package Products

import (
	"net/http"

	"github.com/GAQF202/servidor-rest/dot"
)

func GetAnios(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	hello := "digraph \n { rankdir=UD concentrate=true \n"
	hello += Pedidos.Raiz.GetCodigoInterno(Pedidos.Raiz)
	dot.CrearArchivo(hello+"}", "txt")
	dot.GraphEvery("EstructuraAnios", "jpg")
	//fmt.Println(Pedidos.Raiz.GetCodigoInterno(Pedidos.Raiz))
}

//OBTENER EL STRING DE REPORTE DE ANIOS
func reporteAnios() {
	hello := "digraph \n { rankdir=UD concentrate=true \n"
	hello += Pedidos.Raiz.GetCodigoInterno(Pedidos.Raiz)
	dot.CrearArchivo(hello+"}", "txt")
	dot.GraphEvery("EstructuraAnios", "jpg")
}

//OBTENER STRING DE REPORTE DE MESES
func reporteMeses() {

	hello := "digraph \n { rankdir=LR concentrate=false \n"
	hello += Pedidos.RecorrerInOrder(Pedidos.Raiz)
	dot.CrearArchivo(hello+"}", "txt")
	dot.GraphEvery("EstructuraMeses", "jpg")
	/*hello := "digraph \n { rankdir=UD concentrate=true \n"
	hello += Pedidos.Raiz.GetCodigoInterno(Pedidos.Raiz)
	dot.CrearArchivo(hello+"}", "txt")
	dot.GraphEvery("EstructuraAnios", "jpg")*/
}
