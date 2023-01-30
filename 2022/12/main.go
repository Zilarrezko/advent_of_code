
package main;

import (
    "os"
    "log"
    "strings"
    "fmt"
)


var (
    printing bool;
)


func main() {
    printing = get_arg("-p");
    data, err := os.ReadFile("input.txt");
    if err != nil {
        log.Fatal(err);
    }
    rows := strings.Split(string(data), "\n");
    height := len(rows);
    width := len(rows[0]);
    var grid []byte = make([]byte, height*width);
    for y, row := range rows {
        for x, h := range row {
            grid[y*width + x] = byte(h);
        }
    }
    fmt.Print("\x1b[?25l");
    path := pathfind(grid, width);
    print_path(grid, width, path);
    fmt.Println(len(path) - 1, "steps");
    fmt.Print("\x1b[?25h");
}


type Node struct {
    pos int;
    parent int;
    g_cost int;
    f_cost int;
}


func print_path(grid []byte, width int, path []int) {
    for i, h := range grid {
        x := i%width;
        if x == 0 {
            fmt.Println();
        }
        var f bool;
        for _, n := range path {
            if i == n {
                f = true;
                break;
            }
        }
        if !f {
            fmt.Print("\x1b[0;39m", string(h));
        } else {
            fmt.Print("\x1b[0;31m", string(h));
        }
    }
    fmt.Println();
}


func pathfind(grid []byte, width int) []int {
    var open []Node; // Could make a min heap
    var visited map[int]Node = make(map[int]Node, 0);

    var start int = -1;
    var end int = -1;
    for i, x := range grid {
        if x == 'S' {
            start = i;
        }
        if x == 'E' {
            end = i;
        }
        if end > -1 && start > -1 {
            break;
        }
    }

    var node Node;
    node.pos = start;
    node.parent = -1;
    open = append(open, node);
    for len(open) > 0 {
        if printing && len(visited)%4 == 0 {
            print_visited_open_grid(grid, width, open, visited);
        }

        var closest int;
        for i, n := range open {
            if n.f_cost < open[closest].f_cost {
                closest = i;
            }
        }
        node = open[closest];
        if node.pos == end {
            var path []int;
            path = append(path, node.pos);
            var c int;
            for node.parent >= 0 {
                node = visited[node.parent];
                path = append(path, node.pos);
                c += 1;
            }
            return path;
        }
        open = remove(open, closest);
        visited[node.pos] = node;

        for y := -1; y <= 1; y += 1 {
            ny := node.pos/width;
            ny += y;
            if ny < 0 || ny > (len(grid) - 1)/width {
                continue;
            }
            for x := -1; x <= 1; x += 1 {
                if (y == 0 && x == 0) || abs(x) == abs(y) {
                    continue;
                }
                nx := node.pos%width;
                nx += x;
                if nx < 0 || nx >= width {
                    continue;
                }
                var neighbor Node;
                neighbor.pos = ny*width + nx;
                neighbor.g_cost = node.f_cost + 10; // cardinal tile distance
                neighbor.f_cost = neighbor.g_cost + integer_distance(neighbor.pos, end, width);

                if exists(open, neighbor.pos) {
                    var idx int;
                    for i, n := range open {
                        if n.pos == neighbor.pos {
                            idx = i;
                            break;
                        }
                    }
                    cmpr := open[idx];
                    if cmpr.f_cost > neighbor.f_cost {
                        neighbor.parent = node.pos;
                        open[idx] = neighbor;
                    }
                } else {
                    _, is_visited := visited[neighbor.pos];
                    var adding bool = walkable(grid[neighbor.pos], grid[node.pos]) && !is_visited;
                    if adding {
                        neighbor.parent = node.pos;
                        open = append(open, neighbor);
                    }
                }
            }
        }
    }
    return []int{};
}


func print_visited_open_grid(grid []byte, width int, open []Node, visited map[int]Node) {
    fmt.Print("\x1b[42F");
    for i, h := range grid {
        x := i%width;
        if x == 0 {
            fmt.Println();
        }
        _, is_visited := visited[i];
        if is_visited {
            fmt.Print("\x1b[0;31m", string(h));
        } else {
            var f bool;
            for _, m := range open {
                if m.pos == i {
                    f = true;
                    break;
                }
            }
            if f {
                fmt.Print("\x1b[0;33m", string(h));
            } else {
                fmt.Print("\x1b[0;39m", string(h));
            }
        }
    }
    fmt.Println();
}


func integer_distance(a, b int, width int) int {
    ax := a%width;
    ay := a/width;
    bx := b%width;
    by := b/width;
    var d int;
    dx := abs(bx - ax);
    dy := abs(by - ay);
    if dx > dy {
        d = 14*dy + 10*(dx - dy);
    } else {
        d = 14*dx + 10*(dy - dx);
    }
    return d;
}


func walkable(to, from byte) bool {
    gradient := gematria(to) - gematria(from);
    return gradient <= 1;
}


func abs(n int) int {
    if n < 0 {
        return -n;
    }
    return n;
}


func gematria(n byte) int {
    if n == 'S' {
        n = 'a'
    }
    if n == 'E' {
        n = 'z'
    }
    return int(n - 'a');
}


func exists(arr []Node, pos int) bool {
    for _, n := range arr {
        if n.pos == pos {
            return true;
        }
    }
    return false;
}


// unordered
func remove(s []Node, i int) []Node {
    s[i] = s[len(s) - 1];
    return s[:len(s) - 1];
}


func get_arg(s string) bool {
    for _, arg := range os.Args[1:] {
        if arg == s {
            return true;
        }
    }
    return false;
}
