package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/calc", CalcSum)
	http.HandleFunc("/now", Now)
	http.ListenAndServe(":8080", nil)
}
