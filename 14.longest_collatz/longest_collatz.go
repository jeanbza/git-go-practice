package main

import "fmt"

func main() {
    largestChainCount := 0
    largestNum := 0

    for i := 1; i < 1000000; i++ {
        fmt.Println(i)

        chainCountForI := chainCount(i, 1)

        if (chainCountForI > largestChainCount) {
            fmt.Printf("%v: %v\n\n", i, chainCountForI)

            largestNum = i
            largestChainCount = chainCountForI
        }
   }

   fmt.Printf("The largest chain is %v numbers long and is produced by %v \n", largestChainCount, largestNum)
}

func chainCount(someNum int, countChains int) (int) {
    if (someNum == 1) {
        return countChains
    }

    if (someNum % 2 == 0) {
        return chainCount(someNum/2, countChains+1)
    } else {
        return chainCount(3*someNum+1, countChains+1)
    }
}