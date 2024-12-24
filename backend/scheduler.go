package backend

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"
	"time"
)

func StartPredictionLoop() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		go fetchAndBroadcastPrediction()
	}
}

func fetchAndBroadcastPrediction() {
	cmd := exec.Command("D:\\Semester 5\\Tugas\\mqtt-go\\venv\\Scripts\\python.exe", "ai_inference.py")
	cmd.Dir = "D:\\Semester 5\\Tugas\\mqtt-go\\backend" // Ganti dengan path Anda

	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error running Python script: %v", err)
		return
	}

	// Bersihkan output
	outputString := strings.TrimSpace(string(output))

	// Print output untuk debugging
	log.Printf("Python output: %s", outputString)

	var predictionData map[string]interface{}
	if err := json.Unmarshal(output, &predictionData); err != nil {
		log.Printf("Error parsing Python script output: %v", err)
		return
	}

	Broadcast(predictionData)
}
