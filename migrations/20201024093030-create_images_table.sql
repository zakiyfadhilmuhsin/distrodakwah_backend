-- +migrate Up
CREATE TABLE IF NOT EXISTS `images` (
    `id` bigint UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `url` VARCHAR(255) NOT NULL
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE images;