// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/crowdsecurity/crowdsec/pkg/database/ent/machine"
	"github.com/facebook/ent/dialect/sql"
)

// Machine is the model entity for the Machine schema.
type Machine struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// MachineId holds the value of the "machineId" field.
	MachineId string `json:"machineId,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// IpAddress holds the value of the "ipAddress" field.
	IpAddress string `json:"ipAddress,omitempty"`
	// IsValidated holds the value of the "isValidated" field.
	IsValidated bool `json:"isValidated,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MachineQuery when eager-loading is set.
	Edges MachineEdges `json:"edges"`
}

// MachineEdges holds the relations/edges for other nodes in the graph.
type MachineEdges struct {
	// Alerts holds the value of the alerts edge.
	Alerts []*Alert
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AlertsOrErr returns the Alerts value or an error if the edge
// was not loaded in eager-loading.
func (e MachineEdges) AlertsOrErr() ([]*Alert, error) {
	if e.loadedTypes[0] {
		return e.Alerts, nil
	}
	return nil, &NotLoadedError{edge: "alerts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Machine) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullString{}, // machineId
		&sql.NullString{}, // password
		&sql.NullString{}, // ipAddress
		&sql.NullBool{},   // isValidated
		&sql.NullString{}, // status
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Machine fields.
func (m *Machine) assignValues(values ...interface{}) error {
	if m, n := len(values), len(machine.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[0])
	} else if value.Valid {
		m.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[1])
	} else if value.Valid {
		m.UpdatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field machineId", values[2])
	} else if value.Valid {
		m.MachineId = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[3])
	} else if value.Valid {
		m.Password = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field ipAddress", values[4])
	} else if value.Valid {
		m.IpAddress = value.String
	}
	if value, ok := values[5].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field isValidated", values[5])
	} else if value.Valid {
		m.IsValidated = value.Bool
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[6])
	} else if value.Valid {
		m.Status = value.String
	}
	return nil
}

// QueryAlerts queries the alerts edge of the Machine.
func (m *Machine) QueryAlerts() *AlertQuery {
	return (&MachineClient{config: m.config}).QueryAlerts(m)
}

// Update returns a builder for updating this Machine.
// Note that, you need to call Machine.Unwrap() before calling this method, if this Machine
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Machine) Update() *MachineUpdateOne {
	return (&MachineClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Machine) Unwrap() *Machine {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Machine is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Machine) String() string {
	var builder strings.Builder
	builder.WriteString("Machine(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", machineId=")
	builder.WriteString(m.MachineId)
	builder.WriteString(", password=<sensitive>")
	builder.WriteString(", ipAddress=")
	builder.WriteString(m.IpAddress)
	builder.WriteString(", isValidated=")
	builder.WriteString(fmt.Sprintf("%v", m.IsValidated))
	builder.WriteString(", status=")
	builder.WriteString(m.Status)
	builder.WriteByte(')')
	return builder.String()
}

// Machines is a parsable slice of Machine.
type Machines []*Machine

func (m Machines) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
