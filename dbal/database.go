package dbal

import (
	"context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
)

func NewService(config Config) *Service {
	return &Service{
		db: sqlx.MustOpen(
			"mysql",
			config.DSN(),
		),
	}
}

type Service struct {
	db *sqlx.DB
}

func (s *Service) SelectAll(_ context.Context, target any, query string, bind any) error {
	if bind == nil {
		bind = map[string]any{}
	}

	statement, err := s.db.PrepareNamed(query)
	if err != nil {
		return err
	}

	return statement.Select(target, bind)
}

func (s *Service) SelectSingle(_ context.Context, target any, query string, bind any) error {
	if bind == nil {
		bind = map[string]any{}
	}

	statement, err := s.db.PrepareNamed(query)
	if err != nil {
		return err
	}

	return statement.Get(target, bind)
}

func (s *Service) TruncateTables(t *testing.T, tables []string) {
	if _, err := s.db.Exec("SET FOREIGN_KEY_CHECKS = 0;"); err != nil {
		t.Fatal(err.Error())
	}

	for _, table := range tables {
		_, err := s.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", table))
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	if _, err := s.db.Exec("SET FOREIGN_KEY_CHECKS = 1;"); err != nil {
		t.Fatal(err.Error())
	}
}
