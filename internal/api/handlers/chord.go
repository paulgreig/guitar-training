package handlers

import (
	"encoding/json"
	"net/http"
)

// ChordHandler handles chord-related API requests
func ChordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// TODO: Implement GET /chords
		response := map[string]interface{}{
			"message": "Chords endpoint - not yet implemented",
		}
		json.NewEncoder(w).Encode(response)
	case http.MethodPost:
		// TODO: Implement POST /chords
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
