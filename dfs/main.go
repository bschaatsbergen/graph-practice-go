package main

import (
	"fmt"
)

// Graph represents a graph with nodes and edges
type Graph struct {
	Nodes []*Node
	Edges map[*Node][]*Edge
}

// Node represents a node in the graph with a value
type Node struct {
	Value int
}

// Edge represents a connection between two nodes with a weight
type Edge struct {
	Start  *Node
	End    *Node
	Weight int
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

// AddEdge adds an edge between two nodes
func (g *Graph) AddEdge(start, end *Node, weight int) {
	if g.Edges == nil {
		g.Edges = make(map[*Node][]*Edge)
	}
	edge := &Edge{Start: start, End: end, Weight: weight}
	g.Edges[start] = append(g.Edges[start], edge)
}

func main() {
	graph := Graph{}

	// Creating and adding nodes
	nodes := []*Node{
		{Value: 3}, {Value: 9}, {Value: 15}, {Value: 2}, {Value: 10},
	}
	for _, node := range nodes {
		graph.AddNode(node)
	}

	// Adding edges
	graph.AddEdge(nodes[0], nodes[1], 1) // 3 -> 9
	graph.AddEdge(nodes[1], nodes[2], 1) // 9 -> 15
	graph.AddEdge(nodes[1], nodes[3], 1) // 9 -> 2
	graph.AddEdge(nodes[2], nodes[3], 1) // 15 -> 2
	graph.AddEdge(nodes[2], nodes[4], 1) // 15 -> 10
	graph.AddEdge(nodes[3], nodes[2], 1) // 2 -> 15

	// Perform DFS starting from node 3
	visited := make(map[*Node]bool)
	graph.DFS(nodes[0], visited)
}

// DFS performs a Depth-First Search on the graph from a given node
func (g *Graph) DFS(node *Node, visited map[*Node]bool) {
	if visited[node] {
		return
	}

	// Mark node as visited
	visited[node] = true
	fmt.Println("Visiting node", node.Value)

	// Explore adjacent edges
	for _, edge := range g.Edges[node] {
		if !visited[edge.End] {
			fmt.Println("Traversing edge", edge.Start.Value, "->", edge.End.Value)
			g.DFS(edge.End, visited)
			fmt.Println("Backtracked to node", edge.Start.Value)
		}
	}
}

