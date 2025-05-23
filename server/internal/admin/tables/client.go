package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetClientTable(ctx *context.Context) table.Table {
	cfg := table.
		DefaultConfigWithDriverAndConnection("mysql", "default").
		SetPrimaryKey("client_id", db.Int)

	clients := table.NewDefaultTable(ctx, cfg)

	info := clients.GetInfo().HideFilterArea()
	info.
		SetTable("client").
		SetTitle("Client").
		SetDescription("Client").
		SetPrimaryKey("client_id", db.Int)
	info.AddField("Client ID", "client_id", db.Int)
	info.AddField("First Name", "first_name", db.Varchar)
	info.AddField("Last Name", "last_name", db.Varchar)
	info.AddField("Email", "email", db.Varchar)
	info.AddField("Telephone", "telephone", db.Varchar)
	info.AddField("Password Hash", "password_hash", db.Varchar)
	info.AddField("Birth Date", "birth_date", db.Date)
	info.AddField("Is Blocked", "is_blocked", db.Tinyint)
	info.AddField("Is VIP", "is_vip", db.Tinyint)
	info.AddField("Created At", "created_at", db.Timestamp)
	info.AddField("Updated At", "updated_at", db.Timestamp)

	formList := clients.GetForm()
	formList.
		SetTable("client").
		SetTitle("Client").
		SetDescription("Client").
		SetPrimaryKey("client_id", db.Int)
	formList.AddField("Client ID", "client_id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("First Name", "first_name", db.Varchar, form.Text)
	formList.AddField("Last Name", "last_name", db.Varchar, form.Text)
	formList.AddField("Email", "email", db.Varchar, form.Text)
	formList.AddField("Telephone", "telephone", db.Varchar, form.Text)
	formList.AddField("Password Hash", "password_hash", db.Varchar, form.Text)
	formList.AddField("Birth Date", "birth_date", db.Date, form.Text)
	formList.AddField("Is Blocked", "is_blocked", db.Tinyint, form.Number)
	formList.AddField("Is VIP", "is_vip", db.Tinyint, form.Number)
	formList.AddField("Created At", "created_at", db.Timestamp, form.Text)
	formList.AddField("Updated At", "updated_at", db.Timestamp, form.Text)

	return clients
}
