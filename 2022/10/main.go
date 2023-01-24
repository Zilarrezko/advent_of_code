
package main;

import (
    "os"
    "log"
    "strings"
    "strconv"
    "fmt"
)


const (
    width = 40;
    height = 6;
)


func main() {
    data, err := os.ReadFile("input.txt");
    if err != nil {
        log.Fatal(err);
    }
    var trace []int;
    var x int = 1;
    lines := strings.Split(string(data), "\n");
    for _, line := range lines {
        op := line[:4];
        switch op {
            case "addx":
                val, err := strconv.Atoi(line[5:]);
                if err != nil {
                    log.Fatal(err);
                }
                trace = append(trace, x, x);
                x += val;
            case "noop":
                trace = append(trace, x);
            default:
                log.Fatal("unknown mnemonic"); // unreachable
        }
    }

    var screen [height*width]rune;
    for c, v := range trace {
        x := c%width;
        r := '.';
        if v - 1 == x || v == x || v + 1 == x {
            r = '#';
        }
        screen[c] = r;
    }

    for i, r := range screen {
        x := i%width;
        y := i/width;
        if x == 0 && y != 0 {
            fmt.Println();
        }
        fmt.Print(string(r));
    }
}
