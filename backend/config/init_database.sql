CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    role_id INT NOT NULL,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles(role_id) 
);

CREATE INDEX idx_user_id ON users(user_id);

CREATE TABLE roles (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(10) NOT NULL
);


CREATE IF NOT EXISTS TABLE bugs (

);

CREATE IF NOT EXISTS TABLE status (

);



