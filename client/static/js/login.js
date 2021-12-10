async function sendForm(event) {
    event.preventDefault();
    const email = document.querySelector('#email').value;
    const password = document.querySelector('#password').value;

    const data = {
        email,
        password
    }

    try {
        alert(data)

        const res = await fetch('http://localhost:3000/auth/sign-in', {
            method: "POST",
            body: JSON.stringify(data)
        });

        console.log(res)

        // location.replace("http://localhost:3000/");
    } catch (error) {
        alert(error);
    }
    return false;
}