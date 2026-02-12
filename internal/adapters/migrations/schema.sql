-- This file is used by sqlc to generate Go code.
-- Keep in sync with schema.hcl (the Atlas source of truth).

CREATE TABLE IF NOT EXISTS products (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text    NOT NULL,
  price  DECIMAL(10, 2) NOT NULL CHECK (price >= 0),
  quantity INT NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
