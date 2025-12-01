package handlers

import (
	"encoding/json"
	"net/http"
	"question-answer/models"
	"strings"
)

func errorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
		Code:    statusCode,
	})
}

func ExtractID(path string, prefix string) string {
	return strings.TrimPrefix(path, prefix)
}
