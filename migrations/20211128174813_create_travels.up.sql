CREATE TABLE places(
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE travels(
    id SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(255) NOT NULL,
    duration_days INT NOT NULL,
    price INT NOT NULL,
    party_size INT,
    complexity INT NOT NULL,
    description TEXT NOT NULL,
    date DATE NOT NULL,
    place INT NOT NULL,

    FOREIGN KEY(place) REFERENCES places(id) ON DELETE SET NULL
);

CREATE TABLE users_travels(
    user_id INT NOT NULL,
    travel_id INT NOT NULL,

    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE SET NULL,
    FOREIGN KEY(travel_id) REFERENCES travels(id) ON DELETE SET NULL
);