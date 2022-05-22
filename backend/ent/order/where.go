// Code generated by entc, DO NOT EDIT.

package order

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ChainExpressbill/coldchain/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Orderer applies equality check predicate on the "orderer" field. It's identical to OrdererEQ.
func Orderer(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderer), v))
	})
}

// Receiver applies equality check predicate on the "receiver" field. It's identical to ReceiverEQ.
func Receiver(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReceiver), v))
	})
}

// DrugName applies equality check predicate on the "drug_name" field. It's identical to DrugNameEQ.
func DrugName(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDrugName), v))
	})
}

// DrugStandard applies equality check predicate on the "drug_standard" field. It's identical to DrugStandardEQ.
func DrugStandard(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDrugStandard), v))
	})
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// RegisterName applies equality check predicate on the "register_name" field. It's identical to RegisterNameEQ.
func RegisterName(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterName), v))
	})
}

// StorageCondition applies equality check predicate on the "storage_condition" field. It's identical to StorageConditionEQ.
func StorageCondition(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStorageCondition), v))
	})
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// Updated applies equality check predicate on the "updated" field. It's identical to UpdatedEQ.
func Updated(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdated), v))
	})
}

// OrdererEQ applies the EQ predicate on the "orderer" field.
func OrdererEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderer), v))
	})
}

// OrdererNEQ applies the NEQ predicate on the "orderer" field.
func OrdererNEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderer), v))
	})
}

// OrdererIn applies the In predicate on the "orderer" field.
func OrdererIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOrderer), v...))
	})
}

// OrdererNotIn applies the NotIn predicate on the "orderer" field.
func OrdererNotIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrderer), v...))
	})
}

// OrdererGT applies the GT predicate on the "orderer" field.
func OrdererGT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderer), v))
	})
}

// OrdererGTE applies the GTE predicate on the "orderer" field.
func OrdererGTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderer), v))
	})
}

// OrdererLT applies the LT predicate on the "orderer" field.
func OrdererLT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderer), v))
	})
}

// OrdererLTE applies the LTE predicate on the "orderer" field.
func OrdererLTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderer), v))
	})
}

// OrdererContains applies the Contains predicate on the "orderer" field.
func OrdererContains(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOrderer), v))
	})
}

// OrdererHasPrefix applies the HasPrefix predicate on the "orderer" field.
func OrdererHasPrefix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOrderer), v))
	})
}

// OrdererHasSuffix applies the HasSuffix predicate on the "orderer" field.
func OrdererHasSuffix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOrderer), v))
	})
}

// OrdererEqualFold applies the EqualFold predicate on the "orderer" field.
func OrdererEqualFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOrderer), v))
	})
}

// OrdererContainsFold applies the ContainsFold predicate on the "orderer" field.
func OrdererContainsFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOrderer), v))
	})
}

// ReceiverEQ applies the EQ predicate on the "receiver" field.
func ReceiverEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReceiver), v))
	})
}

// ReceiverNEQ applies the NEQ predicate on the "receiver" field.
func ReceiverNEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReceiver), v))
	})
}

// ReceiverIn applies the In predicate on the "receiver" field.
func ReceiverIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReceiver), v...))
	})
}

// ReceiverNotIn applies the NotIn predicate on the "receiver" field.
func ReceiverNotIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReceiver), v...))
	})
}

// ReceiverGT applies the GT predicate on the "receiver" field.
func ReceiverGT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReceiver), v))
	})
}

// ReceiverGTE applies the GTE predicate on the "receiver" field.
func ReceiverGTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReceiver), v))
	})
}

// ReceiverLT applies the LT predicate on the "receiver" field.
func ReceiverLT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReceiver), v))
	})
}

// ReceiverLTE applies the LTE predicate on the "receiver" field.
func ReceiverLTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReceiver), v))
	})
}

// ReceiverContains applies the Contains predicate on the "receiver" field.
func ReceiverContains(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReceiver), v))
	})
}

// ReceiverHasPrefix applies the HasPrefix predicate on the "receiver" field.
func ReceiverHasPrefix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReceiver), v))
	})
}

// ReceiverHasSuffix applies the HasSuffix predicate on the "receiver" field.
func ReceiverHasSuffix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReceiver), v))
	})
}

// ReceiverEqualFold applies the EqualFold predicate on the "receiver" field.
func ReceiverEqualFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReceiver), v))
	})
}

// ReceiverContainsFold applies the ContainsFold predicate on the "receiver" field.
func ReceiverContainsFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReceiver), v))
	})
}

