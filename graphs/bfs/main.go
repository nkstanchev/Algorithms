package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

type Pair struct {
	val    interface{}
	weight int
}

func setup() ([][]interface{}, []bool) {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	graph := make([][]interface{}, n)
	visited := make([]bool, n)
	var u, v, w int
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d", &u, &v, &w)
		graph[u] = append(graph[u], Pair{val: v, weight: w})
	}
	return graph, visited
}

func bfs(graph [][]interface{}, visited []bool) {
	var q deque.Deque
	start := 0
	q.PushFront(Pair{val: 0, weight: 0})
	visited[start] = true
	for q.Len() != 0 {
		u := q.PopFront().(Pair)
		val := u.val.(int)
		weight := u.weight
		fmt.Printf("%d (%d)\n", val, weight)
		neighbours := graph[val]
		for _, v := range neighbours {
			neighbourPair := v.(Pair)
			neighbour := neighbourPair.val.(int)
			if !visited[neighbour] {
				q.PushFront(neighbourPair)
			}
		}
	}
}

func main() {
	graph, visited := setup()
	bfs(graph, visited)
}
