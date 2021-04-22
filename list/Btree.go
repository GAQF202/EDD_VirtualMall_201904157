package list

import (
	"strconv"

	"github.com/GAQF202/servidor-rest/Structs"
)

type Usuario struct {
	Dpi      int
	Nombre   string
	Correo   string
	Password string
	Cuenta   string
}

type NodoBTree struct {
	Leaf  bool
	N     int
	Key   []Structs.Usuario
	Hijo  []*NodoBTree
	Padre *NodoBTree
}

type BTree struct {
	grado int
	Root  *NodoBTree
}

func NewNode(grado int, parent *NodoBTree) *NodoBTree {
	var key []Structs.Usuario
	var hijos []*NodoBTree
	for i := 0; i < grado+1; i++ {
		hijos = append(hijos, nil)
	}
	for i := 0; i < grado; i++ {
		key = append(key, Structs.Usuario{})
	}

	return &NodoBTree{true, 0, key, hijos, parent}
}

func NewBTree(grado int) *BTree {
	return &BTree{grado, NewNode(grado, nil)}
}

func (act *NodoBTree) Insert(key Structs.Usuario) {
	act.Key[act.N] = key
	act.N++
	if act.N > 1 {
		ordenarPagina(act)
	}
}

/*func (act *NodoBTree) InsertNode(key *NodoBTree) {
	act.Key[act.N] = key
	act.Hijo[act.N]
	act.N++
	if act.N > 1 {
		ordenarPagina(act)
	}
}*/

