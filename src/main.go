package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz/{num}", fizzbuzzHandler)

	fmt.Println("server is running...")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", "8080"), r); err != nil {
		log.Fatal(err)
	}
}

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num, err := strconv.Atoi(vars["num"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid number")
		return
	}

	ans := fizzbuzz(num)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, ans)
}

func fizzbuzz(n int) string {
	switch {
	case n%15 == 0:
		return "Fizz Buzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(n)
}
