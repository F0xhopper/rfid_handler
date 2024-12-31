package handler

import (
	"fmt"
	"net/http"
	"rfid_handler/pkg/state"
	"rfid_handler/pkg/utils"
	"runtime"
	"strconv"
	"time"
)

func clearConsole() {
	
	if runtime.GOOS == "windows" {
		
		fmt.Print("\x1b[H\x1b[2J") 
	} else {
		
		fmt.Print("\x1b[H\x1b[2J")
	}
}

func HandleUpdateCollected(s *state.AppState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method. Only POST is allowed.", http.StatusMethodNotAllowed)
			return
		}

		
		id := r.URL.Query().Get("id")
		if id == "" {
            utils.SendJSONResponse(w, http.StatusBadRequest, "ID is required")
			return
		}

		
		s.Mutex.Lock()
		defer s.Mutex.Unlock()

		var item *state.Item
		for i := range s.Items {
			if s.Items[i].ID == id {
				item = &s.Items[i]
				break
			}
		}

		
		if item == nil {
			utils.SendJSONResponse(w, http.StatusNotFound,"Item not found")
			return
		}

		
		if item.Collected {
			utils.SendJSONResponse(w, http.StatusConflict, "Item is already collected")
			return
		}

		
		now := time.Now()
		item.Collected = true
		item.CollectionDate = &now
		
		collectedCount := 0
		for _, itm := range s.Items {
			if itm.Collected {
				collectedCount++
			}
		}

		
		consoleUpdate := strconv.Itoa(collectedCount) + "/" + strconv.Itoa(len(s.Items)) + " Items collected"
		clearConsole()
		fmt.Println(consoleUpdate)

		
        utils.SendJSONResponse(w, http.StatusOK, fmt.Sprintf("Item ID: %s was marked as collected.", id))
	}
}