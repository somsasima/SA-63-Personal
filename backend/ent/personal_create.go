// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/ssaatw/app/ent/department"
	"github.com/ssaatw/app/ent/jobtitle"
	"github.com/ssaatw/app/ent/personal"
	"github.com/ssaatw/app/ent/systemmember"
)

// PersonalCreate is the builder for creating a Personal entity.
type PersonalCreate struct {
	config
	mutation *PersonalMutation
	hooks    []Hook
}

// SetPersonalName sets the PersonalName field.
func (pc *PersonalCreate) SetPersonalName(s string) *PersonalCreate {
	pc.mutation.SetPersonalName(s)
	return pc
}

// SetPersonalMail sets the PersonalMail field.
func (pc *PersonalCreate) SetPersonalMail(s string) *PersonalCreate {
	pc.mutation.SetPersonalMail(s)
	return pc
}

// SetPersonalPhone sets the PersonalPhone field.
func (pc *PersonalCreate) SetPersonalPhone(s string) *PersonalCreate {
	pc.mutation.SetPersonalPhone(s)
	return pc
}

// SetPersonalDob sets the PersonalDob field.
func (pc *PersonalCreate) SetPersonalDob(s string) *PersonalCreate {
	pc.mutation.SetPersonalDob(s)
	return pc
}

// SetAdded sets the Added field.
func (pc *PersonalCreate) SetAdded(t time.Time) *PersonalCreate {
	pc.mutation.SetAdded(t)
	return pc
}

// SetNillableAdded sets the Added field if the given value is not nil.
func (pc *PersonalCreate) SetNillableAdded(t *time.Time) *PersonalCreate {
	if t != nil {
		pc.SetAdded(*t)
	}
	return pc
}

// SetJobtitleID sets the jobtitle edge to Jobtitle by id.
func (pc *PersonalCreate) SetJobtitleID(id int) *PersonalCreate {
	pc.mutation.SetJobtitleID(id)
	return pc
}

// SetNillableJobtitleID sets the jobtitle edge to Jobtitle by id if the given value is not nil.
func (pc *PersonalCreate) SetNillableJobtitleID(id *int) *PersonalCreate {
	if id != nil {
		pc = pc.SetJobtitleID(*id)
	}
	return pc
}

// SetJobtitle sets the jobtitle edge to Jobtitle.
func (pc *PersonalCreate) SetJobtitle(j *Jobtitle) *PersonalCreate {
	return pc.SetJobtitleID(j.ID)
}

// SetDepartmentID sets the department edge to Department by id.
func (pc *PersonalCreate) SetDepartmentID(id int) *PersonalCreate {
	pc.mutation.SetDepartmentID(id)
	return pc
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (pc *PersonalCreate) SetNillableDepartmentID(id *int) *PersonalCreate {
	if id != nil {
		pc = pc.SetDepartmentID(*id)
	}
	return pc
}

// SetDepartment sets the department edge to Department.
func (pc *PersonalCreate) SetDepartment(d *Department) *PersonalCreate {
	return pc.SetDepartmentID(d.ID)
}

// SetSystemmemberID sets the systemmember edge to Systemmember by id.
func (pc *PersonalCreate) SetSystemmemberID(id int) *PersonalCreate {
	pc.mutation.SetSystemmemberID(id)
	return pc
}

// SetNillableSystemmemberID sets the systemmember edge to Systemmember by id if the given value is not nil.
func (pc *PersonalCreate) SetNillableSystemmemberID(id *int) *PersonalCreate {
	if id != nil {
		pc = pc.SetSystemmemberID(*id)
	}
	return pc
}

// SetSystemmember sets the systemmember edge to Systemmember.
func (pc *PersonalCreate) SetSystemmember(s *Systemmember) *PersonalCreate {
	return pc.SetSystemmemberID(s.ID)
}

// Mutation returns the PersonalMutation object of the builder.
func (pc *PersonalCreate) Mutation() *PersonalMutation {
	return pc.mutation
}

// Save creates the Personal in the database.
func (pc *PersonalCreate) Save(ctx context.Context) (*Personal, error) {
	if _, ok := pc.mutation.PersonalName(); !ok {
		return nil, &ValidationError{Name: "PersonalName", err: errors.New("ent: missing required field \"PersonalName\"")}
	}
	if _, ok := pc.mutation.PersonalMail(); !ok {
		return nil, &ValidationError{Name: "PersonalMail", err: errors.New("ent: missing required field \"PersonalMail\"")}
	}
	if _, ok := pc.mutation.PersonalPhone(); !ok {
		return nil, &ValidationError{Name: "PersonalPhone", err: errors.New("ent: missing required field \"PersonalPhone\"")}
	}
	if _, ok := pc.mutation.PersonalDob(); !ok {
		return nil, &ValidationError{Name: "PersonalDob", err: errors.New("ent: missing required field \"PersonalDob\"")}
	}
	if _, ok := pc.mutation.Added(); !ok {
		v := personal.DefaultAdded()
		pc.mutation.SetAdded(v)
	}
	var (
		err  error
		node *Personal
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PersonalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PersonalCreate) SaveX(ctx context.Context) *Personal {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PersonalCreate) sqlSave(ctx context.Context) (*Personal, error) {
	pe, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pe.ID = int(id)
	return pe, nil
}

func (pc *PersonalCreate) createSpec() (*Personal, *sqlgraph.CreateSpec) {
	var (
		pe    = &Personal{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: personal.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: personal.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.PersonalName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personal.FieldPersonalName,
		})
		pe.PersonalName = value
	}
	if value, ok := pc.mutation.PersonalMail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personal.FieldPersonalMail,
		})
		pe.PersonalMail = value
	}
	if value, ok := pc.mutation.PersonalPhone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personal.FieldPersonalPhone,
		})
		pe.PersonalPhone = value
	}
	if value, ok := pc.mutation.PersonalDob(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personal.FieldPersonalDob,
		})
		pe.PersonalDob = value
	}
	if value, ok := pc.mutation.Added(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: personal.FieldAdded,
		})
		pe.Added = value
	}
	if nodes := pc.mutation.JobtitleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   personal.JobtitleTable,
			Columns: []string{personal.JobtitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: jobtitle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   personal.DepartmentTable,
			Columns: []string{personal.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SystemmemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   personal.SystemmemberTable,
			Columns: []string{personal.SystemmemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: systemmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pe, _spec
}
