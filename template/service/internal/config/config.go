package config

type App struct {
	Debug    bool     `json:"debug" yaml:"debug" env:"DEBUG" default:"false"`
	Server   Server   `json:"server" yaml:"server"`
	Database Database `json:"database" yaml:"database"`
}
