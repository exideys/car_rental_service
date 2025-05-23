package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCarsTable(ctx *context.Context) table.Table {

	cfg := table.
		DefaultConfigWithDriverAndConnection("mysql", "default").
		SetPrimaryKey("car_id", db.Int)

	cars := table.NewDefaultTable(ctx, cfg)

	info := cars.GetInfo().HideFilterArea()

	info.SetTable("cars").SetTitle("Cars").SetDescription("Cars").SetPrimaryKey("car_id", db.Int)
	info.AddField("Car ID", "car_id", db.Int)
	info.AddField("Brand", "brand", db.Varchar)
	info.AddField("Model", "model", db.Varchar)
	info.AddField("Color", "color", db.Varchar)
	info.AddField("Plate", "plate_number", db.Varchar)
	info.AddField("Year", "year_of_issue", db.Int)
	info.AddField("Class", "car_class", db.Varchar)
	info.AddField("Status", "status", db.Varchar)
	info.AddField("Price", "daily_price", db.Int)
	formList := cars.GetForm()

	formList.SetTable("cars").SetTitle("Cars").SetDescription("Cars").SetPrimaryKey("car_id", db.Int)

	formList.AddField("Car ID", "car_id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Brand", "brand", db.Varchar, form.Text)
	formList.AddField("Model", "model", db.Varchar, form.Text)
	formList.AddField("Color", "color", db.Varchar, form.Text)
	formList.AddField("Plate", "plate_number", db.Varchar, form.Text)
	formList.AddField("Year", "year_of_issue", db.Int, form.Number)
	formList.AddField("Class", "car_class", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Varchar, form.Text)
	formList.AddField("Price", "daily_price", db.Int, form.Number)

	return cars
}
