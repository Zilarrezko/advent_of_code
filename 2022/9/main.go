
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
    var knots [10]v2;
    visited[knots[0]] = true;
    lines := strings.Split(string(data), "\n");
    for _, line := range lines {
        face := line[0];
        amnt, _ := strconv.Atoi(line[2:]);
        for d := 0; d < amnt; d += 1 {
            switch face {
                case 'U': knots[0].y += 1;
                case 'D': knots[0].y -= 1;
                case 'L': knots[0].x -= 1;
                case 'R': knots[0].x += 1;
            }
            for i := 1; i < len(knots); i += 1 {
                if chebyshev_distance(knots[i], knots[i - 1]) > 1 {
                    dx := knots[i - 1].x - knots[i].x;
                    dy := knots[i - 1].y - knots[i].y;
                    knots[i].x += sign(dx);
                    knots[i].y += sign(dy);
                }
            }
            last_knot := knots[len(knots) - 1];
            visited[last_knot] = true;
        }
    }
    fmt.Println(len(visited));
}


func sign(x int) int {
    if x < 0 {
        return -1;
    }
    if x > 0 {
        return 1;
    }
    return 0;
}


func chebyshev_distance(a, b v2) int {
    return max(abs(b.x - a.x), abs(b.y - a.y));
}


func abs(x int) int {
    if x < 0 {
        return -x;
    }
    return x;
}


func max(a, b int) int {
    if a > b {
        return a;
    }
    return b;
}
