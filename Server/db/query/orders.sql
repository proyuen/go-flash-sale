-- name: CreateOrder :one

INSERT INTO orders(
  user_id,seckill_event_id,amount,status
) VALUES(
  $1,$2,$3,$4
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;