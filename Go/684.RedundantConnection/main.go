package main

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	rank := make([]int, n+1)
	parent := []int{}
	for i := range n + 1 {
		parent = append(parent, i)
	}

	find := func(n int) int {
		par := parent[n]
		for par != parent[par] {
			parent[par] = parent[parent[par]]
			par = parent[par]
		}
		return par
	}

	union := func(src, dest int) bool {
		par1 := find(src)
		par2 := find(dest)

		if par1 == par2 {
			return false
		}

		if rank[par1] > rank[par2] {
			parent[par2] = par1
			rank[par1] += rank[par2]
		} else {
			parent[par1] = par2
			rank[par2] += rank[par1]
		}
		return true
	}
	for _, edge := range edges {
		src, dest := edge[0], edge[1]
		if !union(src, dest) {
			return []int{src, dest}
		}
	}
	return nil
}
