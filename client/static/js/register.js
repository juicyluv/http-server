// Handle register form
async function sendForm(event) {
    event.preventDefault();
    // Get input values
    const email = document.querySelector('#email').value;
    const username = document.querySelector('#username').value;
    const password = document.querySelector('#password').value;

    // Prepare data to send
    const data = {
        email,
        username,
        password
    };

    // Try to register user
    try {
        const res = await fetch('http://localhost:3000/auth/sign-up', {
            method: "POST",
            body: JSON.stringify(data)
        });

        // If an error occurred, show error message
        if (!res.ok) {
            error = await res.json();
            alert(error.error);
            return;
        }

        // If success, go to sign in page
        location.replace("http://localhost:3000/sign-in");
    } catch (error) {
        console.log(error)
    }
    return false;
}