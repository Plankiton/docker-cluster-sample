package main

import (
	"net/http"

	Sex "github.com/Plankiton/SexPistol"
	"github.com/rs/cors"
)

func main() {
	h := cors.AllowAll()
	r := Sex.NewPistol().
		Add("/{name}", func(r Sex.Request) string {
			return Sex.Fmt("Hello, %s!", r.PathVars["name"])
		})

	s := h.Handler(r)
	http.ListenAndServe(":8000", s)
}
