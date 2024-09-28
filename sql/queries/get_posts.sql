-- name: GetPostsForUser :many
SELECT posts.title, posts.description, posts.url, posts.published_at, feeds.name FROM posts
JOIN feeds ON posts.feed_id = feeds.id
JOIN feed_follows ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC LIMIT $2;

