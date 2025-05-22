package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetPaymentTable(ctx *context.Context) table.Table {

	payment := table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := payment.GetInfo().HideFilterArea()

	info.SetTable("payment").SetTitle("Payment").SetDescription("Payment")

	formList := payment.GetForm()

	formList.SetTable("payment").SetTitle("Payment").SetDescription("Payment")

	return payment
}
