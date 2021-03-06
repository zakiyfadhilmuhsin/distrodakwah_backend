package productmodel

type ProductsProductImage struct {
	ProductID      uint64        `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	ProductImageID uint64        `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_image_id"`
	ProductImage   *ProductImage `gorm:"foreignKey:ProductImageID" json:"product_image"`
}

// func (ProductHasManyImage) TableName() string {
// 	return "products_product_images"
// }
