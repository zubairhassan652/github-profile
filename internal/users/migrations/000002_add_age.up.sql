-- Start of transaction

BEGIN;

-- DB state will be updated here

ALTER Table users RENAME COLUMN username TO name;

ALTER TABLE users ADD COLUMN age INT;

-- Commit transaction changes

COMMIT;