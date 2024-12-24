// Chart.js instances
const phChart = new Chart(document.getElementById('phChart'), {
    type: 'line',
    data: {
        labels: [], // Waktu atau label data
        datasets: [{
            label: 'PH',
            data: [],
            borderColor: 'rgba(75, 192, 192, 1)',
            backgroundColor: 'rgba(75, 192, 192, 0.2)',
            tension: 0.1
        }]
    },
    options: {
        responsive: true,
        maintainAspectRatio: false
    }
});

const turbidityChart = new Chart(document.getElementById('turbidityChart'), {
    type: 'line',
    data: {
        labels: [],
        datasets: [{
            label: 'Turbidity',
            data: [],
            borderColor: 'rgba(255, 99, 132, 1)',
            backgroundColor: 'rgba(255, 99, 132, 0.2)',
            tension: 0.1
        }]
    },
    options: {
        responsive: true,
        maintainAspectRatio: false
    }
});

const temperatureChart = new Chart(document.getElementById('temperatureChart'), {
    type: 'line',
    data: {
        labels: [],
        datasets: [{
            label: 'Temperature',
            data: [],
            borderColor: 'rgba(54, 162, 235, 1)',
            backgroundColor: 'rgba(54, 162, 235, 0.2)',
            tension: 0.1
        }]
    },
    options: {
        responsive: true,
        maintainAspectRatio: false
    }
});

const conductivityChart = new Chart(document.getElementById('conductivityChart'), {
    type: 'line',
    data: {
        labels: [],
        datasets: [{
            label: 'Conductivity',
            data: [],
            borderColor: 'rgba(153, 102, 255, 1)',
            backgroundColor: 'rgba(153, 102, 255, 0.2)',
            tension: 0.1
        }]
    },
    options: {
        responsive: true,
        maintainAspectRatio: false
    }
});

// Function to update charts
function updateChart(chart, label, data) {
    chart.data.labels.push(label);
    chart.data.datasets[0].data.push(data);

    if (chart.data.labels.length > 10) { // Limit data points to 10
        chart.data.labels.shift();
        chart.data.datasets[0].data.shift();
    }

    chart.update();
}
