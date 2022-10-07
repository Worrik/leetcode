// 88.18% 100%
func findRestaurant(list1 []string, list2 []string) []string {
    var res []string
    min := len(list1) * len(list2)

    for i, x := range list1 {
        if i > min {
            return res
        }
        for j, y := range list2 {
            if x == y {
                if i + j < min || min == -1 {
                    res = []string{x}
                    min = i + j
                } else if i + j == min {
                    res = append(res, x)
                }
                break
            }
        }
    }
    return res
}
