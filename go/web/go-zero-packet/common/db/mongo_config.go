package db

import (
	"fmt"
)

type MongoConfig struct {
	User           string
	Passwd         string
	Addr           string
	DBName         string
	ReadPreference string `yaml:"ReadPreference,omitempty" json:"read_preference,omitempty"`
}

func (c MongoConfig) GetURL() string {

	// mongodb://user:passwd@127.0.0.1:27017/platform_test?authSource=admin&readPreference=secondaryPreferred
	return fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=admin&readPreference=%s",
		c.User,
		c.Passwd,
		c.Addr,
		c.DBName,
		c.ReadPreference,
	)
}

func (c MongoConfig) GetDBName() string {
	return c.DBName
}
