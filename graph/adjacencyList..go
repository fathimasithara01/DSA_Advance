package main

import "fmt"

type Graph struct {
	adjacencyList map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[string][]string),
	}
}

func (g *Graph) AddEdge(u, v string) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	g.adjacencyList[v] = append(g.adjacencyList[v], u)
}

func (g *Graph) Print() {
	fmt.Println("Adjacency List:")
	for node, neighbors := range g.adjacencyList {
		fmt.Printf("%s -> %v\n", node, neighbors)
	}
}

func main() {
	graph := NewGraph()

	//Adding Edges
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")

	//print graph
	graph.Print()
}
