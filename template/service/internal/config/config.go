package config

type App struct {
	Server   Server   `json:"server" yaml:"server"`
	Database Database `json:"database" yaml:"database"`
}
