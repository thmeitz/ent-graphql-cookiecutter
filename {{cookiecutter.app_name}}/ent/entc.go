//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithConfigPath("../gqlgen.yml"),
		// Generate the filters to a separate schema
		// file and load it in the gqlgen.yml config.
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("../graph/schema/ent.graphql"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(ex),
		entc.TemplateDir("./template"),
		entc.FeatureNames("sql/upsert"),
		// doesnt work on first 'go generate ./...' with cookiecutter
		// entc.FeatureNames("sql/schemaconfig"),
		entc.FeatureNames("sql/lock"),
		entc.FeatureNames("sql/modifier"),
		entc.FeatureNames("privacy"),
		// doesnt work on first 'go generate ./...' with cookiecutter
		// entc.FeatureNames("schema/snapshot"),
		entc.FeatureNames("entql"),
	}
	if err := entc.Generate("./schema", &gen.Config{
		Templates: entgql.AllTemplates,
		Header: `
			// Copyright 2022-present {{cookiecutter.full_name}}
			//
			// Licensed under the Apache License, Version 2.0 (the "License");
			// you may not use this file except in compliance with the License.
			// You may obtain a copy of the License at
			//
			//      http://www.apache.org/licenses/LICENSE-2.0
			//
			// Unless required by applicable law or agreed to in writing, software
			// distributed under the License is distributed on an "AS IS" BASIS,
			// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
			// See the License for the specific language governing permissions and
			// limitations under the License.
			//
			// Code generated by entc, DO NOT EDIT.
		`,
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
