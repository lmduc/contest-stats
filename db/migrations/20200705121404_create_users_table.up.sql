BEGIN;

CREATE TABLE IF NOT EXISTS users(
  id uuid DEFAULT gen_random_uuid() NOT NULL,
  created_at timestamp without time zone NOT NULL,
  leetcode_handle character varying
);

ALTER TABLE ONLY users ADD CONSTRAINT users_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX users_leetcode_handle ON users USING btree (leetcode_handle);

COMMIT;
