package config

type App struct {
	Debug    bool     `json:"debug" yaml:"debug" env:"DEBUG" default:"false"`
	Server   Server   `json:"server" yaml:"server"`
	Database Database `json:"database" yaml:"database"`
}

type Server struct {
	Addr string `json:"addr" yaml:"addr" env:"SERVER_ADDR" default:"0.0.0.0:8080"`
}

type Database struct {
	URI string `json:"uri" yaml:"uri" env:"DATABASE_URI" required:"true"`
}
