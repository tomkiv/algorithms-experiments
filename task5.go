package main

import (
  "bufio"
	"flag"
	"fmt"
	"io"
	"os"
	//"sort"
	"strconv"
	"strings"
)

type VertexType uint

type Edge struct {
	weight int64
	vertex *Vertex
}

type Vertex struct {
	index    VertexType
	edges    []*Edge
	visited  bool
	distance int64
}

func (v *Vertex) addEdge(to *Vertex, weight int64) {
	toEdge := new(Edge)
	toEdge.vertex = to
	toEdge.weight = weight
	v.edges = append(v.edges, toEdge)
}

type Graph struct {
	vertexes map[VertexType]*Vertex
}

func newGraph() *Graph {
	g := new(Graph)
	g.vertexes = make(map[VertexType]*Vertex)
	return g
}

func (g *Graph) addEdge(from VertexType, to VertexType, weight int64) {
	vFrom := g.getOrCreateVertex(from)
	vTo := g.getOrCreateVertex(to)
	vFrom.addEdge(vTo, weight)
}

func (g *Graph) getOrCreateVertex(id VertexType) *Vertex {
	if v, exists := g.vertexes[id]; exists {
		return v
	} else {
		v := new(Vertex)
		v.index = id
		v.distance = 1000000
		g.vertexes[id] = v
		return v
	}
	return nil
}

func (g *Graph) reset() {
	for _, vert := range g.vertexes {
		vert.visited = false
	}
}

var leader VertexType
var secondOrder []VertexType

func dfsLoop(g *Graph) {
	var j int
	for j = len(secondOrder) - 1; j >= 0; j-- {
		index := secondOrder[VertexType(j)]
		vert := g.vertexes[index]
		if !vert.visited {
			leader = vert.index
			dfs(g, vert.index)
		}
	}
}

func dfs(g *Graph, index VertexType) {
	var stack []VertexType
	stack = append(stack, index)
	var topIndex VertexType
	for len(stack) > 0 {
		topIndex, stack = stack[len(stack)-1], stack[:len(stack)-1]
		g.vertexes[topIndex].visited = true
		for _, v := range g.vertexes[topIndex].edges {
			if !v.vertex.visited {
				stack = append(stack, v.vertex.index)
			}
		}
	}
}

func dijkstra(g *Graph, start VertexType) {
	var processed []VertexType
	g.vertexes[start].distance = 0
	g.vertexes[start].visited = true
	processed = append(processed, start)
	for {
		v, weight := shortest(g, processed)
		if v == 0 {
			break
		}
		processed = append(processed, v)
		g.vertexes[v].distance = weight
		g.vertexes[v].visited = true
	}
}

func shortest(g *Graph, from []VertexType) (VertexType, int64) {
	var vRes VertexType
	vRes = 0
	var curMin int64
	curMin = 10000000000
	for _, v := range from {
		for _, e := range g.vertexes[v].edges {
			if !e.vertex.visited && g.vertexes[v].distance+e.weight < curMin {
				curMin = g.vertexes[v].distance + e.weight
				vRes = e.vertex.index
			}
		}
	}
	return vRes, curMin
}

func readFile() *Graph {
	g := newGraph()
	fi, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		inp := strings.Split(line, "\t")
		from, err := strconv.ParseUint(inp[0], 10, 0)
		if err != nil {
			break
		}
		inp = inp[1:]
		for _, pair := range inp {
			parsed := strings.Split(strings.Trim(pair, " \n"), ",")
			to, err := strconv.ParseUint(parsed[0], 10, 0)
			if err != nil {
				break
			}
			weight, err := strconv.ParseInt(parsed[1], 10, 0)
			if err != nil {
				break
			}
			g.addEdge(VertexType(from), VertexType(to), weight)
		}
	}
	return g
}

var fileName = flag.String("file", "", "File name")

type Sint []int

func (s Sint) Len() int {
	return len(s)
}
func (s Sint) Less(i, j int) bool {
	return s[i] > s[j]
}
func (s Sint) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	flag.Parse()
	g := readFile()
	dijkstra(g, 1)
	fmt.Println(
		g.vertexes[7].distance, ",",
		g.vertexes[37].distance, ",",
		g.vertexes[59].distance, ",",
		g.vertexes[82].distance, ",",
		g.vertexes[99].distance, ",",
		g.vertexes[115].distance, ",",
		g.vertexes[133].distance, ",",
		g.vertexes[165].distance, ",",
		g.vertexes[188].distance, ",",
		g.vertexes[197].distance, ",",
	)
}
