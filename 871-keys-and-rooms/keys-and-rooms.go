func canVisitAllRooms(rooms [][]int) bool {
    visited := make(map[int]bool)

    var dfs func(int)

    dfs = func(roomNum int) {
        if visited[roomNum] {
            return
        }
        visited[roomNum] = true
        for _, key := range rooms[roomNum] {
            dfs(key)
        }
    }
    dfs(0)

    return len(visited) == len(rooms)
}