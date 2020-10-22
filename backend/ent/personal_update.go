// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/ssaatw/app/ent/department"
	"github.com/ssaatw/app/ent/jobtitle"
	"github.com/ssaatw/app/ent/personal"
	"github.com/ssaatw/app/ent/predicate"
	"github.com/ssaatw/app/ent/systemmember"
)

// PersonalUpdate is the builder for updating Personal entities.
type PersonalUpdate struct {
	config
	hooks      []Hook
	mutation   *PersonalMutation
	predicates []predicate.Personal
}

// Where adds a new predicate for the builder.
func (pu *PersonalUpdate) Where(ps ...predicate.Personal) *PersonalUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetPersonalName sets the PersonalName field.
func (pu *PersonalUpdate) SetPersonalName(s string) *PersonalUpdate {
	pu.mutation.SetPersonalName(s)
	return pu
}

// SetPersonalMail sets the PersonalMail field.
func (pu *PersonalUpdate) SetPersonalMail(s string) *PersonalUpdate {
	pu.mutation.SetPersonalMail(s)
	return pu
}

// SetPersonalPhone sets the PersonalPhone field.
func (pu *PersonalUpdate) SetPersonalPhone(s string) *PersonalUpdate {
	pu.mutation.SetPersonalPhone(s)
	return pu
}

// SetPersonalDob sets the PersonalDob field.
func (pu *PersonalUpdate) SetPersonalDob(s string) *PersonalUpdate {
	pu.mutation.SetPersonalDob(s)
	return pu
}

// SetAdded sets the Added field.
func (pu *PersonalUpdate) SetAdded(t time.Time) *PersonalUpdate {
	pu.mutation.SetAdded(t)
	return pu
}

// SetNillableAdded sets the Added field if the given value is not nil.
func (pu *PersonalUpdate) SetNillableAdded(t *time.Time) *PersonalUpdate {
	if t != nil {
		pu.SetAdded(*t)
	}
	return pu
}

// SetJobtitleID sets the jobtitle edge to Jobtitle by id.
func (pu *PersonalUpdate) SetJobtitleID(id int) *PersonalUpdate {
	pu.mutation.SetJobtitleID(id)
	return pu
}

// SetNillableJobtitleID sets the jobtitle edge to Jobtitle by id if the given value is not nil.
func (pu *PersonalUpdate) SetNillableJobtitleID(id *int) *PersonalUpdate {
	if id != nil {
		pu = pu.SetJobtitleID(*id)
	}
	return pu
}

// SetJobtitle sets the jobtitle edge to Jobtitle.
func (pu *PersonalUpdate) SetJobtitle(j *Jobtitle) *PersonalUpdate {
	return pu.SetJobtitleID(j.ID)
}

