// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/endot1231/ec-backend/ent/reviews"
)

// Reviews is the model entity for the Reviews schema.
type Reviews struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Contents holds the value of the "contents" field.
	Contents string `json:"contents,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReviewsQuery when eager-loading is set.
	Edges        ReviewsEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ReviewsEdges holds the relations/edges for other nodes in the graph.
type ReviewsEdges struct {
	// User holds the value of the user edge.
	User []*Users `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e ReviewsEdges) UserOrErr() ([]*Users, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Reviews) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case reviews.FieldID, reviews.FieldUserID:
			values[i] = new(sql.NullInt64)
		case reviews.FieldContents:
			values[i] = new(sql.NullString)
		case reviews.FieldCreatedAt, reviews.FieldUpdatedAt, reviews.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Reviews fields.
func (r *Reviews) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reviews.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case reviews.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				r.UserID = int(value.Int64)
			}
		case reviews.FieldContents:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contents", values[i])
			} else if value.Valid {
				r.Contents = value.String
			}
		case reviews.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case reviews.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case reviews.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				r.DeletedAt = new(time.Time)
				*r.DeletedAt = value.Time
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Reviews.
// This includes values selected through modifiers, order, etc.
func (r *Reviews) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Reviews entity.
func (r *Reviews) QueryUser() *UsersQuery {
	return NewReviewsClient(r.config).QueryUser(r)
}

// Update returns a builder for updating this Reviews.
// Note that you need to call Reviews.Unwrap() before calling this method if this Reviews
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Reviews) Update() *ReviewsUpdateOne {
	return NewReviewsClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Reviews entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Reviews) Unwrap() *Reviews {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Reviews is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Reviews) String() string {
	var builder strings.Builder
	builder.WriteString("Reviews(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", r.UserID))
	builder.WriteString(", ")
	builder.WriteString("contents=")
	builder.WriteString(r.Contents)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := r.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// ReviewsSlice is a parsable slice of Reviews.
type ReviewsSlice []*Reviews
