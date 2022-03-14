package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

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
		log.WithFields(
			log.Fields{"method": req.Method}).Debugf("valid method %s for post handler", req.Method)
		writer.WriteHeader(200)
	default:
		log.WithFields(
			log.Fields{"method": req.Method}).Debugf("method %s not supported for post handler", req.Method)
		writer.WriteHeader(404)
	}
}

func putHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "PUT":
		log.WithFields(
			log.Fields{"method": req.Method}).Debugf("valid method %s for put handler", req.Method)
		writer.WriteHeader(200)
	default:
		log.WithFields(
			log.Fields{"method": req.Method}).Debugf("method %s not supported for put handler", req.Method)
		writer.WriteHeader(404)
	}
}

func deleteHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "DELETE":
		log.WithFields(
			log.Fields{"method": req.Method}).Debugf("valid method %s for delete handler", req.Method)
		writer.WriteHeader(200)
	default:
		log.WithFields(
			log.Fields{"method": req.Method}).Debugf("method %s not supported for delete handler", req.Method)
		writer.WriteHeader(404)
	}
}

func rootHandler(_ http.ResponseWriter, req *http.Request) {
	result, _ := httputil.DumpRequest(req, true)
	log.WithFields(log.Fields{"request": fmt.Sprintf("%s", result)}).Infof("received")
}

func main() {
	port := 8010
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/404", notFoundHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/put", putHandler)
	http.HandleFunc("/delete", deleteHandler)

	log.WithFields(log.Fields{"port": port}).Infof("starting http server")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to start the server, reason %v", err)
	}
}
