CREATE TABLE "loans" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "loan_amount" bigint NOT NULL,
  "interest_rate" bigint NOT NULL,
  "status" varchar NOT NULL,
  "start_date" timestamp NOT NULL DEFAULT (now()),
  "end_date" timestamp NOT NULL
);

CREATE INDEX ON "loans" ("account_id");

COMMENT ON COLUMN "loans"."loan_amount" IS 'must be positive';

COMMENT ON COLUMN "loans"."interest_rate" IS 'must be positive';

ALTER TABLE "loans" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");