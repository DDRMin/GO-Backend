-- name: ListProducts :many
SELECT * FROM products
ORDER BY name;

-- name: FindProductByID :one
SELECT * FROM products
WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO products (name, price, quantity, created_at)  
VALUES ($1, $2, $3, NOW())
RETURNING id;

-- name: CreateOrderItem :exec
INSERT INTO order_items (order_id, product_id, quantity)
VALUES ($1, $2, $3);

-- name: CreateOrder :one
INSERT INTO orders (user_id, created_at)
VALUES ($1, NOW())
RETURNING id;

-- name: ReduceProductQuantity :one
UPDATE products
SET quantity = quantity - $2
WHERE id = $1 AND quantity >= $2
RETURNING quantity;