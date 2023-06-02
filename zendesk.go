package gopher

import "github.com/nukosuke/go-zendesk/zendesk"

func (s *Service) ZD() *zendesk.Client {
	return s.zd
}
