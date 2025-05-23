package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetClientProfileTable(ctx *context.Context) table.Table {
	cfg := table.
		DefaultConfigWithDriverAndConnection("mysql", "default").
		SetPrimaryKey("client_id", db.Int)

	profiles := table.NewDefaultTable(ctx, cfg)

	info := profiles.GetInfo().HideFilterArea()
	info.
		SetTable("client_profile").
		SetTitle("Client Profile").
		SetDescription("Client Profile").
		SetPrimaryKey("client_id", db.Int)
	info.AddField("Client ID", "client_id", db.Int)
	info.AddField("First Name", "first_name", db.Varchar)
	info.AddField("Last Name", "last_name", db.Varchar)
	info.AddField("Email", "email", db.Varchar)
	info.AddField("Birth Date", "birth_date", db.Date)
	info.AddField("Age", "age", db.Int)
	info.AddField("Is Blocked", "is_blocked", db.Tinyint)
	info.AddField("Is VIP", "is_vip", db.Tinyint)
	info.AddField("Created At", "created_at", db.Timestamp)
	info.AddField("Updated At", "updated_at", db.Timestamp)
	info.AddField("Total Orders", "total_orders", db.Int)
	info.AddField("Total Paid", "total_paid", db.Decimal)
	info.AddField("Total Evaluations", "total_evaluations", db.Int)
	info.AddField("Avg Rating", "avg_rating", db.Decimal)
	info.AddField("Last Evaluation Date", "last_evaluation_date", db.Date)
	info.AddField("Late Return Count", "late_return_count", db.Decimal)

	return profiles
}
