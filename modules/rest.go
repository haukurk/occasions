package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/haukurk/occasions/utils"
	"net/http"
)

type ODate struct {
	Summary   string
	DateStart string
	DateEnd   string
}

func GetOccasions(w rest.ResponseWriter, req *rest.Request) {
	datesor := ODate{
		Summary:   "blehhh",
		DateStart: "Antoine",
		DateEnd:   "hehe",
	}

	w.WriteJson(&datesor)
}

func initRestInterface(string Port) {
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		rest.Route{"GET", "/api", GetOccasions},
	)
	listenStr := []string{"", ":", Port}
	http.ListenAndServe(listenStr, &handler)
}
