package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetPaymentDetailsViewTable(ctx *context.Context) table.Table {
	cfg := table.
		DefaultConfigWithDriverAndConnection("mysql", "default").
		SetPrimaryKey("payment_id", db.Int)

	details := table.NewDefaultTable(ctx, cfg)

	info := details.GetInfo().HideFilterArea()
	info.
		SetTable("payment_details_view").
		SetTitle("Payment Details").
		SetDescription("Payment Details")
	info.AddField("Payment ID", "payment_id", db.Int)
	info.AddField("Payment Date", "payment_date", db.Date)
	info.AddField("Amount", "amount", db.Decimal)
	info.AddField("Order ID", "order_id", db.Int)
	info.AddField("First Name", "first_name", db.Varchar)
	info.AddField("Last Name", "last_name", db.Varchar)

	return details
}
