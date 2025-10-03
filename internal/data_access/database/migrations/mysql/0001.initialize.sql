-- +migrate Up
CREATE TABLE IF NOT EXISTS `accounts` (
    `id` UNSIGNED BIGINT AUTO_INCREMENT PRIMARY KEY,
    `account_name` VARCHAR(256) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS `account_passwords` (
    `of_account_id` UNSIGNED BIGINT AUTO_INCREMENT PRIMARY KEY,
    `password_hash` VARCHAR(128) NOT NULL,
    FOREIGN KEY (`of_account_id`) REFERENCES `accounts`(`id`)
)

CREATE TABLE IF NOT EXISTS `token_public_keys` (
    `id` UNSIGNED BIGINT AUTO_INCREMENT PRIMARY KEY,
    `public_key` TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS `download_tasks` (
    `id` UNSIGNED BIGINT AUTO_INCREMENT PRIMARY KEY,
    `of_account_id` UNSIGNED BIGINT,
    `download_type` SMALLINT NOT NULL,
    `file_url` TEXT NOT NULL,
    `download_status` SMALLINT NOT NULL,
    `metadata` TEXT NOT NULL,
    FOREIGN KEY (`of_account_id`) REFERENCES `accounts`(`id`)
);

-- +migrate Down

DROP TABLE IF EXISTS download_tasks;

DROP TABLE IF EXISTS token_public_keys;

DROP TABLE IF EXISTS account_passwords;

DROP TABLE IF EXISTS accounts;
