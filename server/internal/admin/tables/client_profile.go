package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetClientprofileTable(ctx *context.Context) table.Table {

	clientProfile := table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := clientProfile.GetInfo().HideFilterArea()

	info.SetTable("client_profile").SetTitle("Clientprofile").SetDescription("Clientprofile")

	formList := clientProfile.GetForm()

	formList.SetTable("client_profile").SetTitle("Clientprofile").SetDescription("Clientprofile")

	return clientProfile
}
