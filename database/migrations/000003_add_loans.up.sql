CREATE TABLE "loan_transfers" (
  "id" BIGSERIAL PRIMARY KEY,
  "account_id" BIGINT NOT NULL,
  "amount" BIGINT NOT NULL,
  "interest_rate" NUMERIC NOT NULL DEFAULT 1.5,
  "open" BOOLEAN NOT NULL DEFAULT true,
  "start_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now()),
  "end_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now()) + INTERVAL '1' MINUTE
);

CREATE INDEX ON "loan_transfers" ("account_id");

COMMENT ON COLUMN "loan_transfers"."amount" IS 'must be positive';

COMMENT ON COLUMN "loan_transfers"."interest_rate" IS 'must be positive';

ALTER TABLE "loan_transfers" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");