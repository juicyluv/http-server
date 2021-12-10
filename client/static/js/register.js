async function sendForm(event) {
    event.preventDefault();
    const email = document.querySelector('#email').value;
    const username = document.querySelector('#username').value;
    const password = document.querySelector('#password').value;

    const data = {
        email,
        username,
        password
    }

    try {
        await fetch('http://localhost:3000/auth/sign-up', {
            method: "POST",
            body: JSON.stringify(data)
        });

        location.replace("http://localhost:3000/");
    } catch (error) {
        alert(error)
    }
    return false;
}