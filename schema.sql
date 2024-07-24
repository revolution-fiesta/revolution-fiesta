CREATE TABLE info (
    version TEXT NOT NULL
);

CREATE TABLE users (
    id  SERIAL PRIMARY KEY,
    passwd_hash TEXT NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT NOT NULL,
    type TEXT CHECK ( type IN ('ADMIN', 'REGULAR'))
);