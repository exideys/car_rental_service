package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetPaymentTable(ctx *context.Context) table.Table {

	payment := table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := payment.GetInfo().HideFilterArea()

	info.SetTable("payment").SetTitle("Payment").SetDescription("Payment")

	formList := payment.GetForm()

	formList.SetTable("payment").SetTitle("Payment").SetDescription("Payment")

	return payment
}
