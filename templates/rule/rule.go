package rule

import (
	"context"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent/predicate"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent/privacy"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent/tenant"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent/user"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/viewer"
)

// DenyIfNoViewer is a rule that returns deny decision if the viewer is missing in the context.
func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view == nil {
			return privacy.Denyf("viewer-context is missing")
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// AllowIfAdmin is a rule that returns allow decision if the viewer is admin.
func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view.Admin() {
			return privacy.Allow
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// FilterTenantRule is a query/mutation rule that filters out entities that are not in the tenant.
func FilterTenantRule() privacy.QueryMutationRule {
	// TenantsFilter is an interface to wrap WhereHasTenantWith()
	// predicate that is used by both `Group` and `User` schemas.
	type TenantsFilter interface {
		WhereHasTenantWith(...predicate.Tenant)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		view := viewer.FromContext(ctx)
		if view.Tenant() == "" {
			return privacy.Denyf("missing tenant information in viewer")
		}
		tf, ok := f.(TenantsFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}
		// Make sure that a tenant reads only entities that has an edge to it.
		tf.WhereHasTenantWith(tenant.Name(view.Tenant()))
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// DenyMismatchedTenants is a rule that runs only on create operations and returns a deny
// decision if the operation tries to add users to groups that are not in the same tenant.
func DenyMismatchedTenants() privacy.MutationRule {
	return privacy.GroupMutationRuleFunc(func(ctx context.Context, m *ent.GroupMutation) error {
		tid, exists := m.TenantID()
		if !exists {
			return privacy.Denyf("missing tenant information in mutation")
		}
		users := m.UsersIDs()
		// If there are no users in the mutation, skip this rule-check.
		if len(users) == 0 {
			return privacy.Skip
		}
		// Query the tenant-id of all users. Expect to have exact 1 result,
		// and it matches the tenant-id of the group above.
		id, err := m.Client().User.Query().Where(user.IDIn(users...)).QueryTenant().OnlyID(ctx)
		if err != nil {
			return privacy.Denyf("querying the tenant-id %v", err)
		}
		if id != tid {
			return privacy.Denyf("mismatch tenant-ids for group/users %d != %d", tid, id)
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}
