const API_URL = "/users";

const form = document.getElementById("userForm");
const tableBody = document.getElementById("usersTable");
const message = document.getElementById("message");
const refreshBtn = document.getElementById("refreshBtn");

async function loadUsers() {
    tableBody.innerHTML = `
        <tr>
            <td colspan="4">Loading...</td>
        </tr>
    `;

    try {
        const response = await fetch(API_URL);

        if (!response.ok) {
            throw new Error("Failed to fetch users.");
        }

        const users = await response.json();

        tableBody.innerHTML = "";

        if (!users.length) {
            tableBody.innerHTML = `
                <tr>
                    <td colspan="4">No users found.</td>
                </tr>
            `;
            return;
        }

        users.forEach(user => {
            const row = document.createElement("tr");

            row.innerHTML = `
        <td>${user.id ?? "-"}</td>
        <td>${user.firstName}</td>
        <td>${user.lastName}</td>
        <td>${user.birthDate}</td>
    `;

            tableBody.appendChild(row);
        });

    } catch (err) {
        tableBody.innerHTML = `
            <tr>
                <td colspan="4">${err.message}</td>
            </tr>
        `;
    }
}

form.addEventListener("submit", async (e) => {
    e.preventDefault();

    message.textContent = "";
    message.className = "";

    const user = {
        firstName: document.getElementById("firstName").value.trim(),
        lastName: document.getElementById("lastName").value.trim(),
        birthDate: `${document.getElementById("birthDate").value}T00:00:00Z`
    };

    try {
        const response = await fetch(API_URL, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(user)
        });

        if (!response.ok) {
            throw new Error("Failed to add user.");
        }

        message.textContent = "User successfully added.";
        message.classList.add("success");

        form.reset();

        loadUsers();

    } catch (err) {
        message.textContent = err.message;
        message.classList.add("error");
    }
});

refreshBtn.addEventListener("click", loadUsers);

loadUsers();

async function deleteUser(id) {
    try {
        const response = await fetch(API_URL, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ id })
        });

        if (!response.ok) {
            throw new Error(await response.text());
        }

        message.textContent = "User deleted.";
        message.className = "success";

        loadUsers();
    } catch (err) {
        message.textContent = err.message;
        message.className = "error";
    }
}



async function updateUser(user) {
    try {
        const response = await fetch(API_URL, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(user)
        });

        if (!response.ok) {
            throw new Error(await response.text());
        }

        message.textContent = "User updated.";
        message.className = "success";

        loadUsers();
    } catch (err) {
        message.textContent = err.message;
        message.className = "error";
    }
}



async function findUsersByName(firstName, lastName) {
    try {
        const response = await fetch(API_URL + "/search", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                firstName,
                lastName
            })
        });

        if (!response.ok) {
            throw new Error(await response.text());
        }

        const users = await response.json();

        tableBody.innerHTML = "";

        if (users.length === 0) {
            tableBody.innerHTML = `
                <tr>
                    <td colspan="4">No users found.</td>
                </tr>
            `;
            return;
        }

        users.forEach(user => {
            const row = document.createElement("tr");

            row.innerHTML = `
                <td>${user.id ?? "-"}</td>
                <td>${user.firstName}</td>
                <td>${user.lastName}</td>
                <td>${user.birthDate}</td>
            `;

            tableBody.appendChild(row);
        });

    } catch (err) {
        message.textContent = err.message;
        message.className = "error";
    }
}