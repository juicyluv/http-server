// Get order button
const orderBtn = document.querySelector('#order-btn');
// Get travel id from attributes
const travelIdString = document.querySelector('.travel-container').id;
const travelId = travelIdString.split('-')[1];

// Order travel
orderBtn.addEventListener('click', async (e) => {
    const userId = localStorage.getItem("user_id");
    const url = `http://localhost:3000/api/v1/users/${userId}/travels/${travelId}`

    try {
        const res = await fetch(url, {
            method: "POST"
        });

        // If an error occurred, show error message
        if (!res.ok) {
            const error = await res.json();
            alert(error.error);
            return;
        }

        orderBtn.value = "Отказаться";
        alert("Путешествие заказано");
    } catch (error) {
        console.log(error);
    }  
})