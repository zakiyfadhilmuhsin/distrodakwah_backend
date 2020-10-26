-- +migrate Up
CREATE TABLE IF NOT EXISTS brands(
    `id` int UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `image_id` BIGINT UNSIGNED DEFAULT NULL,
    CONSTRAINT brands_image_id foreign key(image_id) references images(id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;


-- +migrate Down
DROP TABLE brands;