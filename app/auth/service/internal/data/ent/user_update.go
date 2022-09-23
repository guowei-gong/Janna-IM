// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Janna-IM/app/auth/service/internal/data/ent/predicate"
	"Janna-IM/app/auth/service/internal/data/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUserID sets the "user_id" field.
func (uu *UserUpdate) SetUserID(s string) *UserUpdate {
	uu.mutation.SetUserID(s)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetFaceURL sets the "face_url" field.
func (uu *UserUpdate) SetFaceURL(s string) *UserUpdate {
	uu.mutation.SetFaceURL(s)
	return uu
}

// SetPhoneNumber sets the "phone_number" field.
func (uu *UserUpdate) SetPhoneNumber(s string) *UserUpdate {
	uu.mutation.SetPhoneNumber(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetEx sets the "ex" field.
func (uu *UserUpdate) SetEx(s string) *UserUpdate {
	uu.mutation.SetEx(s)
	return uu
}

// SetCreateIP sets the "create_ip" field.
func (uu *UserUpdate) SetCreateIP(s string) *UserUpdate {
	uu.mutation.SetCreateIP(s)
	return uu
}

// SetLastLoginIP sets the "last_login_ip" field.
func (uu *UserUpdate) SetLastLoginIP(s string) *UserUpdate {
	uu.mutation.SetLastLoginIP(s)
	return uu
}

// SetInvitationCode sets the "invitation_code" field.
func (uu *UserUpdate) SetInvitationCode(s string) *UserUpdate {
	uu.mutation.SetInvitationCode(s)
	return uu
}

// SetGender sets the "gender" field.
func (uu *UserUpdate) SetGender(i int32) *UserUpdate {
	uu.mutation.ResetGender()
	uu.mutation.SetGender(i)
	return uu
}

// AddGender adds i to the "gender" field.
func (uu *UserUpdate) AddGender(i int32) *UserUpdate {
	uu.mutation.AddGender(i)
	return uu
}

// SetLoginTimes sets the "login_times" field.
func (uu *UserUpdate) SetLoginTimes(i int32) *UserUpdate {
	uu.mutation.ResetLoginTimes()
	uu.mutation.SetLoginTimes(i)
	return uu
}

// AddLoginTimes adds i to the "login_times" field.
func (uu *UserUpdate) AddLoginTimes(i int32) *UserUpdate {
	uu.mutation.AddLoginTimes(i)
	return uu
}

// SetLoginLimit sets the "login_limit" field.
func (uu *UserUpdate) SetLoginLimit(i int32) *UserUpdate {
	uu.mutation.ResetLoginLimit()
	uu.mutation.SetLoginLimit(i)
	return uu
}

// AddLoginLimit adds i to the "login_limit" field.
func (uu *UserUpdate) AddLoginLimit(i int32) *UserUpdate {
	uu.mutation.AddLoginLimit(i)
	return uu
}

// SetAppMangerLevel sets the "app_manger_level" field.
func (uu *UserUpdate) SetAppMangerLevel(i int32) *UserUpdate {
	uu.mutation.ResetAppMangerLevel()
	uu.mutation.SetAppMangerLevel(i)
	return uu
}

// AddAppMangerLevel adds i to the "app_manger_level" field.
func (uu *UserUpdate) AddAppMangerLevel(i int32) *UserUpdate {
	uu.mutation.AddAppMangerLevel(i)
	return uu
}

// SetGlobalRecvMsgOpt sets the "global_recv_msg_opt" field.
func (uu *UserUpdate) SetGlobalRecvMsgOpt(i int32) *UserUpdate {
	uu.mutation.ResetGlobalRecvMsgOpt()
	uu.mutation.SetGlobalRecvMsgOpt(i)
	return uu
}

// AddGlobalRecvMsgOpt adds i to the "global_recv_msg_opt" field.
func (uu *UserUpdate) AddGlobalRecvMsgOpt(i int32) *UserUpdate {
	uu.mutation.AddGlobalRecvMsgOpt(i)
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(i int32) *UserUpdate {
	uu.mutation.ResetStatus()
	uu.mutation.SetStatus(i)
	return uu
}

// AddStatus adds i to the "status" field.
func (uu *UserUpdate) AddStatus(i int32) *UserUpdate {
	uu.mutation.AddStatus(i)
	return uu
}

// SetBirth sets the "birth" field.
func (uu *UserUpdate) SetBirth(t time.Time) *UserUpdate {
	uu.mutation.SetBirth(t)
	return uu
}

// SetCreateTime sets the "create_time" field.
func (uu *UserUpdate) SetCreateTime(t time.Time) *UserUpdate {
	uu.mutation.SetCreateTime(t)
	return uu
}

// SetLastLoginTime sets the "last_login_time" field.
func (uu *UserUpdate) SetLastLoginTime(t time.Time) *UserUpdate {
	uu.mutation.SetLastLoginTime(t)
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.PhoneNumber(); ok {
		if err := user.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`ent: validator failed for field "User.phone_number": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Ex(); ok {
		if err := user.ExValidator(v); err != nil {
			return &ValidationError{Name: "ex", err: fmt.Errorf(`ent: validator failed for field "User.ex": %w`, err)}
		}
	}
	if v, ok := uu.mutation.CreateIP(); ok {
		if err := user.CreateIPValidator(v); err != nil {
			return &ValidationError{Name: "create_ip", err: fmt.Errorf(`ent: validator failed for field "User.create_ip": %w`, err)}
		}
	}
	if v, ok := uu.mutation.LastLoginIP(); ok {
		if err := user.LastLoginIPValidator(v); err != nil {
			return &ValidationError{Name: "last_login_ip", err: fmt.Errorf(`ent: validator failed for field "User.last_login_ip": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUserID,
		})
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.FaceURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFaceURL,
		})
	}
	if value, ok := uu.mutation.PhoneNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhoneNumber,
		})
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.Ex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEx,
		})
	}
	if value, ok := uu.mutation.CreateIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldCreateIP,
		})
	}
	if value, ok := uu.mutation.LastLoginIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLastLoginIP,
		})
	}
	if value, ok := uu.mutation.InvitationCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldInvitationCode,
		})
	}
	if value, ok := uu.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldGender,
		})
	}
	if value, ok := uu.mutation.AddedGender(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldGender,
		})
	}
	if value, ok := uu.mutation.LoginTimes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldLoginTimes,
		})
	}
	if value, ok := uu.mutation.AddedLoginTimes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldLoginTimes,
		})
	}
	if value, ok := uu.mutation.LoginLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldLoginLimit,
		})
	}
	if value, ok := uu.mutation.AddedLoginLimit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldLoginLimit,
		})
	}
	if value, ok := uu.mutation.AppMangerLevel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldAppMangerLevel,
		})
	}
	if value, ok := uu.mutation.AddedAppMangerLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldAppMangerLevel,
		})
	}
	if value, ok := uu.mutation.GlobalRecvMsgOpt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldGlobalRecvMsgOpt,
		})
	}
	if value, ok := uu.mutation.AddedGlobalRecvMsgOpt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldGlobalRecvMsgOpt,
		})
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldStatus,
		})
	}
	if value, ok := uu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldStatus,
		})
	}
	if value, ok := uu.mutation.Birth(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldBirth,
		})
	}
	if value, ok := uu.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreateTime,
		})
	}
	if value, ok := uu.mutation.LastLoginTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldLastLoginTime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUserID sets the "user_id" field.
