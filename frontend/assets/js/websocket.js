//WebSocket URL
const wsUrl = "ws://localhost:8080/ws";
// Inisasi Websocket
const socket = new WebSocket(wsUrl);
// Status element
const statusElement = document.getElementById('status');
const dataItemsContainer = document.getElementById('data-items');

// Fungsi untuk menampilkan notifikasi
function showNotification(message, type = "info") {
    Toastify({
        text: message,
        duration: 3000,
        gravity: "top",
        position: "right",
        style: {
            background: type === "error" ? "rgba(220, 38, 38, 0.8)" : "rgba(5, 150, 105, 0.8)" // Merah untuk error, hijau untuk info
        }
    }).showToast();
}
//buka koneksi Websocket
socket.onopen = () =>{
    showNotification("Terhubung ke server");
};

//menerima pesan
socket.onmessage = (event) => {
    try {
        const data = JSON.parse(event.data);

        //bersihkan data sebelumnya bila ada
        dataItemsContainer.innerHTML = "";

        //menampilkan data baru
        const turbidityItem = document.createElement('div');
        turbidityItem.className = "data-item";
        turbidityItem.textContent =`Turbidity: ${data.turbidity}`;

        const phItem = document.createElement('div');
        phItem.className = "data-item";
        phItem.textContent = `PH: ${data.ph}`;

        const conductivityItem = document.createElement('div');
        conductivityItem.className = "data-item";
        conductivityItem.textContent = `Conductivity: ${data.conductivity}`;

        // Prediction logic
        let predictionText;
        switch (data.prediction) {
            case 0:
                predictionText = "Layak Minum";
                break;
            case 1:
                predictionText = "Agak Layak Minum";
                break;
            case 2:
                predictionText = "Tidak Layak Minum";
                break;
            default:
                predictionText = "Unknown";
        }

        const predictionItem = document.createElement('div');
        predictionItem.className = "data-item prediction";
        predictionItem.textContent = `Prediction: ${predictionText}`;

        const confidenceItem = document.createElement('div');
        confidenceItem.className = "data-item confidence";
        confidenceItem.textContent = `Confidence: ${(data.confidence * 100).toFixed(2)}%`;

        dataItemsContainer.appendChild(turbidityItem);
        dataItemsContainer.appendChild(phItem);
        dataItemsContainer.appendChild(conductivityItem);
        dataItemsContainer.appendChild(predictionItem);
        dataItemsContainer.appendChild(confidenceItem);
    }catch (error){
        console.log("Error parsing message:", error);
    }
};

//websocket saat error
socket.onerror = () =>{
    showNotification("Terjadi kesalahan saat terhubung ke server", "error");
    statusElement.className = "error";
};

//websocket tutup
socket.onclose = () =>{
    showNotification("Koneksi ke server putus", "error");
    statusElement.className = "error";
};
