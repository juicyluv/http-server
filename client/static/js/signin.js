async function sendForm(event) {
    event.preventDefault();
    // Get input values
    const email = document.querySelector('#email').value;
    const password = document.querySelector('#password').value;

    // Prepare data to send
    const data = {
        email,
        password
    };

    // Try to sign in
    try {
        const res = await fetch('http://localhost:3000/auth/sign-in', {
            method: "POST",
            body: JSON.stringify(data)
        });

        // If an error occurred, show message
        if (!res.ok) {
            error = await res.json();
            alert(error.error);
            return;
        }

        // Decode response to JSON
        const obj = await res.json();
        
        // Save user info
        localStorage.setItem("user_id", obj.id);
        localStorage.setItem("email", obj.email);
        localStorage.setItem("username", obj.username);
        localStorage.setItem("role", obj.role);

        // Go to index page
        location.replace("http://localhost:3000/");
    } catch (error) {
        console.log(error)
    }

    return false;
}