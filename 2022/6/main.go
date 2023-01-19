
package main

import (
    "fmt"
    "os"
    "log"
)

const g_window_size int = 14;

func main () {
    data, err := os.ReadFile("input.txt");
    if err != nil {
        log.Fatal(err);
    }
    for i := 0; i < len(data) - (g_window_size - 1); i += 1 {
        window := data[i : i + g_window_size];
        if !has_duplicate(window) {
            fmt.Println(i + g_window_size);
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
