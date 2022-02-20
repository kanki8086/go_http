package main

import (
	"fmt"
	"net/http"
	"time"
)

func Now(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, time.Now())
}
