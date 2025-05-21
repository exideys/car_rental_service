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

document.addEventListener('DOMContentLoaded', () => {
    
    document.querySelector('.signup form')
        .addEventListener('submit', async e => {
            e.preventDefault();
            const resp = await fetch('/authorisation', {
                method: 'POST',
                body: new FormData(e.target)
            });

            if (resp.redirected) {
                return;
            }

            const data = await resp.json();
            if (resp.ok) {
                sessionStorage.setItem('user', JSON.stringify(data.user));
                window.location.replace('/html/profile.html');
            } else {
                alert(data.error);
            }
        });

});
function fillProfile(user) {
    // На любой странице, где есть эти id, заполняем значения
    const nameField  = document.getElementById('profile-name');
    const emailField = document.getElementById('profile-email');

    if (nameField)  nameField.textContent  = user.username ?? (user.first_name + ' ' + user.last_name);
    if (emailField) emailField.textContent = user.email;
}

document.addEventListener("DOMContentLoaded", () => {
    fetch("/api/auth/check_session.go")
        .then(res => res.json())
        .then(user => {
            if (user && user.loggedIn) {
                showUserMenu(user);   // как раньше
                fillProfile(user);    // ← новое
            }
        })
        .catch(err => console.error("Session check failed:", err));
});