// SetDepartmentID sets the department edge to Department by id.
func (pu *PersonalUpdate) SetDepartmentID(id int) *PersonalUpdate {
	pu.mutation.SetDepartmentID(id)
	return pu
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (pu *PersonalUpdate) SetNillableDepartmentID(id *int) *PersonalUpdate {
	if id != nil {
		pu = pu.SetDepartmentID(*id)
	}
	return pu
}

// SetDepartment sets the department edge to Department.
func (pu *PersonalUpdate) SetDepartment(d *Department) *PersonalUpdate {
	return pu.SetDepartmentID(d.ID)
}

// SetSystemmemberID sets the systemmember edge to Systemmember by id.
func (pu *PersonalUpdate) SetSystemmemberID(id int) *PersonalUpdate {
	pu.mutation.SetSystemmemberID(id)
	return pu
}

// SetNillableSystemmemberID sets the systemmember edge to Systemmember by id if the given value is not nil.
func (pu *PersonalUpdate) SetNillableSystemmemberID(id *int) *PersonalUpdate {
	if id != nil {
		pu = pu.SetSystemmemberID(*id)
	}
	return pu
}

// SetSystemmember sets the systemmember edge to Systemmember.
func (pu *PersonalUpdate) SetSystemmember(s *Systemmember) *PersonalUpdate {
	return pu.SetSystemmemberID(s.ID)
}

// Mutation returns the PersonalMutation object of the builder.
func (pu *PersonalUpdate) Mutation() *PersonalMutation {
	return pu.mutation
}

// ClearJobtitle clears the jobtitle edge to Jobtitle.
func (pu *PersonalUpdate) ClearJobtitle() *PersonalUpdate {
	pu.mutation.ClearJobtitle()
	return pu
}

// ClearDepartment clears the department edge to Department.
func (pu *PersonalUpdate) ClearDepartment() *PersonalUpdate {
	pu.mutation.ClearDepartment()
	return pu
}

// ClearSystemmember clears the systemmember edge to Systemmember.
func (pu *PersonalUpdate) ClearSystemmember() *PersonalUpdate {
	pu.mutation.ClearSystemmember()
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PersonalUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PersonalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PersonalUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PersonalUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PersonalUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PersonalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   personal.Table,
			Columns: personal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: personal.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.PersonalName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personal.FieldPersonalName,
		})
	}
	if value, ok := pu.mutation.PersonalMail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personal.FieldPersonalMail,
		})
	}
	if value, ok := pu.mutation.PersonalPhone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personal.FieldPersonalPhone,
		})
	}
	if value, ok := pu.mutation.PersonalDob(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personal.FieldPersonalDob,
		})
	}
	if value, ok := pu.mutation.Added(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: personal.FieldAdded,
		})
	}
	if pu.mutation.JobtitleCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.JobtitleIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.DepartmentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.DepartmentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.SystemmemberCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SystemmemberIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{personal.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PersonalUpdateOne is the builder for updating a single Personal entity.
type PersonalUpdateOne struct {
	config
	hooks    []Hook
	mutation *PersonalMutation
}

// SetPersonalName sets the PersonalName field.
func (puo *PersonalUpdateOne) SetPersonalName(s string) *PersonalUpdateOne {
	puo.mutation.SetPersonalName(s)
	return puo
}

// SetPersonalMail sets the PersonalMail field.
func (puo *PersonalUpdateOne) SetPersonalMail(s string) *PersonalUpdateOne {
	puo.mutation.SetPersonalMail(s)
	return puo
}

// SetPersonalPhone sets the PersonalPhone field.
func (puo *PersonalUpdateOne) SetPersonalPhone(s string) *PersonalUpdateOne {
	puo.mutation.SetPersonalPhone(s)
	return puo
}

// SetPersonalDob sets the PersonalDob field.
func (puo *PersonalUpdateOne) SetPersonalDob(s string) *PersonalUpdateOne {
	puo.mutation.SetPersonalDob(s)
	return puo
}

// SetAdded sets the Added field.
func (puo *PersonalUpdateOne) SetAdded(t time.Time) *PersonalUpdateOne {
	puo.mutation.SetAdded(t)
	return puo
}

// SetNillableAdded sets the Added field if the given value is not nil.
func (puo *PersonalUpdateOne) SetNillableAdded(t *time.Time) *PersonalUpdateOne {
	if t != nil {
		puo.SetAdded(*t)
	}
	return puo
}

// SetJobtitleID sets the jobtitle edge to Jobtitle by id.
func (puo *PersonalUpdateOne) SetJobtitleID(id int) *PersonalUpdateOne {
	puo.mutation.SetJobtitleID(id)
	return puo
}

// SetNillableJobtitleID sets the jobtitle edge to Jobtitle by id if the given value is not nil.
func (puo *PersonalUpdateOne) SetNillableJobtitleID(id *int) *PersonalUpdateOne {
	if id != nil {
		puo = puo.SetJobtitleID(*id)
	}
	return puo
}

// SetJobtitle sets the jobtitle edge to Jobtitle.
func (puo *PersonalUpdateOne) SetJobtitle(j *Jobtitle) *PersonalUpdateOne {
	return puo.SetJobtitleID(j.ID)
}

// SetDepartmentID sets the department edge to Department by id.
func (puo *PersonalUpdateOne) SetDepartmentID(id int) *PersonalUpdateOne {
	puo.mutation.SetDepartmentID(id)
	return puo
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (puo *PersonalUpdateOne) SetNillableDepartmentID(id *int) *PersonalUpdateOne {
	if id != nil {
		puo = puo.SetDepartmentID(*id)
	}
	return puo
}

// SetDepartment sets the department edge to Department.
func (puo *PersonalUpdateOne) SetDepartment(d *Department) *PersonalUpdateOne {
	return puo.SetDepartmentID(d.ID)
}

// SetSystemmemberID sets the systemmember edge to Systemmember by id.
func (puo *PersonalUpdateOne) SetSystemmemberID(id int) *PersonalUpdateOne {
	puo.mutation.SetSystemmemberID(id)
	return puo
}

// SetNillableSystemmemberID sets the systemmember edge to Systemmember by id if the given value is not nil.
func (puo *PersonalUpdateOne) SetNillableSystemmemberID(id *int) *PersonalUpdateOne {
	if id != nil {
		puo = puo.SetSystemmemberID(*id)
	}
	return puo
}

// SetSystemmember sets the systemmember edge to Systemmember.
func (puo *PersonalUpdateOne) SetSystemmember(s *Systemmember) *PersonalUpdateOne {
	return puo.SetSystemmemberID(s.ID)
}

// Mutation returns the PersonalMutation object of the builder.
func (puo *PersonalUpdateOne) Mutation() *PersonalMutation {
	return puo.mutation
}

// ClearJobtitle clears the jobtitle edge to Jobtitle.
func (puo *PersonalUpdateOne) ClearJobtitle() *PersonalUpdateOne {
	puo.mutation.ClearJobtitle()
	return puo
}

// ClearDepartment clears the department edge to Department.
func (puo *PersonalUpdateOne) ClearDepartment() *PersonalUpdateOne {
	puo.mutation.ClearDepartment()
	return puo
}

// ClearSystemmember clears the systemmember edge to Systemmember.
func (puo *PersonalUpdateOne) ClearSystemmember() *PersonalUpdateOne {
	puo.mutation.ClearSystemmember()
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PersonalUpdateOne) Save(ctx context.Context) (*Personal, error) {

	var (
		err  error
		node *Personal
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PersonalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PersonalUpdateOne) SaveX(ctx context.Context) *Personal {
	pe, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pe
}

// Exec executes the query on the entity.
func (puo *PersonalUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PersonalUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PersonalUpdateOne) sqlSave(ctx context.Context) (pe *Personal, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   personal.Table,
			Columns: personal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: personal.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Personal.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.PersonalName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personal.FieldPersonalName,
		})
	}
	if value, ok := puo.mutation.PersonalMail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personal.FieldPersonalMail,
		})
	}
	if value, ok := puo.mutation.PersonalPhone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personal.FieldPersonalPhone,
		})
	}
	if value, ok := puo.mutation.PersonalDob(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personal.FieldPersonalDob,
		})
	}
	if value, ok := puo.mutation.Added(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: personal.FieldAdded,
		})
	}
	if puo.mutation.JobtitleCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.JobtitleIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.DepartmentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.DepartmentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.SystemmemberCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SystemmemberIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pe = &Personal{config: puo.config}
	_spec.Assign = pe.assignValues
	_spec.ScanValues = pe.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{personal.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pe, nil
}
