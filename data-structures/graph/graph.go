package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"sync"
)

// ID is unique identifier
type ID interface {
	// String returns the string ID
	String() string
}

type StringID string

func (s StringID) String() string {
	return string(s)
}

// Node is vertex. The ID must be unique within the graph
type Node interface {
	ID() ID
	String() string
}

type node struct {
	id string
}

func (n *node) ID() ID {
	return StringID(n.id)
}

func (n *node) String() string {
	return n.id
}

func NewNode(id string) Node {
	return &node{id: id}
}

// Edge connects between two Nodes
type Edge interface {
	Source() Node
	Target() Node
	Weight() float64
	String() string
}

type edge struct {
	src Node
	tgt Node
	wgt float64
}

func NewEdge(src, tgt Node, wgt float64) Edge {
	return &edge{
		src: src,
		tgt: tgt,
		wgt: wgt,
	}
}

func (e *edge) Source() Node {
	return e.src
}

func (e *edge) Target() Node {
	return e.tgt
}

func (e *edge) Weight() float64 {
	return e.wgt
}

func (e *edge) String() string {
	return fmt.Sprintf("%s -- %.3f -→ %s\n", e.src, e.wgt, e.tgt)
}

// Graph describes the methods of graph operations.
// It assumes that the identifier of a Node is unique.
// And weight values is float64
type Graph interface {
	// GetNodeCount returns the total number of nodes in the graph.
	GetNodeCount() int

	// GetNode finds the Node
	GetNode(id ID) (Node, error)

	// GetNodes returns a map from node ID to empty struct value.
	// Graph dose not allow duplicate node ID or name.
	GetNodes() map[ID]Node

	// AddNode adds a node to a graph, and returns false
	// if the node already existed in the graph.
	AddNode(nd Node) bool

	// DeleteNode deletes a node from a graph.
	// It returns true if it got deleted.
	// And false if it didn't get deleted.
	DeleteNode(id ID) bool

	// AddEdge adds an edge from id1 to id2 with the weight.
	// It returns error if a node does not exist.
	AddEdge(id1, id2 ID, weight float64) error

	// ReplaceEdge replaces an edge from id1 to id2 with the weight.
	ReplaceEdge(id1, id2 ID, weight float64) error

	// DeleteEdge deletes an edge from id1 to id2.
	DeleteEdge(id1, id2 ID) error

	// GetWeight returns the weight from id1 to id2.
	GetWeight(id1, id2 ID) (float64, error)

	// GetSources returns the map of child Nodes.
	// (Nodes that come towards the argument vertex.)
	GetSources(id ID) (map[ID]Node, error)

	// GetTargets returns the map of child Nodes.
	// (Nodes that go out of the argument vertex.)
	GetTargets(id ID) (map[ID]Node, error)

	// String describes the Graph.
	String() string

	// GetAllEdges returns the all edges
	GetAllEdges() []Edge

	GetNodeEdges(id1 ID) []Edge
}

// graph is an internal default graph type that
// implements all methods in Graph interface.
type graph struct {
	mu sync.RWMutex
	// idToNodes stores all nodes.
	idToNodes map[ID]Node

	// nodeToSources maps a Node identifier to sources(parents) with edge weight.
	nodeToSources map[ID]map[ID]float64

	// nodeToTargets maps a Node identifier to targets(children) with edge weight.
	nodeToTargets map[ID]map[ID]float64

	// edges stores all edges
	//edges []Edge
}

func (g *graph) GetNodeCount() int {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return len(g.idToNodes)
}

func (g *graph) GetNode(id ID) (Node, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if !g.existNodeLocked(id) {
		return nil, fmt.Errorf("s% does not exist in the graph.", id)
	}

	return g.idToNodes[id], nil
}

func (g *graph) GetNodes() map[ID]Node {
	return g.idToNodes
}

func (g *graph) existNodeLocked(id ID) bool {
	_, ok := g.idToNodes[id]
	return ok
}

func (g *graph) AddNode(nd Node) bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.existNodeLocked(nd.ID()) {
		return false
	}

	g.idToNodes[nd.ID()] = nd
	return true
}

