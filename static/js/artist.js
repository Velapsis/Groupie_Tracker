document.addEventListener('DOMContentLoaded', async function() {
    // Initialize map
    const map = L.map('concertMap').setView([0, 0], 2);
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© OpenStreetMap contributors'
    }).addTo(map);

    const markers = [];
    const locationLinks = document.querySelectorAll('.location a');
    
    // Hide all dates initially and set explicit display style
    document.querySelectorAll('.date').forEach(date => {
        date.style.display = 'none';
    });

    // Geocode all locations and add markers
    for (const link of locationLinks) {
        const location = link.textContent.trim().slice(2); // Remove arrow character
        try {
            const response = await fetch(`/geocode?location=${encodeURIComponent(location)}`);
            const data = await response.json();
            
            if (data && data[0]) {
                const marker = L.marker([parseFloat(data[0].lat), parseFloat(data[0].lon)])
                    .bindPopup(location)
                    .addTo(map);
                markers.push(marker);
            }
        } catch (error) {
            console.error(`Error geocoding ${location}:`, error);
        }
    }

    // Add click handlers for showing/hiding dates
    locationLinks.forEach(link => {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            const dateRow = this.closest('table').querySelector('.date');
            const arrow = this.querySelector('.arrow');
            const currentDisplay = window.getComputedStyle(dateRow).display;
            
            if (currentDisplay === 'none') {
                dateRow.style.display = 'table-row';
                arrow.innerHTML = '&#9650;';
            } else {
                dateRow.style.display = 'none';
                arrow.innerHTML = '&#9660;';
            }
        });
    });
});