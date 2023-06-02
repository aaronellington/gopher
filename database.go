package gopher

import "github.com/jmoiron/sqlx"

func (s *Service) DB() *sqlx.DB {
	return s.db
}
