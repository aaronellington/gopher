package gopher

import (
	"context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
)

func (s *Service) DBAL() *DatabaseAbstractionLayer {
	return s.dbal
}

type DatabaseEntity interface {
	TableName() string
}

type DBALConfig struct {
	User string
	Pass string
	Host string
	Port uint
	Name string
}

func (config DBALConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&allowCleartextPasswords=1",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.Name,
	)
}

func NewDBAL(config DBALConfig) *DatabaseAbstractionLayer {
	return &DatabaseAbstractionLayer{
		db: sqlx.MustOpen(
			"mysql",
			config.DSN(),
		),
	}
}

type DatabaseAbstractionLayer struct {
	db *sqlx.DB
}

func (dbal *DatabaseAbstractionLayer) SelectAll(ctx context.Context, target any, query string, bind any) error {
	if bind == nil {
		bind = map[string]any{}
	}

	statement, err := dbal.db.PrepareNamed(query)
	if err != nil {
		return err
	}

	if err := statement.Select(target, bind); err != nil {
		return err
	}

	return nil
}

func (dbal *DatabaseAbstractionLayer) SelectSingle(ctx context.Context, target any, query string, bind any) error {
	if bind == nil {
		bind = map[string]any{}
	}

	statement, err := dbal.db.PrepareNamed(query)
	if err != nil {
		return err
	}

	if err := statement.Get(target, bind); err != nil {
		return err
	}

	return nil
}

func (dbal *DatabaseAbstractionLayer) TruncateTables(t *testing.T, tables []string) {
	if _, err := dbal.db.Exec("SET FOREIGN_KEY_CHECKS = 0;"); err != nil {
		t.Fatal(err.Error())
	}

	for _, table := range tables {
		_, err := dbal.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", table))
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	if _, err := dbal.db.Exec("SET FOREIGN_KEY_CHECKS = 1;"); err != nil {
		t.Fatal(err.Error())
	}
}
