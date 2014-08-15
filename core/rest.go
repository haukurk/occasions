package occasions

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

func InitRestInterface(p string, dates []ODate) {
	handler := rest.ResourceHandler{}
	err := handler.SetRoutes(
		&rest.Route{"GET", "/api/occasions", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(&dates)
		}},
		&rest.Route{"GET", "/api/occasions/upcoming", func(w rest.ResponseWriter, req *rest.Request) {
			_, filteredDates, _ := UpcomingDates(dates)
			w.WriteJson(&filteredDates)
		}},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[Occasions] Starting REST Interface on port " + p + "..")
	log.Fatal(http.ListenAndServe(":"+p, &handler))

}
