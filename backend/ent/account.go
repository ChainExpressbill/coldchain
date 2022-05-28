// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ChainExpressbill/coldchain/ent/account"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	// 계정 id
	ID string `json:"id,omitempty"`
	// Password holds the value of the "password" field.
	// 계정 비밀번호
	Password string `json:"-"`
	// Name holds the value of the "name" field.
	// 관리자 이름
	Name string `json:"name,omitempty"`
	// EmailAddress holds the value of the "email_address" field.
	// 이메일 주소
	EmailAddress string `json:"emailAddress,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	// 생성 시간
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	// 수정 시간
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges AccountEdges `json:"edges"`
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// Orders holds the value of the orders edge.
	Orders []*Order `json:"orders,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrdersOrErr returns the Orders value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) OrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[0] {
		return e.Orders, nil
	}
	return nil, &NotLoadedError{edge: "orders"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldID, account.FieldPassword, account.FieldName, account.FieldEmailAddress:
			values[i] = new(sql.NullString)
		case account.FieldCreatedAt, account.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Account", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = value.String
			}
		case account.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				a.Password = value.String
			}
		case account.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case account.FieldEmailAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email_address", values[i])
			} else if value.Valid {
				a.EmailAddress = value.String
			}
		case account.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case account.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryOrders queries the "orders" edge of the Account entity.
func (a *Account) QueryOrders() *OrderQuery {
	return (&AccountClient{config: a.config}).QueryOrders(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return (&AccountClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", password=<sensitive>")
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", email_address=")
	builder.WriteString(a.EmailAddress)
	builder.WriteString(", createdAt=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account

func (a Accounts) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
