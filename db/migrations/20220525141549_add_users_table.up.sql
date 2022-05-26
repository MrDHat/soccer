-- Sequence and defined type
CREATE SEQUENCE users_id_seq;
CREATE TABLE "users" (
  "id" int4 NOT NULL UNIQUE DEFAULT nextval('users_id_seq'::regclass),
  "name" text NOT NULL DEFAULT '',
  "email" text UNIQUE NOT NULL DEFAULT '',
  "password" text NOT NULL DEFAULT '',
  "team_id" int4 NOT NULL DEFAULT 0,
  "created_at" bigint NOT NULL DEFAULT extract(
    epoch
    from now()
  ),
  "updated_at" bigint NOT NULL DEFAULT -62135596800,
  "deleted_at" bigint NULL
);
ALTER TABLE "users"
ADD CONSTRAINT user_team_constrant FOREIGN KEY ("team_id") REFERENCES "teams" ("id");
