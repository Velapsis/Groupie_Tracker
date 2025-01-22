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