// DrugNameEQ applies the EQ predicate on the "drug_name" field.
func DrugNameEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDrugName), v))
	})
}

// DrugNameNEQ applies the NEQ predicate on the "drug_name" field.
func DrugNameNEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDrugName), v))
	})
}

// DrugNameIn applies the In predicate on the "drug_name" field.
func DrugNameIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDrugName), v...))
	})
}

// DrugNameNotIn applies the NotIn predicate on the "drug_name" field.
func DrugNameNotIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDrugName), v...))
	})
}

// DrugNameGT applies the GT predicate on the "drug_name" field.
func DrugNameGT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDrugName), v))
	})
}

// DrugNameGTE applies the GTE predicate on the "drug_name" field.
func DrugNameGTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDrugName), v))
	})
}

// DrugNameLT applies the LT predicate on the "drug_name" field.
func DrugNameLT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDrugName), v))
	})
}

// DrugNameLTE applies the LTE predicate on the "drug_name" field.
func DrugNameLTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDrugName), v))
	})
}

// DrugNameContains applies the Contains predicate on the "drug_name" field.
func DrugNameContains(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDrugName), v))
	})
}

// DrugNameHasPrefix applies the HasPrefix predicate on the "drug_name" field.
func DrugNameHasPrefix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDrugName), v))
	})
}

// DrugNameHasSuffix applies the HasSuffix predicate on the "drug_name" field.
func DrugNameHasSuffix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDrugName), v))
	})
}

// DrugNameEqualFold applies the EqualFold predicate on the "drug_name" field.
func DrugNameEqualFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDrugName), v))
	})
}

// DrugNameContainsFold applies the ContainsFold predicate on the "drug_name" field.
func DrugNameContainsFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDrugName), v))
	})
}

// DrugStandardEQ applies the EQ predicate on the "drug_standard" field.
func DrugStandardEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDrugStandard), v))
	})
}

// DrugStandardNEQ applies the NEQ predicate on the "drug_standard" field.
func DrugStandardNEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDrugStandard), v))
	})
}

// DrugStandardIn applies the In predicate on the "drug_standard" field.
func DrugStandardIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDrugStandard), v...))
	})
}

// DrugStandardNotIn applies the NotIn predicate on the "drug_standard" field.
func DrugStandardNotIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDrugStandard), v...))
	})
}

// DrugStandardGT applies the GT predicate on the "drug_standard" field.
func DrugStandardGT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDrugStandard), v))
	})
}

// DrugStandardGTE applies the GTE predicate on the "drug_standard" field.
func DrugStandardGTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDrugStandard), v))
	})
}

// DrugStandardLT applies the LT predicate on the "drug_standard" field.
func DrugStandardLT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDrugStandard), v))
	})
}

// DrugStandardLTE applies the LTE predicate on the "drug_standard" field.
func DrugStandardLTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDrugStandard), v))
	})
}

// DrugStandardContains applies the Contains predicate on the "drug_standard" field.
func DrugStandardContains(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDrugStandard), v))
	})
}

// DrugStandardHasPrefix applies the HasPrefix predicate on the "drug_standard" field.
func DrugStandardHasPrefix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDrugStandard), v))
	})
}

// DrugStandardHasSuffix applies the HasSuffix predicate on the "drug_standard" field.
func DrugStandardHasSuffix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDrugStandard), v))
	})
}

// DrugStandardEqualFold applies the EqualFold predicate on the "drug_standard" field.
func DrugStandardEqualFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDrugStandard), v))
	})
}

// DrugStandardContainsFold applies the ContainsFold predicate on the "drug_standard" field.
func DrugStandardContainsFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDrugStandard), v))
	})
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQuantity), v))
	})
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldQuantity), v...))
	})
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldQuantity), v...))
	})
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQuantity), v))
	})
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQuantity), v))
	})
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQuantity), v))
	})
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQuantity), v))
	})
}

// RegisterNameEQ applies the EQ predicate on the "register_name" field.
func RegisterNameEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterName), v))
	})
}

// RegisterNameNEQ applies the NEQ predicate on the "register_name" field.
func RegisterNameNEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegisterName), v))
	})
}

// RegisterNameIn applies the In predicate on the "register_name" field.
func RegisterNameIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRegisterName), v...))
	})
}

// RegisterNameNotIn applies the NotIn predicate on the "register_name" field.
func RegisterNameNotIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRegisterName), v...))
	})
}

// RegisterNameGT applies the GT predicate on the "register_name" field.
func RegisterNameGT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegisterName), v))
	})
}

