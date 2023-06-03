package gopher

import (
	"github.com/aaronellington/gopher/dbal"
	"github.com/go-sql-driver/mysql"
	"github.com/kyberbits/forge/forge"
	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/sashabaranov/go-openai"
)

func NewServiceWithConfig() *Service {
	config := Config{}
	environment := forge.NewEnvironment()

	if err := environment.Decode(&config); err != nil {
		panic(err)
	}

	return NewService(config)
}

func NewService(config Config) *Service {
	zdClient, _ := zendesk.NewClient(nil)
	_ = zdClient.SetSubdomain(config.ZendeskSubdomain)
	zdClient.SetCredential(zendesk.NewAPITokenCredential(config.ZendeskEmail, config.ZendeskToken))

	_ = mysql.MySQLError{}

	return &Service{
		zd: zdClient,
		ai: openai.NewClient(config.OpenAIKey),
		dbal: dbal.NewService(dbal.Config{
			User: config.DatabaseUser,
			Pass: config.DatabasePass,
			Host: config.DatabaseHost,
			Port: config.DatabasePort,
			Name: config.DatabaseName,
		}),
	}
}

type Service struct {
	dbal *dbal.Service
	zd   *zendesk.Client
	ai   *openai.Client
}
