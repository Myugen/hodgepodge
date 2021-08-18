package config

import "github.com/myugen/utils/internal/gmail"

type AppConfig struct {
	Version string
	Gmail   gmail.GmailConfig
}
