CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(256),
    price int
);

INSERT INTO books (title, price) VALUES ('Cracking the coding interview', 20);
INSERT INTO books (title, price) VALUES ('Never split the difference', 25);