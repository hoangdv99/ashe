func canVisitAllRooms(rooms [][]int) bool {
    visited := []int{}

    var dfs func(int)

    dfs = func(roomNum int) {
        if slices.Contains(visited, roomNum) {
            return
        }
        visited = append(visited, roomNum)
        for _, key := range rooms[roomNum] {
            dfs(key)
        }
    }
    dfs(0)
    
    return len(visited) == len(rooms)
}