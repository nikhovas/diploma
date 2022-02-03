SELECT *
FROM telegram_staff_account
WHERE telegram_user_id = 5;

SELECT *
FROM shop
INNER JOIN telegram_staff_account ON telegram_staff_account.user_id = shop.creator
WHERE telegram_staff_account.telegram_user_id = 5;

BEGIN;

SELECT user_id
FROM telegram_staff_account
WHERE telegram_user_id = 5;

WITH demo as (
    INSERT INTO staff
    DEFAULT VALUES
    RETURNING id
)
INSERT INTO telegram_staff_account(user_id, telegram_user_id, telegram_chat_id)
VALUES ((SELECT id FROM demo), 55, 55)
RETURNING *;



SELECT user_id
FROM telegram_staff_account
WHERE telegram_user_id = 54584535;