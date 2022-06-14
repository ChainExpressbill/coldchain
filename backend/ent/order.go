// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ChainExpressbill/coldchain/ent/account"
	"github.com/ChainExpressbill/coldchain/ent/order"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Oid holds the value of the "oid" field.
	// 주문 ID
	Oid string `json:"oid,omitempty"`
	// Orderer holds the value of the "orderer" field.
	// 주문 업체(제약업체. 위탁배송을 요청한 업체)
	Orderer string `json:"orderer,omitempty"`
	// Receiver holds the value of the "receiver" field.
	// 수령 업체 (주문 요청 업체(약국, 병원 등))
	Receiver string `json:"receiver,omitempty"`
	// DrugName holds the value of the "drug_name" field.
	// 제품명
	DrugName string `json:"drugName,omitempty"`
	// DrugStandard holds the value of the "drug_standard" field.
	// 제품 규격
	DrugStandard string `json:"drugStandard,omitempty"`
	// Quantity holds the value of the "quantity" field.
	// 제품 수량
	Quantity int `json:"quantity,omitempty"`
	// RegisterName holds the value of the "register_name" field.
	// 주문 등록자
	RegisterName string `json:"registerName,omitempty"`
	// StorageCondition holds the value of the "storage_condition" field.
	// 보관조건
	StorageCondition string `json:"storageCondition,omitempty"`
	// DeliveryDriverName holds the value of the "delivery_driver_name" field.
	// 배송 기사 이름
	DeliveryDriverName string `json:"deliveryDriverName,omitempty"`
	// DeliveryDriverTelNo holds the value of the "delivery_driver_tel_no" field.
	// 배송 기사 연락처
	DeliveryDriverTelNo string `json:"deliveryDriverTelNo,omitempty"`
	// Memo holds the value of the "memo" field.
	// 메모
	Memo *string `json:"memo,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	// 생성 시간
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	// 수정 시간
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges          OrderEdges `json:"edges"`
	account_orders *string
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// Manager holds the value of the manager edge.
	Manager *Account `json:"manager,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ManagerOrErr returns the Manager value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) ManagerOrErr() (*Account, error) {
	if e.loadedTypes[0] {
		if e.Manager == nil {
			// The edge manager was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Manager, nil
	}
	return nil, &NotLoadedError{edge: "manager"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldID, order.FieldQuantity:
			values[i] = new(sql.NullInt64)
		case order.FieldOid, order.FieldOrderer, order.FieldReceiver, order.FieldDrugName, order.FieldDrugStandard, order.FieldRegisterName, order.FieldStorageCondition, order.FieldDeliveryDriverName, order.FieldDeliveryDriverTelNo, order.FieldMemo:
			values[i] = new(sql.NullString)
		case order.FieldCreatedAt, order.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case order.ForeignKeys[0]: // account_orders
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Order", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case order.FieldOid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field oid", values[i])
			} else if value.Valid {
				o.Oid = value.String
			}
		case order.FieldOrderer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field orderer", values[i])
			} else if value.Valid {
				o.Orderer = value.String
			}
		case order.FieldReceiver:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receiver", values[i])
			} else if value.Valid {
				o.Receiver = value.String
			}
		case order.FieldDrugName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field drug_name", values[i])
			} else if value.Valid {
				o.DrugName = value.String
			}
		case order.FieldDrugStandard:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field drug_standard", values[i])
			} else if value.Valid {
				o.DrugStandard = value.String
			}
		case order.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				o.Quantity = int(value.Int64)
			}
		case order.FieldRegisterName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field register_name", values[i])
			} else if value.Valid {
				o.RegisterName = value.String
			}
		case order.FieldStorageCondition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field storage_condition", values[i])
			} else if value.Valid {
				o.StorageCondition = value.String
			}
		case order.FieldDeliveryDriverName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field delivery_driver_name", values[i])
			} else if value.Valid {
				o.DeliveryDriverName = value.String
			}
		case order.FieldDeliveryDriverTelNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field delivery_driver_tel_no", values[i])
			} else if value.Valid {
				o.DeliveryDriverTelNo = value.String
			}
		case order.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				o.Memo = new(string)
				*o.Memo = value.String
			}
		case order.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case order.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		case order.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_orders", values[i])
			} else if value.Valid {
				o.account_orders = new(string)
				*o.account_orders = value.String
			}
		}
	}
	return nil
}

// QueryManager queries the "manager" edge of the Order entity.
func (o *Order) QueryManager() *AccountQuery {
	return (&OrderClient{config: o.config}).QueryManager(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return (&OrderClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", oid=")
	builder.WriteString(o.Oid)
	builder.WriteString(", orderer=")
	builder.WriteString(o.Orderer)
	builder.WriteString(", receiver=")
	builder.WriteString(o.Receiver)
	builder.WriteString(", drug_name=")
	builder.WriteString(o.DrugName)
	builder.WriteString(", drug_standard=")
	builder.WriteString(o.DrugStandard)
	builder.WriteString(", quantity=")
	builder.WriteString(fmt.Sprintf("%v", o.Quantity))
	builder.WriteString(", register_name=")
	builder.WriteString(o.RegisterName)
	builder.WriteString(", storage_condition=")
	builder.WriteString(o.StorageCondition)
	builder.WriteString(", delivery_driver_name=")
	builder.WriteString(o.DeliveryDriverName)
	builder.WriteString(", delivery_driver_tel_no=")
	builder.WriteString(o.DeliveryDriverTelNo)
	if v := o.Memo; v != nil {
		builder.WriteString(", memo=")
		builder.WriteString(*v)
	}
	builder.WriteString(", createdAt=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order

func (o Orders) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
