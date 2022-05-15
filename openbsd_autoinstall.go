package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [autoinstall file]", os.Args[0])
		os.Exit(1)
	}
	autoconfFile := os.Args[1]
	if _, err := os.Stat(autoconfFile); err != nil {
		fmt.Fprintf(os.Stderr, "stat file '%s': %v", autoconfFile, err)
		os.Exit(2)
	}
	http.HandleFunc("/install.conf", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, autoconfFile)
	})
	http.ListenAndServe("0.0.0.0:8000", nil)
}
