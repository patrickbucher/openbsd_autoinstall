package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	endpoint = "/install.conf"
	port     = 8000
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [autoinstall file]\n", os.Args[0])
		os.Exit(1)
	}
	autoconfFile := os.Args[1]
	if _, err := os.Stat(autoconfFile); err != nil {
		fmt.Fprintf(os.Stderr, "stat file '%s': %v\n", autoconfFile, err)
		os.Exit(2)
	}
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("serve %s to %s\n", endpoint, r.RemoteAddr)
		http.ServeFile(w, r, autoconfFile)
	})
	fmt.Printf("serving %s on port %d\n", endpoint, port)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
}
