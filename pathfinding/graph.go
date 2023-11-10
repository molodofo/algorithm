package pathfinding

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
