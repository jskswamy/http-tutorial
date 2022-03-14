package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

var defaultLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

func notFoundHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(404)
}

func jsonHandler(writer http.ResponseWriter, _ *http.Request) {
	data := map[string]string{"name": "basic example"}
	body, _ := json.Marshal(data)

	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(body)
}

func postHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		defaultLogger.Printf("valid method %s for post handler", req.Method)
		writer.WriteHeader(200)
	default:
		defaultLogger.Printf("method %s not supported for post handler", req.Method)
		writer.WriteHeader(404)
	}
}

func putHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "PUT":
		defaultLogger.Printf("valid method %s for put handler", req.Method)
		writer.WriteHeader(200)
	default:
		defaultLogger.Printf("method %s not supported for put handler", req.Method)
		writer.WriteHeader(404)
	}
}

func deleteHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "DELETE":
		defaultLogger.Printf("valid method %s for delete handler", req.Method)
		writer.WriteHeader(200)
	default:
		defaultLogger.Printf("method %s not supported for delete handler", req.Method)
		writer.WriteHeader(404)
	}
}

func rootHandler(_ http.ResponseWriter, req *http.Request) {
	result, _ := httputil.DumpRequest(req, true)
	defaultLogger.Printf("%s", result)
}

func main() {
	port := 8010
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/404", notFoundHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/put", putHandler)
	http.HandleFunc("/delete", deleteHandler)

	defaultLogger.Printf("starting http server using port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to start the server, reason %v", err)
	}
}
