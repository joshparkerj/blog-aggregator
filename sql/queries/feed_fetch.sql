-- name: MarkFeedFetched :exec
-- url, time
UPDATE feeds SET last_fetched_at = $2, updated_at = $2
WHERE url = $1;

