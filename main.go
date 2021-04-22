package main

import "net/http"

type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar called!!!"))
}

func main() {
	http.Handle("/foo", &fooHandler{Message: "foo called!!!"})
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":5000", nil)
}
