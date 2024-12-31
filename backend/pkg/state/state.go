package state

import (
	"net/http"
	"rfid_handler/pkg/utils"
	"sync"
	"time"
)


type Item struct {
	ID       string
	Collected bool
	CollectionDate *time.Time 
}


type AppState struct {
	Items []Item
	Mutex sync.Mutex
}


func NewAppState() *AppState {
	return &AppState{
		Items: []Item{	
		   	{ID: "a3b2d1", Collected: false, CollectionDate: nil},
			{ID: "d6e4f7", Collected: false, CollectionDate: nil},
			{ID: "z1x3c8", Collected: false, CollectionDate: nil},
			{ID: "j5k7l2", Collected: false, CollectionDate: nil},
			{ID: "m9n0o4", Collected: false, CollectionDate: nil},
			{ID: "p8q3r1", Collected: false, CollectionDate: nil},
			{ID: "u5v2w9", Collected: false, CollectionDate: nil},
		},
	}
}

func HandleGetStatus(s *AppState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method. Only GET is allowed.", http.StatusMethodNotAllowed)
			return
		}		
		s.Mutex.Lock()
		defer s.Mutex.Unlock()
		utils.SendJSONResponse(w, http.StatusOK, s.Items)
	}
}