-- Sequence and defined type
CREATE SEQUENCE player_transfers_id_seq;
CREATE TABLE "player_transfers" (
  "id" int4 NOT NULL UNIQUE DEFAULT nextval('player_transfers_id_seq'::regclass),
  "player_id" int4 NOT NULL,
  "owner_team_id" int4 NOT NULL,
  "amount_in_dollars" bigint NOT NULL DEFAULT 0,
  "created_at" bigint NOT NULL DEFAULT extract(
    epoch
    from now()
  ),
  "updated_at" bigint NOT NULL DEFAULT -62135596800,
  "deleted_at" bigint NULL
);
ALTER TABLE "player_transfers"
ADD CONSTRAINT player_transfer_team_constrant FOREIGN KEY ("owner_team_id") REFERENCES "teams" ("id");
ALTER TABLE "player_transfers"
ADD CONSTRAINT player_transfer_player_constrant FOREIGN KEY ("player_id") REFERENCES "players" ("id");