func (a *BTree) Insert(key Structs.Usuario) {
	a.Root = a._Insert(key, a.Root)
}
func (a *BTree) split(key Structs.Usuario, tmp *NodoBTree) {

	//SI ES UNA HOJA Y NO TIENE PADRE ES LA RAIZ SIN HIJOS
	if tmp.N == a.grado && tmp.Padre != nil {
		//SI EL PADRE AUN NO ESTA LLENO
		if tmp.Padre.N != a.grado {
			mkey := tmp.Key[(a.grado-1)/2]
			tmp.Padre.Insert(mkey)
			index := 0

			for index = 0; index < tmp.Padre.N; index++ {
				if tmp.Padre.Key[index] == mkey {
					break
				}
			}
			for i := tmp.Padre.N; i > index+1; i-- {
				tmp.Padre.Hijo[i] = tmp.Padre.Hijo[i-1]
			}
			tmp.Padre.Hijo[index+1] = NewNode(a.grado, tmp.Padre)
			for i := ((a.grado - 1) / 2) + 1; i < a.grado; i++ {
				tmp.Padre.Hijo[index+1].Insert(tmp.Key[i])
			}
			aux := tmp
			tmp.Padre.Hijo[index] = NewNode(a.grado, tmp.Padre)
			for i := 0; i < (a.grado-1)/2; i++ {
				tmp.Padre.Hijo[index].Insert(aux.Key[i])
			}
		}

	}
}
func (a *BTree) _Insert(key Structs.Usuario, tmp *NodoBTree) *NodoBTree {

	//SI ES HOJA
	if tmp.Leaf {
		tmp.Insert(key)
		//SI ES UNA HOJA Y NO TIENE PADRE ES LA RAIZ SIN HIJOS
		if tmp.N == a.grado && tmp.Padre == nil {
			c := tmp
			tmp = NewNode(a.grado, nil)
			tmp.Insert(c.Key[(a.grado-1)/2])
			tmp.Hijo[0] = NewNode(a.grado, tmp)
			tmp.Hijo[1] = NewNode(a.grado, tmp)
			tmp.Hijo[0].Padre = tmp
			tmp.Hijo[1].Padre = tmp
			tmp.Leaf = false

			for i := 0; i < (a.grado-1)/2; i++ {
				tmp.Hijo[0].Insert(c.Key[i])
			}
			for i := ((a.grado - 1) / 2) + 1; i < a.grado; i++ {
				tmp.Hijo[1].Insert(c.Key[i])
			}
			//SI ES UNA HOJA CON PADRE Y SE LLENA SE PARTE Y SE SUBE EL NODO MEDIO AL PADRE
		} else if tmp.N == a.grado && tmp.Padre != nil {

			a.split(key, tmp)
			//SI UNA RAMA SE LLENA JAJA
			if tmp.Padre.N == a.grado && tmp.Padre.Padre != nil {

				padre := tmp.Padre
				tmp.Padre = NewNode(a.grado, tmp.Padre.Padre)
				mkey := padre.Key[(a.grado-1)/2]
				tmp.Padre.Insert(mkey)

				derecho := NewNode(a.grado, nil)
				izquierdo := NewNode(a.grado, nil)

				derecho.Insert(padre.Key[3])
				derecho.Insert(padre.Key[4])
				derecho.Hijo[0] = padre.Hijo[3]
				derecho.Hijo[1] = padre.Hijo[4]
				derecho.Hijo[2] = padre.Hijo[5]
				padre.Hijo[3].Padre = derecho
				padre.Hijo[4].Padre = derecho
				padre.Hijo[5].Padre = derecho

				izquierdo.Insert(padre.Key[0])
				izquierdo.Insert(padre.Key[1])
				izquierdo.Hijo[0] = padre.Hijo[0]
				izquierdo.Hijo[1] = padre.Hijo[1]
				izquierdo.Hijo[2] = padre.Hijo[2]
				padre.Hijo[0].Padre = izquierdo
				padre.Hijo[1].Padre = izquierdo
				padre.Hijo[2].Padre = izquierdo
				izquierdo.Leaf = false
				derecho.Leaf = false

				tmp.Padre.Hijo[0] = izquierdo
				tmp.Padre.Hijo[1] = derecho
				//SOLO FALTA INSERTAR EN EL PADRE EL tmp.padre
				tmp.Padre.Padre.Insert(tmp.Padre.Key[0])
				a.Root.N++
				index := 0

				for index = 0; index < tmp.Padre.Padre.N; index++ {
					if tmp.Padre.Padre.Key[index] == mkey {
						break
					}
				}
				tmp.Padre.Padre.Hijo[index] = izquierdo
				tmp.Padre.Padre.Hijo[index+1] = derecho

				izquierdo.Padre = tmp.Padre.Padre
				derecho.Padre = tmp.Padre.Padre
			}
			//SI LA RAIZ SE LLENA
			if a.Root.N == a.grado {
				padre := tmp.Padre
				tmp.Padre = NewNode(a.grado, nil)
				mkey := tmp.Key[(a.grado-1)/2]

				tmp.Padre.Insert(mkey)

				derecho := NewNode(a.grado, nil)
				izquierdo := NewNode(a.grado, nil)

				derecho.Insert(padre.Key[3])
				derecho.Insert(padre.Key[4])
				derecho.Hijo[0] = padre.Hijo[3]
				derecho.Hijo[1] = padre.Hijo[4]
				derecho.Hijo[2] = padre.Hijo[5]
				padre.Hijo[3].Padre = derecho
				padre.Hijo[4].Padre = derecho
				padre.Hijo[5].Padre = derecho

				izquierdo.Insert(padre.Key[0])
				izquierdo.Insert(padre.Key[1])
				izquierdo.Hijo[0] = padre.Hijo[0]
				izquierdo.Hijo[1] = padre.Hijo[1]
				izquierdo.Hijo[2] = padre.Hijo[2]
				padre.Hijo[0].Padre = izquierdo
				padre.Hijo[1].Padre = izquierdo
				padre.Hijo[2].Padre = izquierdo
				izquierdo.Leaf = false
				derecho.Leaf = false

				tmp.Padre.Hijo[0] = izquierdo
				tmp.Padre.Hijo[1] = derecho

				izquierdo.Padre = tmp.Padre
				derecho.Padre = tmp.Padre

				*a.Root = *tmp.Padre
				a.Root.Leaf = false
			}
		}
	} else {

		found := false
		for i := 0; i < tmp.N; i++ {
			if key.Dpi < tmp.Key[i].Dpi {
				found = true
				a._Insert(key, tmp.Hijo[i])
				break
			}
		}
		if !found {
			a._Insert(key, tmp.Hijo[tmp.N])
		}

	}
	return tmp
}

var encontrado Structs.Usuario

