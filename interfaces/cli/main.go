
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("🔧 OptiTerra CLI")
    args := os.Args
    if len(args) > 1 {
        fmt.Println("Command:", args[1])
    } else {
        fmt.Println("Usage: optiterra-cli [command]")
    }
}
