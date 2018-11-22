package main

import (
	"net/http"

	"github.com/SergioML9/caliogo/routers"
	"github.com/urfave/negroni"
)

func main() {
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
