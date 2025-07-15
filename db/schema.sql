CREATE TABLE mockdata (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT,
    gender TEXT,
    ip_address TEXT,
    date DATE
);