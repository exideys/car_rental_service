package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetClientratingviewTable(ctx *context.Context) table.Table {

	clientRatingView := table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := clientRatingView.GetInfo().HideFilterArea()

	info.SetTable("client_rating_view").SetTitle("Clientratingview").SetDescription("Clientratingview")

	formList := clientRatingView.GetForm()

	formList.SetTable("client_rating_view").SetTitle("Clientratingview").SetDescription("Clientratingview")

	return clientRatingView
}
