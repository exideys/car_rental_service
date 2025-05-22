package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetClientevaluationTable(ctx *context.Context) table.Table {

	clientEvaluation := table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := clientEvaluation.GetInfo().HideFilterArea()

	info.SetTable("client_evaluation").SetTitle("Clientevaluation").SetDescription("Clientevaluation")

	formList := clientEvaluation.GetForm()

	formList.SetTable("client_evaluation").SetTitle("Clientevaluation").SetDescription("Clientevaluation")

	return clientEvaluation
}
