// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommandsColumns holds the columns for the "commands" table.
	CommandsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "keyword", Type: field.TypeString},
		{Name: "language", Type: field.TypeString, Default: "javascript"},
		{Name: "detail", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "creator", Type: field.TypeString},
		{Name: "server", Type: field.TypeString},
		{Name: "code", Type: field.TypeString},
	}
	// CommandsTable holds the schema information for the "commands" table.
	CommandsTable = &schema.Table{
		Name:       "commands",
		Columns:    CommandsColumns,
		PrimaryKey: []*schema.Column{CommandsColumns[0]},
	}
	// ResultLogsColumns holds the columns for the "result_logs" table.
	ResultLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "level", Type: field.TypeString},
		{Name: "log", Type: field.TypeString},
	}
	// ResultLogsTable holds the schema information for the "result_logs" table.
	ResultLogsTable = &schema.Table{
		Name:       "result_logs",
		Columns:    ResultLogsColumns,
		PrimaryKey: []*schema.Column{ResultLogsColumns[0]},
	}
	// CommandLogsColumns holds the columns for the "command_logs" table.
	CommandLogsColumns = []*schema.Column{
		{Name: "command_id", Type: field.TypeString},
		{Name: "result_log_id", Type: field.TypeInt},
	}
	// CommandLogsTable holds the schema information for the "command_logs" table.
	CommandLogsTable = &schema.Table{
		Name:       "command_logs",
		Columns:    CommandLogsColumns,
		PrimaryKey: []*schema.Column{CommandLogsColumns[0], CommandLogsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "command_logs_command_id",
				Columns:    []*schema.Column{CommandLogsColumns[0]},
				RefColumns: []*schema.Column{CommandsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "command_logs_result_log_id",
				Columns:    []*schema.Column{CommandLogsColumns[1]},
				RefColumns: []*schema.Column{ResultLogsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommandsTable,
		ResultLogsTable,
		CommandLogsTable,
	}
)

func init() {
	CommandLogsTable.ForeignKeys[0].RefTable = CommandsTable
	CommandLogsTable.ForeignKeys[1].RefTable = ResultLogsTable
}
