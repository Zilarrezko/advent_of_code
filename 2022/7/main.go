
package main

import (
    "os"
    "log"
    "fmt"
    "strings"
    "strconv"
    "sort"
)

type Node struct {
    name   string
    is_dir bool
    size   int
    children  []Node
}

const (
    disk_capacity  int = 70000000;
    needed_space   int = 30000000;
)


func main() {
    data, err := os.ReadFile("input.txt");
    if err != nil {
        log.Fatal(err);
    }
    root := parse_filesystem(string(data));
    var dir_sizes []int;
    size := find_candidates(root, &dir_sizes, "");
    empty_space := disk_capacity - size;
    space_to_go := needed_space - empty_space;
    dot_print("total size", 32);
    fmt.Println(size);
    dot_print("empty space", 32);
    fmt.Println(empty_space);
    dot_print("space to go", 32);
    fmt.Println(space_to_go);
    sort.Ints(dir_sizes);
    for i, x := range dir_sizes {
        if x >= space_to_go {
            fmt.Println(i, x);
            break;
        }
    }
}


func find_candidates(root Node, dir_sizes *[]int, path string) int {
    var dir_size int;
    for _, f := range root.children {
        if f.is_dir {
            new_path := strings.Join([]string{path, f.name}, "/");
            dir_size += find_candidates(f, dir_sizes, new_path);
        } else {
            dir_size += f.size;
        }
    }
    *dir_sizes = append(*dir_sizes, dir_size);
    return dir_size;
}


func parse_filesystem(stream string) Node {
    var root Node;
    root.name = "";
    root.is_dir = true;
    var stack []*Node;
    stack = append(stack, &root);
    var token string;
    for {
        if stream == "" {
            break; // End Of Stream
        }
        token, stream = get_token(stream);
        if token == "$" {
            token, stream = get_token(stream);
            switch token {
                case "cd":
                    token, stream = get_token(stream);
                    if token == ".." {
                        pop(&stack);
                    } else if token == "/" {
                        stack = stack[:1];
                    } else {
                        if !cd(&stack, token) {
                            log.Fatal("dir not found ", token);
                        }
                    }
                case "ls": // Nothing
                default:
                    log.Fatalf("unexpected token %q", token);
            }
        } else if token == "dir" {
            token, stream = get_token(stream);
            var node Node;
            node.name   = token;
            node.is_dir = true;
            cwd := get_cwd(stack);
            cwd.children = append(cwd.children, node);
        } else {
            // Expecting number
            size, err := strconv.Atoi(token);
            if err != nil {
                log.Fatalf("expected filesize, got %q", token);
            }
            token, stream = get_token(stream);
            var node Node;
            node.name   = token;
            node.is_dir = false;
            node.size   = size;
            cwd := get_cwd(stack);
            cwd.children = append(cwd.children, node);
        }
    }
    return root;
}


func get_cwd(stack []*Node) *Node {
    return stack[len(stack) - 1];
}


func cd(stack *[]*Node, dir string) bool {
    node := get_cwd(*stack);
    for i, child := range node.children {
        if child.name == dir {
            *stack = append(*stack, &node.children[i]);
            return true;
        }
    }
    return false;
}


func pop(stack *[]*Node) {
    *stack = (*stack)[:len(*stack) - 1];
}


func get_token(str string) (before, after string) {
    for i, c := range str {
        if c == ' ' || c == '\n' {
            return str[:i], str[i + 1:];
        }
    }
    return str, "";
}


func dot_print(str string, width int) {
    var dots string = ".";
    dots = strings.Repeat(dots, width - len(str));
    fmt.Print(str, dots);
}
