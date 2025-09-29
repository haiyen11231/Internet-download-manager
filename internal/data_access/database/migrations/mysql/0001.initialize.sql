-- +migrate Up

-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS `accounts` (
    `account_id` UNSIGNED BIGINT PRIMARY KEY,
    `account_name` VARCHAR(256) NOT NULL
);

CREATE TABLE IF NOT EXISTS `account_passwords` (
    `account_id` UNSIGNED BIGINT PRIMARY KEY,
    `password_hash` VARCHAR(128) NOT NULL,
    FOREIGN KEY (`account_id`) REFERENCES `accounts`(`account_id`)
)

CREATE TABLE IF NOT EXISTS `token_public_keys` (
    `key_id` UNSIGNED BIGINT PRIMARY KEY,
    `public_key` VARBINARY(4096) NOT NULL
);

CREATE TABLE IF NOT EXISTS `download_tasks` (
    `task_id` UNSIGNED BIGINT PRIMARY KEY,
    `account_id` UNSIGNED BIGINT,
    `download_type` SMALLINT NOT NULL,
    `file_url` TEXT NOT NULL,
    `download_status` SMALLINT NOT NULL,
    `metadata` TEXT NOT NULL,
    FOREIGN KEY (`account_id`) REFERENCES `accounts`(`account_id`)
);
-- +migrate StatementEnd

-- +migrate Down

-- +migrate StatementBegin
DROP TABLE IF EXISTS download_tasks;
DROP TABLE IF EXISTS token_public_keys;
DROP TABLE IF EXISTS account_passwords;
DROP TABLE IF EXISTS accounts;
-- +migrate StatementEnd
