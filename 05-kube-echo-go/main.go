package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	var host string
	var port int
	flag.StringVar(&host, "host", "0.0.0.0", "host to listen")
	flag.IntVar(&port, "port", 9090, "port to listen")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		buffer := bytes.NewBuffer(nil)
		//print request method and URI
		buffer.WriteString(fmt.Sprintln(req.Method, req.RequestURI))

		buffer.WriteString(fmt.Sprintln())
		//print headers
		for k, v := range req.Header {
			buffer.WriteString(fmt.Sprintln(k+":", strings.Join(v, ";")))
		}
		buffer.WriteString(fmt.Sprintln())

		//print request body
		reqBody, readErr := ioutil.ReadAll(req.Body)
		if readErr != nil {
			buffer.WriteString(fmt.Sprintln(readErr))
		} else {
			buffer.WriteString(fmt.Sprintln(string(reqBody)))
		}

		outputData := buffer.Bytes()
		fmt.Println(string(outputData))
		w.Write(outputData)
	})

	//listen and serve
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	fmt.Println("listen err,", err)
}

//$ go build echo-go.go
//
//$ ./echo-go -host 0.0.0.0 -port 9090
//
//# 新开一个终端，执行如下命令
//
//$ curl http://localhost:9090/
//GET /
//
//User-Agent: curl/7.47.0
//Accept: */*
//
//$ curl http://localhost:9090/  -X POST -d '{"ok":true}' -H 'Content-Type: application/json'
//POST /
//
//User-Agent: curl/7.47.0
//Accept: */*
//Content-Type: application/json
//Content-Length: 11
//
//{"ok":true}
