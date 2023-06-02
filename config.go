package gopher

type Config struct {
	ZendeskSubdomain string `env:"GOPHER_ZENDESK_SUBDOMAIN"`
	ZendeskEmail     string `env:"GOPHER_ZENDESK_EMAIL"`
	ZendeskToken     string `env:"GOPHER_ZENDESK_TOKEN"`
	OpenAIKey        string `env:"GOPHER_OPENAI_KEY"`
	DatabaseUser     string `env:"GOPHER_DATABASE_USER"`
	DatabasePass     string `env:"GOPHER_DATABASE_PASS"`
	DatabaseHost     string `env:"GOPHER_DATABASE_HOST"`
	DatabasePort     uint   `env:"GOPHER_DATABASE_PORT"`
	DatabaseName     string `env:"GOPHER_DATABASE_NAME"`
}
