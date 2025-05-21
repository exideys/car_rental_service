document.addEventListener("DOMContentLoaded", () => {
    fetch("/api/auth/check_session.go")
        .then(res => res.json())
        .then(user => {
            if (user && user.loggedIn) {
                showUserMenu(user);
            }
        })
        .catch(err => {
            console.error("Session check failed:", err);
        });
});

function showUserMenu(user) {
    const authBtn = document.querySelector(".auth-btn");
    if (authBtn) authBtn.remove();

    const header = document.querySelector("header");
    const profileContainer = document.createElement("div");
    profileContainer.className = "profile-menu";

    profileContainer.innerHTML = `
        <img src="${user.avatar || '/assets/img/default-avatar.png'}" alt="Avatar" class="avatar-icon" />
        <ul class="dropdown-menu hidden">
            <li><a href="/html/profile.html">My Profile</a></li>
            <li><a href="/html/orders.html">My Orders</a></li>
            <li id="logout-btn">Logout</li>
        </ul>
    `;

    header.appendChild(profileContainer);

    const avatar = profileContainer.querySelector(".avatar-icon");
    const menu = profileContainer.querySelector(".dropdown-menu");

    avatar.addEventListener("click", () => {
        menu.classList.toggle("hidden");
    });

    document.getElementById("logout-btn").addEventListener("click", () => {
        fetch("/api/auth/logout.go") 
            .then(() => location.reload());
    });
}
