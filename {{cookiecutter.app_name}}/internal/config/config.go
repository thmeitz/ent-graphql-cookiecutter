package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type SrvConfig struct {
	Port        int    `mapstructure:"port" yaml:"port"`
	HostIp      string `mapstructure:"hostip" yaml:"hostip"`
	Environment string `mapstructure:"env" yaml:"env"`
}

func (srv SrvConfig) GetHostWithPort() string {
	return fmt.Sprintf("%s:%d", srv.HostIp, srv.Port)
}

type DatabaseConfig struct {
	User     string `mapstructure:"user" yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	Name     string `mapstructure:"name" yaml:"name"`
	Type     string `mapstructure:"type" yaml:"type"`
	Debug    bool   `mapstructure:"debug" yaml:"debug"`
}

func (c DatabaseConfig) GetDsn() string {
	switch c.Type {
	// TODO: missing Sqlite3
	case "postgres":
		return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", c.Host, c.User, c.Password, c.Name, c.Port)
	case "mysql":
		return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True", c.User, c.Password, c.Host, c.Port, c.Name)
	default:
		return ""
	}
}

type Config struct {
	Server   SrvConfig      `mapstructure:"api" yaml:"api"`
	Database DatabaseConfig `mapstructure:"database" yaml:"database"`
}

func NewDefaultConfig() *Config {
	dbConfig := DatabaseConfig{
		User:     "root",
		Password: "",
		Host:     "db",
		Port:     3306,
		Name:     "mygreatdb",
		Type:     "mysql",
	}

	return &Config{Database: dbConfig}
}

// ReadConfig reads in the config
func ReadConfig(file string) (*Config, error) {
	var config *Config = NewDefaultConfig()
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal config file: %v", err)
	}
	return config, nil
}
