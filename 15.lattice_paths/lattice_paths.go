package main

import "fmt"

type loc struct {
    x, y int
}

func main() {
    myLoc := loc{0, 0}
    total := 0
    traverse(myLoc, 20, &total)
    fmt.Println(total)
}

func traverse(curLoc loc, maxSize int, total *int) {
    fmt.Println(curLoc)

    if (curLoc.x == maxSize && curLoc.y == maxSize) {
        *total += 1
    }

    if (curLoc.x < maxSize) {
        traverse(loc{curLoc.x+1, curLoc.y}, maxSize, total)
    }

    if (curLoc.y < maxSize) {
        traverse(loc{curLoc.x, curLoc.y+1}, maxSize, total)
    }
}