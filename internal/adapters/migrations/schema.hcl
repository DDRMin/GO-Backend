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

table "orders" {
  schema = schema.public

  column "id" {
    type = bigint
    identity {
      generated = ALWAYS
    }
  }

  column "user_id" {
    type = bigint
    null = false
  }

  column "created_at" {
    type    = timestamptz
    null    = false
    default = sql("CURRENT_TIMESTAMP")
  }

  primary_key {
    columns = [column.id]
  }
}

table "order_items" {
  schema = schema.public

  column "id" {
    type = bigint
    identity {
      generated = ALWAYS
    }
  }

  column "order_id" {
    type = bigint
    null = false
  }

  column "product_id" {
    type = bigint
    null = false
  }

  column "quantity" {
    type = int
    null = false
    default = 1
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "order_items_order_id_fkey" {
    columns     = [column.order_id]
    ref_columns = [table.orders.column.id]
    on_delete   = "CASCADE"
  }

  foreign_key "order_items_product_id_fkey" {
    columns     = [column.product_id]
    ref_columns = [table.products.column.id]
    on_delete   = "CASCADE"
  }
}
