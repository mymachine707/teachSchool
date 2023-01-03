CREATE TABLE "accaunts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "enteries" (
  "id" bigserial PRIMARY KEY,
  "accaunts_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_accaunts_id" bigint NOT NULL,
  "to_accaunts_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accaunts" ("owner");

CREATE INDEX ON "enteries" ("accaunts_id");

CREATE INDEX ON "transfers" ("from_accaunts_id");

CREATE INDEX ON "transfers" ("to_accaunts_id");

CREATE INDEX ON "transfers" ("from_accaunts_id", "to_accaunts_id");

COMMENT ON COLUMN "enteries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "enteries" ADD FOREIGN KEY ("accaunts_id") REFERENCES "accaunts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_accaunts_id") REFERENCES "accaunts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_accaunts_id") REFERENCES "accaunts" ("id");