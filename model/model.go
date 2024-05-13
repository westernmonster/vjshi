package db

import (
	types "github.com/jmoiron/sqlx/types"
)

type Sale struct {
	ID               int64         `db:"id" json:"id,string"`
	Title            string        `db:"title" json:"title"`
	CreatedAt        int64         `db:"created_at" json:"created_at"`
	DownloadTimes    int32         `db:"download_times" json:"download_times"`
	SoftwareTypeName string        `db:"software_type_name" json:"software_type_name"`
	IsRecommended    types.BitBool `db:"is_recommended" json:"is_recommended"`
	SoftwareType     int32         `db:"software_type" json:"software_type"`
	SellerID         int32         `db:"seller_id" json:"seller_id"`
	SellerName       string        `db:"seller_name" json:"seller_name"`
	LicType          string        `db:"lic_type" json:"lic_type"`
	Keyword          string        `db:"keyword" json:"keyword"`
	SaleTime         int32         `db:"sale_time" json:"sale_time"`
	VideoID          int32         `db:"video_id" json:"video_id"`
}
