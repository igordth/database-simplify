package pggorm

import (
	"fmt"
	"strings"
	"time"
)

type Config struct {
	Name            string
	User            string `default:"postgres"`
	Password        string `default:""`
	Host            string `default:"localhost"`
	Schema          string
	MaxOpenConn     int               `yaml:"max_open_conn" default:"0"`
	MaxIdleConn     int               `yaml:"max_idle_conn" default:"2"`
	ConnMaxLifetime time.Duration     `yaml:"conn_max_lifetime" default:"1m"`
	Options         map[string]string `yaml:"options"`
}

func (c Config) String() (string, error) {
	if c.Host == "" {
		return "", ErrSetupHost
	}
	if c.Name == "" {
		return "", ErrSetupName
	}
	str := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable %s %s",
		c.Host,
		c.User,
		c.Password,
		c.Name,
		c.getSchema(),
		c.getOptions(),
	)
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, "  ", " ")
	return str, nil
}

func (c Config) getSchema() string {
	if c.Schema == "" {
		return ""
	}
	return fmt.Sprintf("search_path=%s", c.Schema)
}

func (c Config) getOptions() (s string) {
	var tmp []string
	for k, v := range c.Options {
		tmp = append(tmp, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(tmp, " ")
}
