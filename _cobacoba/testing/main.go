package testing

import (
	"fmt"
	"net/http"
	// equipments "/equipments"
)

func main() {
	http.HandleFunc("/equipments", equipments)
	http.HandleFunc("/events", equipments)
	fmt.Println(("starting web server at http:/localhost:8000/"))
	http.ListenAndServe(":8000", nil)
}