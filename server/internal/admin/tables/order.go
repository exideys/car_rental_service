package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetOrderTable(ctx *context.Context) table.Table {

	order := table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := order.GetInfo().HideFilterArea()

	info.SetTable("order").SetTitle("Order").SetDescription("Order")

	formList := order.GetForm()

	formList.SetTable("order").SetTitle("Order").SetDescription("Order")

	return order
}
