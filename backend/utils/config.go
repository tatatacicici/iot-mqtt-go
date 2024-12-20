package backend

type Config struct {
	MQTTBROKER    string
	MQTTTOPIC     string
	WebSocketPort int
}

func LoadConfig() Config {
	return Config{
		MQTTBROKER:    "tcp://192.168.0.108:1883",
		MQTTTOPIC:     "polines/data/#",
		WebSocketPort: 8080,
	}
}
