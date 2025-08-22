package config

type Server struct {
	Debug            bool     `json:"debug" yaml:"debug" env:"SERVER_DEBUG" default:"false"`
	Addr             string   `json:"addr" yaml:"addr" env:"SERVER_ADDR" default:"0.0.0.0:8080"`
	AllowedOrigins   []string `json:"allowedOrigins" yaml:"allowedOrigins" env:"SERVER_ALLOWED_ORIGINS" default:"*"`
	AllowedMethods   []string `json:"allowedMethods" yaml:"allowedMethods" env:"SERVER_ALLOWED_METHODS" default:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   []string `json:"allowedHeaders" yaml:"allowedHeaders" env:"SERVER_ALLOWED_HEADERS" default:"Content-Type,Authorization"`
	AllowCredentials bool     `json:"allowCredentials" yaml:"allowCredentials" env:"SERVER_ALLOW_CREDENTIALS" default:"true"`
	MaxAge           string   `json:"maxAge" yaml:"maxAge" env:"SERVER_MAX_AGE" default:"1h"`
}
