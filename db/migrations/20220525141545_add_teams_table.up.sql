-- Sequence and defined type
CREATE SEQUENCE teams_id_seq;
CREATE TABLE "teams" (
  "id" int4 NOT NULL UNIQUE DEFAULT nextval('teams_id_seq'::regclass),
  "name" text NOT NULL DEFAULT '',
  "country" text,
  "remaining_budget_in_dollars" bigint NOT NULL DEFAULT 0,
  "created_at" bigint NOT NULL DEFAULT extract(
    epoch
    from now()
  ),
  "updated_at" bigint NOT NULL DEFAULT -62135596800,
  "deleted_at" bigint NULL
);
