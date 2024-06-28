package main

import "net/http"

func res(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", res)
	http.ListenAndServe(":8080", router)
}
