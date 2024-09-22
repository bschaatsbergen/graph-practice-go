package main

import (
	"testing"
)

func TestNodeAddition(t *testing.T) {
	graph := Graph{}

	node4 := &Node{Value: 4}
	graph.AddNode(node4)

	if len(graph.Nodes) != 1 {
		t.Errorf("Expected 1 node, got %d", len(graph.Nodes))
	}

	if graph.Nodes[0] != node4 {
		t.Errorf("Expected node %v, got %v", node4, graph.Nodes[0])
	}
}

func TestEdgeAddition(t *testing.T) {
	graph := Graph{}

	node1 := &Node{Value: 1}
	graph.AddNode(node1)
	node2 := &Node{Value: 2}
	graph.AddNode(node2)

	graph.AddEdge(node1, node2, 0)

	if len(graph.EdgesMap) != 1 {
		t.Errorf("Expected 1 edge, got %d", len(graph.EdgesMap))
	}

	if len(graph.EdgesMap[node1]) != 1 {
		t.Errorf("Expected 1 edge for node1, got %d", len(graph.EdgesMap[node1]))
	}

	if graph.EdgesMap[node1][0].Start != node1 {
		t.Errorf("Expected edge start to be node1, got %v", graph.EdgesMap[node1][0].Start)
	}

	if graph.EdgesMap[node1][0].End != node2 {
		t.Errorf("Expected edge end to be node2, got %v", graph.EdgesMap[node1][0].End)
	}

	if graph.EdgesMap[node1][0].Weight != 0 {
		t.Errorf("Expected edge weight to be 0, got %d", graph.EdgesMap[node1][0].Weight)
	}
}
