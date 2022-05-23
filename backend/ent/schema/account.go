package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Comment("계정 id"),
		field.String("password").Sensitive().NotEmpty().Comment("계정 비밀번호"),
		field.String("name").NotEmpty().Comment("관리자 이름"),
		field.String("email_address").
			NotEmpty().
			StructTag(`json:"emailAddress,omitempty"`).
			Comment("이메일 주소"),
		field.Time("created").
			Default(time.Now).
			Immutable().
			Comment("생성 시간"),
		field.Time("updated").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("수정 시간"),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		// ent에서는 To를 정의하는 스키마, 즉 여기선 Account 가 참조 관계의 주인이라고 정의합니다.
		edge.To("orders", Order.Type),
	}
}
