# Weather API (Go)

API REST en Go construida con arquitectura hexagonal (Ports & Adapters).

## 🚀 Features

- Health check endpoint
- Weather por coordenadas
- Integración con API externa (Open-Meteo)
- Arquitectura desacoplada y testeable

## 🧱 Arquitectura

Estructura basada en Hexagonal Architecture:

## 📡 Endpoints

### Health check
GET /health

Response: ok

---

### Weather por coordenadas
GET /weather?lat={lat}&lon={lon}

Ejemplo: /weather?lat=-34.61&lon=-58.38


Response:
```json
{
  "temperature": 25.6,
  "description": "clear",
  "wind_speed": 12.3,
  "source": "open-meteo"
}