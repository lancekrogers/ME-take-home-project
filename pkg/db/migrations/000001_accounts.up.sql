CREATE TABLE "accounts" (
  "id" varchar PRIMARY KEY,
  "account_type" varchar NOT NULL,
  "tokens" bigint NOT NULL DEFAULT 0,
  "data" jsonb,
  "version" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "account_updates" (
  "id" varchar NOT NULL,
  "account_type" varchar NOT NULL,
  "tokens" bigint NOT NULL DEFAULT 0,
  "data" jsonb,
  "version" int,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
