package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetPaymentTable(ctx *context.Context) table.Table {
	cfg := table.
		DefaultConfigWithDriverAndConnection("mysql", "default").
		SetPrimaryKey("payment_id", db.Int)

	payments := table.NewDefaultTable(ctx, cfg)

	info := payments.GetInfo().HideFilterArea()
	info.
		SetTable("payment").
		SetTitle("Payment").
		SetDescription("Payment")
	info.AddField("Payment ID", "payment_id", db.Int)
	info.AddField("Payment Date", "payment_date", db.Date)
	info.AddField("Amount", "amount", db.Decimal)
	info.AddField("Order ID", "order_id", db.Int)

	formList := payments.GetForm()
	formList.
		SetTable("payment").
		SetTitle("Payment").
		SetDescription("Payment")
	formList.AddField("Payment ID", "payment_id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Payment Date", "payment_date", db.Date, form.Text)
	formList.AddField("Amount", "amount", db.Decimal, form.Number)
	formList.AddField("Order ID", "order_id", db.Int, form.Number)

	return payments
}
