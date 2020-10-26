-- +migrate Up
CREATE TABLE IF NOT EXISTS variants(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `product_id` BIGINT UNSIGNED NOT NULL,
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,

    CONSTRAINT variants_product_id foreign key(product_id) references products(id)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



-- +migrate Down
DROP TABLE `variants`;