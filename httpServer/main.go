package main

import (
	"io"
	"net/http"
)

type stringHandler struct {
	msg string
}

func (sh stringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	MyPrintlin("processing req: %v", r.URL.Path)
	io.WriteString(w, sh.msg)
}

func main() {
	http.Handle("/msg", stringHandler{"Default page"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/msg", http.StatusTemporaryRedirect))
	go func() {

		err := http.ListenAndServeTLS(":5001", "localhost-pub.cer", "localhost.key", nil)
		if err != nil {
			MyPrintlin("TLS Error Occured: %v", err.Error())
		}
	}()
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		MyPrintlin("Error: %v", err.Error())
	}
}
