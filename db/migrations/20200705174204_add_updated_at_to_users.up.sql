BEGIN;

ALTER TABLE users ADD COLUMN updated_at timestamp without time zone NOT NULL;

COMMIT;
