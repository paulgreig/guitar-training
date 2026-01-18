package handlers

import (
	"encoding/json"
	"net/http"
)

// ScaleHandler handles scale-related API requests
func ScaleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// TODO: Implement GET /scales
		response := map[string]interface{}{
			"message": "Scales endpoint - not yet implemented",
		}
		json.NewEncoder(w).Encode(response)
	case http.MethodPost:
		// TODO: Implement POST /scales
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
