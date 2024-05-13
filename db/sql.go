package db

import (
	"context"
	"database/sql"
	"vjshi/model"

	"github.com/charmbracelet/log"
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

func GetSaleByID(c context.Context, node sqalx.Node, id int64) (item *model.Sale, err error) {
	item = new(model.Sale)
	sqlSelect := "SELECT a.id,a.title,a.created_at,a.download_times,a.software_type_name,a.is_recommended,a.software_type,a.seller_id,a.seller_name,a.lic_type,a.keyword,a.sale_time,a.video_id FROM d_sale a WHERE a.id=?"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.Errorf("dao.GetSaleByID err(%+v), id(%+v)", err, id)
	}

	return
}

func AddSales(c context.Context, node sqalx.Node, item *model.Sale) (err error) {
	sqlInsert := "INSERT INTO d_sale( id,title,created_at,download_times,software_type_name,is_recommended,software_type,seller_id,seller_name,lic_type,keyword,sale_time,video_id) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Title, item.CreatedAt, item.DownloadTimes, item.SoftwareTypeName, item.IsRecommended, item.SoftwareType, item.SellerID, item.SellerName, item.LicType, item.Keyword, item.SaleTime, item.VideoID); err != nil {
		log.Errorf("dao.AddSale err(%+v), item(%+v)", err, item)
		return
	}

	return
}

func UpdateSale(c context.Context, node sqalx.Node, item *model.Sale) (err error) {
	sqlUpdate := "UPDATE d_sale SET title=?,download_times=?,software_type_name=?,is_recommended=?,software_type=?,seller_id=?,seller_name=?,lic_type=?,keyword=?,sale_time=?,video_id=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Title, item.DownloadTimes, item.SoftwareTypeName, item.IsRecommended, item.SoftwareType, item.SellerID, item.SellerName, item.LicType, item.Keyword, item.SaleTime, item.VideoID, item.ID)
	if err != nil {
		log.Errorf("dao.UpdateSale err(%+v), item(%+v)", err, item)
		return
	}

	return
}

func DelSale(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "DELETE FROM d_sale WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.Errorf("dao.DelSale err(%+v), item(%+v)", err, id)
		return
	}

	return
}
