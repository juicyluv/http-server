const refuseBtns = document.querySelectorAll('.order-price__btn');

refuseBtns.forEach((btn) => {
    // Delete user travel from orders
    btn.addEventListener('click', async e => {
        const agree = confirm("Вы действительно хотите отказаться от путешествия?");
        if (!agree) {
            return;
        }
        
        const travelId = btn.id.split('-')[1];
        const url = `http://localhost:3000/api/v1/users/travels/${travelId}`
        try {
            const res = await fetch(url, {
                method: "DELETE"
            });

            // If error occured
            if (!res.ok) {
                const error = await res.json();
                alert(error.error);
                return
            }

            // If success, delete the travel
            const travelToDelete = document.querySelector(`#travel-${travelId}`);
            travelToDelete.parentNode.removeChild(travelToDelete);
        } catch (error) {
            console.log(error)
        }  
    });
});