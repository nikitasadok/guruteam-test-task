package main

import (
	"net/http"
	"testTaskGuru/globals"
	"testTaskGuru/routes"
)

func main() {
	globals.InitApp()
	router := routes.NewRouter()
	http.ListenAndServe(":8080", router)
}
