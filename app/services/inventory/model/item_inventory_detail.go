package model

import "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/user/model"

type ItemInventoryDetail struct {
	ID              uint64            `gorm:"primaryKey;autoIncrement;not null"`
	ItemInventoryID uint64            `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"item_inventory_id"`
	VendorID        uint32            `gorm:"type:INT;UNSIGNED;NOT NULL" json:"vendor_id"`
	UserVendor      *model.UserVendor `gorm:"foreignKey:VendorID" json:"vendor"`
}