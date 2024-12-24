import paho.mqtt.client as mqtt
import json
import time
import random
from datetime import datetime

class MQTTTester:
    def __init__(self, broker="192.168.1.24", port=1883, topic="polines/data/sensor1"):
        self.broker = broker
        self.port = port
        self.topic = topic
        self.client = mqtt.Client()
        self.client.username_pw_set("jurnal123", "jurnal123")

    def connect(self):
        try:
            self.client.connect(self.broker, self.port, 60)
            print(f"Connected to MQTT broker at {self.broker}:{self.port}")
        except Exception as e:
            print(f"Failed to connect to MQTT broker: {e}")
            return False
        return True

    def generate_sample_data(self):
        """Generate realistic sample data for water quality"""
        data = {
            "turbidity": round(random.uniform(0.5, 10.0), 2),  # NTU
            "ph": round(random.uniform(6.0, 9.0), 2),          # pH scale
            "temperature": round(random.uniform(20.0, 35.0), 2),  # Celsius
            "timestamp": datetime.now().isoformat()
        }
        return data

    def serialize_data(self, data):
        """Convert all int64 values to int for JSON serialization"""
        for key, value in data.items():
            if isinstance(value, (int, int)):
                data[key] = int(value)  # Convert int64 to int
        return data

    def send_data(self, data):
        """Send data to MQTT broker"""
        try:
            # Convert data to serializable format
            data = self.serialize_data(data)
            message = json.dumps(data)
            result = self.client.publish(self.topic, message)
            if result.rc == 0:
                print(f"Published: {message}")
            else:
                print(f"Failed to publish message")
        except Exception as e:
            print(f"Error sending data: {e}")

    def run_test(self, num_samples=10, delay=1):
        """Run test by sending multiple samples"""
        if not self.connect():
            return

        print(f"\nStarting test: Sending {num_samples} samples with {delay}s delay")
        print("-" * 50)

        try:
            for i in range(num_samples):
                data = self.generate_sample_data()
                self.send_data(data)
                time.sleep(delay)

        except KeyboardInterrupt:
            print("\nTest interrupted by user")
        finally:
            self.client.disconnect()
            print("\nTest completed. MQTT connection closed.")

def main():
    # Konfigurasi tester
    tester = MQTTTester(
        broker="192.168.0.104",    # Sesuaikan dengan broker Anda
        port=1883,                 # Port default MQTT
        topic="polines/data/sensor1"  # Sesuaikan dengan topic Anda
    )

    # Jalankan test dengan 10 sampel, delay 2 detik
    tester.run_test(num_samples=10, delay=2)

if __name__ == "__main__":
    main()
