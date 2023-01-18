
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

    var max_calorie_elf int;
    var max_calories int;
    var elf int = 1;
    var calories int;
    lines := strings.Split(string(data), "\n");
    for _, line := range lines {
        if line == "" {
            if calories > max_calories {
                max_calories = calories;
                max_calorie_elf = elf;
            }
            calories = 0;
            elf += 1;
        } else {
            calorie, err := strconv.Atoi(line);
            if err != nil {
                log.Fatal(err);
            }
            calories += calorie;
        }
    }
    fmt.Println("elf:", max_calorie_elf, "calories:", max_calories);
}
