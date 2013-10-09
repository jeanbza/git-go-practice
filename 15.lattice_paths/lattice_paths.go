package main

import "fmt"

type loc struct {
    x, y int
}

func main() {
    myLoc := loc{0, 0}
    total := traverse(myLoc, 3, 0)
    fmt.Println(total)
}

func traverse(curLoc loc, maxSize int, total int) (int) {
    if (curLoc.x == maxSize-1 && curLoc.y == maxSize-1) {
        return 1;
    }

    fmt.Println(curLoc)

    if (curLoc.x < maxSize-1) {
        total += traverse(loc{curLoc.x+1, curLoc.y}, maxSize, total)
    }

    if (curLoc.y < maxSize-1) {
        total += traverse(loc{curLoc.x, curLoc.y+1}, maxSize, total)
    }

    return total
}