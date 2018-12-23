package main

import (
	"fmt";
	"log";
	"net/http"
)

func main() {
	fmt.Printf("hello, world\n")
	routers();
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func routers() {
	httpGet("/login");
}

func httpGet(route string) {
	backend := "http://localhost"
	http.Get(backend + route);
}
