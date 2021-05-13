package Products

import (
	"net/http"

	"github.com/GAQF202/servidor-rest/dot"
)

func GetAnios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	hello := "digraph \n { rankdir=UD concentrate=true \n"
	hello += Pedidos.Raiz.GetCodigoInterno(Pedidos.Raiz)
	dot.CrearArchivoEvery(hello+"}", "txt", "DotAnios")
	dot.GraphEvery("EstructuraAnios", "jpg", "DotAnios")
}

//OBTENER EL STRING DE REPORTE DE ANIOS
func reporteAnios() {
	hello := "digraph \n { rankdir=UD concentrate=true \n"
	hello += Pedidos.Raiz.GetCodigoInterno(Pedidos.Raiz)
	dot.CrearArchivoEvery(hello+"}", "txt", "DotAnios")
	dot.GraphEvery("EstructuraAnios", "jpg", "DotAnios")
}

//OBTENER STRING DE REPORTE DE MESES
func reporteMeses() {

	hello := "digraph \n { rankdir=LR concentrate=false \n"
	hello += Pedidos.RecorrerInOrder(Pedidos.Raiz)
	dot.CrearArchivo(hello+"}", "txt")
	dot.GraphEvery("EstructuraMeses", "jpg", "DotFile")
}

func ReporteRecorrido(recorridoDot string) {
	hello := "graph \n { rankdir=LR concentrate=false \n"
	hello += recorridoDot
	dot.CrearArchivo(hello+"}", "txt")
	dot.GraphEvery("Recorrido", "jpg", "DotFile")

}

func ReporteUsuarios(recorridoDot string) {
	hello := "digraph \n { node[style=\"filled\",fillcolor=\"#8df7ef\",shape=\"record\"] \n"
	hello += recorridoDot
	dot.CrearArchivo(hello+"}", "txt")
	dot.GraphEvery("Usuarios", "jpg", "DotFile")

}
