package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetOrderTable(ctx *context.Context) table.Table {
	cfg := table.
		DefaultConfigWithDriverAndConnection("mysql", "default").
		SetPrimaryKey("order_id", db.Int)

	orders := table.NewDefaultTable(ctx, cfg)

	info := orders.GetInfo().HideFilterArea()
	info.
		SetTable("orders").
		SetTitle("Order").
		SetDescription("Order").
		SetPrimaryKey("order_id", db.Int)
	info.AddField("Order ID", "order_id", db.Int)
	info.AddField("Start Date", "start_date", db.Date)
	info.AddField("End Date", "end_date", db.Date)
	info.AddField("Total Cost", "total_cost", db.Decimal)
	info.AddField("Status", "status", db.Enum)
	info.AddField("Is Paid", "is_paid", db.Tinyint)
	info.AddField("Tariff ID", "tariff_id", db.Int)
	info.AddField("Client ID", "client_id", db.Int)
	info.AddField("Car ID", "car_id", db.Int)

	formList := orders.GetForm()
	formList.
		SetTable("orders").
		SetTitle("Order").
		SetDescription("Order").
		SetPrimaryKey("order_id", db.Int)
	formList.AddField("Order ID", "order_id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Start Date", "start_date", db.Date, form.Text)
	formList.AddField("End Date", "end_date", db.Date, form.Text)
	formList.AddField("Total Cost", "total_cost", db.Decimal, form.Number)
	formList.AddField("Status", "status", db.Enum, form.Text)
	formList.AddField("Is Paid", "is_paid", db.Tinyint, form.Number)
	formList.AddField("Tariff ID", "tariff_id", db.Int, form.Number)
	formList.AddField("Client ID", "client_id", db.Int, form.Number)
	formList.AddField("Car ID", "car_id", db.Int, form.Number)

	return orders
}
