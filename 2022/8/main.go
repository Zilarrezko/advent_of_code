
package main

import (
    "fmt"
    "os"
    "log"
    "strings"
)

func main() {
    data, err := os.ReadFile("input.txt");
    if err != nil {
        log.Fatal(err);
    }
    var grid []int;
    var width int;
    lines := strings.Split(string(data), "\n");
    for linenumber, line := range lines {
        for _, x := range line {
            if linenumber < 1 {
                width += 1;
            }
            grid = append(grid, int(x - '0'));
        }
    }

    var max_score int;
    for i, _ := range grid {
        x := i%width;
        y := i/width;
        score := tree_score(grid, width, x, y);
        if score > max_score {
            max_score = score;
        }
    }
    fmt.Println("max scenic score", max_score);
}


func tree_score(grid []int, width int, col, row int) int {
    h := grid[row*width + col];

    var scores [4]int;
    // left
    for i := col - 1; i >= 0; i -= 1 {
        scores[0] += 1;
        if grid[row*width + i] >= h {
            break;
        }
    }
    // right
    for i := col + 1; i < width; i += 1 {
        scores[1] += 1;
        if grid[row*width + i] >= h {
            break;
        }
    }
    // up
    for i := row - 1; i >= 0; i -= 1 {
        scores[2] += 1;
        if grid[i*width + col] >= h {
            break;
        }
    }
    // bottom
    for i := row + 1; i < width; i += 1 {
        scores[3] += 1;
        if grid[i*width + col] >= h {
            break;
        }
    }

    var sum int = 1;
    for _, x := range scores {
        sum *= x;
    }

    return sum;
}
