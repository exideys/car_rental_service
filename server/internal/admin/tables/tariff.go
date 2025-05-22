package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetTariffTable(ctx *context.Context) table.Table {

	tariff := table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := tariff.GetInfo().HideFilterArea()

	info.SetTable("tariff").SetTitle("Tariff").SetDescription("Tariff")

	formList := tariff.GetForm()

	formList.SetTable("tariff").SetTitle("Tariff").SetDescription("Tariff")

	return tariff
}
