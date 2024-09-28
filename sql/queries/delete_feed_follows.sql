-- name: DeleteFeedFollow :exec
WITH feed_id AS (
    SELECT id FROM feeds WHERE feeds.url = $2
)
DELETE FROM feed_follows
WHERE feed_follows.user_id = $1
AND feed_id =  feed_follows.feed_id;

