package main

import (
	"api-gateway/internal/config"
	"api-gateway/internal/middleware"
	"api-gateway/internal/proxy"
	"api-gateway/internal/router"
	"log"
	"net/http"
)

func main() {
	// 1. Load Configurations
	cfg, err := config.LoadConfig("configs/routes.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Println("Successfully loaded routing configurations.")

	// 2. Initialize Router
	rt := router.NewRouter(cfg.Routes)

	// 3. Setup Gateway Handler & Middleware Chain
	gatewayHandler := proxy.GatewayHandler(rt)
	handlerPipeline := middleware.Chain(gatewayHandler, middleware.Logging)

	// 4. Start Server
	port := ":8080"
	log.Printf("API Gateway running on port %s", port)
	if err := http.ListenAndServe(port, handlerPipeline); err != nil {
		log.Fatalf("Server stopped: %v", err)
	}
}
