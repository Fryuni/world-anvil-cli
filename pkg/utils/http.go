package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, map[string]string{"error": message})
}

func ValidateRequired(fields map[string]string) error {
	for field, value := range fields {
		if value == "" {
			return fmt.Errorf("%s is required", field)
		}
	}
	return nil
}
