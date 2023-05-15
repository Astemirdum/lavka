package config

import (
	"log"
	"sync"
	"time"

	"github.com/Astemirdum/lavka/pkg/logger"

	"github.com/kelseyhightower/envconfig"
)

type HTTPServer struct {
	Host         string        `yaml:"host" envconfig:"HTTP_HOST"`
	Port         string        `yaml:"port" envconfig:"HTTP_PORT"`
	ReadTimeout  time.Duration `yaml:"readTimeout" envconfig:"HTTP_READ"`
	WriteTimeout time.Duration
}

type DB struct {
	Host     string `yaml:"host" envconfig:"DB_HOST"`
	Port     int    `yaml:"port" envconfig:"DB_PORT"`
	Username string `yaml:"user" envconfig:"DB_USER"`
	Password string `yaml:"password" envconfig:"DB_PASSWORD"`
	NameDB   string `yaml:"dbname" envconfig:"DB_NAME"`
}

type Config struct {
	Server   HTTPServer `yaml:"server"`
	Database DB         `yaml:"db"`
	Log      logger.Log `yaml:"log"`
}

var (
	once sync.Once
	cfg  *Config
)

// NewConfig reads config from environment.
func NewConfig(ops ...Option) *Config {
	once.Do(func() {
		var config Config
		for _, op := range ops {
			op(&config)
		}
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal("NewConfig ", err)
		}
		cfg = &config
		// printConfig(cfg)
	})

	return cfg
}

// func printConfig(cfg *Config) {
//	jscfg, _ := json.MarshalIndent(cfg, "", "	") //nolint:errcheck
//	fmt.Println(string(jscfg))
//}
