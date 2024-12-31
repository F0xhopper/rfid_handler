package utils

import (
	"encoding/json"
	"net/http"
)


func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	switch v := data.(type) {
	case string:
		data = map[string]string{"message": v}
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}