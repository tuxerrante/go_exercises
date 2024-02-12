package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	bars "server2/api/v1"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	newId    uuid.UUID               = uuid.New()
	BarAlex  *bars.Bar               = &bars.Bar{Name: "Alex", Id: newId}
	MyBars   bars.BarList            = bars.BarList{}
	MyBarsId map[uuid.UUID]*bars.Bar = make(map[uuid.UUID]*bars.Bar)
)

func main() {

	MyBarsId[newId] = BarAlex
	log.Println(MyBars)
	log.Println(BarAlex)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/bar/{id}", bar)

	log.Fatal(http.ListenAndServe(":8888", router))
}

func bar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	barId := vars["id"]
	parsedId, err := uuid.Parse(barId)
	if err != nil {
		fmt.Fprintf(w, "Error converting the UUID: %s", err)
		return
	}
	json.NewEncoder(w).Encode(MyBars.Id[parsedId])
}
