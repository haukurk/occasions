package modules

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/haukurk/occasions/utils"
	"log"
	"net/http"
)

func getOccasions(w rest.ResponseWriter, req *rest.Request) {
	datesor := utils.ODate{
		Summary:   "blehhh",
		DateStart: req.PathParam("num"),
		DateEnd:   "hehe",
	}

	w.WriteJson(&datesor)
}

func InitRestInterface(p string) {
	handler := rest.ResourceHandler{}
	err := handler.SetRoutes(
		&rest.Route{"GET", "/api/occasions", getOccasions},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[Occasions] Starting REST Interface on port " + p + "..")
	log.Fatal(http.ListenAndServe(":"+p, &handler))

}
