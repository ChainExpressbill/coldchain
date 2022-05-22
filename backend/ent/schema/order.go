package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("oid", uuid.UUID{}).
			Unique().
			Default(uuid.New).
			Comment("주문 ID"),
		field.String("orderer").NotEmpty().Comment("주문자(제약업체. 위탁배송을 요청한 업체)"),
		field.String("receiver").NotEmpty().Comment("수령자 (주문 요청 업체(약국, 병원 등))"),
		field.String("drug_name").NotEmpty().
			StructTag(`json:"drugName,omitempty"`).
			Comment("약품명"),
		field.String("drug_standard").NotEmpty().
			StructTag(`json:"drugStandard,omitempty"`).
			Comment("규격"),
		field.Int("quantity").Comment("수량"),
		field.String("register_name").
			StructTag(`json:"registerName,omitempty"`).Comment("주문 등록자"),
		field.String("storage_condition").NotEmpty().
			StructTag(`json:"storageCondition,omitempty"`).Comment("보관조건"),
		field.Time("created").Default(time.Now()).Immutable().Comment("생성 시간"),
		field.Time("updated").Default(time.Now()).Comment("수정 시간"),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("manager", Account.Type).
			Ref("orders").
			Required().
			Unique(),
	}
}
