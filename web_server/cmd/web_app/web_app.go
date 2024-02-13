package main

import (
	"mvp-2-spms/web_server/routes"
	"net/http"
)

func main() {
	router := routes.SetupRouter()
	http.ListenAndServe(":8080", router.Router())
}
