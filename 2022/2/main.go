
package main

import (
    "os"
    "log"
    "fmt"
    "strings"
)

func main() {
    data, err := os.ReadFile("input.txt");
    if err != nil {
        log.Fatal(err);
    }
    lines := strings.Split(string(data), "\n");
    var sum int;
    for _, line := range lines {
        if line != "" {
            hands := strings.Split(line, " ");
            opponent := get_hand_index(hands[0]);
            player   := get_hand_index(hands[1]);
            sum += reward[opponent][player];
        }
    }
    fmt.Println("sum:", sum);
}

// A, X = rock
// B, Y = paper
// C, Z = scissors

var reward [3][3]int = [3][3]int{
//   X,     Y,     Z
    {1 + 3, 2 + 6, 3 + 0}, // A
    {1 + 0, 2 + 3, 3 + 6}, // B
    {1 + 6, 2 + 0, 3 + 3}, // C
}

func get_hand_index(hand string) int {
    switch hand {
        case "A": fallthrough
        case "X": return 0;
        case "B": fallthrough
        case "Y": return 1;
        case "C": fallthrough
        case "Z": return 2;
    }
    return -1;
}
