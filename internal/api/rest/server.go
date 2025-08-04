
package rest

import (
    "encoding/json"
    "log"
    "net/http"
)

func StartRESTServer(port string) {
    http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
    })

    log.Println("✅ REST API listening on port", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
