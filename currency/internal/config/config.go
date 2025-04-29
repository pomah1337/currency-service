package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type GrpcCfg struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type DbCfg struct {
	Url string `mapstructure:"url"`
}

type ExternalApiCfg struct {
	Url string `mapstructure:"url"`
}

type WorkerCfg struct {
	Cron       string `mapstructure:"cron"`
	Currencies struct {
		BaseCurrency   string `mapstructure:"base_currency"`
		TargetCurrency string `mapstructure:"target_currency"`
	} `mapstructure:"currencies"`
}
type Config struct {
	Grpc        GrpcCfg        `mapstructure:"grpc"`
	DB          DbCfg          `mapstructure:"db"`
	ExternalAPI ExternalApiCfg `mapstructure:"external_api"`
	Worker      WorkerCfg      `mapstructure:"worker"`
}

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}
