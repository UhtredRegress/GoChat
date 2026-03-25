CREATE TABLE users (
    id serial PRIMARY KEY,
    name TEXT NOT NULL, 
    password TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPZ DEFAULT now(),   
    updated_at TIMESTAMPZ default now()
);