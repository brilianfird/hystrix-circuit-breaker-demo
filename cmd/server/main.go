package main

import (
	"hystrix-circuit-breaker-demo/internal/api"
	"hystrix-circuit-breaker-demo/internal/cb"
	"hystrix-circuit-breaker-demo/internal/http"
	"log"
	stdhttp "net/http"

	hystrixgo "github.com/afex/hystrix-go/hystrix"
)

func main() {
	// Initialize components
	apiClient := api.NewClient()
	handler := http.NewHandler(apiClient)

	// Configure Hystrix
	cb.ConfigureBreaker("mock_api_call", cb.DefaultBreakerConfig())

	// Start Hystrix stream handler
	hystrixStreamHandler := hystrixgo.NewStreamHandler()
	hystrixStreamHandler.Start()
	go func() {
		log.Println("Starting Hystrix stream handler on port 8081...")
		if err := stdhttp.ListenAndServe(":8081", hystrixStreamHandler); err != nil {
			log.Printf("Error starting Hystrix stream handler: %v", err)
		}
	}()

	// Register routes
	stdhttp.HandleFunc("/invoke", handler.HandleInvoke)
	stdhttp.HandleFunc("/mock-api", handler.HandleMockAPI)

	// Start main server
	log.Println("Starting main server on port 8080...")
	if err := stdhttp.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
