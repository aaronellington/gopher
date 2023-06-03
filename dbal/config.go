package dbal

import "fmt"

type Config struct {
	User string
	Pass string
	Host string
	Port uint
	Name string
}

func (config Config) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&allowCleartextPasswords=1",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.Name,
	)
}