func (uuo *UserUpdateOne) SetUserID(s string) *UserUpdateOne {
	uuo.mutation.SetUserID(s)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetFaceURL sets the "face_url" field.
func (uuo *UserUpdateOne) SetFaceURL(s string) *UserUpdateOne {
	uuo.mutation.SetFaceURL(s)
	return uuo
}

// SetPhoneNumber sets the "phone_number" field.
func (uuo *UserUpdateOne) SetPhoneNumber(s string) *UserUpdateOne {
	uuo.mutation.SetPhoneNumber(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetEx sets the "ex" field.
func (uuo *UserUpdateOne) SetEx(s string) *UserUpdateOne {
	uuo.mutation.SetEx(s)
	return uuo
}

// SetCreateIP sets the "create_ip" field.
func (uuo *UserUpdateOne) SetCreateIP(s string) *UserUpdateOne {
	uuo.mutation.SetCreateIP(s)
	return uuo
}

// SetLastLoginIP sets the "last_login_ip" field.
func (uuo *UserUpdateOne) SetLastLoginIP(s string) *UserUpdateOne {
	uuo.mutation.SetLastLoginIP(s)
	return uuo
}

// SetInvitationCode sets the "invitation_code" field.
func (uuo *UserUpdateOne) SetInvitationCode(s string) *UserUpdateOne {
	uuo.mutation.SetInvitationCode(s)
	return uuo
}

// SetGender sets the "gender" field.
func (uuo *UserUpdateOne) SetGender(i int32) *UserUpdateOne {
	uuo.mutation.ResetGender()
	uuo.mutation.SetGender(i)
	return uuo
}

// AddGender adds i to the "gender" field.
func (uuo *UserUpdateOne) AddGender(i int32) *UserUpdateOne {
	uuo.mutation.AddGender(i)
	return uuo
}

// SetLoginTimes sets the "login_times" field.
func (uuo *UserUpdateOne) SetLoginTimes(i int32) *UserUpdateOne {
	uuo.mutation.ResetLoginTimes()
	uuo.mutation.SetLoginTimes(i)
	return uuo
}

// AddLoginTimes adds i to the "login_times" field.
func (uuo *UserUpdateOne) AddLoginTimes(i int32) *UserUpdateOne {
	uuo.mutation.AddLoginTimes(i)
	return uuo
}

// SetLoginLimit sets the "login_limit" field.
func (uuo *UserUpdateOne) SetLoginLimit(i int32) *UserUpdateOne {
	uuo.mutation.ResetLoginLimit()
	uuo.mutation.SetLoginLimit(i)
	return uuo
}

// AddLoginLimit adds i to the "login_limit" field.
func (uuo *UserUpdateOne) AddLoginLimit(i int32) *UserUpdateOne {
	uuo.mutation.AddLoginLimit(i)
	return uuo
}

// SetAppMangerLevel sets the "app_manger_level" field.
func (uuo *UserUpdateOne) SetAppMangerLevel(i int32) *UserUpdateOne {
	uuo.mutation.ResetAppMangerLevel()
	uuo.mutation.SetAppMangerLevel(i)
	return uuo
}

// AddAppMangerLevel adds i to the "app_manger_level" field.
func (uuo *UserUpdateOne) AddAppMangerLevel(i int32) *UserUpdateOne {
	uuo.mutation.AddAppMangerLevel(i)
	return uuo
}

// SetGlobalRecvMsgOpt sets the "global_recv_msg_opt" field.
func (uuo *UserUpdateOne) SetGlobalRecvMsgOpt(i int32) *UserUpdateOne {
	uuo.mutation.ResetGlobalRecvMsgOpt()
	uuo.mutation.SetGlobalRecvMsgOpt(i)
	return uuo
}

// AddGlobalRecvMsgOpt adds i to the "global_recv_msg_opt" field.
func (uuo *UserUpdateOne) AddGlobalRecvMsgOpt(i int32) *UserUpdateOne {
	uuo.mutation.AddGlobalRecvMsgOpt(i)
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(i int32) *UserUpdateOne {
	uuo.mutation.ResetStatus()
	uuo.mutation.SetStatus(i)
	return uuo
}

// AddStatus adds i to the "status" field.
func (uuo *UserUpdateOne) AddStatus(i int32) *UserUpdateOne {
	uuo.mutation.AddStatus(i)
	return uuo
}

// SetBirth sets the "birth" field.
func (uuo *UserUpdateOne) SetBirth(t time.Time) *UserUpdateOne {
	uuo.mutation.SetBirth(t)
	return uuo
}

// SetCreateTime sets the "create_time" field.
func (uuo *UserUpdateOne) SetCreateTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreateTime(t)
	return uuo
}

// SetLastLoginTime sets the "last_login_time" field.
func (uuo *UserUpdateOne) SetLastLoginTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetLastLoginTime(t)
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.PhoneNumber(); ok {
		if err := user.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`ent: validator failed for field "User.phone_number": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Ex(); ok {
		if err := user.ExValidator(v); err != nil {
			return &ValidationError{Name: "ex", err: fmt.Errorf(`ent: validator failed for field "User.ex": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.CreateIP(); ok {
		if err := user.CreateIPValidator(v); err != nil {
			return &ValidationError{Name: "create_ip", err: fmt.Errorf(`ent: validator failed for field "User.create_ip": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.LastLoginIP(); ok {
		if err := user.LastLoginIPValidator(v); err != nil {
			return &ValidationError{Name: "last_login_ip", err: fmt.Errorf(`ent: validator failed for field "User.last_login_ip": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUserID,
		})
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.FaceURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFaceURL,
		})
	}
	if value, ok := uuo.mutation.PhoneNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhoneNumber,
		})
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.Ex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEx,
		})
	}
	if value, ok := uuo.mutation.CreateIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldCreateIP,
		})
	}
	if value, ok := uuo.mutation.LastLoginIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLastLoginIP,
		})
	}
	if value, ok := uuo.mutation.InvitationCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldInvitationCode,
		})
	}
	if value, ok := uuo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldGender,
		})
	}
	if value, ok := uuo.mutation.AddedGender(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldGender,
		})
	}
	if value, ok := uuo.mutation.LoginTimes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldLoginTimes,
		})
	}
	if value, ok := uuo.mutation.AddedLoginTimes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldLoginTimes,
		})
	}
	if value, ok := uuo.mutation.LoginLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldLoginLimit,
		})
	}
	if value, ok := uuo.mutation.AddedLoginLimit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldLoginLimit,
		})
	}
	if value, ok := uuo.mutation.AppMangerLevel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldAppMangerLevel,
		})
	}
	if value, ok := uuo.mutation.AddedAppMangerLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldAppMangerLevel,
		})
	}
	if value, ok := uuo.mutation.GlobalRecvMsgOpt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldGlobalRecvMsgOpt,
		})
	}
	if value, ok := uuo.mutation.AddedGlobalRecvMsgOpt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldGlobalRecvMsgOpt,
		})
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldStatus,
		})
	}
	if value, ok := uuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldStatus,
		})
	}
	if value, ok := uuo.mutation.Birth(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldBirth,
		})
	}
	if value, ok := uuo.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreateTime,
		})
	}
	if value, ok := uuo.mutation.LastLoginTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldLastLoginTime,
		})
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}