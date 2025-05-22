package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetImagesTable(ctx *context.Context) table.Table {

	images := table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := images.GetInfo().HideFilterArea()

	info.SetTable("images").SetTitle("Images").SetDescription("Images")

	formList := images.GetForm()

	formList.SetTable("images").SetTitle("Images").SetDescription("Images")

	return images
}
