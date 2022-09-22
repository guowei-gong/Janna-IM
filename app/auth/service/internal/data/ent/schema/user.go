package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").Unique(),
		field.String("name"),
		field.String("face_url"),
		field.String("phone_number").MaxLen(32),
		field.String("email").MaxLen(64),
		field.String("ex").MaxLen(1024),
		field.String("create_ip").MaxLen(15),
		field.String("last_login_ip").MaxLen(15),
		field.String("invitation_code"),
		field.Int32("gender"),
		field.Int32("login_times"),
		field.Int32("login_limit"),
		field.Int32("app_manger_level"),
		field.Int32("global_recv_msg_opt"),
		field.Int32("status"),
		field.Time("birth"),
		field.Time("create_time"),
		field.Time("last_login_time"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
