const title = document.querySelector('#title')
const file = document.querySelector('#image');
const durationDays = document.querySelector('#duration_days');
const partySize = document.querySelector('#party_size');
const description = document.querySelector('#description');
const date = document.querySelector('#date');
const place = document.querySelector('#place');
const complexity = document.querySelector('#complexity');
const price = document.querySelector('#price');

async function sendForm(e) {
    e.preventDefault();
    const data = new FormData();
    data.append("travel_image", file.files[0]);
    data.append("title", title.value);

    const res = await fetch("/api/v1/travels/image", {
        method: "POST",
        body: data,
    });

    const resJSON = await res.json();
    
    if(!res.ok) {
        console.log(resJSON.error)
        return false;
    }

    const imageURL = resJSON.URL;
    
    try {
        await createTravel(imageURL);
    } catch (error) {
        console.log(error);
    }

    
    return false;
}

async function createTravel(imageURL) {
    const travel = {
        "title": title.value,
        "duration_days": parseInt(durationDays.value),
        "price": parseInt(price.value),
        "party_size": parseInt(party_size.value),
        "complexity": parseInt(complexity.value),
        "place": parseInt(place.value),
        "description": description.value,
        "date": date.value,
        "URL": imageURL,
    }

    try {
        const res = await fetch("/api/v1/travels", {
            method: "POST",
            body: JSON.stringify(travel),
        });
        
        const resJSON = await res.json();
        if(!res.ok) {
            console.log(resJSON)
            return;
        }
        alert("Travel created");
    } catch(error) {
        console.log(error);
    }
    
    console.log(resJSON);
}