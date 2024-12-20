package backend

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type SensorData struct {
	Turbidity    float64 `json:"turbidity"`
	PH           float64 `json:"ph"`
	Conductivity float64 `json:"conductivity"`
}

func ConnectAndSubscribe(broker string, topic string) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("tubes-mqtt-client")
	opts.SetUsername("jurnal123")
	opts.SetPassword("jurnal123")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Gagal terhubung ke broker MQTT:%v", token.Error())
	}
	log.Println("Berhasil terhubung ke broker MQTT")

	if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Pesan diterima : %s", msg.Payload())

		var data SensorData
		if err := json.Unmarshal(msg.Payload(), &data); err != nil {
			log.Printf("Gagal mem-parsing data sensor: %v", err)
			return
		}

		data.Conductivity = Alpha*data.Turbidity + Beta*data.PH + Gamma

		Broadcast(data)

		SaveSensorData(data)
	}); token.Wait() && token.Error() != nil {
		log.Fatalf("Gagal berlangganan ke topik: %v", token.Error())
	}
}
