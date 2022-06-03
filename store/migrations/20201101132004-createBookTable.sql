
-- +migrate Up
CREATE TABLE books (
    id uuid PRIMARY KEY,
    title VARCHAR(256),
    price int
);

INSERT INTO books (id, title, price) VALUES ('030febee-2d4c-4aa6-b578-ae2d084b1b31', 'Cracking the coding interview', 20);
INSERT INTO books (id, title, price) VALUES ('62b9c00e-221f-4894-b7e2-3f03bd7f727f', 'Never split the difference', 25);

-- +migrate Down
DROP TABLE books;