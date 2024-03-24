package graphDfs

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

type Graph struct {
	Nodes map[int][]int
}

var indexNode map[int][]int
var indexNodeC map[int][]int

func (g *Graph) DFS(start int, target int, visited map[int]bool, outstr *string, path *[]int) bool {
	if start == target {
		*path = append(*path, target)
		for i, n := range *path {
			if g.Nodes[n] != nil {
				if i == 0 {
					*outstr += fmt.Sprintln("Full path: ")
				}
				*outstr += fmt.Sprintf("%v, ", n)
			}
		}
		return true
	}

	visited[start] = true
	*path = append(*path, start)
	*outstr += fmt.Sprintf("%v: %v-> *", start, g.Nodes[start])

	for _, neighbor := range g.Nodes[start] {
		if !visited[neighbor] {
			if g.DFS(neighbor, target, visited, outstr, path) {
				return true
			}
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func generateSVG(graph *Graph, path []int, generate bool, strOut string) string {
	svgTemplate := ""
	svgTemplate += fmt.Sprintln(`<defs><marker id="arrowhead" markerWidth="10" markerHeight="7" refX="30" refY="3.5" orient="auto"><polygon points="0 0, 10 3.5, 0 7" /></marker></defs>`)

	centerX := 250
	centerY := 250
	radius := 200
	radiusDop := radius + 30

	totalNodesX := len(graph.Nodes)
	index := 1

	if generate {
		indexNode = make(map[int][]int)
		indexNodeC = make(map[int][]int)
		for i := range graph.Nodes {
			childAngle := 2 * math.Pi * float64(index) / float64(totalNodesX)
			childX := centerX + int(float64(radius)*math.Cos(childAngle))
			childY := centerY + int(float64(radius)*math.Sin(childAngle))
			indexNode[i] = []int{childX, childY}

			childX = centerX + int(float64(radiusDop)*math.Cos(childAngle))
			childY = centerY + int(float64(radiusDop)*math.Sin(childAngle))
			indexNodeC[i] = []int{childX, childY}
			index++
		}
	}

	if path != nil {
		pathLen := len(path)
		if pathLen == 1 {
			if path[0] <= totalNodesX {
				svgTemplate += fmt.Sprintf(`<circle cx="%d" cy="%d" r="23" fill="Plum"/>`, indexNode[path[0]][0], indexNode[path[0]][1])
			}
		} else {
			for i := 0; i < pathLen-1; i++ {
				svgTemplate += fmt.Sprintf(`<circle cx="%d" cy="%d" r="23" fill="Plum"/>`, indexNode[path[i]][0], indexNode[path[i]][1])
				svgTemplate += fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="Plum" stroke-width="5"/>`, indexNode[path[i]][0], indexNode[path[i]][1], indexNode[path[i+1]][0], indexNode[path[i+1]][1])
			}
			svgTemplate += fmt.Sprintf(`<circle cx="%d" cy="%d" r="23" fill="Plum"/>`, indexNode[path[pathLen-1]][0], indexNode[path[pathLen-1]][1])
		}

	}

	strArr := strings.Split(strOut, "*")
	for i, s := range strArr {
		svgTemplate += fmt.Sprintf(`<text x="%d" y="%d" fill="black" font-size="16" text-anchor="left" alignment-baseline="middle">%s</text>`, centerX*2+300, 10+i*20, fmt.Sprintf("%v", s))
	}

	for node := range graph.Nodes {
		children, ok := graph.Nodes[node]
		if ok {
			for _, child := range children {
				svgTemplate += fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="1" marker-end="url(#arrowhead)"/>`, indexNode[node][0], indexNode[node][1], indexNode[child][0], indexNode[child][1])
				if node == child {
					svgTemplate += fmt.Sprintf(`<circle cx="%d" cy="%d" r="25" fill="transparent" stroke="black" stroke-width="1"/>`, indexNodeC[node][0], indexNodeC[node][1])
				}
			}
		}
	}

	svgTemplate += fmt.Sprintf(`<text x="%d" y="%d" fill="black" font-size="16" text-anchor="left" alignment-baseline="middle">%s</text>`, centerX*2+50, 10, "Ориентированный граф:")
	for node := range indexNode {
		children, ok := indexNode[node]
		if ok {
			svgTemplate += fmt.Sprintf(`<circle cx="%d" cy="%d" r="20" fill="RoyalBlue" stroke="black" stroke-width="1"/>`, children[0], children[1])
			svgTemplate += fmt.Sprintf(`<text x="%d" y="%d" fill="white" font-size="16" text-anchor="middle" alignment-baseline="middle">%d</text>`, children[0], children[1], node)
			svgTemplate += fmt.Sprintf(`<text x="%d" y="%d" fill="black" font-size="16" text-anchor="left" alignment-baseline="middle">%s</text>`, centerX*2+50, 10+node*20, fmt.Sprintf("%v: %v\n", node, graph.Nodes[node]))
		}
	}

	return svgTemplate
}

func createGraph(numNodes, numEdges int) Graph {
	graph := Graph{Nodes: make(map[int][]int)}

	for i := 1; i <= numNodes; i++ {
		graph.Nodes[i] = make([]int, 0)
	}

	for i := 0; i < numEdges; i++ {
		node1 := rand.Intn(numNodes) + 1
		node2 := rand.Intn(numNodes) + 1

		// Проверяем, что node2 не содержится уже в списке ребер node1
		contains := false
		for _, v := range graph.Nodes[node1] {
			if v == node2 {
				contains = true
				break
			}
		}

		if !contains && node1 != node2 {
			graph.Nodes[node1] = append(graph.Nodes[node1], node2)
		}
	}

	return graph
}

func CreateElementRenderGraphSVG(numElements int) (Graph, string) {

	graph := createGraph(numElements, 15)
	output := generateSVG(&graph, nil, true, "")
	return graph, output
}

func FindElementRenderGraphSVG(graph Graph, startNode int, targetNode int) string {

	visited := make(map[int]bool)
	outstr := ""
	var path []int

	result := graph.DFS(startNode, targetNode, visited, &outstr, &path)
	outstr = fmt.Sprintln("Поиск в глубину (DFS): *") + outstr
	if !result {
		path = nil
		if graph.Nodes[startNode] != nil {
			path = append(path, startNode)
		}
	}
	output := generateSVG(&graph, path, false, outstr)

	return output
}
