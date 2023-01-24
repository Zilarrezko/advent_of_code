
package main;

import (
    "os"
    "log"
    "fmt"
    "strings"
    "strconv"
    "sort"
)

type Monkey struct {
    items    []int;
    op       byte;
    op_val   int; // 0 means itself
    test     int;
    on_true  int;
    on_false int;
}


func main() {
    data, err := os.ReadFile("input.txt");
    if err != nil {
        log.Fatal(err);
    }
    var monkeys []Monkey;
    var stream string = string(data);
    var token string;
    for {
        token, stream = get_token(stream);
        if stream == "" {
            break;
        }
        var monkey Monkey;
        expect(token, "Monkey");
        token, stream = get_token(stream); // #:
        token = token[:len(token) - 1];

        // Starting items: #, #, #
        token, stream = get_token(stream);
        expect(token, "Starting");
        token, stream = get_token(stream);
        expect(token, "items:");
        var another bool = true;
        for another {
            token, stream = get_token(stream);
            another = false;
            if token[len(token) - 1] == ',' {
                another = true;
                token = token[:len(token) - 1];
            }
            worry_level, err := strconv.Atoi(token);
            if err != nil {
                log.Fatal(err);
            }
            monkey.items = append(monkey.items, worry_level);
        }

        // Operation: new = old * 12
        token, stream = get_token(stream);
        expect(token, "Operation:");
        token, stream = get_token(stream);
        expect(token, "new");
        token, stream = get_token(stream);
        expect(token, "=");
        token, stream = get_token(stream);
        expect(token, "old");
        token, stream = get_token(stream); // * or +, could make expect_set
        monkey.op = token[0];
        token, stream = get_token(stream);
        if token == "old" {
            monkey.op_val = 0; // 0 will indicate to double/square
        } else {
            val, err := strconv.Atoi(token);
            if err != nil {
                log.Fatal(err);
            }
            monkey.op_val = val;
        }

        // Test: divisible by #
        token, stream = get_token(stream);
        expect(token, "Test:");
        token, stream = get_token(stream);
        expect(token, "divisible");
        token, stream = get_token(stream);
        expect(token, "by");

        token, stream = get_token(stream);
        div, err := strconv.Atoi(token);
        if err != nil {
            log.Fatal(err);
        }
        monkey.test = div;

        // If true: throw to monkey #
        token, stream = get_token(stream);
        expect(token, "If");
        token, stream = get_token(stream);
        expect(token, "true:");
        token, stream = get_token(stream);
        expect(token, "throw");
        token, stream = get_token(stream);
        expect(token, "to");
        token, stream = get_token(stream);
        expect(token, "monkey");
        token, stream = get_token(stream);
        to_monkey, err := strconv.Atoi(token);
        if err != nil {
            log.Fatal(err);
        }
        monkey.on_true = to_monkey;

        // If false: throw to monkey #
        token, stream = get_token(stream);
        expect(token, "If");
        token, stream = get_token(stream);
        expect(token, "false:");
        token, stream = get_token(stream);
        expect(token, "throw");
        token, stream = get_token(stream);
        expect(token, "to");
        token, stream = get_token(stream);
        expect(token, "monkey");
        token, stream = get_token(stream);
        to_monkey, err = strconv.Atoi(token);
        if err != nil {
            log.Fatal(err);
        }
        monkey.on_false = to_monkey;

        monkeys = append(monkeys, monkey);
    }


    var inspections []int = make([]int, len(monkeys));
    var total_rounds int = 20;
    for round := 1; round <= total_rounds; round += 1 {
        for i, m := range monkeys {
            for j := len(m.items) - 1; j >= 0; j -= 1 {
                inspections[i] += 1;
                v := m.items[j];
                u := m.op_val;
                if u == 0 {
                    u = v;
                }
                if m.op == '*' {
                    v = v*u;
                } else {
                    v = v + u;
                }
                v = v/3;
                if v % m.test == 0 {
                    monkeys[m.on_true].items = append(monkeys[m.on_true].items, v);
                } else {
                    monkeys[m.on_false].items = append(monkeys[m.on_false].items, v);
                }
            }
            monkeys[i].items = []int{};
        }
    }

    for i, it := range inspections {
        fmt.Println("Monkey", i, "inspected", it, "items");
    }

    sort.Ints(inspections);
    fmt.Println("ans:", inspections[len(inspections) - 1]*inspections[len(inspections) - 2]);
}


func expect(tok string, e string) {
    if tok != e {
        log.Fatalf("expected %q got %q", e, tok);
    }
}


func get_token(str string) (before, after string) {
    str = strings.TrimLeft(str, " \r\t\n");
    for i, c := range str {
        if c == ' ' || c == '\n' || c == '\t' || c == '\r' {
            return str[:i], str[i + 1:];
        }
    }
    return str, "";
}
