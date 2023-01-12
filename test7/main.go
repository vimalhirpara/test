package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/test", testfunc)

	fmt.Println("9090 port is running")
	http.ListenAndServe(":9090", router)
}

func testfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test")
}
