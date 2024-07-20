package main

import (
    "log"
    "os"

    "github.com/ceramicnetwork/go-ipfs-healthcheck/healthcheck"
    coreiface "github.com/ipfs/kubo/core/coreiface"
)

const portEnvVar = "HEALTHCHECK_PORT"
var port = "8011"

func main() {
    envPort := os.Getenv(portEnvVar)
    if envPort != "" {
        port = envPort
    }

    log.Printf("Using port: %s\n", port)

    // Placeholder for initializing IPFS CoreAPI instance
    var ipfs coreiface.CoreAPI
    // Initialize IPFS CoreAPI instance here

    log.Println("Starting health check server")
    healthcheck.StartServer(port, ipfs)
}
