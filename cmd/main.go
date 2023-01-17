package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Something struct {
	Ids []string
}

func getItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ids := strings.Split(ps.ByName("ids"), ",")
	// Here I could access the database and select records by their IDs
	// instead of reading a lousy-formed excel file and even more so
	// return an array of json files instead of a single json file with many fields.

	// I just want to say that the author of the task is a genius (no)
	// and I love his mother very much.

	s := Something{Ids: ids}
	body, _ := json.Marshal(s)

	w.Write(body)
	w.WriteHeader(200)
}

func main() {
	router := httprouter.New()
	router.GET("/get_items/:ids", getItems)

	log.Fatal(http.ListenAndServe(":8080", router))
}
