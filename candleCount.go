package main

func PlatesBetweenCandles(s string, queries [][]int) []int {
    result := make([]int, len(queries))

    for i, query := range queries {
        left, right := query[0], query[1]
        count := 0
        tempCount := 0
        betweenPipes := false

        for j := left; j <= right; j++ {
            if s[j] == '|' {
                if betweenPipes {
                    count += tempCount
                    tempCount = 0
                } else {
                    betweenPipes = true
                }
            } else if betweenPipes && s[j] == '*' {
                tempCount++
            }
        }

        result[i] = count
    }

    return result
}





