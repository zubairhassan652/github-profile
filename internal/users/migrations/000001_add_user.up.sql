-- Start of transaction

BEGIN;

-- DB state will be updated here

CREATE TABLE
    users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) NOT NULL
    );

-- Commit transaction changes

COMMIT;