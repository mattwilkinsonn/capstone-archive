-- Add migration script here
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);
INSERT INTO users
VALUES (1, 'Matt');