package main

import (
	"container/list"
	"fmt"
)

var maze, visited, from [][]int
var n, m int
var path [][]int

var moveDirs [][]int = [][]int{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func main() {

	fmt.Scan(&n, &m)
	maze = make([][]int, n)
	visited = make([][]int, n)
	path = make([][]int, 0)
	for i := 0; i < n; i++ {
		maze[i] = make([]int, m)
		visited[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&maze[i][j])
		}
	}
	if DFS(0, 0) {
		for i := len(path) - 1; i >= 0; i-- {
			fmt.Printf("(%d,%d)\n", path[i][0], path[i][1])
		}
	}
	if BFS(0, 0) {
		// reconstruct the path from the end
		r, c := n-1, m-1
		path = append(path, []int{r, c})
		for from[r][c] >= 0 {
			fromDir := from[r][c]
			r = r - moveDirs[fromDir][0]
			c = c - moveDirs[fromDir][1]
			path = append(path, []int{r, c})
		}
		for i := len(path) - 1; i >= 0; i-- {
			fmt.Printf("(%d,%d)\n", path[i][0], path[i][1])
		}
	}
}

func withinMaze(r, c int) bool {
	return 0 <= r && r < n && 0 <= c && c < m
}

func DFS(r, c int) bool {
	visited[r][c] = 1
	if r == n-1 && c == m-1 {
		path = append(path, []int{r, c})
		return true
	} else {
		// try 4 move directions
		for _, dir := range moveDirs {
			newr := r + dir[0]
			newc := c + dir[1]
			// the new position is still in the maze
			// the new position is not a wall, that is, maze value is 0
			// the new position is not visisted
			if withinMaze(newr, newc) && maze[newr][newc] == 0 && visited[newr][newc] == 0 {
				if DFS(newr, newc) {
					path = append(path, []int{r, c})
					return true
				}
			}
		}
		return false
	}
}

func BFS(r, c int) bool {

	// record the direction where the previous position goes to this position
	// the value is the index in the array moveDirs
	from = make([][]int, n)
	for i := 0; i < n; i++ {
		from[i] = make([]int, m)
		for j := 0; j < m; j++ {
			from[i][j] = -1
		}
	}
	queue := list.New()
	queue.PushBack([]int{r, c})
	for queue.Len() > 0 {
		this := queue.Remove(queue.Front()).([]int)
		r = this[0]
		c = this[1]
		visited[r][c] = 1
		if r == n-1 && c == m-1 {
			return true
		}
		for i, dir := range moveDirs {
			newr := r + dir[0]
			newc := c + dir[1]
			if withinMaze(newr, newc) && maze[newr][newc] == 0 && visited[newr][newc] == 0 {
				visited[newr][newc] = 1 // mark visited, so don't repeat the work
				from[newr][newc] = i
				queue.PushBack([]int{newr, newc})
			}
		}
	}
	return false
}
