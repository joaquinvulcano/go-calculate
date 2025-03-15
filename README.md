# Calculadora API en Go

Una API HTTP y CLI en Go para realizar operaciones matemáticas básicas. El proyecto está estructurado siguiendo las mejores prácticas de Go y proporciona tanto una interfaz HTTP como una interfaz de línea de comandos.

## Características

- **API RESTful**
  - Endpoint `/calculate` para operaciones matemáticas
  - Formato JSON para solicitudes y respuestas
  - Manejo de errores robusto
  - Logging estructurado con `slog`
  - Middleware para request ID
  - Soporte CORS

- **CLI Interactiva**
  - Interfaz de línea de comandos amigable
  - Soporte para símbolos matemáticos estándar (+, -, *, /)
  - Validación de entrada en tiempo real
  - Mensajes de error claros

- **Operaciones Soportadas**
  - Suma (+)
  - Resta (-)
  - Multiplicación (*)
  - División (/) con validación de división por cero

## Estructura del Proyecto

```
go-calculate/
├── cmd/
│   └── go-calculate/    # Punto de entrada principal
├── internal/
│   ├── api/            # Handlers HTTP
│   ├── calculator/     # Lógica de cálculo
│   └── cli/           # Interfaz de línea de comandos
```

## Requisitos

- Go 1.21 o superior
- Dependencias:
  - `github.com/google/uuid`
  - `github.com/rs/cors`

## Instalación

1. Clona el repositorio:
   ```bash
   git clone https://github.com/joaquinvulcano/go-calculate.git
   ```

2. Navega al directorio del proyecto:
   ```bash
   cd go-calculate
   ```

3. Instala las dependencias:
   ```bash
   go mod tidy
   ```

## Uso

### Modo API

1. Inicia el servidor:
   ```bash
   go run ./cmd/go-calculate
   ```

2. Realiza solicitudes POST al endpoint `/calculate`:
   ```bash
   curl -X POST http://localhost:8080/calculate \
   -H "Content-Type: application/json" \
   -d '{"operation": "add", "num1": 5, "num2": 10}'
   ```

   Ejemplo de respuesta:
   ```json
   {
     "result": 15
   }
   ```

### Modo CLI

1. Ejecuta la aplicación en modo CLI:
   ```bash
   go run ./cmd/go-calculate -cli
   ```

2. Ingresa operaciones usando símbolos matemáticos estándar:
   ```
   5 + 3
   10 - 4
   6 * 8
   15 / 3
   ```

3. Para salir, escribe `exit`

## Ejemplos de Operaciones

- Suma: `5 + 3` → `8`
- Resta: `10 - 4` → `6`
- Multiplicación: `6 * 8` → `48`
- División: `15 / 3` → `5`

## Manejo de Errores

- División por cero: Error apropiado
- Operación inválida: Mensaje descriptivo
- Formato JSON inválido: Error de validación
- Números inválidos: Mensaje de error claro

## Desarrollo

### Estructura de Paquetes

- `calculator`: Lógica central de cálculo
- `api`: Manejo de solicitudes HTTP
- `cli`: Interfaz de línea de comandos

### Compilación

```bash
go build ./cmd/go-calculate
```

## Contribución

1. Haz un fork del proyecto
2. Crea una rama para tu feature (`git checkout -b feature/nueva-funcionalidad`)
3. Realiza tus cambios y haz commit (`git commit -am 'Agrega nueva funcionalidad'`)
4. Push a la rama (`git push origin feature/nueva-funcionalidad`)
5. Crea un Pull Request

## Licencia

Este proyecto está licenciado bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.
