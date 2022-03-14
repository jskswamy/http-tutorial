package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

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
		writer.WriteHeader(200)
	default:
		writer.WriteHeader(404)
	}
}

func putHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "PUT":
		writer.WriteHeader(200)
	default:
		writer.WriteHeader(404)
	}
}

func deleteHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "DELETE":
		writer.WriteHeader(200)
	default:
		writer.WriteHeader(404)
	}
}

func rootHandler(_ http.ResponseWriter, req *http.Request) {
	result, _ := httputil.DumpRequest(req, true)
	_, _ = fmt.Fprintf(os.Stdout, "%s", result)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/404", notFoundHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/put", putHandler)
	http.HandleFunc("/delete", deleteHandler)

	err := http.ListenAndServe(":8010", nil)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to start the server, reason %v", err)
	}
}
