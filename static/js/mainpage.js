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
    
document.addEventListener("DOMContentLoaded", () => {
    const button = document.getElementById('back-to-top');

    if (button) {
        window.addEventListener('scroll', function () {
            if (window.scrollY > 300) {
                button.classList.add('show');
            } else {
                button.classList.remove('show');
            }
        });

        button.addEventListener('click', function () {
            window.scrollTo({
                top: 0,
                behavior: 'smooth'
            });
        });
    } else {
        console.error("Element with ID 'back-to-top' not found.");
    }
});



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


    
       
        const creationDateMin = document.getElementById('creationDateMin');
        const creationDateMax = document.getElementById('creationDateMax');
        const albumDateMin = document.getElementById('albumDateMin');
        const albumDateMax = document.getElementById('albumDateMax');
        const locationFilter = document.getElementById('locationFilter');
        const memberCheckboxes = document.querySelectorAll('#memberFilter input[type="checkbox"]');

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

        function getActiveFilters() {
            const filters = {
                query: searchInput.value.toLowerCase().trim(),
                creationDateMin: creationDateMin.value,
                creationDateMax: creationDateMax.value,
                albumDateMin: albumDateMin.value,
                albumDateMax: albumDateMax.value,
                location: locationFilter.value.trim(),
                members: Array.from(memberCheckboxes)
                    .filter(cb => cb.checked)
                    .map(cb => cb.value)
            };
            return filters;
        }

        async function performSearch() {
            const filters = getActiveFilters();
            
            try {
                const queryString = new URLSearchParams({
                    query: filters.query,
                    creationDateMin: filters.creationDateMin,
                    creationDateMax: filters.creationDateMax,
                    albumDateMin: filters.albumDateMin,
                    albumDateMax: filters.albumDateMax,
                    location: filters.location,
                    members: filters.members.join(',')
                }).toString();

                const response = await fetch(`/search?${queryString}`, {
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

        // Attachez l'événement à tous les filtres
        
        
        searchInput.addEventListener('input', debouncedSearch);
        creationDateMin.addEventListener('input', debouncedSearch);
        creationDateMax.addEventListener('input', debouncedSearch);
        albumDateMin.addEventListener('input', debouncedSearch);
        albumDateMax.addEventListener('input', debouncedSearch);
        locationFilter.addEventListener('input', debouncedSearch);
        memberCheckboxes.forEach(cb => cb.addEventListener('change', debouncedSearch));
          