
package main

import (
    "log"
    "os"

    "optiterra/internal/api/rest"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Println("🚀 Starting OptiTerra...")
    rest.StartRESTServer(port)
}
