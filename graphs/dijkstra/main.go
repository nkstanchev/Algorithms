package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	val    interface{}
	weight int
}

type PQ []Pair

func (self PQ) Len() int { return len(self) }

func (self PQ) Less(i, j int) bool {
	return self[i].weight < self[j].weight
}
func (self PQ) Swap(i, j int) { self[i], self[j] = self[j], self[i] }

func (self *PQ) Push(x interface{}) { *self = append(*self, x.(Pair)) }

func (self *PQ) Pop() (popped interface{}) {
	popped = (*self)[len(*self)-1]
	*self = (*self)[:len(*self)-1]
	return
}

func setup() [][]interface{} {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	graph := make([][]interface{}, n)
	var u, v, w int
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d", &u, &v, &w)
		graph[u] = append(graph[u], Pair{val: v, weight: w})
	}
	return graph
}

func dijkstra(graph [][]interface{}) {
	start := 0

	pq := &PQ{}
	heap.Init(pq)
	heap.Push(pq, Pair{start, 0})

	dist := make([]int, len(graph))
	//visited[start] = true
	for k, _ := range dist {
		dist[k] = 100
	}
	dist[start] = 0

	for pq.Len() != 0 {
		u := pq.Pop().(Pair)
		val := u.val.(int)
		neighbours := graph[val]
		//fmt.Printf("%d \n", val)
		for _, v := range neighbours {
			neighbourPair := v.(Pair)
			neighbour := neighbourPair.val.(int)

			if dist[val]+neighbourPair.weight < dist[neighbour] {
				// update dist
				dist[neighbour] = dist[val] + neighbourPair.weight
				// push to q
				pq.Push(Pair{val: neighbour, weight: dist[neighbour]})
			}
		}
	}

	for k, v := range dist {
		fmt.Printf("%d (%d)\n", k, v)
	}
}

func main() {
	graph := setup()
	dijkstra(graph)
}
