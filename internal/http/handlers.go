package http

import (
	"fmt"
	"hystrix-circuit-breaker-demo/internal/api"
	"log"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

// Handler handles HTTP requests
type Handler struct {
	apiClient *api.Client
}

// NewHandler creates a new instance of HTTP Handler
func NewHandler(apiClient *api.Client) *Handler {
	return &Handler{
		apiClient: apiClient,
	}
}

// HandleMockAPI handles requests to the mock API endpoint
func (h *Handler) HandleMockAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("Mock API called")
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// HandleInvoke handles requests to the invoke endpoint
func (h *Handler) HandleInvoke(w http.ResponseWriter, r *http.Request) {
	output := make(chan string, 1)
	errors := hystrix.Go("mock_api_call", func() error {
		log.Println("Calling mock API")
		result, err := h.apiClient.CallMockEndpoint()
		if err != nil {
			return err
		}
		output <- result
		return nil
	}, func(err error) error {
		// Fallback function
		output <- fmt.Sprintf("Fallback: %v", err)
		return nil
	})

	select {
	case result := <-output:
		fmt.Fprintln(w, result)
	case err := <-errors:
		log.Printf("Error: %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
	}
}
