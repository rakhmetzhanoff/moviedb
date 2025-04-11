CREATE TABLE directors (
    id SERIAL PRIMARY KEY,
    firstname TEXT,
    lastname TEXT
);

CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title TEXT,
    director_id INT,
    CONSTRAINT fk_director FOREIGN KEY (director_id) REFERENCES directors(id)
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);