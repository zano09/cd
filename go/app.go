package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Get the file path from the "FILE_PATH" environment variable, or use a default value
	filePath := os.Getenv("FILE_PATH")
	if filePath == "" {
		filePath = "/var/secret/secret.txt"
	}

	// Get the greeting from the "GREETING" environment variable, or use a default value
	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = "Hello World!"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			// File not found or error reading the file
			fmt.Fprintf(w, "%s\nServer hostname: %s\n", greeting, getHostname())
		} else {
			// File found, send its content along with the server hostname
			fmt.Fprintf(w, "File content:\n%s\nServer hostname: %s\n", data, getHostname())
		}
	})

	port := 3000
	addr := fmt.Sprintf(":%d", port)

	server := &http.Server{Addr: addr}

	go func() {
		fmt.Printf("Server is running on http://localhost:%d\n", port)
		fmt.Printf("File path is set to: %s\n", filePath)
		handleSignals(server)
	}()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error: %s\n", err)
	}
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "Unknown"
	}
	return hostname
}

func handleSignals(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	sig := <-stop

	fmt.Printf("Received signal: %v. Shutting down gracefully...\n", sig)

	err := server.Shutdown(nil)
	if err != nil {
		fmt.Printf("Error during server shutdown: %s\n", err)
	}

	fmt.Println("Server has been closed.")
}
