import sys
import json
import joblib
from pyexpat import features

try:
    model = joblib.load("../model/xgboost_water_quality_3labels.pkl")
    print(json.dumps({"status":"Model loaded succesfully"}), file=sys.stderr)
except Exception as e:
    print(json.dumps({"error": f"Failed to load model: {str(e)}"}), file=sys.stderr)
    sys.exit(1)

def main():
    try:
        input_data = sys.stdin.read()
        data = json.loads(input_data)

        if 'turbidity' not in data or 'ph' not in data:
            raise ValueError("Missing required input fields: 'turbidity' and 'ph'")

        features = [[data['turbidity'], data['ph']]]

        prediction = model.predict(features)
        confidence = model.predict_proba(features).max()

        output = {
            "prediction": prediction[0],
            "confidence": confidence
        }
        print(json.dumps(output))

    except json.JSONDecodeError:
        print(json.dumps({"error":"Invalid JSON input"}))
    except Exception as e:
        error_response = {
            "error": str(e)
        }
        print(json.dumps(error_response))
if __name__ == "__main__":
    main()
