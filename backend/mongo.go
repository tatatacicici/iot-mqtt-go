package backend

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI        = "mongodb://localhost:27017"
	databaseName    = "iot_data"
	collectionName  = "sensor_readings"
	mongoCollection *mongo.Collection
)

func InitMongoDB() {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Gagal terhubung ke MongoDB: %v", err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("MongoDB tidak merespon: %v", err)
	}

	log.Println("Berhasil terhubung ke MongoDB")
	mongoCollection = client.Database(databaseName).Collection(collectionName)
}

func SaveSensorData(data SensorData) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	document := bson.M{
		"turbidity":    data.Turbidity,
		"ph":           data.PH,
		"conductivity": data.Conductivity,
		"timestamp":    time.Now(),
	}

	_, err := mongoCollection.InsertOne(ctx, document)
	if err != nil {
		log.Printf("Gagal menyimpan data sensor ke MongoDB: %v", err)
		return
	}
	log.Println("Data sensor berhasil disimpan ke MongoDB")
}
