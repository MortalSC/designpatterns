package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resultCh := make(chan string)
	go fetchDetaFromDB(ctx, resultCh)

	select {
	case result := <-resultCh:
		fmt.Fprintf(w, "Result: %s", result)
	case <-ctx.Done():
		http.Error(w, "Request timed out", http.StatusGatewayTimeout)
	}
}

func fetchDetaFromDB(ctx context.Context, resultCh chan<- string) {
	select {
	case <-time.After(3 * time.Second):
		resultCh <- "Data from DB"
	case <-ctx.Done():
		fmt.Printf("Request cancelled: %v\n", ctx.Err())
	}
}

func main() {
	http.HandleFunc("/", handler)
	server := &http.Server{Addr: ":8080"}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("Server error:", err)
		}
	}()

	// after 5 seconds, shutdown the server
	time.Sleep(5 * time.Second)
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(shutdownCtx)
	fmt.Println("Server shutdown")
}
