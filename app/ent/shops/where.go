// Code generated by ent, DO NOT EDIT.

package shops

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/endot1231/ec-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Shops {
	return predicate.Shops(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Shops {
	return predicate.Shops(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Shops {
	return predicate.Shops(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Shops {
	return predicate.Shops(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Shops {
	return predicate.Shops(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Shops {
	return predicate.Shops(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Shops {
	return predicate.Shops(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldName, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldAddress, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldEmail, v))
}

// EmailVerified applies equality check predicate on the "email_verified" field. It's identical to EmailVerifiedEQ.
func EmailVerified(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldEmailVerified, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldPassword, v))
}

// RememberToken applies equality check predicate on the "remember_token" field. It's identical to RememberTokenEQ.
func RememberToken(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldRememberToken, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldDeletedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Shops {
	return predicate.Shops(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Shops {
	return predicate.Shops(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Shops {
	return predicate.Shops(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Shops {
	return predicate.Shops(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Shops {
	return predicate.Shops(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Shops {
	return predicate.Shops(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Shops {
	return predicate.Shops(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Shops {
	return predicate.Shops(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Shops {
	return predicate.Shops(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Shops {
	return predicate.Shops(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Shops {
	return predicate.Shops(sql.FieldContainsFold(FieldName, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Shops {
	return predicate.Shops(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Shops {
	return predicate.Shops(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Shops {
	return predicate.Shops(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Shops {
	return predicate.Shops(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Shops {
	return predicate.Shops(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Shops {
	return predicate.Shops(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Shops {
	return predicate.Shops(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Shops {
	return predicate.Shops(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Shops {
	return predicate.Shops(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Shops {
	return predicate.Shops(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Shops {
	return predicate.Shops(sql.FieldContainsFold(FieldAddress, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Shops {
	return predicate.Shops(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Shops {
	return predicate.Shops(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Shops {
	return predicate.Shops(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Shops {
	return predicate.Shops(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Shops {
	return predicate.Shops(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Shops {
	return predicate.Shops(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Shops {
	return predicate.Shops(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Shops {
	return predicate.Shops(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Shops {
	return predicate.Shops(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Shops {
	return predicate.Shops(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Shops {
	return predicate.Shops(sql.FieldContainsFold(FieldEmail, v))
}

// EmailVerifiedEQ applies the EQ predicate on the "email_verified" field.
func EmailVerifiedEQ(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldEmailVerified, v))
}

// EmailVerifiedNEQ applies the NEQ predicate on the "email_verified" field.
func EmailVerifiedNEQ(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldNEQ(FieldEmailVerified, v))
}

// EmailVerifiedIn applies the In predicate on the "email_verified" field.
func EmailVerifiedIn(vs ...time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldIn(FieldEmailVerified, vs...))
}

// EmailVerifiedNotIn applies the NotIn predicate on the "email_verified" field.
func EmailVerifiedNotIn(vs ...time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldNotIn(FieldEmailVerified, vs...))
}

// EmailVerifiedGT applies the GT predicate on the "email_verified" field.
func EmailVerifiedGT(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldGT(FieldEmailVerified, v))
}

// EmailVerifiedGTE applies the GTE predicate on the "email_verified" field.
func EmailVerifiedGTE(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldGTE(FieldEmailVerified, v))
}

// EmailVerifiedLT applies the LT predicate on the "email_verified" field.
func EmailVerifiedLT(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldLT(FieldEmailVerified, v))
}

// EmailVerifiedLTE applies the LTE predicate on the "email_verified" field.
func EmailVerifiedLTE(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldLTE(FieldEmailVerified, v))
}

// EmailVerifiedIsNil applies the IsNil predicate on the "email_verified" field.
func EmailVerifiedIsNil() predicate.Shops {
	return predicate.Shops(sql.FieldIsNull(FieldEmailVerified))
}

// EmailVerifiedNotNil applies the NotNil predicate on the "email_verified" field.
func EmailVerifiedNotNil() predicate.Shops {
	return predicate.Shops(sql.FieldNotNull(FieldEmailVerified))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Shops {
	return predicate.Shops(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Shops {
	return predicate.Shops(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Shops {
	return predicate.Shops(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Shops {
	return predicate.Shops(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Shops {
	return predicate.Shops(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Shops {
	return predicate.Shops(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Shops {
	return predicate.Shops(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Shops {
	return predicate.Shops(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Shops {
	return predicate.Shops(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Shops {
	return predicate.Shops(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordIsNil applies the IsNil predicate on the "password" field.
func PasswordIsNil() predicate.Shops {
	return predicate.Shops(sql.FieldIsNull(FieldPassword))
}

// PasswordNotNil applies the NotNil predicate on the "password" field.
func PasswordNotNil() predicate.Shops {
	return predicate.Shops(sql.FieldNotNull(FieldPassword))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Shops {
	return predicate.Shops(sql.FieldContainsFold(FieldPassword, v))
}

// RememberTokenEQ applies the EQ predicate on the "remember_token" field.
func RememberTokenEQ(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldRememberToken, v))
}

// RememberTokenNEQ applies the NEQ predicate on the "remember_token" field.
func RememberTokenNEQ(v string) predicate.Shops {
	return predicate.Shops(sql.FieldNEQ(FieldRememberToken, v))
}

// RememberTokenIn applies the In predicate on the "remember_token" field.
func RememberTokenIn(vs ...string) predicate.Shops {
	return predicate.Shops(sql.FieldIn(FieldRememberToken, vs...))
}

// RememberTokenNotIn applies the NotIn predicate on the "remember_token" field.
func RememberTokenNotIn(vs ...string) predicate.Shops {
	return predicate.Shops(sql.FieldNotIn(FieldRememberToken, vs...))
}

// RememberTokenGT applies the GT predicate on the "remember_token" field.
func RememberTokenGT(v string) predicate.Shops {
	return predicate.Shops(sql.FieldGT(FieldRememberToken, v))
}

// RememberTokenGTE applies the GTE predicate on the "remember_token" field.
func RememberTokenGTE(v string) predicate.Shops {
	return predicate.Shops(sql.FieldGTE(FieldRememberToken, v))
}

// RememberTokenLT applies the LT predicate on the "remember_token" field.
func RememberTokenLT(v string) predicate.Shops {
	return predicate.Shops(sql.FieldLT(FieldRememberToken, v))
}

// RememberTokenLTE applies the LTE predicate on the "remember_token" field.
func RememberTokenLTE(v string) predicate.Shops {
	return predicate.Shops(sql.FieldLTE(FieldRememberToken, v))
}

// RememberTokenContains applies the Contains predicate on the "remember_token" field.
func RememberTokenContains(v string) predicate.Shops {
	return predicate.Shops(sql.FieldContains(FieldRememberToken, v))
}

// RememberTokenHasPrefix applies the HasPrefix predicate on the "remember_token" field.
func RememberTokenHasPrefix(v string) predicate.Shops {
	return predicate.Shops(sql.FieldHasPrefix(FieldRememberToken, v))
}

// RememberTokenHasSuffix applies the HasSuffix predicate on the "remember_token" field.
func RememberTokenHasSuffix(v string) predicate.Shops {
	return predicate.Shops(sql.FieldHasSuffix(FieldRememberToken, v))
}

// RememberTokenIsNil applies the IsNil predicate on the "remember_token" field.
func RememberTokenIsNil() predicate.Shops {
	return predicate.Shops(sql.FieldIsNull(FieldRememberToken))
}

// RememberTokenNotNil applies the NotNil predicate on the "remember_token" field.
func RememberTokenNotNil() predicate.Shops {
	return predicate.Shops(sql.FieldNotNull(FieldRememberToken))
}

// RememberTokenEqualFold applies the EqualFold predicate on the "remember_token" field.
func RememberTokenEqualFold(v string) predicate.Shops {
	return predicate.Shops(sql.FieldEqualFold(FieldRememberToken, v))
}

// RememberTokenContainsFold applies the ContainsFold predicate on the "remember_token" field.
func RememberTokenContainsFold(v string) predicate.Shops {
	return predicate.Shops(sql.FieldContainsFold(FieldRememberToken, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Shops {
	return predicate.Shops(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Shops {
	return predicate.Shops(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Shops {
	return predicate.Shops(sql.FieldNotNull(FieldDeletedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Shops) predicate.Shops {
	return predicate.Shops(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Shops) predicate.Shops {
	return predicate.Shops(func(s *sql.Selector) {
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
func Not(p predicate.Shops) predicate.Shops {
	return predicate.Shops(func(s *sql.Selector) {
		p(s.Not())
	})
}
