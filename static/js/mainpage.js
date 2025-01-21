document.addEventListener("DOMContentLoaded", () => {
    const canvas = document.getElementById("sakuraCanvas");
    const ctx = canvas.getContext("2d");
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;

    const petals = [];
    for (let i = 0; i < 100; i++) {
        petals.push({
            x: Math.random() * canvas.width,
            y: Math.random() * canvas.height,
            radius: Math.random() * 3 + 1,
            dx: Math.random() * 2 - 1,
            dy: Math.random() * 1 + 0.5,
        });
    }

    function drawPetals() {
        ctx.clearRect(0, 0, canvas.width, canvas.height);
        petals.forEach((p) => {
            ctx.beginPath();
            ctx.arc(p.x, p.y, p.radius, 0, Math.PI * 2);
            ctx.fillStyle = "rgba(255, 183, 197, 0.8)";
            ctx.fill();
            p.x += p.dx;
            p.y += p.dy;
            if (p.y > canvas.height) p.y = 0;
            if (p.x > canvas.width || p.x < 0) p.x = Math.random() * canvas.width;
        });
    }

    function animate() {
        drawPetals();
        requestAnimationFrame(animate);
    }

    animate();
});
    


//NEEDS TO BE EDITED BY MAHAN !!!!!!





        const searchInput = document.getElementById('searchInput');
        
        function debounce(func, wait) {
            let timeout;
            return function executedFunction(...args) {
                const later = () => {
                    clearTimeout(timeout);
                    func(...args);
                };
                clearTimeout(timeout);
                timeout = setTimeout(later, wait);
            };
        }

        async function performSearch() {
    const query = searchInput.value.toLowerCase().trim(); // trim() pour enlever les espaces
    
    try {
        // Si la recherche est vide, récupérer tous les artistes
        if (query === '') {
            const response = await fetch('/index', {
                method: 'GET',
                headers: {
                    'X-Requested-With': 'XMLHttpRequest'
                }
            });
            
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            
            const artists = await response.json();
            updateArtistsList(artists);
            return;
        }

        // Sinon, effectuer la recherche normale
        const response = await fetch('/search?query=' + encodeURIComponent(query), {
            method: 'GET',
            headers: {
                'X-Requested-With': 'XMLHttpRequest'
            }
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const artists = await response.json();
        updateArtistsList(artists);
        
    } catch (error) {
        console.error('Error:', error);
    }
}
function updateArtistsList(artists) {
    console.log('Artists reçus:', artists); // Pour voir les données exactes
    
    const artistsList = document.getElementById('artistsList');
    artistsList.innerHTML = '';
    
    artists.forEach(artist => {
        console.log('Traitement artiste:', artist); // Pour voir chaque artiste
        
        const li = document.createElement('li');
        li.className = 'articles__article';
        
        // Vérification des données avant affichage
        const id = artist.Id || artist.id || 'ID manquant';
        const name = artist.Name || artist.name || 'Nom manquant';
        const image = artist.Image || artist.image || '/default-image.jpg';
        
        console.log('Données utilisées:', { id, name, image }); // Pour voir les données utilisées
        
        li.innerHTML = `
            <a href="/artist?id=${id}" class="articles__link">
                <div class="articles__content">
                    <img class="articles__img" src="${image}" alt="${name}">
                    <div class="title">${name}</div>
                </div>
            </a>
        `;
        artistsList.appendChild(li);
    });
}

        const debouncedSearch = debounce(performSearch, 300);
        searchInput.addEventListener('input', debouncedSearch);
