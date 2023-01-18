
package main

import (
    "os"
    "log"
    "fmt"
    "strings"
)

func main() {
    data, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    var sum int
    lines := strings.Split(string(data), "\n")
    for _, line := range lines {
        if line == "" {
            continue;
        }
        pivot := len(line)/2
        left  := line[:pivot]
        right := line[pivot:]
        var common rune
        loop:
        for _, l := range left {
            for _, r := range right {
                if l == r {
                    common = l
                    break loop
                }
            }
        }
        sum += gamatria(common)
    }
    fmt.Println("sum:", sum)
}

func gamatria(r rune) int {
    var result int
    if r >= 'A' && r <= 'Z' {
        result = 26
        r += 'a' - 'A'
    }
    result += int(r - 'a' + 1)
    return result
}
