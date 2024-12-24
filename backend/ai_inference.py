import sys
import json
import joblib
import pymongo
import numpy as np

# Load model
try:
    model = joblib.load("../model/xgboost_water_quality_3labels.pkl")
    print(json.dumps({"status": "Model loaded successfully"}), file=sys.stderr)
except Exception as e:
    print(json.dumps({"error": f"Failed to load model: {str(e)}"}), file=sys.stderr)
    sys.exit(1)

# Koneksi MongoDB
try:
    mongo_uri = "mongodb://localhost:27017"
    client = pymongo.MongoClient(mongo_uri)
    db = client["iot_data"]
    collection = db["sensor_readings"]
    print(json.dumps({"status": "Connected to MongoDB successfully"}), file=sys.stderr)
except Exception as e:
    print(json.dumps({"error": f"Failed to connect to MongoDB: {str(e)}"}), file=sys.stderr)
    sys.exit(1)

def fetch_latest_data():
    """Ambil dokumen terbaru dari MongoDB."""
    return collection.find_one(sort=[("timestamp", pymongo.DESCENDING)])

def convert_value(value):
    """Convert numpy types (int64, float64) to native Python types (int, float)."""
    if isinstance(value, np.int64):
        return int(value)  # Convert numpy.int64 to int
    elif isinstance(value, np.float64):
        return float(value)  # Convert numpy.float64 to float
    elif isinstance(value, (int, float, str, list, dict)):  # Already serializable
        return value
    elif isinstance(value, bytes):
        return value.decode("utf-8")
    # else:
    #     raise TypeError(f"Object of type {type(value)} is not JSON serializable")

def convert_data(data):
    """Convert all values in the data to serializable types."""
    return {key: convert_value(value) for key, value in data.items()}

def main():
    try:
        # Ambil data terbaru dari MongoDB
        data = fetch_latest_data()
        if not data:
            raise ValueError("No data found in MongoDB")

        # Pastikan semua field yang diperlukan tersedia
        required_fields = ['turbidity', 'ph', 'conductivity', 'temperature']
        if not all(field in data for field in required_fields):
            missing_fields = [field for field in required_fields if field not in data]
            raise ValueError(f"Missing required input fields: {', '.join(missing_fields)}")

        # Convert all fields in the data to serializable types
        data = convert_data(data)

        # Format data untuk model
        features = [[data['turbidity'], data['ph'], data['conductivity'], data['temperature']]]

        # Prediksi menggunakan model
        prediction = model.predict(features)
        confidence = model.predict_proba(features).max()
        # print(data)
        # Format output JSON
        output = {
            "prediction": int(prediction[0]),  # Convert prediction to int
            "confidence": float(confidence)  # Convert confidence to float
        }

        # Menampilkan output JSON
        print(json.dumps(output, indent=None))

    except pymongo.errors.ConnectionFailure as e:
        print(json.dumps({"error": f"MongoDB connection error: {e}"}))
        sys.exit(1)
    except KeyError as e:
        print(json.dumps({"error": f"Missing key in MongoDB data: {e}"}))
        sys.exit(1)
    except ValueError as e:
        print(json.dumps({"error": f"Invalid data in MongoDB: {e}"}))
        sys.exit(1)
    except Exception as e:
        print(json.dumps({"error": f"An unexpected error occurred: {e}"}))
        sys.exit(1)
if __name__ == "__main__":
    main()