// fakeWeb project main.go
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	address := "0.0.0.0:8082"
	statusCode := 200
	logging := false
	flag.StringVar(&address, "address", address, "address to listen on")
	flag.IntVar(&statusCode, "status", statusCode, "status code to return")
	flag.BoolVar(&logging, "log", logging, "enable request logging to stdout")
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if logging {
			fmt.Printf("%s %s %d\n", time.Now().Format(time.RFC3339), r.URL.RequestURI(), statusCode)
		}
		w.WriteHeader(statusCode)
	})
	server := http.Server{ReadTimeout: time.Second, Addr: address}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Problem listening - check for conflicting processes and permissions " + err.Error())
		os.Exit(1)
	}
}
