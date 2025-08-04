package web

import (
    "log"
    "net/http"
)

func StartWebUI() {
    log.Println("🖥️ Starting Web Dashboard at http://localhost:8080")
    http.Handle("/", http.FileServer(http.Dir("internal/web/static")))
    http.ListenAndServe(":8080", nil)
}
