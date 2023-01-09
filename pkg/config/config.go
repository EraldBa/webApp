package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig is the app configuration
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	LogInfo       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
