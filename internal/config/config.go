package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
	DB          DataBase      `yaml:"database"`
}
type DataBase struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Database string `yaml:"dbname"`
	Password string `yaml:"password"`
	SSL      string `yaml:"ssl"`
}
type GRPCConfig struct {
	Port    int    `yaml:"port"`
	Timeout string `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path not implemented")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("CONFIG_PATH does not exist:" + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}
	return &cfg
}
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "config path file")
	flag.Parse()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
