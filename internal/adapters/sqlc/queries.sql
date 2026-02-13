-- name: ListProducts :many
SELECT * FROM products
ORDER BY name;

-- name: FindProductByID :one
SELECT * FROM products
WHERE id = $1;

-- name: CreateOrderItem :exec
INSERT INTO order_items (order_id, product_id, quantity)
VALUES ($1, $2, $3);

-- name: CreateOrder :exec
INSERT INTO orders (user_id, created_at)
VALUES ($1, $2)
RETURNING id; 