package backend

import (
	"bytes"
	"encoding/json"
	"log"
	"os/exec"
)

type AIInput struct {
	Turbidity float64 `json:"turbidity"`
	PH        float64 `json:"ph"`
}

type AIOutput struct {
	Prediction string  `json:"prediction"`
	Confidence float64 `json:"confidence"`
}

func RunInference(input AIInput) (*AIOutput, error) {
	inputData, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("python3", "ai_inference.py")
	cmd.Stdin = bytes.NewReader(inputData)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Printf("Error running inference: %v", err)
		return nil, err
	}

	var result AIOutput
	if err := json.Unmarshal(out.Bytes(), &result); err != nil {
		return nil, err
	}

	return &result, nil
}
