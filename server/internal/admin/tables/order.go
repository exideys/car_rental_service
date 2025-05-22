package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetOrderTable(ctx *context.Context) table.Table {

	order := table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := order.GetInfo().HideFilterArea()

	info.SetTable("order").SetTitle("Order").SetDescription("Order")

	formList := order.GetForm()

	formList.SetTable("order").SetTitle("Order").SetDescription("Order")

	return order
}
