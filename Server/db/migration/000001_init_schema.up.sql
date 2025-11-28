CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" text NOT NULL,
  "image_url" varchar NOT NULL,
  "price" decimal(10,2) NOT NULL,
  "created_at" timestamptz NOT NULL
);

CREATE TABLE "seckill_events" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "start_time" timestamptz NOT NULL,
  "end_time" timestamptz NOT NULL,
  "stock_count" bigint NOT NULL,
  "remaining_stock" bigint NOT NULL,
  "seckill_price" decimal(10,2) NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "seckill_event_id" bigint NOT NULL,
  "amount" decimal(10,2) NOT NULL,
  "status" varchar NOT NULL DEFAULT 'created',
  "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "seckill_events" ("start_time", "end_time");

CREATE INDEX ON "orders" ("user_id");

CREATE UNIQUE INDEX ON "orders" ("user_id", "seckill_event_id");

COMMENT ON COLUMN "users"."password" IS 'bcrypt Hash';

COMMENT ON COLUMN "seckill_events"."stock_count" IS 'The total inventory for this activity';

COMMENT ON COLUMN "seckill_events"."remaining_stock" IS 'Remaining inventory';

COMMENT ON COLUMN "orders"."status" IS 'created, paid, cancelled';

ALTER TABLE "seckill_events" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("seckill_event_id") REFERENCES "seckill_events" ("id");
