
//create user table with fields name, surname, email, password postgre sql

CREATE TABLE IF NOT EXISTS user
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);


// create script to insert data to user table