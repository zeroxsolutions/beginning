package config

type Database struct {
	Debug bool         `json:"debug" yaml:"debug" env:"DATABASE_DEBUG" default:"false"`
	URI   string       `json:"uri" yaml:"uri" env:"DATABASE_URI" required:"true"`
	Pool  DatabasePool `json:"pool" yaml:"pool"`
}

type DatabasePool struct {
	Enabled         bool `json:"enabled" yaml:"enabled" env:"DATABASE_POOL_ENABLED" default:"false"`
	MaxIdleConns    int  `json:"maxIdleConns" yaml:"maxIdleConns" env:"DATABASE_POOL_MAX_IDLE_CONNS" default:"10"`
	MaxOpenConns    int  `json:"maxOpenConns" yaml:"maxOpenConns" env:"DATABASE_POOL_MAX_OPEN_CONNS" default:"100"`
	ConnMaxLifetime int  `json:"connMaxLifetime" yaml:"connMaxLifetime" env:"DATABASE_POOL_CONN_MAX_LIFETIME" default:"0"`
}