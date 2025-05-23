package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetClientRatingViewTable(ctx *context.Context) table.Table {
	cfg := table.
		DefaultConfigWithDriverAndConnection("mysql", "default").
		SetPrimaryKey("first_name", db.Varchar) // у VIEW нет автоинкремента, но укажем условно
	ratings := table.NewDefaultTable(ctx, cfg)

	info := ratings.GetInfo().HideFilterArea()
	info.
		SetTable("client_rating_view").
		SetTitle("Client Rating View").
		SetDescription("Client Rating View")
	info.AddField("First Name", "first_name", db.Varchar)
	info.AddField("Last Name", "last_name", db.Varchar)
	info.AddField("Date", "evaluation_date", db.Date)
	info.AddField("Rating", "rating", db.Int)
	info.AddField("Comment", "comment", db.Text)
	info.AddField("Is Late Return", "is_late_return", db.Tinyint)

	return ratings
}