func (a BTree) Buscar(key int, tmp *NodoBTree) Structs.Usuario {

	for j := 0; j < len(tmp.Key); j++ {
		if tmp.Key[j].Dpi == key {
			encontrado = tmp.Key[j]
			break
		}
	}
	for j := 0; j < len(tmp.Hijo); j++ {
		if tmp.Hijo[j] != nil {
			a.Buscar(key, tmp.Hijo[j])
		}
	}
	return encontrado
}

var GraficaArbol string = ""
var GraficaArbolDatosSensibles string = ""
var GraficaArbolEncriptado string = ""
var contadorPagina int = 0

func VerElementos(n *NodoBTree) {
	dot := "\npagina" + strconv.Itoa(n.Key[0].Dpi) + "[label=\""
	dotS := "\npagina" + strconv.Itoa(n.Key[0].Dpi) + "[label=\""
	//dotEncriptado := "\npagina" + strconv.Itoa(n.Key[0].Dpi) + "[label=\""

	contadorPagina++
	for j := 0; j < len(n.Key); j++ {
		if j < len(n.Key)-1 {
			if n.Key[j].Dpi != 0 {
				dot += "<" + strconv.Itoa(n.Key[j].Dpi) + "h" + strconv.Itoa(j) + ">" + "|"
				dot += "<" + strconv.Itoa(n.Key[j].Dpi) + ">" + strconv.Itoa(n.Key[j].Dpi) + "\\n" + n.Key[j].Cuenta + "\\n" + n.Key[j].Nombre + "\\n" + n.Key[j].Correo + "\\n" + n.Key[j].Password + "|"

				dotS += "<" + strconv.Itoa(n.Key[j].Dpi) + "h" + strconv.Itoa(j) + ">" + "|"
				dotS += "<" + strconv.Itoa(n.Key[j].Dpi) + ">" + Encrypt([]byte("123"), strconv.Itoa(n.Key[j].Dpi)) + "\\n" + n.Key[j].Cuenta + "\\n" + n.Key[j].Nombre + "\\n" + Encrypt([]byte("123"), n.Key[j].Correo) + "\\n" + Encrypt([]byte("123"), n.Key[j].Password) + "|"
				//dot += "<" + strconv.Itoa(n.Key[j].Dpi) + "h" + strconv.Itoa(j) + ">" + "|"
			}
		} else {
			dot += "<" + strconv.Itoa(n.Key[j].Dpi) + "h" + strconv.Itoa(j) + ">" + "|"
			dot += "<" + strconv.Itoa(n.Key[j].Dpi) + ">" + strconv.Itoa(n.Key[j].Dpi) + "|"
			dot += "<" + strconv.Itoa(n.Key[j].Dpi) + "h" + strconv.Itoa(j) + ">"

			dotS += "<" + strconv.Itoa(n.Key[j].Dpi) + "h" + strconv.Itoa(j) + ">" + "|"
			dotS += "<" + strconv.Itoa(n.Key[j].Dpi) + ">" + strconv.Itoa(n.Key[j].Dpi) + "|"
			dotS += "<" + strconv.Itoa(n.Key[j].Dpi) + "h" + strconv.Itoa(j) + ">"
		}
		if n.Hijo[j] != nil {
			GraficaArbol += "\npagina" + strconv.Itoa(n.Key[0].Dpi) + ":\"" + strconv.Itoa(n.Key[j].Dpi) + "h" + strconv.Itoa(j) + "\" -> pagina" + strconv.Itoa(n.Hijo[j].Key[0].Dpi)
		}
	}
	for j := 0; j < len(n.Hijo); j++ {
		if n.Hijo[j] != nil {
			//fmt.Println(n.Hijo[j])
			VerElementos(n.Hijo[j])
		}
	}
	//fmt.Println(dot + "\"]")
	GraficaArbol += dot + "\"]\n"
	GraficaArbolDatosSensibles += dotS + "\"]\n"
}

func ordenarPagina(paginaDesordenada *NodoBTree) {
	var aux Structs.Usuario
	for i := 0; i < paginaDesordenada.N-1; i++ {
		for j := i + 1; j < paginaDesordenada.N; j++ {
			if paginaDesordenada.Key[i].Dpi > paginaDesordenada.Key[j].Dpi {
				aux = paginaDesordenada.Key[i]
				paginaDesordenada.Key[i] = paginaDesordenada.Key[j]
				paginaDesordenada.Key[j] = aux
			}
		}
	}
}
