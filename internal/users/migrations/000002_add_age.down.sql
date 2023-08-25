-- Start of transaction

BEGIN;

-- DB state will be updated here

ALTER Table users RENAME COLUMN name TO username;

ALTER TABLE users DROP COLUMN age;

-- Commit transaction changes

COMMIT;