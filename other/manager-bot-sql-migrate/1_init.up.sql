CREATE TABLE IF NOT EXISTS staff (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS telegram_staff_account (
    user_id BIGINT REFERENCES staff (id),
    telegram_user_id BIGINT NOT NULL,
    telegram_chat_id BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS shop (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(64) UNIQUE NOT NULL,
    creator BIGINT REFERENCES staff (id),
    created_at TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS telegram_staff_shop_group (
    shop_id BIGINT REFERENCES shop (id),
    telegram_group_id BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS shop_staff_roles (
    shop_id BIGINT REFERENCES shop (id),
    user_id BIGINT REFERENCES staff (id),
    role VARCHAR(32) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS vk_client_group (
    shop_id BIGINT REFERENCES shop (id),
    vk_group_id BIGINT NOT NULL,
    access_token VARCHAR(256) NOT NULL
);
