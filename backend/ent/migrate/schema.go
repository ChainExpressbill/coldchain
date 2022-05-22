// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "email_address", Type: field.TypeString},
		{Name: "created", Type: field.TypeTime},
		{Name: "updated", Type: field.TypeTime},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "oid", Type: field.TypeUUID, Unique: true},
		{Name: "orderer", Type: field.TypeString},
		{Name: "receiver", Type: field.TypeString},
		{Name: "drug_name", Type: field.TypeString},
		{Name: "drug_standard", Type: field.TypeString},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "register_name", Type: field.TypeString},
		{Name: "storage_condition", Type: field.TypeString},
		{Name: "created", Type: field.TypeTime},
		{Name: "updated", Type: field.TypeTime},
		{Name: "account_orders", Type: field.TypeString},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "orders_accounts_orders",
				Columns:    []*schema.Column{OrdersColumns[11]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		OrdersTable,
	}
)

func init() {
	OrdersTable.ForeignKeys[0].RefTable = AccountsTable
}
