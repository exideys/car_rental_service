package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetClientTable(ctx *context.Context) table.Table {

	client := table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := client.GetInfo().HideFilterArea()

	info.SetTable("client").SetTitle("Client").SetDescription("Client")

	formList := client.GetForm()

	formList.SetTable("client").SetTitle("Client").SetDescription("Client")

	return client
}
