-- Add new schema named "public"
CREATE SCHEMA IF NOT EXISTS "public";
-- Set comment to schema: "public"
COMMENT ON SCHEMA "public" IS 'standard public schema';
-- Create "orders" table
CREATE TABLE "public"."orders" (
  "id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create "products" table
CREATE TABLE "public"."products" (
  "id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
  "name" text NOT NULL,
  "price" numeric(10,2) NOT NULL,
  "quantity" integer NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "products_price_check" CHECK (price >= (0)::numeric)
);
-- Create "order_items" table
CREATE TABLE "public"."order_items" (
  "id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
  "order_id" bigint NOT NULL,
  "quantity" integer NOT NULL DEFAULT 1,
  "product_ids" bigint[] NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "order_items_order_id_fkey" FOREIGN KEY ("order_id") REFERENCES "public"."orders" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
