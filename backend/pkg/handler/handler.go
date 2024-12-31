package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rfid_handler/pkg/state"
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

func sendJSONResponse(w http.ResponseWriter, statusCode int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func HandleUpdateCollected(s *state.AppState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method. Only POST is allowed.", http.StatusMethodNotAllowed)
			return
		}

		
		id := r.URL.Query().Get("id")
		if id == "" {
            sendJSONResponse(w, http.StatusBadRequest, "ID is required")
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
			sendJSONResponse(w, http.StatusNotFound,"Item not found")
			return
		}

		
		if item.Collected {
			sendJSONResponse(w, http.StatusConflict, "Item is already collected")
			return
		}

		
		item.Collected = true
		item.CollectionDate = time.Now()
		
		collectedCount := 0
		for _, itm := range s.Items {
			if itm.Collected {
				collectedCount++
			}
		}

		
		consoleUpdate := strconv.Itoa(collectedCount) + "/" + strconv.Itoa(len(s.Items)) + " Items collected"
		clearConsole()
		fmt.Println(consoleUpdate)

		
        sendJSONResponse(w, http.StatusOK, fmt.Sprintf("Item ID: %s was marked as collected.", id))
	}
}
