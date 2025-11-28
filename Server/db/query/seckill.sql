-- name: CreateSeckillEvent :one
INSERT INTO seckill_events (
  product_id,start_time,end_time,stock_count,remaining_stock,seckill_price
) VALUES(
   $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetSeckillEvent :one
SELECT * FROM seckill_events
WHERE id = $1 LIMIT 1;

-- name: ListActiveSeckillEvents :many
SELECT
  se.*,
  p.name as product_name,
  p.image_url as product_image
FROM seckill_events se
JOIN products p ON se.product_id = p.id
WHERE se.start_time <= now() AND se.end_time >= now()
ORDER BY se.start_time ASC;

-- name: ReduceInventory :execresult
UPDATE seckill_events
SET remaining_stock = remaining_stock - 1
WHERE id = $1 AND remaining_stock > 0;