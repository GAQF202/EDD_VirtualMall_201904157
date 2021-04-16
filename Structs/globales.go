package Structs

//STRUCT PARA GUARDAR PRODUCTOS
type Product struct {
	Nombre         string
	Codigo         int
	Descripcion    string
	Precio         float64
	Cantidad       int
	Imagen         string
	Almacenamiento string
}

type CodeProduct struct {
	//CODIGO DE CADA PRODUCTO A GUARDAR EN LA COLA
	Codigo_producto int
}

func Get_month(mes int) string {
	if mes == 1 {
		return "Enero"
	} else if mes == 2 {
		return "Febrero"
	} else if mes == 3 {
		return "Marzo"
	} else if mes == 4 {
		return "Abril"
	} else if mes == 5 {
		return "Mayo"
	} else if mes == 6 {
		return "Junio"
	} else if mes == 7 {
		return "Juio"
	} else if mes == 8 {
		return "Agosto"
	} else if mes == 9 {
		return "Septiembre"
	} else if mes == 10 {
		return "Octubre"
	} else if mes == 11 {
		return "Noviembre"
	} else if mes == 12 {
		return "Diciembre"
	}
	return "Mes desconocido"
}

//ESTRUCTS PARA SELECCION Y BUSQUEDA DE MES Y ANIO
type Anio struct {
	Anio  int
	Meses []Mes
}

type Mes struct {
	Mes string
}

//STRUCTS PARA CUANDO YA SE HAYA SELECCIONADO EL MES Y EL ANIO
type MesSeleccionado struct {
	Mes    string
	Matriz []Cola
}

type Cola struct {
	Dia            int
	Categoria      string
	CodigoProducto []int
}

type CodigoDeProducto []int

//TYPE GLOBAL PARA MANEJAR LOS USUARIOS
type Usuario struct {
	Dpi      int
	Nombre   string
	Correo   string
	Password string
	Cuenta   string
}

var EstacionesDePedidos []string

func ExisteEstacion(list []string, val string) bool {
	res := false
	for _, element := range list {
		if element == val {
			res = true
		}
	}
	return res
}

type Caminos struct {
	Inicio string
	Fin    string
	Puntos []string
}

//STRUCT PARA RECIBIR EL GRAFO
type Rec struct {
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

/*type Year struct {
	Year int
}

type Month struct {
	Month string
}*/
