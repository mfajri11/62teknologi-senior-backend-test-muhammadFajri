package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

var cfg Config

type postgres struct {
	Host         string        `yaml:"host" env-default:"localhost"`
	Port         int           `yaml:"port" env-default:"5432" env-layout:"int"`
	Username     string        `yaml:"username"`
	Secret       string        `yaml:"secret" env:"PG_SECRET"`
	DatabaseName string        `yaml:"database-name"`
	MaxPoolSize  int32         `yaml:"max-pool-size"`
	MinPoolSize  int32         `yaml:"min-pool-size"`
	MaxLifeTime  time.Duration `yaml:"max-lifetime"`
	MaxIdleTime  time.Duration `yaml:"max-idletime"`
}

func (p postgres) URL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", p.Username, p.Secret, p.Host, p.Port, p.DatabaseName)
}

type app struct {
	Host               string `yaml:"host" env-default:"localhost"`
	Port               int    `yaml:"port" env-default:"8080" env-layout:"int"`
	SecretAPIKey       string `yaml:"secret-api-key" env:"SECRET_KEY_API_KEY"`
	MaxLimitPagination int    `yaml:"max-limit-pagination" env-default:"10" env-layout:"int"`
}

func (a app) Address() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

type Config struct {
	App      app      `yaml:"app"`
	Postgres postgres `yaml:"postgres"`
}

func Cfg() Config {
	return cfg
}

func init() {
	env := os.Getenv("BS_ENV")
	if env == "" {
		env = "development"
	}

	configFile := "config." + env + ".yaml"
	cfg = Config{}
	_, fname, _, _ := runtime.Caller(0)
	dir := filepath.Join(fname, "..")
	configPath := filepath.Join(dir, configFile)
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		panic(err)
	}
}
