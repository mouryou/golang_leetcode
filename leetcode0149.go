/**
 * Definition for a point.
 * type Point struct {
 *     X int
 *     Y int
 * }
 */
func maxPoints(points []Point) int {
    if len(points) <= 1 {
        return len(points)
    }
    count := map[[3]int]map[int]bool {}
    res := 0
    for i := 0; i < len(points) - 1; i++ {
        for j := i + 1; j < len(points); j++ {
            k := getLine(points[i], points[j])
            if _, ok := count[k]; !ok {
                count[k] = make(map[int]bool)
            }
            count[k][i] = true
            count[k][j] = true
        }
    }
    for _, v := range count {
        if len(v) > res {
            res = len(v)
        }
    }
    return res
}

func getGCD(a, b int) int {
    if a < 0 {
        a = -a
    }
    if b < 0 {
        b = -b
    }
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

func getLine(p1, p2 Point) [3]int {
    a, b, c := p2.Y - p1.Y, p1.X - p2.X, p2.X * p1.Y - p1.X * p2.Y
    gcd := getGCD(a, getGCD(b, c))
    if gcd == 0 {
        return [3]int {-1, p1.X, p1.Y}
    }
    if a < 0 {
        gcd = -gcd
    }
    return [3]int {a / gcd, b / gcd, c / gcd}
}