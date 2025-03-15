# Calculadora API en Go

Este proyecto implementa una API HTTP en Go para realizar operaciones matemáticas básicas, con soporte para suma, resta, multiplicación y división.

## Características

- API RESTful con endpoint `/calculate`
- Soporte para operaciones básicas: suma, resta, multiplicación y división
- Modo CLI interactivo
- Validación de entradas
- Manejo de errores
- Registro estructurado de logs
- Configuración de CORS
- Autenticación JWT (próximamente)

## Requisitos

- Go 1.21 o superior

## Instalación

1. Clona el repositorio:
   ```bash
   git clone https://github.com/joaquinvulcano/go-calculate.git
   ```
2. Navega al directorio del proyecto:
   ```bash
   cd go-calculate
   ```
3. Descarga las dependencias:
   ```bash
   go mod tidy
   ```

## Uso

### Modo API

1. Inicia el servidor:
   ```bash
   go run .
   ```
2. Envía solicitudes POST al endpoint `/calculate` con un cuerpo JSON:
   ```json
   {
     "operation": "add",
     "num1": 5,
     "num2": 10
   }
   ```

### Modo CLI

Ejecuta el programa con el flag `-cli`:
```bash
go run . -cli
```

Ingresa las operaciones en el formato:
```
num1 operacion num2
```

Ejemplos:
```
5 + 3
10 / 2
15 * 4
20 - 7
```

## Contribución

1. Haz un fork del proyecto
2. Crea una nueva rama (`git checkout -b feature/nueva-funcionalidad`)
3. Haz commit de tus cambios (`git commit -am 'Agrega nueva funcionalidad'`)
4. Haz push a la rama (`git push origin feature/nueva-funcionalidad`)
5. Abre un Pull Request

## Licencia

MIT
