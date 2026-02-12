-- Add new schema named "public"
CREATE SCHEMA IF NOT EXISTS "public";
-- Set comment to schema: "public"
COMMENT ON SCHEMA "public" IS 'standard public schema';
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
