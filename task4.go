package main

import (
  "bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type VertexType uint

type Vertex struct {
	index       VertexType
	edges       []*Vertex
	revertEdges []*Vertex
	visited     bool
	leader      VertexType
}

func (v *Vertex) addEdge(to *Vertex) {
	v.edges = append(v.edges, to)
	to.revertEdges = append(to.revertEdges, v)
}

type Graph struct {
	vertexes map[VertexType]*Vertex
}

func newGraph() *Graph {
	g := new(Graph)
	g.vertexes = make(map[VertexType]*Vertex)
	return g
}

func (g *Graph) addEdge(from VertexType, to VertexType) {
	vFrom := g.getOrCreateVertex(from)
	vTo := g.getOrCreateVertex(to)
	vFrom.addEdge(vTo)
}

func (g *Graph) getOrCreateVertex(id VertexType) *Vertex {
	if v, exists := g.vertexes[id]; exists {
		return v
	} else {
		v := new(Vertex)
		v.index = id
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

func dfsLoopReverse(g *Graph) {
	length := len(g.vertexes)
	for ind := 1; ind <= length; ind++ {
		vert := g.vertexes[VertexType(ind)]
		if !vert.visited {
			leader = vert.index
			dfsReverse(g, vert.index)
		}
	}
}

func dfsReverse(g *Graph, index VertexType) {
	var stack []VertexType
	stack = append(stack, index)
	var topIndex VertexType
	for len(stack) > 0 {
		topIndex, stack = stack[len(stack)-1], stack[:len(stack)-1]
		g.vertexes[topIndex].visited = true
		for _, v := range g.vertexes[topIndex].revertEdges {
			if !v.visited {
				stack = append(stack, v.index)
			}
		}
		secondOrder = append(secondOrder, topIndex)
	}
}

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
			if !v.visited {
				stack = append(stack, v.index)
			}
		}
		g.vertexes[topIndex].leader = leader
	}
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
		from, err := r.ReadString(' ')
		if err != nil && err != io.EOF {
			panic(err)
		}
		to, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		fromI, err := strconv.ParseUint(strings.Trim(from, " "), 10, 0)
		if err != nil {
			break
		}
		toI, err := strconv.ParseUint(strings.Trim(to, " \n"), 10, 0)
		if err != nil {
			break
		}
		g.addEdge(VertexType(fromI), VertexType(toI))
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
	dfsLoopReverse(g)
	g.reset()
	dfsLoop(g)
	var summing map[VertexType]int
	summing = make(map[VertexType]int)
	for _, v := range g.vertexes {
		if _, exist := summing[v.leader]; !exist {
			summing[v.leader] = 0
		}
		summing[v.leader]++
	}

	var cleaned Sint
	for _, vl := range summing {
		cleaned = append(cleaned, vl)
	}
	sort.Sort(cleaned)
	for len(cleaned) < 5 {
		cleaned = append(cleaned, 0)
	}

	fmt.Println(cleaned[:5])
}
