-- name: GetFeedFollowsForUser :many
SELECT users.name AS user_name, feeds.name AS feed_name, feeds.url AS feed_url FROM feed_follows
JOIN users ON feed_follows.user_id = users.id
JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE users.name = $1;

