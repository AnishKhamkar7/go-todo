// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/AnishKhamkar7/todo-api/ent/todo"
	"github.com/AnishKhamkar7/todo-api/ent/user"
)

// Todo is the model entity for the Todo schema.
type Todo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Types holds the value of the "types" field.
	Types []string `json:"types,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Status holds the value of the "status" field.
	Status todo.Status `json:"status,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority todo.Priority `json:"priority,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TodoQuery when eager-loading is set.
	Edges        TodoEdges `json:"edges"`
	user_todos   *int
	selectValues sql.SelectValues
}

// TodoEdges holds the relations/edges for other nodes in the graph.
type TodoEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TodoEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Todo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case todo.FieldTypes:
			values[i] = new([]byte)
		case todo.FieldID:
			values[i] = new(sql.NullInt64)
		case todo.FieldTitle, todo.FieldStatus, todo.FieldPriority:
			values[i] = new(sql.NullString)
		case todo.ForeignKeys[0]: // user_todos
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Todo fields.
func (t *Todo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case todo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case todo.FieldTypes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field types", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Types); err != nil {
					return fmt.Errorf("unmarshal field types: %w", err)
				}
			}
		case todo.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				t.Title = value.String
			}
		case todo.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = todo.Status(value.String)
			}
		case todo.FieldPriority:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				t.Priority = todo.Priority(value.String)
			}
		case todo.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_todos", value)
			} else if value.Valid {
				t.user_todos = new(int)
				*t.user_todos = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Todo.
// This includes values selected through modifiers, order, etc.
func (t *Todo) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Todo entity.
func (t *Todo) QueryUser() *UserQuery {
	return NewTodoClient(t.config).QueryUser(t)
}

// Update returns a builder for updating this Todo.
// Note that you need to call Todo.Unwrap() before calling this method if this Todo
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Todo) Update() *TodoUpdateOne {
	return NewTodoClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Todo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Todo) Unwrap() *Todo {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Todo is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Todo) String() string {
	var builder strings.Builder
	builder.WriteString("Todo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("types=")
	builder.WriteString(fmt.Sprintf("%v", t.Types))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(t.Title)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", t.Priority))
	builder.WriteByte(')')
	return builder.String()
}

// Todos is a parsable slice of Todo.
type Todos []*Todo
