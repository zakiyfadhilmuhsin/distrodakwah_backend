package repository

import (
	"fmt"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"gorm.io/gorm"
)

func (r *ProductRepository) SaveSingleProductPricesTx(singleProductPrices []*model.SingleProductsPrices, tx *gorm.DB) error {
	err := tx.Model(&model.SingleProductsPrices{}).Create(&singleProductPrices).Error
	if err != nil {
		fmt.Printf("error creating Single Product Prices\n %+v \n", err)
		tx.Rollback()
		return err
	}
	return nil

}