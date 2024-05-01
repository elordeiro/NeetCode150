package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	mapping := make(map[int]*Node)
	mapping[node.Val] = &Node{Val: node.Val, Neighbors: []*Node{}}
	queue := []*Node{node}
	idx := 0

	for idx < len(queue) {
		original := queue[idx]
		idx++
		copy := mapping[original.Val]
		for _, neighbor := range original.Neighbors {
			if _, ok := mapping[neighbor.Val]; !ok {
				mapping[neighbor.Val] = &Node{Val: neighbor.Val, Neighbors: []*Node{}}
				queue = append(queue, neighbor)
			}
			copy.Neighbors = append(copy.Neighbors, mapping[neighbor.Val])
		}
	}
	return mapping[node.Val]
}

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}
	x := cloneGraph(node1)
	_ = x
}