func (g *graph) DeleteNode(id ID) bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.existNodeLocked(id) {
		return false
	}

	delete(g.idToNodes, id)
	// TODO delete edges

	return true
}

// this is a bug : not compare id1 == id2
func (g *graph) AddEdge(id1, id2 ID, weight float64) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.existNodeLocked(id1) {
		return fmt.Errorf("s% does not exist in the graph.", id1)
	}

	if !g.existNodeLocked(id2) {
		return fmt.Errorf("s% does not exist in the graph.", id2)
	}

	if _, ok := g.nodeToTargets[id1]; ok {
		if v, ok2 := g.nodeToTargets[id1][id2]; ok2 {
			g.nodeToTargets[id1][id2] = v + weight
		} else {
			g.nodeToTargets[id1][id2] = weight
		}
	} else { // id1,id2 first edge
		g.nodeToTargets[id1] = make(map[ID]float64)
		g.nodeToTargets[id1][id2] = weight
	}

	if _, ok := g.nodeToSources[id2]; ok {
		if v, ok2 := g.nodeToSources[id1][id2]; ok2 {
			g.nodeToSources[id2][id1] = v + weight
		} else {
			g.nodeToSources[id2][id1] = weight
		}
	} else { // id1,id2 first edge
		g.nodeToSources[id2] = make(map[ID]float64)
		g.nodeToSources[id2][id1] = weight
	}

	return nil
}

func (g *graph) ReplaceEdge(id1, id2 ID, weight float64) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.existNodeLocked(id1) {
		return fmt.Errorf("s% does not exist in the graph.", id1)
	}

	if !g.existNodeLocked(id2) {
		return fmt.Errorf("s% does not exist in the graph.", id2)
	}

	if _, ok := g.nodeToTargets[id1]; ok {
		g.nodeToTargets[id1][id2] = weight
	} else { // id1,id2 first edge
		g.nodeToTargets[id1] = make(map[ID]float64)
		g.nodeToTargets[id1][id2] = weight
	}

	if _, ok := g.nodeToSources[id2]; ok {
		g.nodeToSources[id2][id1] = weight
	} else { // id1,id2 first edge
		g.nodeToSources[id2] = make(map[ID]float64)
		g.nodeToSources[id2][id1] = weight
	}

	return nil
}

func (g *graph) GetAllEdges() []Edge {
	allEdges := make([]Edge, 0, g.GetNodeCount())
	foundEdge := make(map[string]bool)
	for id1, nd1 := range g.GetNodes() {
		tm, err := g.GetTargets(id1)
		if err != nil {
			continue
		}
		for id2, nd2 := range tm {
			weight, err := g.GetWeight(id1, id2)
			if err != nil {
				continue
			}
			edge := NewEdge(nd1, nd2, weight)
			if _, ok := foundEdge[edge.String()]; !ok {
				allEdges = append(allEdges, edge)
				foundEdge[edge.String()] = true
			}
		}

		sm, err := g.GetSources(id1)
		if err != nil {
			continue
		}
		for id3, nd3 := range sm {
			weight, err := g.GetWeight(id3, id1)
			if err != nil {
				continue
			}
			edge := NewEdge(nd3, nd1, weight)
			if _, ok := foundEdge[edge.String()]; !ok {
				allEdges = append(allEdges, edge)
				foundEdge[edge.String()] = true
			}
		}
	}

	return allEdges
}

func (g *graph) GetNodeEdges(id1 ID) []Edge {
	allEdges := make([]Edge, 0, g.GetNodeCount())
	foundEdge := make(map[string]bool)
	nd1 := g.idToNodes[id1]

	tm, err := g.GetTargets(id1)
	if err == nil {
		for id2, nd2 := range tm {
			weight, err := g.GetWeight(id1, id2)
			if err != nil {
				continue
			}
			edge := NewEdge(nd1, nd2, weight)
			if _, ok := foundEdge[edge.String()]; !ok {
				allEdges = append(allEdges, edge)
				foundEdge[edge.String()] = true
			}
		}
	}

	sm, err := g.GetSources(id1)
	if err == nil {
		for id3, nd3 := range sm {
			weight, err := g.GetWeight(id3, id1)
			if err != nil {
				continue
			}
			edge := NewEdge(nd3, nd1, weight)
			if _, ok := foundEdge[edge.String()]; !ok {
				allEdges = append(allEdges, edge)
				foundEdge[edge.String()] = true
			}
		}
	}

	return allEdges
}

