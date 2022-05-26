-- Sequence and defined type
CREATE SEQUENCE players_id_seq;
CREATE TABLE "players" (
  "id" int4 NOT NULL UNIQUE DEFAULT nextval('players_id_seq'::regclass),
  "first_name" text NOT NULL DEFAULT '',
  "last_name" text NOT NULL DEFAULT '',
  "age" int4 NOT NULL DEFAULT 0,
  "current_value_in_dollars" int4 NOT NULL DEFAULT 0,
  "player_type" text NOT NULL DEFAULT '',
  "team_id" int4 NOT NULL DEFAULT 0,
  "created_at" bigint NOT NULL DEFAULT extract(
    epoch
    from now()
  ),
  "updated_at" bigint NOT NULL DEFAULT -62135596800,
  "deleted_at" bigint NULL
);
ALTER TABLE "players"
ADD CONSTRAINT user_team_constrant FOREIGN KEY ("team_id") REFERENCES "teams" ("id");
