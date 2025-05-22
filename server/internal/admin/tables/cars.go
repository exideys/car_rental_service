package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetCarsTable(ctx *context.Context) table.Table {

	cars := table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := cars.GetInfo().HideFilterArea()

	info.SetTable("cars").SetTitle("Cars").SetDescription("Cars")

	formList := cars.GetForm()

	formList.SetTable("cars").SetTitle("Cars").SetDescription("Cars")

	return cars
}
