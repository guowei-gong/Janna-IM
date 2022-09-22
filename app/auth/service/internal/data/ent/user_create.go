// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Janna-IM/app/auth/service/internal/data/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (uc *UserCreate) SetUserID(s string) *UserCreate {
	uc.mutation.SetUserID(s)
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetFaceURL sets the "face_url" field.
func (uc *UserCreate) SetFaceURL(s string) *UserCreate {
	uc.mutation.SetFaceURL(s)
	return uc
}

// SetPhoneNumber sets the "phone_number" field.
func (uc *UserCreate) SetPhoneNumber(s string) *UserCreate {
	uc.mutation.SetPhoneNumber(s)
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetEx sets the "ex" field.
func (uc *UserCreate) SetEx(s string) *UserCreate {
	uc.mutation.SetEx(s)
	return uc
}

// SetCreateIP sets the "create_ip" field.
func (uc *UserCreate) SetCreateIP(s string) *UserCreate {
	uc.mutation.SetCreateIP(s)
	return uc
}

// SetLastLoginIP sets the "last_login_ip" field.
func (uc *UserCreate) SetLastLoginIP(s string) *UserCreate {
	uc.mutation.SetLastLoginIP(s)
	return uc
}

// SetInvitationCode sets the "invitation_code" field.
func (uc *UserCreate) SetInvitationCode(s string) *UserCreate {
	uc.mutation.SetInvitationCode(s)
	return uc
}

// SetGender sets the "gender" field.
func (uc *UserCreate) SetGender(i int32) *UserCreate {
	uc.mutation.SetGender(i)
	return uc
}

// SetLoginTimes sets the "login_times" field.
func (uc *UserCreate) SetLoginTimes(i int32) *UserCreate {
	uc.mutation.SetLoginTimes(i)
	return uc
}

// SetLoginLimit sets the "login_limit" field.
func (uc *UserCreate) SetLoginLimit(i int32) *UserCreate {
	uc.mutation.SetLoginLimit(i)
	return uc
}

// SetAppMangerLevel sets the "app_manger_level" field.
func (uc *UserCreate) SetAppMangerLevel(i int32) *UserCreate {
	uc.mutation.SetAppMangerLevel(i)
	return uc
}

// SetGlobalRecvMsgOpt sets the "global_recv_msg_opt" field.
func (uc *UserCreate) SetGlobalRecvMsgOpt(i int32) *UserCreate {
	uc.mutation.SetGlobalRecvMsgOpt(i)
	return uc
}

// SetStatus sets the "status" field.
func (uc *UserCreate) SetStatus(i int32) *UserCreate {
	uc.mutation.SetStatus(i)
	return uc
}

// SetBirth sets the "birth" field.
func (uc *UserCreate) SetBirth(t time.Time) *UserCreate {
	uc.mutation.SetBirth(t)
	return uc
}

// SetCreateTime sets the "create_time" field.
func (uc *UserCreate) SetCreateTime(t time.Time) *UserCreate {
	uc.mutation.SetCreateTime(t)
	return uc
}

// SetLastLoginTime sets the "last_login_time" field.
func (uc *UserCreate) SetLastLoginTime(t time.Time) *UserCreate {
	uc.mutation.SetLastLoginTime(t)
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "User.user_id"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if _, ok := uc.mutation.FaceURL(); !ok {
		return &ValidationError{Name: "face_url", err: errors.New(`ent: missing required field "User.face_url"`)}
	}
	if _, ok := uc.mutation.PhoneNumber(); !ok {
		return &ValidationError{Name: "phone_number", err: errors.New(`ent: missing required field "User.phone_number"`)}
	}
	if v, ok := uc.mutation.PhoneNumber(); ok {
		if err := user.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`ent: validator failed for field "User.phone_number": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "User.email"`)}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Ex(); !ok {
		return &ValidationError{Name: "ex", err: errors.New(`ent: missing required field "User.ex"`)}
	}
	if v, ok := uc.mutation.Ex(); ok {
		if err := user.ExValidator(v); err != nil {
			return &ValidationError{Name: "ex", err: fmt.Errorf(`ent: validator failed for field "User.ex": %w`, err)}
		}
	}
	if _, ok := uc.mutation.CreateIP(); !ok {
		return &ValidationError{Name: "create_ip", err: errors.New(`ent: missing required field "User.create_ip"`)}
	}
	if v, ok := uc.mutation.CreateIP(); ok {
		if err := user.CreateIPValidator(v); err != nil {
			return &ValidationError{Name: "create_ip", err: fmt.Errorf(`ent: validator failed for field "User.create_ip": %w`, err)}
		}
	}
	if _, ok := uc.mutation.LastLoginIP(); !ok {
		return &ValidationError{Name: "last_login_ip", err: errors.New(`ent: missing required field "User.last_login_ip"`)}
	}
	if v, ok := uc.mutation.LastLoginIP(); ok {
		if err := user.LastLoginIPValidator(v); err != nil {
			return &ValidationError{Name: "last_login_ip", err: fmt.Errorf(`ent: validator failed for field "User.last_login_ip": %w`, err)}
		}
	}
	if _, ok := uc.mutation.InvitationCode(); !ok {
		return &ValidationError{Name: "invitation_code", err: errors.New(`ent: missing required field "User.invitation_code"`)}
	}
	if _, ok := uc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "User.gender"`)}
	}
	if _, ok := uc.mutation.LoginTimes(); !ok {
		return &ValidationError{Name: "login_times", err: errors.New(`ent: missing required field "User.login_times"`)}
	}
	if _, ok := uc.mutation.LoginLimit(); !ok {
		return &ValidationError{Name: "login_limit", err: errors.New(`ent: missing required field "User.login_limit"`)}
	}
	if _, ok := uc.mutation.AppMangerLevel(); !ok {
		return &ValidationError{Name: "app_manger_level", err: errors.New(`ent: missing required field "User.app_manger_level"`)}
	}
	if _, ok := uc.mutation.GlobalRecvMsgOpt(); !ok {
		return &ValidationError{Name: "global_recv_msg_opt", err: errors.New(`ent: missing required field "User.global_recv_msg_opt"`)}
	}
	if _, ok := uc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "User.status"`)}
	}
	if _, ok := uc.mutation.Birth(); !ok {
		return &ValidationError{Name: "birth", err: errors.New(`ent: missing required field "User.birth"`)}
	}
	if _, ok := uc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "User.create_time"`)}
	}
	if _, ok := uc.mutation.LastLoginTime(); !ok {
		return &ValidationError{Name: "last_login_time", err: errors.New(`ent: missing required field "User.last_login_time"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
		_node.Name = value
	}
	if value, ok := uc.mutation.FaceURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFaceURL,
		})
		_node.FaceURL = value
	}
	if value, ok := uc.mutation.PhoneNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhoneNumber,
		})
		_node.PhoneNumber = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := uc.mutation.Ex(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEx,
		})
		_node.Ex = value
	}
	if value, ok := uc.mutation.CreateIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldCreateIP,
		})
		_node.CreateIP = value
	}
	if value, ok := uc.mutation.LastLoginIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLastLoginIP,
		})
		_node.LastLoginIP = value
	}
	if value, ok := uc.mutation.InvitationCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldInvitationCode,
		})
		_node.InvitationCode = value
	}
	if value, ok := uc.mutation.Gender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldGender,
		})
		_node.Gender = value
	}
	if value, ok := uc.mutation.LoginTimes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldLoginTimes,
		})
		_node.LoginTimes = value
	}
	if value, ok := uc.mutation.LoginLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldLoginLimit,
		})
		_node.LoginLimit = value
	}
	if value, ok := uc.mutation.AppMangerLevel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldAppMangerLevel,
		})
		_node.AppMangerLevel = value
	}
	if value, ok := uc.mutation.GlobalRecvMsgOpt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldGlobalRecvMsgOpt,
		})
		_node.GlobalRecvMsgOpt = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := uc.mutation.Birth(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldBirth,
		})
		_node.Birth = value
	}
	if value, ok := uc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := uc.mutation.LastLoginTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldLastLoginTime,
		})
		_node.LastLoginTime = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
