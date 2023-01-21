
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

    var visible_trees int;
    for i, _ := range grid {
        x := i%width;
        y := i/width;
        if is_visible(grid, width, x, y) {
            visible_trees += 1;
        }
    }
    fmt.Println("visible_trees", visible_trees);
}


func is_visible(grid []int, width int, col, row int) bool {
    if row == 0 || row == width - 1 || col == 0 || col == width - 1 {
        return true;
    }
    h := grid[row*width + col];

    var visible bool = true;
    // left
    for i := col - 1; i >= 0; i -= 1 {
        if grid[row*width + i] >= h {
            visible = false;
            break;
        }
    }
    if visible {
        return true;
    }
    visible = true;

    // right
    for i := col + 1; i < width; i += 1 {
        if grid[row*width + i] >= h {
            visible = false;
            break;
        }
    }
    if visible {
        return true;
    }
    visible = true;

    // up
    for i := row - 1; i >= 0; i -= 1 {
        if grid[i*width + col] >= h {
            visible = false;
            break;
        }
    }
    if visible {
        return true;
    }
    visible = true;

    // bottom
    for i := row + 1; i < width; i += 1 {
        if grid[i*width + col] >= h {
            visible = false;
            break;
        }
    }

    return visible;
}
