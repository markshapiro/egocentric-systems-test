package main

import (
	"egocentric-systems-test/calculator"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	repo := calculator.NewOperationRepo()
	service := calculator.NewOperationService(repo)
	handler := calculator.NewOperationHandler(service)

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	handler.MountEndpoints(mux)

	go func() {
		fmt.Println("Server is running on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error:", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	fmt.Println("Stopping Server...")
}
