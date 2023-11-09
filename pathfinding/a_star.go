package pathfinding

import (
	"container/heap"
	"github.com/molodofo/algorithm-go/datastructure"
)

type Graph interface {
	GetStart() any
	IsEnd(position any) bool
	Neighbors(position any) []any
	Cost(position1, position2 any) int
	Heuristic(position any) int
}

func getPath(current any, cameFrom map[any]any) []any {
	path := make([]any, 0)
	path = append(path, current)
	current, ok := cameFrom[current]
	for ; ok && current != nil; current, ok = cameFrom[current] {
		path = append(path, current)
	}
	return path
}

func AStar(graph Graph) []any {
	frontier := make(datastructure.PriorityQueue, 0)
	cameFrom := make(map[any]any)
	soFar := make(map[any]int)

	heap.Push(&frontier, &datastructure.Item{
		Value:    graph.GetStart(),
		Priority: 0,
	})
	cameFrom[graph.GetStart()] = nil
	soFar[graph.GetStart()] = 0

	for frontier.Len() > 0 {
		item := heap.Pop(&frontier).(*datastructure.Item)
		current := item.Value

		if graph.IsEnd(current) {
			return getPath(current, cameFrom)
		}

		for _, next := range graph.Neighbors(current) {
			newCost := soFar[current] + graph.Cost(current, next)
			if cost, ok := soFar[next]; !ok || cost > newCost {
				soFar[next] = newCost
				priority := newCost + graph.Heuristic(next)
				heap.Push(&frontier, &datastructure.Item{
					Value:    next,
					Priority: priority,
				})
				cameFrom[next] = current
			}
		}
	}
	return nil
}
