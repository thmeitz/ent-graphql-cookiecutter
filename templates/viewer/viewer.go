package viewer

import (
	"context"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent"
)

// Role for viewer actions.
type Role int

// List of roles.
const (
	_ Role = 1 << iota
	Admin
	View
)

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	Admin() bool    // If viewer is admin.
	Tenant() string // Tenant name.
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	T    *ent.Tenant
	Role Role // Attached roles.
}

func (v UserViewer) Admin() bool {
	return v.Role&Admin != 0
}

func (v UserViewer) Tenant() string {
	if v.T != nil {
		return v.T.Name
	}
	return ""
}

type ctxKey struct{}

// FromContext returns the Viewer stored in a context.
func FromContext(ctx context.Context) Viewer {
	v, _ := ctx.Value(ctxKey{}).(Viewer)
	return v
}

// NewContext returns a copy of parent context with the given Viewer attached with it.
func NewContext(parent context.Context, v Viewer) context.Context {
	return context.WithValue(parent, ctxKey{}, v)
}
