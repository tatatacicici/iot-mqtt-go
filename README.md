# 🌊 IoT Water Quality Monitoring System

> Real-time water quality monitoring backend built with **Go**, **Python (XGBoost)**, **MQTT**, **MongoDB**, and **WebSocket**.

[![Go Version](https://img.shields.io/badge/Go-1.21.4-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![Python](https://img.shields.io/badge/Python-3.x-3776AB?style=flat-square&logo=python)](https://python.org/)
[![MongoDB](https://img.shields.io/badge/MongoDB-NoSQL-47A248?style=flat-square&logo=mongodb)](https://mongodb.com/)

---

## 📖 Overview

This system collects sensor data (turbidity, pH, conductivity, temperature) from IoT devices via MQTT, persists it to MongoDB, and runs an **XGBoost ML model** every 10 seconds to classify water quality. Results are instantly pushed to frontends via **WebSocket** for real-time visualization.

---

## 🏗️ System Architecture

```
[IoT Sensor / MQTT Tester]
        │  MQTT Publish
        ▼
[MQTT Broker (Eclipse Paho)]
        │  Subscribe & Parse
        ▼
[Go Backend]
  ├── mqtt.go       → receive & parse sensor JSON, calculate conductivity
  ├── mongo.go      → persist sensor data to MongoDB
  ├── scheduler.go  → every 10s, call Python inference subprocess
  │       └──→ [ai_inference.py] XGBoost Model → prediction JSON
  └── websocket.go  → broadcast sensor data & predictions to clients
        │  WebSocket (ws://localhost:8080/ws)
        ▼
[Frontend Dashboard (Chart.js)]
  📊 pH  📊 Turbidity  📊 Temperature  📊 Conductivity
```

---

## ⚙️ Tech Stack

| Layer | Technology | Purpose |
|---|---|---|
| **Backend** | Go 1.21 | Core service orchestration |
| **Messaging** | MQTT (Eclipse Paho) | IoT sensor data ingestion |
| **ML Inference** | Python + XGBoost | Water quality classification |
| **Database** | MongoDB | Time-series sensor storage |
| **Real-time** | WebSocket (Gorilla) | Live data push to frontend |
| **Frontend** | HTML + Chart.js | Real-time data visualization |

---

## 📂 Project Structure

```
iot-mqtt-go/
├── main.go                  # Entry point
├── backend/
│   ├── models.go            # SensorData struct
│   ├── mqtt.go              # MQTT subscribe + conductivity formula
│   ├── mongo.go             # MongoDB connection & storage
│   ├── websocket.go         # WebSocket server & broadcast hub
│   ├── scheduler.go         # 10s ticker → Python subprocess call
│   ├── ai_inference.py      # XGBoost inference: MongoDB → predict → JSON stdout
│   └── mqtt_tester.py       # Sensor data simulator for local testing
├── model/
│   └── xgboost_water_quality_3labels.pkl
└── frontend/
    └── index.html           # Real-time dashboard
```

---

## 🔬 Sensor Data & ML Model

### MQTT Message Payload

```json
{
  "turbidity": 12.5,
  "ph": 7.2,
  "conductivity": 450.0,
  "temperature": 28.3
}
```

> `conductivity` is auto-calculated using: `0.6659 × turbidity + 0.9581 × pH + 416.78`

### ML Prediction Output (every 10s)

```json
{ "prediction": 1, "confidence": 0.87 }
```

| Label | Water Quality |
|---|---|
| `0` | Baik (Good) |
| `1` | Sedang (Moderate) |
| `2` | Buruk (Poor) |

---

## 📡 WebSocket Endpoint

**`ws://localhost:8080/ws`** — broadcasts sensor readings and ML predictions in real-time.

---

## 🚀 Getting Started

```bash
# Clone
git clone https://github.com/tatatacicici/iot-mqtt-go.git
cd iot-mqtt-go

# Go dependencies
go mod tidy

# Python dependencies
pip install joblib pymongo numpy xgboost scikit-learn

# Run backend
go run main.go

# Simulate sensor data (separate terminal)
python backend/mqtt_tester.py
```

---

## 👤 Author

**Hussain Tamam Gucci Al Fauzan** — Odoo Developer & Backend Engineer  
[GitHub](https://github.com/tatatacicici) · [LinkedIn](https://www.linkedin.com/in/hussain-tamam-gucci-al-fauzan/)
