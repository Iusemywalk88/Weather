CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE cities (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE favorite_cities (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    city_id INT REFERENCES cities(id) ON DELETE CASCADE
);

CREATE TABLE weather_history (
    id SERIAL PRIMARY KEY,
    city_id INT REFERENCES cities(id),
    temperature FLOAT,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT now()
);