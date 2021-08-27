/*
You have a graph of n nodes labeled from 0 to n - 1. You are given an integer n and a list of edges where edges[i] = [ai, bi] indicates that there is an undirected edge between nodes ai and bi in the graph.
Return true if the edges of the given graph make up a valid tree, and false otherwise.
*/
package main

import "fmt"

func main() {
	fmt.Println(validTree(5, [][]int{{0,1},{0,2},{0,3},{1,4}}))
}

/*
DFS iterative (stack)

For the graph to be a valid tree, it must have exactly n - 1 edges. Any less, and it can't possibly be fully connected.
Any more, and it has to contain cycles. Additionally, if the graph is fully connected and contains exactly n - 1 edges,
it can't possibly contain a cycle, and therefore must be a tree!

Algorithm:
1. Check number of edges.
2. Check if graph fully connected with DFS.

Time Complexity : O(N).
Space Complexity : O(N).
*/
func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	// Create an adjacency list
	adjList := make(map[int][]int)
	for _, e := range edges {
		adjList[e[0]] = append(adjList[e[0]], e[1])
		adjList[e[1]] = append(adjList[e[1]], e[0])
	}
	fmt.Println("Adjacency list:", adjList)

	seen := map[int]bool{
		0:true,
	}
	stack := []int{0}
	fmt.Println(seen, stack)


	for len(stack)>0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, neighbour := range adjList[node] {
			fmt.Println("Node:", node, "neighbour", neighbour)
			if _, ok := seen[neighbour]; ok {
				continue
			}
			seen[neighbour] = true
			stack = append(stack, neighbour)
		}
	}


	return len(seen) == n
}