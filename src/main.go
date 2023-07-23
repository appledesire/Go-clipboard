package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/atotto/clipboard"
)

var (
	clipboardContent string
	mutex            sync.Mutex
)

func main() {

	http.HandleFunc("/clipboard", handleClipboard)
	fmt.Println("Server is runing on Port 8080:")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleClipboard(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Retrieve the clipboard content
		mutex.Lock()
		defer mutex.Unlock()

		content, err := clipboard.ReadAll()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		modifiedContent := strings.ReplaceAll(content, "\r\n", "<br>")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// Start the HTML document
		fmt.Fprintf(w, "<html><body>")
		// contents
		fmt.Fprintf(w, "<p>%s</p>", modifiedContent)
		// End the HTML document
		fmt.Fprintf(w, "</body></html>")

		fmt.Println(modifiedContent)

		clipboardContent = content

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
