package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type equipment struct {
	ID 		int		`json:"id"`
	Name 	string	`json:"name"`
	Desc 	string	`json:"desc"`
	Atk		int		`json:"atk"`
	Def		int		`json:"def"`
	HP 		int		`json:"hp"`
}

var data = []equipment{
	{1, "Blade", "A basic blade", 5, 2, 0},
	{2, "Leather Armor", "A basic armor made from leather", 0, 5, 10},
}

func equipments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal((data))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write((result))
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/equipments", equipments)
	fmt.Println(("starting web server at http:/localhost:8000/"))
	http.ListenAndServe(":8000", nil)
}

