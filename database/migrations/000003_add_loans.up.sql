CREATE TABLE "loans" (
  "id" BIGSERIAL PRIMARY KEY,
  "account_id" BIGINT NOT NULL,
  "loan_amount" BIGINT NOT NULL,
  "interest_rate" BIGINT NOT NULL,
  "status" VARCHAR NOT NULL,
  "start_date" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now()),
  "end_date" TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX ON "loans" ("account_id");

COMMENT ON COLUMN "loans"."loan_amount" IS 'must be positive';

COMMENT ON COLUMN "loans"."interest_rate" IS 'must be positive';

ALTER TABLE "loans" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");