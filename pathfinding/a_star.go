// refer to https://www.redblobgames.com/pathfinding/a-star/introduction.html

package pathfinding

import (
	"container/heap"
	"github.com/molodofo/algorithm-go/datastructure"
)

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
