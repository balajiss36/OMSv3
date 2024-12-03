package common

import "github.com/spf13/viper"

type Config struct {
	HTTPAddress        string `mapstructure:"HTTP_ADDR"`
	GRPCAddress        string `mapstructure:"GRPC_ADDR"`
	RABBIT_MQ_PORT     string `mapstructure:"RABBIT_MQ_PORT"`
	RABBIT_MQ_HOST     string `mapstructure:"RABBIT_MQ_HOST"`
	RABBIT_MQ_USER     string `mapstructure:"RABBIT_MQ_USER"`
	RABBIT_MQ_PASSWORD string `mapstructure:"RABBIT_MQ_PASSWORD"`
	MONGO_USER         string `mapstructure:"MONGO_USER"`
	MONGO_PASSWORD     string `mapstructure:"MONGO_PASSWORD"`
	MONGO_HOST         string `mapstructure:"MONGO_HOST"`
	MONGO_PORT         string `mapstructure:"MONGO_PORT"`
	MONGO_NAMESPACE    string `mapstructure:"MONGO_NAMESPACE"`
	MONGO_SRV          string `mapstructure:"MONGO_SRV"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
