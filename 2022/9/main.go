
package main;

import (
    "os"
    "log"
    "fmt"
    "strings"
    "strconv"
)

type v2 struct {
    x, y int;
}

func main() {
    data, err := os.ReadFile("input.txt");
    if err != nil {
        log.Fatal(err);
    }
    var visited map[v2]bool = make(map[v2]bool, 0);
    var head v2;
    var tail v2;
    visited[tail] = true
    lines := strings.Split(string(data), "\n");
    for _, line := range lines {
        face := line[0];
        amnt, _ := strconv.Atoi(line[2:]);
        for d := 0; d < amnt; d += 1 {
            from := head;
            switch face {
                case 'U': head.y += 1;
                case 'D': head.y -= 1;
                case 'L': head.x -= 1;
                case 'R': head.x += 1;
            }
            // If the head moves cardinally this should be enough
            if chebychev_distance(tail, head) > 1 {
                tail = from;
                visited[tail] = true;
            }
        }
    }
    fmt.Println(len(visited));
}


// Note: Debug
func print_head_tail(head, tail v2) {
    l := min(head.x, tail.x);
    r := max(head.x, tail.x);
    t := min(head.y, tail.y);
    b := max(head.y, tail.y);
    var p v2;
    for y := t; y <= b; y += 1 {
        for x := l; x <= r; x += 1 {
            p.x = x;
            p.y = y;
            if p == head {
                fmt.Print("H");
            } else if p == tail {
                fmt.Print("T");
            } else {
                fmt.Print(".");
            }
        }
        fmt.Println();
    }
    fmt.Println();
}


func chebychev_distance(a, b v2) int {
    return max(abs(b.x - a.x), abs(b.y - a.y));
}


func abs(x int) int {
    if x < 0 {
        return -x;
    }
    return x;
}


func min(a, b int) int {
    if a < b {
        return a;
    }
    return b;
}


func max(a, b int) int {
    if a > b {
        return a;
    }
    return b;
}
