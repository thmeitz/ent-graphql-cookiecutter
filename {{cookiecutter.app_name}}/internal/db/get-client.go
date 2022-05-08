package db

import (
	"fmt"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/config"
)

func GetClient(conf *config.Config) (client *ent.Client, err error) {

	// DB Stuff
	dsn := (*conf).Database.GetDsn()

	switch (*conf).Database.Type {
		// TODO
	case "sqlite":
		client, err = ent.Open(
			"sqlite3",
			"./elk.db?_fk=1",			
		)
		if err != nil {
			return nil, fmt.Errorf("sqlite error: %v", err)
		}
	case "postgres":
		client, err = ent.Open("postgres", dsn)
		if err != nil {
			return nil, fmt.Errorf("postgres error: %v", err)
		}
	case "mysql":
		client, err = ent.Open("mysql", dsn)
		if err != nil {
			return nil, fmt.Errorf("mysql error: %v", err)
		}
	default:
		return nil, fmt.Errorf("unknown database type: %v", (*conf).Database.Type)
	}
	return client, err
}
