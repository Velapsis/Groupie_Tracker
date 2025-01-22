document.addEventListener('DOMContentLoaded', function() {
    document.querySelectorAll('.date').forEach(date => {
        date.style.display = 'none';
    });

    const locationLinks = document.querySelectorAll('.location a');
    
    locationLinks.forEach(link => {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            const dateRow = this.closest('table').querySelector('.date');
            const arrow = this.querySelector('.arrow');
            
            if (dateRow.style.display === 'none') {
                dateRow.style.display = 'table-row';
                arrow.innerHTML = '&#9650;';
            } else {
                dateRow.style.display = 'none';
                arrow.innerHTML = '&#9660;';
            }
        });
    });
});