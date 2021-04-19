-- migrate:up
CREATE TYPE user_role AS ENUM (
    'USER',
    'ADMIN'
);

CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    deleted_at timestamptz,
    username text NOT NULL UNIQUE,
    email text NOT NULL UNIQUE,
    password TEXT NOT NULL,
    ROLE user_role NOT NULL DEFAULT 'USER'
);

CREATE TABLE IF NOT EXISTS capstones (
    id uuid PRIMARY KEY NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    deleted_at timestamptz,
    title text NOT NULL UNIQUE,
    description text NOT NULL,
    author text NOT NULL,
    semester text NOT NULL
);

-- migrate:down
DROP TYPE IF EXISTS user_role;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS capstones;

