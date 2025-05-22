package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetPaymentdetailsviewTable(ctx *context.Context) table.Table {

	paymentDetailsView := table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := paymentDetailsView.GetInfo().HideFilterArea()

	info.SetTable("payment_details_view").SetTitle("Paymentdetailsview").SetDescription("Paymentdetailsview")

	formList := paymentDetailsView.GetForm()

	formList.SetTable("payment_details_view").SetTitle("Paymentdetailsview").SetDescription("Paymentdetailsview")

	return paymentDetailsView
}
