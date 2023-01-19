
package main

import (
    "fmt"
    "os"
    "log"
)

func main () {
    data, err := os.ReadFile("input.txt");
    if err != nil {
        log.Fatal(err);
    }
    for i := 0; i < len(data) - 3; i += 1 {
        window := data[i:i+4];
        if !has_duplicate(window) {
            fmt.Println(i + 4);
            break;
        }
    }
}


func has_duplicate(arr []byte) bool {
    for j, _ := range arr {
        for k := j + 1; k < len(arr); k += 1 {
            if arr[j] == arr[k] {
                return true;
            }
        }
    }
    return false;
}
