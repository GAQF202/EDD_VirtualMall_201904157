package Structs

//STRUCT PARA GUARDAR PRODUCTOS
type Product struct {
	Nombre      string
	Codigo      int
	Descripcion string
	Precio      float64
	Cantidad    int
	Imagen      string
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

/*type Year struct {
	Year int
}

type Month struct {
	Month string
}*/
