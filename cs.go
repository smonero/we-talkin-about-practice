
func digitCount(in int) int {
    count := 0
    for {
        if in == 0 {
            break
        }
        count++
        in /= 10
    }
    
    return count
}


func solution(a []int) int64 {

res := int64(0)
num2DigitCount := make(map[int]int)

for _, x := range a {
    count := digitCount(x)
    num2DigitCount[count]++
}

// 2 -> 1
// 1 -> 1
// 1000
// 10
// 200
// 20
// (10 * 100) + 10
// (10 * 10) + 2
// (2 * 100) + 10
// (2 * 10) + 2
// 100(10 + 2)
    for _, x := range a {
        for k, v := range num2DigitCount {
            factor := int64(1)
            for i := 0; i < k; i++ {
                factor *= 10
            }
            //fmt.Println(fmt.Sprintf("Adding %d * %d * %d", x, v, factor))
            res += int64(x) * (int64(v) * factor)
        }
        res += int64(x) * int64(len(a)) // 20 + 4
    }

return res
}

