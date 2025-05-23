document.addEventListener("DOMContentLoaded", () => {
    fetch("/api/current_user")
        .then(res => {
            if (!res.ok) throw new Error("Not authenticated");
            return res.json();
        })
        .then(user => {
            if (user && user.email) {
                showUserMenu(user);
                fillProfile(user);
            }
        })
        .catch(err => console.error("Session check failed:", err));

    const signupForm = document.querySelector('.signup form');
    if (signupForm) {
        signupForm.addEventListener('submit', async e => {
            e.preventDefault();
            const resp = await fetch('/authorisation', {
                method: 'POST',
                body: new FormData(e.target)
            });

            if (resp.redirected) {
                window.location.href = resp.url;
                return;
            }

            const data = await resp.json();
            if (resp.ok) {
                window.location.replace('/html/profile.html');
            } else {
                alert(data.error);
            }
        });
    }

    const loginForm = document.querySelector('.login form');
    if (loginForm) {
        loginForm.addEventListener('submit', async e => {
            e.preventDefault();
            const resp = await fetch('/login', {
                method: 'POST',
                body: new FormData(e.target)
            });

            if (resp.redirected) {
                window.location.href = resp.url;
                return;
            }

            const data = await resp.json();
            if (resp.ok) {
                window.location.replace('/html/profile.html');
            } else {
                alert(data.error || "Login failed");
            }
        });
    }
});

function showUserMenu(user) {
    const authBtn = document.querySelector(".auth-btn");
    if (authBtn) authBtn.remove();

    const header = document.querySelector("header");
    const profileContainer = document.createElement("div");
    profileContainer.className = "profile-menu";

    profileContainer.innerHTML = `
        <img src="${user.avatar || '/assets/profile.png'}" alt="Avatar" class="avatar-icon" />
        <ul class="dropdown-menu hidden">
            <li><a href="/html/profile.html">Profile</a></li>
            <li><a href="/html/orders.html">Orders</a></li>
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
        fetch("/logout", { method: "POST" })
            .then(() => location.reload());
    });
}

function fillProfile(user) {
    const nameField = document.getElementById('profile-name');
    const emailField = document.getElementById('profile-email');

    if (nameField) nameField.textContent = user.username ?? (user.first_name + ' ' + user.last_name);
    if (emailField) emailField.textContent = user.email;
}
