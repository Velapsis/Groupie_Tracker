document.addEventListener("DOMContentLoaded", () => {
    const button = document.getElementById("main-page-button"); 
    const overlay = document.getElementById("page-transition-overlay"); 

    button.addEventListener("click", (e) => {
        e.preventDefault(); 
        overlay.classList.add("active"); 

        setTimeout(() => { 
            window.location.href = button.getAttribute("href"); 
        }, 500);
    }); 
}); 
