package main

import "fmt"

type graph struct {
	adj map[string][]string
}

// Create a new Graph
func newGraph() *graph {
	return &graph{
		adj: make(map[string][]string),
	}
}

// Add an undirected edge between two nodes
func (g *graph) AddEdge(u, v string) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *graph) DFS(start string, visited map[string]bool) {
	if visited[start] {
		return
	}
	fmt.Print(start, " ")
	visited[start] = true

	for _, neighbor := range g.adj[start] {
		if !visited[neighbor] {
			g.DFS(neighbor, visited)
		}
	}
}

func (g *graph) BFS(start string) {
	visited := make(map[string]bool)
	queue := []string{start}

	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue :=queue[1:]

		fmt.Println(node,"")

		for _,neightbor := range g.adj[node]{
			if !visited[neightbor]=true
			queue =append(queue, neightbor)
		}
	}
}

func (g *graph) Print(){
	fmt.Println("\n Adjacency List:")
	for node , neighbors := range g.adj{
		fmt.Printf("%s -> %v\n",node,neighbors)
	}
}

func main(){
	g := newGraph()

	// add some edges
	g.AddEdge("A","B")
	g.AddEdge("A","C")
	g.AddEdge("B","D")
	g.AddEdge("C","E")
	g.AddEdge("D","E")
	g.AddEdge("E","F")

	g.Print()

	fmt.Print("\nDFS Traversal starting from A:")
	visited := make(map[string]bool)
	g.DFS("A",visited)

	fmt.Print("\nBFS Traversal starting from A:")
	g.DFS("A")
}