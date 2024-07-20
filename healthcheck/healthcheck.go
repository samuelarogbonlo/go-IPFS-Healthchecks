package healthcheck

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "os/exec"
    "time"

    coreiface "github.com/ipfs/kubo/core/coreiface"
)

type HealthCheckResponse struct {
    Status string `json:"status"`
    Stats  string `json:"stats"`
}

func StartServer(port string, ipfs coreiface.CoreAPI) {
    log.Printf("Starting server on port %s\n", port)
    http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
        healthCheckHandler(w, r, ipfs)
    })
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request, ipfs coreiface.CoreAPI) {
    log.Println("Received health check request")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cmd := exec.CommandContext(ctx, "ipfs", "dag", "stat", "QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn") // Replace Qm... with the actual CID
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Printf("Error executing command: %v, output: %s\n", err, string(output))
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := HealthCheckResponse{
        Status: "ok",
        Stats:  string(output),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
    log.Println("Health check response sent")
}
