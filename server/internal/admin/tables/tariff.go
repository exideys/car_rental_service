package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetTariffTable(ctx *context.Context) table.Table {
	cfg := table.
		DefaultConfigWithDriverAndConnection("mysql", "default").
		SetPrimaryKey("tariff_id", db.Int)

	tariffs := table.NewDefaultTable(ctx, cfg)

	info := tariffs.GetInfo().HideFilterArea()
	info.
		SetTable("tariff").
		SetTitle("Tariff").
		SetDescription("Tariff")
	info.AddField("Tariff ID", "tariff_id", db.Int)
	info.AddField("Name", "tariff_name", db.Varchar)
	info.AddField("Daily Rate", "daily_rate", db.Decimal)

	formList := tariffs.GetForm()
	formList.
		SetTable("tariff").
		SetTitle("Tariff").
		SetDescription("Tariff")
	formList.AddField("Tariff ID", "tariff_id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Name", "tariff_name", db.Varchar, form.Text)
	formList.AddField("Daily Rate", "daily_rate", db.Decimal, form.Number)

	return tariffs
}
