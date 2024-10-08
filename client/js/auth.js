document.getElementById('loginForm').addEventListener('submit', function(event){
    event.preventDefault();

    const email = document.getElementById('inputEmail').value;
    const password = document.getElementById('inputPassword').value;

    const obj = {
        'email': email,
        'password': password
    };

    fetch('http://localhost:8081/api-putra-jaya/auth/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(obj)
    })
    .then(response => response.json())
    .then(data => {
        const token = data.data.token_access;
        const role = data.role;
        console.log(data);
        console.log(token);
        localStorage.setItem('token', token);
        localStorage.setItem('role', role);

        if (role === "OWNER") {
            window.location.href = '../admin.html';
        } else if (role === "CUSTOMER") {
            window.location.href = '../index.html';
        } else {
            alert('You don\'t have access to the dashboard');
        }
    })
    .catch(error => {
        console.error(error);
        alert('Email atau Password Salah');
    });
});

function logout(){
    localStorage.removeItem('token');
    localStorage.removeItem('role');
    window.location.href = "index.html";
}
