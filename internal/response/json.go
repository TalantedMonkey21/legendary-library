package response

import (
	"encoding/json"
	"net/http"
)

// вспомогательные функции для хэндлера
func WriteJSONError(w http.ResponseWriter, code int, e string) {
	WriteJSONResponse(w, code, map[string]string{"error": e})
}

func WriteJSONResponse(w http.ResponseWriter, code int, s any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(s)
}
