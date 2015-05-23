package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func testapi(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	fmt.Println("params:", params)
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are get user %s", uid)
}

func main() {
	mux := routes.New()
	mux.Get("/containers/json", testapi)
	http.ListenAndServe(":9090", nil)
}
