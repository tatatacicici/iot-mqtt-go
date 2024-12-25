//WebSocket URL
const wsUrl = "ws://localhost:8080/ws";
// Inisiasi Websocket
const socket = new WebSocket(wsUrl);
// Status element
const dataItemsContainer = document.getElementById('data-items');

// Fungsi untuk menampilkan notifikasi
function showNotification(message, type = "info") {
    Toastify({
        text: message,
        duration: 3000,
        gravity: "top",
        position: "right",
        style: {
            background: type === "error" ? "rgba(220, 38, 38, 0.8)" : "rgba(5, 150, 105, 0.8)"
        }
    }).showToast();
}
//buka koneksi Websocket
socket.onopen = () =>{
    showNotification("Terhubung ke server");
};

// Variabel untuk menyimpan data terakhir
let lastData = {
    turbidity: "-",
    ph: "-",
    conductivity: "-",
    prediction: "-",
    confidence: "-"
};

// Fungsi untuk menampilkan data
function displayData(data) {
    // Update lastData dengan data baru, atau gunakan lastData jika data baru tidak ada
    lastData = { ...lastData, ...data };

    dataItemsContainer.innerHTML = "";

    const turbidityItem = document.createElement('div');
    turbidityItem.className = "data-item";
    turbidityItem.textContent = `Turbidity: ${lastData.turbidity}`;

    const phItem = document.createElement('div');
    phItem.className = "data-item";
    phItem.textContent = `PH: ${lastData.ph}`;

    const conductivityItem = document.createElement('div');
    conductivityItem.className = "data-item";
    conductivityItem.textContent = `Conductivity: ${lastData.conductivity}`;

    let predictionText;
    switch (lastData.prediction) {
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
    confidenceItem.textContent = `Confidence: ${(lastData.confidence * 100).toFixed(2)}%`;

    dataItemsContainer.appendChild(turbidityItem);
    dataItemsContainer.appendChild(phItem);
    dataItemsContainer.appendChild(conductivityItem);
    dataItemsContainer.appendChild(predictionItem);
    dataItemsContainer.appendChild(confidenceItem);
}

//menerima pesan
socket.onmessage = (event) => {
    try {
        const data = JSON.parse(event.data);

        // Tampilkan data
        displayData(data);

        // Update charts
        updateChart(phChart, new Date().toLocaleTimeString(), data.ph);
        updateChart(turbidityChart, new Date().toLocaleTimeString(), data.turbidity);
        updateChart(temperatureChart, new Date().toLocaleTimeString(), data.temperature);
        updateChart(conductivityChart, new Date().toLocaleTimeString(), data.conductivity);

    }catch (error){
        console.log("Error parsing message:", error);
    }
};

//websocket saat error
socket.onerror = () =>{
    showNotification("Terjadi kesalahan saat terhubung ke server", "error");
};

//websocket tutup
socket.onclose = () =>{
    showNotification("Koneksi ke server putus", "error");

    // Tampilkan data terakhir
    displayData({});
};