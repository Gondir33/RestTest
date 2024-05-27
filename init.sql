CREATE TABLE IF NOT EXISTS currency (
    Id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    value FLOAT,
    created_at TIMESTAMP
);