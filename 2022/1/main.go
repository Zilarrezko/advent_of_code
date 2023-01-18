
package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "strconv"
    "sort"
)

func main() {
    data, err := os.ReadFile("input.txt");
    if err != nil {
        log.Fatal(err);
    }

    // Doing it this way obviously isn't the most performant method, but it works, and it's somewhat elegant
    var calories_list []int;
    var calories int;
    lines := strings.Split(string(data), "\n");
    for _, line := range lines {
        if line == "" {
            calories_list = append(calories_list, calories);
            calories = 0;
        } else {
            calorie, err := strconv.Atoi(line);
            if err != nil {
                log.Fatal(err);
            }
            calories += calorie;
        }
    }
    sort.Ints(calories_list); // ascending
    const elves int = 3;
    var sum int;
    for i := len(calories_list) - 1; i >= len(calories_list) - elves; i -= 1 {
        x := calories_list[i];
        sum += x;
        fmt.Println(i, x);
    }
    fmt.Println("sum:", sum);
}
