package api

import (
	"encoding/json"
	"net/http"

	"github.com/user/go-calculate/internal/calculator"
)

// Handler maneja las solicitudes HTTP
func Handler(w http.ResponseWriter, r *http.Request) {
	// Decodificar la solicitud
	var req calculator.Operation
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validar la operación
	if !calculator.IsValidOperation(req.Operation) {
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	// Realizar el cálculo
	result, err := calculator.Calculate(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Enviar la respuesta
	resp := calculator.Result{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
