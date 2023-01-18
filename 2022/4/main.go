
// I wanted to do a bitwise mask method, but looks like the numbers are > 64

package main

import (
    "os"
    "log"
    "strings"
    "fmt"
    // "strconv"
)

func main() {
    data, err := os.ReadFile("input.txt");
    if err != nil {
        log.Fatal(err);
    }
    var sum int;
    lines := strings.Split(string(data), "\n");
    for _, line := range lines {
        if line != "" {
            line = strings.TrimLeft(line, " \r\t\n");

            // format: n-n,n-n

            // First segment
            a0 := get_next_number(&line);
            line = line[1:]; // '-'
            a1 := get_next_number(&line);

            // Second segment
            line = line[1:]; // ','
            b0 := get_next_number(&line);
            line = line[1:]; // '-'
            b1 := get_next_number(&line);

            if a1 >= b0 && a1 <= b1 || a0 <= b1 && a0 >= b0 ||
               b1 >= a0 && b1 <= a1 || b0 <= a1 && b0 >= a0 {
                sum += 1;
            }
        }
    }
    fmt.Println(sum);
}


func get_next_number(str *string) int {
    i, ok := parse_integer(*str);
    if !ok {
        log.Fatal();
    }
    *str = strings.TrimLeft(*str, "0123456789");
    return i;
}


func parse_integer(str string) (int, bool) {
    var ret int;
    var ok bool;
    for _, c := range str {
        if !is_numeric(c) {
            break;
        }
        ok = true;
        ret *= 10;
        ret += int(c - '0');
    }
    return ret, ok;
}


func is_numeric(c rune) bool {
    if c >= '0' && c <= '9' {
        return true;
    }
    return false;
}
