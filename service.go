package gopher

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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
		db: sqlx.MustOpen("mysql",
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?parseTime=true&allowCleartextPasswords=1",
				config.DatabaseUser,
				config.DatabasePass,
				config.DatabaseHost,
				config.DatabasePort,
				config.DatabaseName,
			),
		),
	}
}

type Service struct {
	db *sqlx.DB
	zd *zendesk.Client
	ai *openai.Client
}
