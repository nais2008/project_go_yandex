package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nais2008/project_go_yandex/internal/calc"
)

// CalculateRequest ...
type CalculateRequest struct {
	Expression string `json:"expression"`
}
// CalculateResponse ...
type CalculateResponse struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

// CalculateHandler ...
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalculateRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.Expression == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(CalculateResponse{Error: "Expression is not valid"})
		return
	}

	result, err := calculator.Calc(req.Expression)

	if err != nil {
		if err.Error() == "invalid character in expression" || err.Error() == "invalid expression" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(CalculateResponse{Error: "Expression is not valid"})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(CalculateResponse{Error: "Internal server error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CalculateResponse{Result: result})
}
