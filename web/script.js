// Atualizar as métricas a cada 3 segundos
setInterval(fetchMetrics, 3000);

function fetchMetrics() {
    fetch('/metrics')
        .then(response => response.json())
        .then(data => {
            // Atualiza os gráficos
            updateChart('l1-chart', data.l1);
            updateChart('l2-chart', data.l2);

            // Atualiza o status do cache
            updateCacheStatus(data);
        });
}

function updateChart(chartId, data) {
    const ctx = document.getElementById(chartId).getContext('2d');
    const chart = new Chart(ctx, {
        type: 'bar',
        data: {
            labels: ['Hits', 'Misses'],
            datasets: [{
                label: 'Cache Metrics',
                data: [data.hits, data.misses],
                backgroundColor: ['green', 'red'],
                borderColor: ['darkgreen', 'darkred'],
                borderWidth: 1
            }]
        },
        options: {
            responsive: true,
            scales: {
                y: {
                    beginAtZero: true
                }
            }
        }
    });
}

function updateCacheStatus(data) {
    const statusContainer = document.getElementById('cache-status');
    statusContainer.innerHTML = `
        <p><strong>L1 Cache:</strong> Hits: ${data.l1.hits}, Misses: ${data.l1.misses}</p>
        <p><strong>L2 Cache:</strong> Hits: ${data.l2.hits}, Misses: ${data.l2.misses}</p>
    `;
}
