package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type event struct {
	ID 			int		`json:"id"`
	TypeID		int		`json:"type_id"`
	Name 		string	`json:"name"`
	Desc 		string	`json:"desc"`
	GoldReward 	int		`json:"gold_reward"`
	XPReward	int		`json:"xp_reward"`
}

var data = []event{
	{1, 1, "Windrise Exploration", "Explore Windrise area, some said there's a strage aura around the sacred tree", 10, 10},
	{2, 2, "Hililchurl Hunter", "Hunt Hililchurls that unsettling the farmers", 15, 15},
}

func events(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/events", events)
	fmt.Println(("starting web server at http:/localhost:8000/"))
	http.ListenAndServe(":8000", nil)
}

