CREATE TABLE "loan_transfers" (
  "id" BIGSERIAL PRIMARY KEY,
  "account_id" BIGINT NOT NULL,
  "loan_amount" BIGINT NOT NULL,
  "interest_rate" BIGINT NOT NULL,
  "status" VARCHAR NOT NULL,
  "start_date" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now()),
  "end_date" TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX ON "loan_transfers" ("account_id");

COMMENT ON COLUMN "loan_transfers"."loan_amount" IS 'must be positive';

COMMENT ON COLUMN "loan_transfers"."interest_rate" IS 'must be positive';

ALTER TABLE "loan_transfers" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");