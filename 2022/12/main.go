
package main;

import (
    "os"
    "log"
    "strings"
    "fmt"
    "math"
)


var (
    printing bool;
    optimize bool;
)


func main() {
    printing = get_arg("-p");
    optimize = get_arg("-o");
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
            if h == 'S' {
                grid[y*width + x] = byte('a');
            } else {
                grid[y*width + x] = byte(h);
            }
        }
    }
    fmt.Print("\x1b[?25l");
    if printing {
        print_path(grid, width, []int{});
    }
    var closest []int;
    var close_dist int = math.MaxInt64;
    var working_grid []byte = make([]byte, len(grid)); //
    copy(working_grid, grid);
    for i := 0; i < len(working_grid); i += 1 {
        x := working_grid[i];
        if x == 'a' {
            grid[i] = 'S';
            working_grid[i] = 'S';
            path := pathfind(grid, width);
            if printing {
                fmt.Print("\x1b[42F");
                print_path(working_grid, width, path);
            }
            dist := len(path) - 1;
            if dist > 0 && dist < close_dist {
                closest = path;
                close_dist = dist;
            }
            grid[i] = 'a';
            working_grid[i] = 'a';

            // Never reached the end, so flood fill so we don't check anything here again
            if optimize && len(path) == 0 {
                flood(working_grid, width, i);
            }
        }
    }
    if printing {
        fmt.Print("\x1b[42F");
    }
    print_path(working_grid, width, closest);
    fmt.Println(close_dist, "steps");
    fmt.Print("\x1b[?25h");
}


type Node struct {
    pos int;
    parent int;
    gcost int;
    fcost int;
}


// Flood fill 'a' with '!' to denote not to check the area later
func flood(grid []byte, width int, at int) {
    if at < 0 || at >= len(grid) {
        return;
    }
    if grid[at] == 'a' {
        grid[at] = '!';
        x := at%width;
        y := at/width;
        if x > 0 {
            flood(grid, width, at - 1);
        }
        if x < width - 1 {
            flood(grid, width, at + 1);
        }
        if y > 0 {
            flood(grid, width, at - width);
        }
        if y < len(grid)/width - 1 {
            flood(grid, width, at + width);
        }
    }
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
        if printing && len(visited)%64 == 0 {
            print_visited_open_grid(grid, width, open, visited);
        }

        var closest int;
        for i, n := range open {
            if n.fcost < open[closest].fcost {
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
                neighbor.gcost = node.fcost + 10; // cardinal tile distance
                neighbor.fcost = neighbor.gcost + integer_distance(neighbor.pos, end, width);

                _, is_visited := visited[neighbor.pos];
                if !is_visited {
                    if exists(open, neighbor.pos) {
                        var idx int;
                        for i, n := range open {
                            if n.pos == neighbor.pos {
                                idx = i;
                                break;
                            }
                        }
                        cmpr := open[idx];
                        if cmpr.fcost > neighbor.fcost {
                            neighbor.parent = node.pos;
                            open[idx] = neighbor;
                        }
                    } else {
                        var adding bool = walkable(grid[neighbor.pos], grid[node.pos]) && !is_visited;
                        if adding {
                            neighbor.parent = node.pos;
                            open = append(open, neighbor);
                        }
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
