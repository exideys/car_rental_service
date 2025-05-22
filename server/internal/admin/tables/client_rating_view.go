package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetClientratingviewTable(ctx *context.Context) table.Table {

	clientRatingView := table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := clientRatingView.GetInfo().HideFilterArea()

	info.SetTable("client_rating_view").SetTitle("Clientratingview").SetDescription("Clientratingview")

	formList := clientRatingView.GetForm()

	formList.SetTable("client_rating_view").SetTitle("Clientratingview").SetDescription("Clientratingview")

	return clientRatingView
}
