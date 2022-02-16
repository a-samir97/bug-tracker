CREATE TABLE users IF NOT EXIST(
    user_id SERIAL PRIMARY KEY,
    role VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL UNIQUE,
    first_name VARCHAR(100) NULL,
    last_name VARCHAR(100) NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
);

CREATE INDEX idx_user_id ON users(user_id);
