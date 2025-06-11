func findCircleNum(isConnected [][]int) int {
    visited := make(map[int]bool)
    var dfs func(int)
    n := len(isConnected)

    dfs = func(city int) {
        for neighbor := range n {
            if !visited[neighbor] && isConnected[city][neighbor] == 1 {
                visited[city] = true
                dfs(neighbor)
            }
        }
    }

    provinceCount := 0
    for i := range n {
        if !visited[i] {
            visited[i] = true
            dfs(i)
            provinceCount++
        }
    }

    return provinceCount
}