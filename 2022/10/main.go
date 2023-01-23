
package main;

import (
    "os"
    "log"
    "strings"
    "strconv"
    "fmt"
)


func main() {
    data, err := os.ReadFile("input.txt");
    if err != nil {
        log.Fatal(err);
    }
    var trace []int; // hilarious
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
    var breakpoints []int = []int{20, 60, 100, 140, 180, 220};
    var sum int;
    for _, x := range breakpoints {
        sum += trace[x - 1]*x;
    }
    fmt.Println("sum", sum);
}
