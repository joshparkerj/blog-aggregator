-- name: CreateFeedFollow :one
WITH new_feed_follow AS (
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
SELECT $1, $2, $3, $4, feeds.id as feed_id FROM feeds
WHERE feeds.url = $5

RETURNING *
) SELECT
  new_feed_follow.id as id,
  new_feed_follow.created_at as created_at,
  new_feed_follow.updated_at as updated_at,
  new_feed_follow.user_id as user_id,
  new_feed_follow.feed_id as feed_id,
  feeds.name as feed_name,
  users.name as user_name
FROM new_feed_follow
JOIN feeds on feeds.id = new_feed_follow.feed_id
JOIN users on users.id = new_feed_follow.user_id;

