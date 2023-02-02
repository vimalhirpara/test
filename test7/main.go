package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Test")
	})

	fmt.Println("9090 port is running")
	http.ListenAndServe(":9090", router)
}
