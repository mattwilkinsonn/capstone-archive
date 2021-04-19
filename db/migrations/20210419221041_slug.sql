-- migrate:up
ALTER TABLE capstones
    ADD slug TEXT NOT NULL UNIQUE;

-- migrate:down
