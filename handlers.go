package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type CalculationRequest struct {
	Operation string  `json:"operation"`
	Num1      float64 `json:"num1"`
	Num2      float64 `json:"num2"`
}

type CalculationResponse struct {
	Result float64 `json:"result"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	// Decodificar la solicitud
	var req CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validar la operación
	if !isValidOperation(req.Operation) {
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	// Realizar el cálculo
	result, err := performCalculation(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Enviar la respuesta
	resp := CalculationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func isValidOperation(op string) bool {
	switch op {
	case "add", "subtract", "multiply", "divide":
		return true
	default:
		return false
	}
}

func performCalculation(req CalculationRequest) (float64, error) {
	switch req.Operation {
	case "add":
		return req.Num1 + req.Num2, nil
	case "subtract":
		return req.Num1 - req.Num2, nil
	case "multiply":
		return req.Num1 * req.Num2, nil
	case "divide":
		if req.Num2 == 0 {
			return 0, errors.New("division by zero")
		}
		return req.Num1 / req.Num2, nil
	default:
		return 0, errors.New("unknown operation")
	}
}
