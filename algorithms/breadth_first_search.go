package algorithms

import (
	"DSA/model"
	"fmt"
)

var vertexQueue *model.Queue

type vertex struct {
	num       int
	neighbors []*vertex
	visited   bool
	steps     int
}

func graph() [8]*vertex {
	var graph [8]*vertex
	v1 := vertex{num: 1}
	v2 := vertex{num: 2}
	v3 := vertex{num: 3}
	v4 := vertex{num: 4}
	v5 := vertex{num: 5}
	v6 := vertex{num: 6}
	v7 := vertex{num: 7}
	v8 := vertex{num: 8}

	v1.neighbors = []*vertex{&v2}
	v2.neighbors = []*vertex{&v1, &v4}
	v3.neighbors = []*vertex{&v4, &v8}
	v4.neighbors = []*vertex{&v2, &v3, &v5, &v7}
	v5.neighbors = []*vertex{&v4, &v6, &v7}
	v6.neighbors = []*vertex{&v5, &v7}
	v7.neighbors = []*vertex{&v4, &v5, &v6, &v8}
	v8.neighbors = []*vertex{&v3, &v7}

	graph[0] = &v1
	graph[1] = &v2
	graph[2] = &v3
	graph[3] = &v4
	graph[4] = &v5
	graph[5] = &v6
	graph[6] = &v7
	graph[7] = &v8

	return graph
}

func Bfs() {
	g := graph()
	from := g[7]
	to := g[0]
	from.steps = 0
	from.visited = true
	vertexQueue = model.NewQueue(from)
	calcBFS(to)
	fmt.Printf("from %+v to %+v", from, to)
}

func calcBFS(to *vertex) {
	for {
		if vertexQueue.IsEmpty() {
			break
		}
		vCurr := vertexQueue.Poll().(*vertex)
		for _, nv := range vCurr.neighbors {
			if !nv.visited {
				nv.visited = true
				nv.steps = vCurr.steps + 1
				vertexQueue.Offer(nv)
				if nv == to {
					return
				}
			}
		}
	}
}
