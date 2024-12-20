package backend

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	mongoURI        = "mongodb://localhost:27017"
	databaseName    = "iot_data"
	collectionName  = "sensor_readings"
	mongoCollection *mongo.Collection
	Alpha           = 0.6658821991086864
	Beta            = 0.9580655663104417
	Gamma           = 416.77983298775104
)

func InitMongoDB() {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Gagal terhubung ke MongoDB: %v", err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal("MongoDB tidak merespon: %v", err)
	}

	log.Println("Berhasil terhubung ke MongoDB")
	mongoCollection = client.Database(databaseName).Collection(collectionName)
}

func SaveSensorData(data SensorData) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conductivity := Alpha*data.Turbidity + Beta*data.PH + Gamma*data.PH

	document := bson.M{
		"turbidity":    data.Turbidity,
		"ph":           data.PH,
		"conductivity": conductivity,
		"timestamp":    time.Now(),
	}

	_, err := mongoCollection.InsertOne(ctx, document)
	if err != nil {
		log.Printf("Gagal menyimpan data sensor ke MongoDB: %v", err)
		return
	}
	log.Println("Data sensor berhasil disimpan ke MongoDB")
}