func (g *graph) DeleteEdge(id1, id2 ID) error {
	return nil
}

func (g *graph) GetWeight(id1, id2 ID) (float64, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if !g.existNodeLocked(id1) {
		return 0.0, fmt.Errorf("s% does not exist in the graph.", id1)
	}

	if !g.existNodeLocked(id2) {
		return 0.0, fmt.Errorf("s% does not exist in the graph.", id2)
	}

	if _, ok := g.nodeToTargets[id1]; ok {
		if v, ok := g.nodeToTargets[id1][id2]; ok {
			return v, nil
		}
	}
	return 0.0, fmt.Errorf("there is no edge from %s to %s", id1, id2)
}

func (g *graph) GetSources(id ID) (map[ID]Node, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if !g.existNodeLocked(id) {
		return nil, fmt.Errorf("%s does not exist in the graph.", id)
	}

	// returns the copy of nodeToSources[id]
	rs := make(map[ID]Node)
	if _, ok := g.nodeToSources[id]; ok {
		for n := range g.nodeToSources[id] {
			rs[n] = g.idToNodes[n]
		}
	}
	return rs, nil
}

func (g *graph) GetTargets(id ID) (map[ID]Node, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if !g.existNodeLocked(id) {
		return nil, fmt.Errorf("%s does not exist in the graph.", id)
	}

	rs := make(map[ID]Node)
	if _, ok := g.nodeToTargets[id]; ok {
		for n := range g.nodeToTargets[id] {
			rs[n] = g.idToNodes[n]
		}
	}
	return rs, nil
}

func (g *graph) String() string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	buf := new(bytes.Buffer)
	for id1, nd1 := range g.idToNodes {
		nmap, _ := g.GetTargets(id1)
		for id2, nd2 := range nmap {
			weight, _ := g.GetWeight(id1, id2)
			fmt.Fprintf(buf, "%s -- %.3f -→ %s\n", nd1, weight, nd2)
		}
	}
	return buf.String()
}

// NewGraph returns a new graph
func NewGraph() *graph {
	return &graph{
		idToNodes:     make(map[ID]Node),
		nodeToSources: make(map[ID]map[ID]float64),
		nodeToTargets: make(map[ID]map[ID]float64),
	}
}

// NewGraphFromJSON returns a new Graph from a JSON file.
// Here's the sample JSON data:
//
//	{
//	    "graph_00": {
//	        "S": {
//	            "A": 100,
//	            "B": 14,
//	            "C": 200
//	        },
//	        "A": {
//	            "S": 15,
//	            "B": 5,
//	            "D": 20,
//	            "T": 44
//	        },
//	    },
//	}
//
func NewGraphFromJSON(rd io.Reader, graphID string) (Graph, error) {
	js := make(map[string]map[string]map[string]float64)
	dec := json.NewDecoder(rd)
	for {
		if err := dec.Decode(&js); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}
	if _, ok := js[graphID]; !ok {
		return nil, fmt.Errorf("%s does not exist", graphID)
	}
	gmap := js[graphID]

	g := NewGraph()
	for id1, mm := range gmap {
		nd1, err := g.GetNode(StringID(id1))
		if err != nil {
			nd1 = NewNode(id1)
			g.AddNode(nd1)
		}
		for id2, weight := range mm {
			nd2, err := g.GetNode(StringID(id2))
			if err != nil {
				nd2 = NewNode(id2)
				g.AddNode(nd2)
			}
			g.ReplaceEdge(nd1.ID(), nd2.ID(), weight)
		}
	}

	return g, nil
}
