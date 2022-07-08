CREATE TABLE users(
                    id SERIAL NOT NULL PRIMARY KEY,
                    email VARCHAR NOT NULL UNIQUE,
                    username VARCHAR NOT NULL UNIQUE,
                    encrypted_password VARCHAR NOT NULL
);

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
    image_url VARCHAR(255) NOT NULL DEFAULT '',
    place INT NOT NULL,

    FOREIGN KEY(place) REFERENCES places(id) ON DELETE SET NULL
);

CREATE TABLE users_travels(
    user_id INT NOT NULL,
    travel_id INT NOT NULL,

    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE SET NULL,
    FOREIGN KEY(travel_id) REFERENCES travels(id) ON DELETE SET NULL
);

INSERT INTO places VALUES (DEFAULT, 'Карелия');
INSERT INTO places VALUES (DEFAULT, 'Англия');
INSERT INTO places VALUES (DEFAULT, 'США');
INSERT INTO places VALUES (DEFAULT, 'Россия');

CREATE TABLE user_roles(
    id SERIAL PRIMARY KEY NOT NULL,
    role VARCHAR(20) NOT NULL
);

INSERT INTO user_roles VALUES(DEFAULT, 'User');
INSERT INTO user_roles VALUES(DEFAULT, 'Moderator');
INSERT INTO user_roles VALUES(DEFAULT, 'Admin');

ALTER TABLE users ADD COLUMN role INT NOT NULL DEFAULT 1;
ALTER TABLE users ADD FOREIGN KEY(role) REFERENCES user_roles(id);

INSERT INTO travels VALUES (
    DEFAULT,
    'The first title ever',
    10,
    2000,
    5,
    5,
    'New desc',
    NOW(),
    'https://res.cloudinary.com/dydim8luy/image/upload/v1639474810/travels/zxc.jpg',
    1
);

INSERT INTO travels VALUES (
    DEFAULT,
    'asdfasdfasdf123',
    20,
    5000,
    NULL,
    4,
    'lorem lorem lorem123123',
    NOW(),
    'https://res.cloudinary.com/dydim8luy/image/upload/v1639473395/travels/asdfasdfasdf.jpg',
    4
);

INSERT INTO travels VALUES (
    DEFAULT,
    'HELLO WORLD',
    2,
    5000,
    NULL,
    2,
    'Hello world',
    NOW(),
    'https://res.cloudinary.com/dydim8luy/image/upload/v1639425725/travels/HELLO.jpg',
    2
);

INSERT INTO users VALUES(
    DEFAULT,
    'admin@mail.ru',
    'juicyworld',
    '$2a$10$zqJs2OoTNRcDDYNYxTVYwONGVXkyhLM2WwnU10MZRFXWE0PPeaOIa',
    3
);

INSERT INTO users_travels VALUES(1, 3);
INSERT INTO users_travels VALUES(1, 1);
INSERT INTO users_travels VALUES(1, 2);

