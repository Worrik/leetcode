// 82.35% 64.71%
func maxCount(m int, n int, ops [][]int) int {
    if len(ops) == 0 {
        return m * n
    }
    minX := ops[0][0]
    minY := ops[0][1]
    for _, el := range ops[1:] {
        if minX > el[0] && el[0] != 0 {
            minX = el[0]
        }
        if minY > el[1] && el[1] != 0 {
            minY = el[1]
        }
    }
    return minX * minY
}
