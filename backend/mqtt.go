package backend

import (
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ConnectAndSubscribe(broker string, topic string) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("tubes-mqtt-client")
	opts.SetUsername("jurnal123")
	opts.SetPassword("jurnal123")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Gagal terhubung ke broker MQTT: %v", token.Error())
	}
	log.Println("Berhasil terhubung ke broker MQTT")

	if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Pesan diterima: %s", msg.Payload())

		var data SensorData
		if err := json.Unmarshal(msg.Payload(), &data); err != nil {
			log.Printf("Gagal mem-parsing data sensor: %v", err)
			return
		}

		// Hitung nilai conductivity berdasarkan formula baru
		data.Conductivity = calculateConductivity(data.Turbidity, data.PH)

		Broadcast(data)
		SaveSensorData(data)
	}); token.Wait() && token.Error() != nil {
		log.Fatalf("Gagal berlangganan ke topik: %v", token.Error())
	}
}

// Tambahkan fungsi helper untuk menghitung conductivity
// Fungsi untuk menghitung nilai conductivity berdasarkan rumus terbaru
func calculateConductivity(turbidity, ph float64) float64 {
	alpha := 0.6658821991086864
	beta := 0.9580655663104417
	gamma := 416.77983298775104

	return alpha*turbidity + beta*ph + gamma
}
