CREATE TABLE users (
    id uuid PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    password_hash varchar(255) NOT NULL
);

CREATE TABLE lists (
    id uuid PRIMARY KEY,
    list_type varchar(255) NOT NULL,
    owner_id uuid NOT NULL,
    entity_id uuid NOT NULL
);
