package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetClientEvaluationTable(ctx *context.Context) table.Table {
	cfg := table.
		DefaultConfigWithDriverAndConnection("mysql", "default").
		SetPrimaryKey("evaluation_id", db.Int)

	evals := table.NewDefaultTable(ctx, cfg)

	info := evals.GetInfo().HideFilterArea()
	info.
		SetTable("client_evaluation").
		SetTitle("Client Evaluation").
		SetDescription("Client Evaluation").
		SetPrimaryKey("evaluation_id", db.Int)
	info.AddField("Evaluation ID", "evaluation_id", db.Int)
	info.AddField("Date", "evaluation_date", db.Date)
	info.AddField("Rating", "rating", db.Int)
	info.AddField("Comment", "comment", db.Text)
	info.AddField("Is Late Return", "is_late_return", db.Tinyint)
	info.AddField("Client ID", "client_id", db.Int)

	formList := evals.GetForm()
	formList.
		SetTable("client_evaluation").
		SetTitle("Client Evaluation").
		SetDescription("Client Evaluation").
		SetPrimaryKey("evaluation_id", db.Int)
	formList.AddField("Evaluation ID", "evaluation_id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Date", "evaluation_date", db.Date, form.Text)
	formList.AddField("Rating", "rating", db.Int, form.Number)
	formList.AddField("Comment", "comment", db.Text, form.Text)
	formList.AddField("Is Late Return", "is_late_return", db.Tinyint, form.Number)
	formList.AddField("Client ID", "client_id", db.Int, form.Number)

	return evals
}
