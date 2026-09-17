CREATE TABLE IF NOT EXISTS users(
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    email_address TEXT NOT NULL UNIQUE,
    phone_number TEXT NOT NULL UNIQUE,
    id_number TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL
);