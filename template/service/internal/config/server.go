package config

type Server struct {
	Debug            bool     `json:"debug" yaml:"debug" env:"SERVER_DEBUG" default:"false"`
	Addr             string   `json:"addr" yaml:"addr" env:"SERVER_ADDR" default:"0.0.0.0:8080"`
	AllowedOrigins   []string `json:"allowedOrigins" yaml:"allowedOrigins" env:"SERVER_ALLOWED_ORIGINS" default:"*"`
	AllowedMethods   []string `json:"allowedMethods" yaml:"allowedMethods" env:"SERVER_ALLOWED_METHODS" default:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   []string `json:"allowedHeaders" yaml:"allowedHeaders" env:"SERVER_ALLOWED_HEADERS" default:"Content-Type,Authorization"`
	AllowCredentials bool     `json:"allowCredentials" yaml:"allowCredentials" env:"SERVER_ALLOW_CREDENTIALS" default:"true"`
	MaxHeaderBytes   int      `json:"maxHeaderBytes" yaml:"maxHeaderBytes" env:"SERVER_MAX_HEADER_BYTES" default:"1"`
	ReadTimeout      int      `json:"readTimeout" yaml:"readTimeout" env:"SERVER_READ_TIMEOUT" default:"5"`
	WriteTimeout     int      `json:"writeTimeout" yaml:"writeTimeout" env:"SERVER_WRITE_TIMEOUT" default:"10"`
	IdleTimeout      int      `json:"idleTimeout" yaml:"idleTimeout" env:"SERVER_IDLE_TIMEOUT" default:"60"`
	ShutdownTimeout  int      `json:"shutdownTimeout" yaml:"shutdownTimeout" env:"SERVER_SHUTDOWN_TIMEOUT" default:"10"`
}
