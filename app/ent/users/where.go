// Code generated by ent, DO NOT EDIT.

package users

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/endot1231/ec-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldEmail, v))
}

// EmailVerified applies equality check predicate on the "email_verified" field. It's identical to EmailVerifiedEQ.
func EmailVerified(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldEmailVerified, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldPassword, v))
}

// RememberToken applies equality check predicate on the "remember_token" field. It's identical to RememberTokenEQ.
func RememberToken(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldRememberToken, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldDeletedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldEmail, v))
}

// EmailVerifiedEQ applies the EQ predicate on the "email_verified" field.
func EmailVerifiedEQ(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldEmailVerified, v))
}

// EmailVerifiedNEQ applies the NEQ predicate on the "email_verified" field.
func EmailVerifiedNEQ(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldEmailVerified, v))
}

// EmailVerifiedIn applies the In predicate on the "email_verified" field.
func EmailVerifiedIn(vs ...time.Time) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldEmailVerified, vs...))
}

// EmailVerifiedNotIn applies the NotIn predicate on the "email_verified" field.
func EmailVerifiedNotIn(vs ...time.Time) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldEmailVerified, vs...))
}

// EmailVerifiedGT applies the GT predicate on the "email_verified" field.
func EmailVerifiedGT(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldEmailVerified, v))
}

// EmailVerifiedGTE applies the GTE predicate on the "email_verified" field.
func EmailVerifiedGTE(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldEmailVerified, v))
}

// EmailVerifiedLT applies the LT predicate on the "email_verified" field.
func EmailVerifiedLT(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldEmailVerified, v))
}

// EmailVerifiedLTE applies the LTE predicate on the "email_verified" field.
func EmailVerifiedLTE(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldEmailVerified, v))
}

// EmailVerifiedIsNil applies the IsNil predicate on the "email_verified" field.
func EmailVerifiedIsNil() predicate.Users {
	return predicate.Users(sql.FieldIsNull(FieldEmailVerified))
}

// EmailVerifiedNotNil applies the NotNil predicate on the "email_verified" field.
func EmailVerifiedNotNil() predicate.Users {
	return predicate.Users(sql.FieldNotNull(FieldEmailVerified))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldPassword, v))
}

// RememberTokenEQ applies the EQ predicate on the "remember_token" field.
func RememberTokenEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldRememberToken, v))
}

// RememberTokenNEQ applies the NEQ predicate on the "remember_token" field.
func RememberTokenNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldRememberToken, v))
}

// RememberTokenIn applies the In predicate on the "remember_token" field.
func RememberTokenIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldRememberToken, vs...))
}

// RememberTokenNotIn applies the NotIn predicate on the "remember_token" field.
func RememberTokenNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldRememberToken, vs...))
}

// RememberTokenGT applies the GT predicate on the "remember_token" field.
func RememberTokenGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldRememberToken, v))
}

// RememberTokenGTE applies the GTE predicate on the "remember_token" field.
func RememberTokenGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldRememberToken, v))
}

// RememberTokenLT applies the LT predicate on the "remember_token" field.
func RememberTokenLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldRememberToken, v))
}

// RememberTokenLTE applies the LTE predicate on the "remember_token" field.
func RememberTokenLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldRememberToken, v))
}

// RememberTokenContains applies the Contains predicate on the "remember_token" field.
func RememberTokenContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldRememberToken, v))
}

// RememberTokenHasPrefix applies the HasPrefix predicate on the "remember_token" field.
func RememberTokenHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldRememberToken, v))
}

// RememberTokenHasSuffix applies the HasSuffix predicate on the "remember_token" field.
func RememberTokenHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldRememberToken, v))
}

// RememberTokenIsNil applies the IsNil predicate on the "remember_token" field.
func RememberTokenIsNil() predicate.Users {
	return predicate.Users(sql.FieldIsNull(FieldRememberToken))
}

// RememberTokenNotNil applies the NotNil predicate on the "remember_token" field.
func RememberTokenNotNil() predicate.Users {
	return predicate.Users(sql.FieldNotNull(FieldRememberToken))
}

// RememberTokenEqualFold applies the EqualFold predicate on the "remember_token" field.
func RememberTokenEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldRememberToken, v))
}

// RememberTokenContainsFold applies the ContainsFold predicate on the "remember_token" field.
func RememberTokenContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldRememberToken, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Users {
	return predicate.Users(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Users {
	return predicate.Users(sql.FieldNotNull(FieldDeletedAt))
}

// HasReviews applies the HasEdge predicate on the "reviews" edge.
func HasReviews() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ReviewsTable, ReviewsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReviewsWith applies the HasEdge predicate on the "reviews" edge with a given conditions (other predicates).
func HasReviewsWith(preds ...predicate.Reviews) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newReviewsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
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
func Not(p predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		p(s.Not())
	})
}
