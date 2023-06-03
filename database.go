package gopher

import "github.com/aaronellington/gopher/dbal"

func (s *Service) DBAL() *dbal.Service {
	return s.dbal
}
