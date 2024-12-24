import pickle
import numpy as np  # Jika Anda menggunakan NumPy untuk data input

def predict_with_model(model_path, input_data):
    """
    Melakukan prediksi menggunakan model machine learning.

    Args:
      model_path (str): Path ke file model (.pkl).
      input_data: Data input untuk prediksi.

    Returns:
      Prediksi dari model.
    """
    try:
        # Load model
        with open(model_path, "rb") as f:
            model = pickle.load(f)

        # --- Melakukan Prediksi ---
        # Pastikan input_data memiliki format yang benar (misalnya, array NumPy)
        # Jika perlu, ubah input_data menjadi array NumPy:
        # input_data = np.array(input_data)
        prediction = model.predict(input_data)

        # --- Menampilkan Hasil Prediksi ---
        print("Hasil Klasifikasi:")
        # Jika model mengembalikan label kelas:
        print(f"  Label Kelas: {prediction}")
        # Jika model mengembalikan probabilitas:
        # probabilities = model.predict_proba(input_data)
        # print(f"  Probabilitas: {probabilities}")

        return prediction

    except Exception as e:
        print(f"Error: {e}")

if __name__ == "__main__":
    model_path = "xgboost_water_quality_3labels.pkl"  # Ganti dengan path model Anda
    # Ganti dengan data input yang sesuai
    input_data = [[6.5, 7.2, 1500, 25.5]]
    predict_with_model(model_path, input_data)