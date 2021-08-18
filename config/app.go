package config

import "github.com/myugen/hodgepodge/internal/gmail"

type AppConfig struct {
	Version string
	Gmail   gmail.GmailConfig
}
