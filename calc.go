package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func CalcSum(w http.ResponseWriter, r *http.Request) {
	sum := calc([]int{2, 5, 2, 4, 4, 24})
	fmt.Fprintln(w, "sum="+strconv.Itoa(sum))
}

func calc(array []int) int {
	sum := 0
	for _, i := range array {
		sum += i
	}
	return sum
}
