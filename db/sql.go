package db

import (
	"context"
	"vjshi/parser"

	_ "github.com/go-sql-driver/mysql"
	"github.com/heetch/sqalx"
	"github.com/jmoiron/sqlx"
)

func ConnectDB() (node sqalx.Node, err error) {
	db, err := sqlx.Connect("mysql", "")
	if err != nil {
		return
	}

	return sqalx.New(db)
}

func CreateSales(ctx context.Context, node sqalx.Node, sale *parser.Sale) (err error) {
	sql := `
	INSERT INTO sales ( id, title, created_at, download_times, software_type, software_type_name, is_recommended, seller_id, seller_name, lic_type, keyword, sale_time, video_id ),
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)
	`

	_, err = node.ExecContext(ctx, sql,
		1,
		sale.Video.Title,
		sale.Video.CreatedAt,
		sale.Video.DownloadTimes,
		sale.Video.SoftwareType.ID,
		sale.Video.SoftwareType.Name,
		sale.Video.IsRecommended,
		sale.Seller.UID,
		sale.Seller.UserName,
		sale.LicType,
		sale.Keyword,
		sale.Timestamp,
		sale.Video.ID,
	)
	if err != nil {
		return
	}

	return
}
