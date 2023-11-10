// refer to https://www.redblobgames.com/pathfinding/a-star/introduction.html

package pathfinding

import "github.com/molodofo/algorithm-go/datastructure"

func Bfs(graph Graph) []any {
	frontier := datastructure.Queue{}
	cameFrom := make(map[any]any)

	frontier.Push(graph.GetStart())
	cameFrom[graph.GetStart()] = nil

	for !frontier.IsEmpty() {
		current := frontier.Pop()

		if graph.IsEnd(current) {
			return getPath(current, cameFrom)
		}

		for _, next := range graph.Neighbors(current) {
			if _, ok := cameFrom[next]; !ok {
				frontier.Push(next)
				cameFrom[next] = current
			}
		}
	}
	return nil
}
