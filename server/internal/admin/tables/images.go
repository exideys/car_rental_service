package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetImagesTable(ctx *context.Context) table.Table {
	cfg := table.
		DefaultConfigWithDriverAndConnection("mysql", "default").
		SetPrimaryKey("image_id", db.Int)

	images := table.NewDefaultTable(ctx, cfg)

	info := images.GetInfo().HideFilterArea()
	info.
		SetTable("images").
		SetTitle("Images").
		SetDescription("Images")
	info.AddField("Image ID", "image_id", db.Int)
	info.AddField("Car ID", "car_id", db.Int)
	info.AddField("Image Path", "image_path", db.Varchar)

	formList := images.GetForm()
	formList.
		SetTable("images").
		SetTitle("Images").
		SetDescription("Images")
	formList.AddField("Image ID", "image_id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Car ID", "car_id", db.Int, form.Number)
	formList.AddField("Image Path", "image_path", db.Varchar, form.Text)

	return images
}
