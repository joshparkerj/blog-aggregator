-- name: GetFeeds :many
SELECT
feeds.name as name,
feeds.url as url,
users.name as username
FROM feeds

JOIN users ON feeds.user_id = users.id;