// RegisterNameGTE applies the GTE predicate on the "register_name" field.
func RegisterNameGTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegisterName), v))
	})
}

// RegisterNameLT applies the LT predicate on the "register_name" field.
func RegisterNameLT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegisterName), v))
	})
}

// RegisterNameLTE applies the LTE predicate on the "register_name" field.
func RegisterNameLTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegisterName), v))
	})
}

// RegisterNameContains applies the Contains predicate on the "register_name" field.
func RegisterNameContains(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRegisterName), v))
	})
}

// RegisterNameHasPrefix applies the HasPrefix predicate on the "register_name" field.
func RegisterNameHasPrefix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRegisterName), v))
	})
}

// RegisterNameHasSuffix applies the HasSuffix predicate on the "register_name" field.
func RegisterNameHasSuffix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRegisterName), v))
	})
}

// RegisterNameEqualFold applies the EqualFold predicate on the "register_name" field.
func RegisterNameEqualFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRegisterName), v))
	})
}

// RegisterNameContainsFold applies the ContainsFold predicate on the "register_name" field.
func RegisterNameContainsFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRegisterName), v))
	})
}

// StorageConditionEQ applies the EQ predicate on the "storage_condition" field.
func StorageConditionEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStorageCondition), v))
	})
}

// StorageConditionNEQ applies the NEQ predicate on the "storage_condition" field.
func StorageConditionNEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStorageCondition), v))
	})
}

// StorageConditionIn applies the In predicate on the "storage_condition" field.
func StorageConditionIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStorageCondition), v...))
	})
}

// StorageConditionNotIn applies the NotIn predicate on the "storage_condition" field.
func StorageConditionNotIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStorageCondition), v...))
	})
}

// StorageConditionGT applies the GT predicate on the "storage_condition" field.
func StorageConditionGT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStorageCondition), v))
	})
}

// StorageConditionGTE applies the GTE predicate on the "storage_condition" field.
func StorageConditionGTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStorageCondition), v))
	})
}

// StorageConditionLT applies the LT predicate on the "storage_condition" field.
func StorageConditionLT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStorageCondition), v))
	})
}

// StorageConditionLTE applies the LTE predicate on the "storage_condition" field.
func StorageConditionLTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStorageCondition), v))
	})
}

// StorageConditionContains applies the Contains predicate on the "storage_condition" field.
func StorageConditionContains(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStorageCondition), v))
	})
}

// StorageConditionHasPrefix applies the HasPrefix predicate on the "storage_condition" field.
func StorageConditionHasPrefix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStorageCondition), v))
	})
}

// StorageConditionHasSuffix applies the HasSuffix predicate on the "storage_condition" field.
func StorageConditionHasSuffix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStorageCondition), v))
	})
}

// StorageConditionEqualFold applies the EqualFold predicate on the "storage_condition" field.
func StorageConditionEqualFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStorageCondition), v))
	})
}

// StorageConditionContainsFold applies the ContainsFold predicate on the "storage_condition" field.
func StorageConditionContainsFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStorageCondition), v))
	})
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreated), v))
	})
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreated), v...))
	})
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreated), v...))
	})
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreated), v))
	})
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreated), v))
	})
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreated), v))
	})
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreated), v))
	})
}

// UpdatedEQ applies the EQ predicate on the "updated" field.
func UpdatedEQ(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdated), v))
	})
}

// UpdatedNEQ applies the NEQ predicate on the "updated" field.
func UpdatedNEQ(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdated), v))
	})
}

// UpdatedIn applies the In predicate on the "updated" field.
func UpdatedIn(vs ...time.Time) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdated), v...))
	})
}

// UpdatedNotIn applies the NotIn predicate on the "updated" field.
func UpdatedNotIn(vs ...time.Time) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdated), v...))
	})
}

// UpdatedGT applies the GT predicate on the "updated" field.
func UpdatedGT(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdated), v))
	})
}

// UpdatedGTE applies the GTE predicate on the "updated" field.
func UpdatedGTE(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdated), v))
	})
}

// UpdatedLT applies the LT predicate on the "updated" field.
func UpdatedLT(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdated), v))
	})
}

// UpdatedLTE applies the LTE predicate on the "updated" field.
func UpdatedLTE(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdated), v))
	})
}

// HasManager applies the HasEdge predicate on the "manager" edge.
func HasManager() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ManagerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ManagerTable, ManagerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasManagerWith applies the HasEdge predicate on the "manager" edge with a given conditions (other predicates).
func HasManagerWith(preds ...predicate.Account) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ManagerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ManagerTable, ManagerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		p(s.Not())
	})
}