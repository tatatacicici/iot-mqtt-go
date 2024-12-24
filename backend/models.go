package backend

type SensorData struct {
	Turbidity    float64 `json:"turbidity"`
	PH           float64 `json:"ph"`
	Conductivity float64 `json:"conductivity"`
	Temperature  float64 `json:"temperature"`
}
