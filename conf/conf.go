package conf

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/upgradeskill/beta-team/pkg/envar"
)

type Group struct {
	Server   Server   `json:"server,omitempty"`
	Database Database `json:"database,omitempty"`
	Redis    Redis    `json:"redis,omitempty"`
}

type Server struct {
	ENV string `json:"env"`
}

type Database struct {
	Engine   string `json:"engine,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"-"`
	Schema   string `json:"schema,omitempty"`
	MaxIdle  int    `json:"max_idle,omitempty"`
	MaxConn  int    `json:"max_conn,omitempty"`
}

type Redis struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"-"`
	DB       int    `json:"db,omitempty"`
	UseTLS   bool   `json:"use_tls,omitempty"`
}

func LoadConfig() *Group {
	if err := godotenv.Load(); err != nil {
		// in prod we will not use this,use os env instead
		log.Print(".env notfound")
	}

	env = envar.GetEnv("ENV", "dev")

	return &Group{
		Server: Server{
			ENV: env,
		},
		Database: Database{
			Engine:   envar.GetEnv("DATABASE_ENGINE", "mysqli"),
			Host:     envar.GetEnv("DATABASE_HOST", "localhost"),
			Port:     envar.GetEnv("DATABASE_PORT", 3306),
			Username: envar.GetEnv("DATABASE_USERNAME", "root"),
			Password: envar.GetEnv("DATABASE_PASSWORD", "root"),
			Schema:   envar.GetEnv("DATABASE_SCHEMA", "inventory"),
			MaxIdle:  envar.GetEnv("DATABASE_MAX_IDLE", 20),
			MaxConn:  envar.GetEnv("DATABASE_MAX_CONN", 100),
		},
		Redis: Redis{
			Host:     envar.GetEnv("REDIS_HOST", "localhost"),
			Port:     envar.GetEnv("REDIS_PORT", 31113),
			Username: envar.GetEnv("REDIS_USERNAME", ""),
			Password: envar.GetEnv("REDIS_PASSWORD", ""),
			DB:       envar.GetEnv("REDIS_DB", 4),
			UseTLS:   envar.GetEnv("REDIS_USE_TLS", false),
		},
	}
}
