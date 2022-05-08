//go:build tools
// +build tools

package generate

//go:generate go run -mod=mod github.com/google/addlicense -c "{{cookiecutter.full_name}}" -y 2022-present ./

import (
	// remove it, if you dont want to reverse engineere a legacy database
	_ "ariga.io/entimport/cmd/entimport"
	_ "entgo.io/ent/cmd/ent"
	_ "github.com/99designs/gqlgen"
	_ "github.com/google/addlicense"
)
