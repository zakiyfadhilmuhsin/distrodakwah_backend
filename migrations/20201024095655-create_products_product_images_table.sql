-- +migrate Up
CREATE TABLE IF NOT EXISTS `products_product_images`(
    `product_id` BIGINT UNSIGNED NOT NULL,
    `product_image_id` BIGINT UNSIGNED NOT NULL,
    CONSTRAINT products_product_id FOREIGN KEY(product_id) references products(id) ON DELETE CASCADE,
    CONSTRAINT products_product_image_id FOREIGN KEY(product_image_id) references product_images(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `products_product_images`;