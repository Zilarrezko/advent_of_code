
package main

import (
    "os"
    "log"
    "fmt"
    "strings"
    "strconv"
)

type Node struct {
    name   string
    is_dir bool
    size   int
    children  []Node
}

const (
    candidate_size int = 100000;
)


func main() {
    data, err := os.ReadFile("input.txt");
    if err != nil {
        log.Fatal(err);
    }
    root := parse_filesystem(string(data));
    find_candidates(root);
}


var g_condidate_total int;
func find_candidates(root Node) {
    path := "";
    g_condidate_total = 0;
    for _, f := range root.children {
        if f.is_dir {
            new_path := strings.Join([]string{path, f.name}, "/");
            find_candidate_helper(f, new_path);
        }
    }
    fmt.Println("Total freeable size:", g_condidate_total);
}
func find_candidate_helper(root Node, path string) int {
    var dir_size int;
    for _, f := range root.children {
        if f.is_dir {
            new_path := strings.Join([]string{path, f.name}, "/");
            dir_size += find_candidate_helper(f, new_path);
        } else {
            dir_size += f.size;
        }
    }
    if dir_size <= candidate_size {
        dot_print(path, 72);
        fmt.Println(dir_size);
        g_condidate_total += dir_size;
    }
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
