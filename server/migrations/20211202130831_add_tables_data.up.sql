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