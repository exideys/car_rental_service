package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetPaymentdetailsviewTable(ctx *context.Context) table.Table {

	paymentDetailsView := table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := paymentDetailsView.GetInfo().HideFilterArea()

	info.SetTable("payment_details_view").SetTitle("Paymentdetailsview").SetDescription("Paymentdetailsview")

	formList := paymentDetailsView.GetForm()

	formList.SetTable("payment_details_view").SetTitle("Paymentdetailsview").SetDescription("Paymentdetailsview")

	return paymentDetailsView
}
