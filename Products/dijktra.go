package Products

import (
	hp "container/heap"
	"strconv"
)

func min(v1 float64, v2 float64) float64 {
	if v1 < v2 {
		return v1
	}
	return v2
}

func caminosR(i int, k int, caminosAuxiliares [][]string, caminoRecorrido string) string {
	if caminosAuxiliares[i][k] == "" {
		return ""
	} else {
		b, _ := strconv.Atoi(caminosAuxiliares[i][k])
		caminoRecorrido += caminosR(i, b, caminosAuxiliares, caminoRecorrido) + strconv.Itoa(b+1) + ","
	}
	return caminoRecorrido
}

func AlgoritmoFloyd(mAdy [][]float64) string {
	vertices := len(mAdy)
	matrizAdyacencia := mAdy

	caminos := make([][]string, vertices)
	for i := range caminos {
		caminos[i] = make([]string, vertices)
	}

	caminosAuxiliares := make([][]string, vertices)
	for i := range caminosAuxiliares {
		caminosAuxiliares[i] = make([]string, vertices)
	}
	caminoRecorrido := ""
	caminitos := ""

	var i, j, k int
	var tmp1, tmp2, tmp3, tmp4, minimo float64
	if tmp3 == 0 {
	}
	for i = 0; i < vertices; i++ {
		for j = 0; j < vertices; j++ {
			caminos[i][j] = ""
			caminosAuxiliares[i][j] = ""
		}

		for k = 0; k < vertices; k++ {
			for i = 0; i < vertices; i++ {
				for j = 0; j < vertices; j++ {
					tmp1 = matrizAdyacencia[i][j]
					tmp2 = matrizAdyacencia[i][k]
					tmp3 = matrizAdyacencia[k][j]
					tmp4 = tmp2 + tmp4
					minimo = min(tmp1, tmp4)
					if tmp1 != tmp4 {
						if minimo == tmp4 {
							caminoRecorrido = ""
							caminosAuxiliares[i][j] = strconv.Itoa(k)
							caminos[i][j] = caminosR(i, k, caminosAuxiliares, caminoRecorrido)
						}
					}
					matrizAdyacencia[i][j] = minimo
				}
			}
		}

		for i = 0; i < vertices; i++ {
			for j = 0; j < vertices; j++ {
				if matrizAdyacencia[i][j] != 1000000 {
					if i != j {
						if caminos[i][j] == "" {
							caminitos += "De " + strconv.Itoa(i+1) + " A " + strconv.Itoa(j+1) + " Irse por:" + strconv.Itoa(i+1) + "," + strconv.Itoa(j+1) + "\n"
						} else {
							caminitos += "De " + strconv.Itoa(i+1) + " A " + strconv.Itoa(j+1) + " Irse por:" + strconv.Itoa(i+1) + "," + caminos[i][j] + "," + strconv.Itoa(j+1) + "\n"
						}
					}
				}
			}
		}

	}
	return caminitos
}

type path struct {
	value int
	nodes []string
}

type minPath []path

func (h minPath) Len() int           { return len(h) }
func (h minPath) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minPath) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(path))
}

func (h *minPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type heap struct {
	values *minPath
}

func newHeap() *heap {
	return &heap{values: &minPath{}}
}

func (h *heap) push(p path) {
	hp.Push(h.values, p)
}

func (h *heap) pop() path {
	i := hp.Pop(h.values)
	return i.(path)
}

type edge struct {
	node   string
	weight int
}

type Graph struct {
	nodes map[string][]edge
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]edge)}
}

func (g *Graph) AddEdge(origin, destiny string, weight int) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
	g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
}

func (g *Graph) getEdges(node string) []edge {
	return g.nodes[node]
}

type ByWay struct {
	PesoTotal  int
	Estaciones []string
}

func (g *Graph) GetPath(origin, destiny string) ByWay /*(int, []string)*/ {
	h := newHeap()
	h.push(path{value: 0, nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			//return p.value, p.nodes
			return ByWay{p.value, p.nodes}
		}

		for _, e := range g.getEdges(node) {
			if !visited[e.node] {
				// We calculate the total spent so far plus the cost and the path of getting here
				h.push(path{value: p.value + e.weight, nodes: append([]string{}, append(p.nodes, e.node)...)})
			}
		}

		visited[node] = true
	}

	//return 0, nil
	return ByWay{0, nil}
}
