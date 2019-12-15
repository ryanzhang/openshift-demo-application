package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func demo(w http.ResponseWriter, r *http.Request) {
	val, err := ioutil.ReadFile("greeting.txt")
	if err != nil {
		panic(err)
	}

	// 	val := `
	// 	____   ___  _  ___    ____                  _
	// 	|___ \ / _ \/ |/ _ \  / ___|  ___ _ ____   _(_) ___ ___
	// 		__) | | | | | (_) | \___ \ / _ \ '__\ \ / / |/ __/ _ \
	// 	 / __/| |_| | |\__, |  ___) |  __/ |   \ V /| | (_|  __/
	// 	|_____|\___/|_|  /_/  |____/ \___|_|    \_/ |_|\___\___|

	// 	 _                 _
	// 	| |__   ___   ___ | |_ ___ __ _ _ __ ___  _ __
	// 	| '_ \ / _ \ / _ \| __/ __/ _  | '_   _ \| '_ \
	// 	| |_) | (_) | (_) | || (_| (_| | | | | | | |_) |
	// 	|_.__/ \___/ \___/ \__\___\__,_|_| |_| |_| .__/
	// 																					 |_|

	// `
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(val))
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
