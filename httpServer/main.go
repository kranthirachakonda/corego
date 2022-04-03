package main

import (
	"io"
	"net/http"
)

type stringHandler struct {
	msg string
}

func (sh stringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	MyPrintlin("method: %v", r.Method)
	MyPrintlin("url: %v", r.URL)
	MyPrintlin("host: %v", r.Host)
	MyPrintlin("proto: %v", r.Proto)
	for name, val := range r.Header {
		MyPrintlin("header: %v, value: %v", name, val)
	}
	MyPrintlin("----------")
	io.WriteString(w, sh.msg)
}

func main() {
	err := http.ListenAndServe(":5000", stringHandler{msg: "simple golang httpServer"})
	if err != nil {
		MyPrintlin("Error Occured: %v", err.Error())
	}
}
