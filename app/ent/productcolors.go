// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/endot1231/ec-backend/ent/productcolors"
)

// ProductColors is the model entity for the ProductColors schema.
type ProductColors struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductColors) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case productcolors.FieldID:
			values[i] = new(sql.NullInt64)
		case productcolors.FieldName:
			values[i] = new(sql.NullString)
		case productcolors.FieldCreatedAt, productcolors.FieldUpdatedAt, productcolors.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductColors fields.
func (pc *ProductColors) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productcolors.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pc.ID = int(value.Int64)
		case productcolors.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pc.Name = value.String
			}
		case productcolors.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pc.CreatedAt = value.Time
			}
		case productcolors.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pc.UpdatedAt = value.Time
			}
		case productcolors.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pc.DeletedAt = new(time.Time)
				*pc.DeletedAt = value.Time
			}
		default:
			pc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProductColors.
// This includes values selected through modifiers, order, etc.
func (pc *ProductColors) Value(name string) (ent.Value, error) {
	return pc.selectValues.Get(name)
}

// Update returns a builder for updating this ProductColors.
// Note that you need to call ProductColors.Unwrap() before calling this method if this ProductColors
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *ProductColors) Update() *ProductColorsUpdateOne {
	return NewProductColorsClient(pc.config).UpdateOne(pc)
}

// Unwrap unwraps the ProductColors entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *ProductColors) Unwrap() *ProductColors {
	_tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductColors is not a transactional entity")
	}
	pc.config.driver = _tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *ProductColors) String() string {
	var builder strings.Builder
	builder.WriteString("ProductColors(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pc.ID))
	builder.WriteString("name=")
	builder.WriteString(pc.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := pc.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// ProductColorsSlice is a parsable slice of ProductColors.
type ProductColorsSlice []*ProductColors
