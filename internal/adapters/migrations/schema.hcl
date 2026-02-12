schema "public" {}

table "products" {
  schema = schema.public

  column "id" {
    type = bigint
    identity {
      generated = ALWAYS
    }
  }

  column "name" {
    type = text
    null = false
  }

  column "price" {
    type = decimal(10,2)
    null = false
  }

  column "quantity" {
    type = int
    null = false
    default = 0
  }

  column "created_at" {
    type    = timestamptz
    null    = false
    default = sql("CURRENT_TIMESTAMP")
  }

  primary_key {
    columns = [column.id]
  }

  check "products_price_check" {
    expr = "price >= 0"
  }
}
