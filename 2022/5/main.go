
// This is the most scuffed input so far

package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "strconv"
)

func main() {
    data, err := os.ReadFile("input.txt");
    if err != nil {
        log.Fatal(err);
    }
    stream := string(data);
    var stacks map[int][]rune;
    stacks = read_setup(stream);
    stream = strings.TrimLeft(stream, " \n[]ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"); // Hilarious
    stream = strings.TrimRight(stream, "\n");
    lines := strings.Split(stream, "\n");
    for _, line := range lines {
        line = line[5:]; // "move "

        temp_line := strings.TrimRight(line[:2], " ");
        n, err := strconv.Atoi(temp_line);
        if err != nil {
            log.Fatal(err);
        }
        line = line[2:];
        line = strings.TrimLeft(line, " ");
        line = line[5:]; // "from "

        temp_line = strings.TrimRight(line[:1], " ");
        from, err := strconv.Atoi(temp_line);
        if err != nil {
            log.Fatal(err);
        }

        line = line[1:];
        line = strings.TrimLeft(line, " ");
        line = line[3:]; // "to "

        temp_line = strings.TrimRight(line[:1], " ");
        to, err := strconv.Atoi(temp_line);
        if err != nil {
            log.Fatal(err);
        }

        var temp_stack []rune;
        for i := 0; i < n; i += 1 {
            c := stacks[from][len(stacks[from]) - 1];
            temp_stack = append(temp_stack, c);
            stacks[from] = stacks[from][:len(stacks[from]) - 1];
        }
        for i := n - 1; i >= 0; i -= 1{
            stacks[to] = append(stacks[to], temp_stack[i]);
        }
    }

    for i := 1; i <= len(stacks); i += 1 {
        x := stacks[i];
        fmt.Println(i, string(x[len(x) - 1:]));
    }
}


// Note: This procedure will have to be hacky unfortunately because of the awkward input
func read_setup(in string) map[int][]rune {
    var stacks map[int][]rune = make(map[int][]rune, 0);
    lines := strings.Split(in, "\n");
    var max_height int;
    var stack_label_line string;
    find_stack_indices:
    for linenumber, line := range lines {
        for _, c := range line {
            if c >= '1' && c <= '9' { // Really we just need to look for '1'...
                stack_label_line = line;
                max_height = linenumber;
                break find_stack_indices;
            }
        }
    }
    var stack int = 1;
    var col int = 1;
    for stack_label_line != "" {
        stack_label_line = strings.TrimLeft(stack_label_line, "0123456789");
        stack_label_line = strings.TrimLeft(stack_label_line, " ");
        if stack_label_line == "" {
            break;
        }
        for i := max_height - 1; i >= 0; i -= 1 {
            c := lines[i][col];
            if c == ' ' {
                break;
            }
            stacks[stack] = append(stacks[stack], rune(c));
        }
        col += 4;
        stack += 1;
    }
    return stacks;
}
