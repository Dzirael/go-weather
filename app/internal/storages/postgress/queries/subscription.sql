-- name: CreateSubscription :exec
INSERT INTO subscription (id, confirmation_code, status, frequency, email, city, sended_at)
VALUES (@id, @code, @status, @frequency, @email, @city, NOW());

-- name: UpdateSubscriptionStatusByCode :exec
UPDATE subscription
SET status = @status,
    confirmation_code = NULL,
    valid_til = NULL
WHERE confirmation_code = @confirmation_code
RETURNING id;

-- name: GetWaitingsNotification :many
SELECT id FROM subscription
WHERE status = 'ACTIVE'
  AND (
    (frequency = 'hourly' AND (sended_at IS NULL OR sended_at < NOW() - INTERVAL '1 hour')) OR
    (frequency = 'daily'  AND (sended_at IS NULL OR sended_at < NOW() - INTERVAL '1 day'))
  );

-- name: GetSubscriptionByID :one
SELECT * FROM subscription
WHERE id = @id;

-- name: SetSendedNow :exec
UPDATE subscription
SET sended_at = NOW()
WHERE id = @id;

-- name: DeleteSubscription :exec
DELETE FROM subscription
WHERE confirmation_code = @code;

-- name: HaveActiveSubscription :one
SELECT EXISTS (
  SELECT 1 FROM subscription WHERE email = @email AND status = 'ACTIVE'
) AS exists;