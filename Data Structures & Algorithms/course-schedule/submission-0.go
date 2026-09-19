func canFinish(numCourses int, prerequisites [][]int) bool {
    
	adjList := make(map[int][]int)
	visited := make(map[int]bool)

	for _, pair := range prerequisites {
		course, reqCourse := pair[0], pair[1]
		adjList[course] = append(adjList[course], reqCourse)
	}
	
	var dfs func(int,map[int]bool) bool

	dfs = func(numCourse int, path map[int]bool) bool {
		
		if path[numCourse] {
			return false
		}

		if visited[numCourse] {
			return true
		}

		path[numCourse] = true

		for _, neighbor := range adjList[numCourse] {
			if !dfs(neighbor, path) {
				return false
			}
		}

		delete(path, numCourse)
		visited[numCourse] = true

		return true
	}

	for i := 0; i < numCourses; i ++ {
		if !dfs(i, make(map[int]bool)) {
			return false
		}
	}

	return true
}
