package main

import "sync"

func writeIntsToChan(ints []int, intChan chan int, wg *sync.WaitGroup) {
    for _, num := range ints {
        if num%2 == 0 {
            intChan <- num
        }
    }
    close(intChan)
    wg.Done()
}

func appendFromChan(intChan chan int, wg *sync.WaitGroup, res *[]int) {
    for num := range intChan {
        *res = append(*res, num)
    }
    wg.Done()
}

func ConcurrentReadAndWrite(ints []int) []int {
    intChan := make(chan int, len(ints))
    var res []int
    var wg sync.WaitGroup

    wg.Add(2)
    go writeIntsToChan(ints, intChan, &wg)
    go appendFromChan(intChan, &wg, &res)
    wg.Wait()

    return res
}