package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent/privacy"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/rule"
)

// User holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Mixin of the User schema.
func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		XidGqlMixin{},
		TimeMixin{},
	}
}

// Fields of the User.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("Unknown").
			Annotations(
				entsql.Annotation{Size: 50},
				entgql.OrderField("NAME"),
			),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Annotations(entgql.RelayConnection()).
			Ref("groups"),
	}
}

// Policy defines the privacy policy of the Group.
func (Group) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// Limit DenyMismatchedTenants only for
			// Create operations
			privacy.OnMutationOperation(
				rule.DenyMismatchedTenants(),
				ent.OpCreate,
			),
			// Limit the FilterTenantRule only for
			// UpdateOne and DeleteOne operations.
			privacy.OnMutationOperation(
				rule.FilterTenantRule(),
				ent.OpUpdateOne|ent.OpDeleteOne,
			),
		},
	}
}

// Annotations returns plan category annotations.
func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
