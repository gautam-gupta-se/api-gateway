package model

type Route struct {
	ID       string `yaml:"id"`
	Path     string `yaml:"path"`
	Upstream string `yaml:"upstream"`
	// Placeholders
	Middlewares []string `yaml:"middlewares,omitempty"`
}
