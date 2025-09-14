CREATE TABLE IF NOT EXISTS `accounts` (
    `user_id` UNSIGNED BIGINT PRIMARY KEY,
    `username` VARCHAR(256) NOT NULL,
);

CREATE TABLE IF NOT EXISTS `account_passwords` (
    `of_user_id` UNSIGNED BIGINT PRIMARY KEY,
    `password_hash` VARCHAR(128) NOT NULL,
    FOREIGN KEY (`of_user_id`) REFERENCES `accounts`(`user_id`)
)

CREATE TABLE IF NOT EXISTS `download_tasks` (
    `task_id` UNSIGNED BIGINT PRIMARY KEY,
    `of_user_id` UNSIGNED BIGINT,
    `download_type` SMALLINT NOT NULL,
    `file_url` TEXT NOT NULL,
    `download_status` SMALLINT NOT NULL,
    `metadata` TEXT NOT NULL,
    FOREIGN KEY (`of_user_id`) REFERENCES `accounts`(`user_id`)
);
