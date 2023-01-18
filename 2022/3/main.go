
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
    for i := 0; i < len(lines) && len(lines) - i > 3; i += 3 {
        line1 := lines[i]
        line2 := lines[i + 1]
        line3 := lines[i + 2]
        var common rune
        loop:
        for _, l1 := range line1 {
            for _, l2 := range line2 {
                for _, l3 := range line3 {
                    if l1 == l2 && l2 == l3 {
                        common = l1
                        break loop
                    }
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
