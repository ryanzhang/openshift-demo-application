package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func demo(w http.ResponseWriter, r *http.Request) {
	// val, err := ioutil.ReadFile("greeting.txt")
	// if err != nil {
	// 	panic(err)
	// }

	val := `
 	 _   _      _ _           ____ ____ ____  
	| | | | ___| | | ___     / ___/ ___/ ___| 
	| |_| |/ _ \ | |/ _ \   | |  | |  _\___ \ 
	|  _  |  __/ | | (_) |  | |__| |_| |___) |
	|_| |_|\___|_|_|\___/( ) \____\____|____/ 
	                     |/                    
	 `
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "<marquee bgcolor='red'><h1><pre>"+string(val))
}

func health(w http.ResponseWriter, r *http.Request) {
	// Simulate at least a bit of processing time.
	time.Sleep(100 * time.Millisecond)

	w.WriteHeader(http.StatusOK)
	if reqBytes, err := httputil.DumpRequest(r, true); err == nil {
		log.Printf("Openshift Http Request Dumper received a message: %+v", string(reqBytes))
		w.Write(reqBytes)
	} else {
		log.Printf("Error dumping the request: %+v :: %+v", err, r)
	}
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", demo)
	m.HandleFunc("/health", health)

	http.ListenAndServe(":8080", m)
}
