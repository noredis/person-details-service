CREATE TABLE IF NOT EXISTS persons
(
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    patronymic TEXT,
    age SMALLINT,
    gender TEXT,
    nationality TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL
);

