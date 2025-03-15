package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/user/go-calculate/internal/calculator"
)

// Start inicia la interfaz de línea de comandos
func Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Ingrese la operación (formato: num1 operacion num2): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Formato inválido. Use: num1 operacion num2")
			continue
		}

		num1, err1 := strconv.ParseFloat(parts[0], 64)
		operation := mapOperation(parts[1])
		num2, err2 := strconv.ParseFloat(parts[2], 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Números inválidos")
			continue
		}

		result, err := calculator.Calculate(calculator.Operation{
			Operation: operation,
			Num1:      num1,
			Num2:      num2,
		})

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Resultado: %.2f\n", result)
		}
	}
}

// mapOperation convierte símbolos matemáticos a operaciones
func mapOperation(op string) string {
	switch op {
	case "+":
		return "add"
	case "-":
		return "subtract"
	case "*":
		return "multiply"
	case "/":
		return "divide"
	default:
		return op
	}
}
