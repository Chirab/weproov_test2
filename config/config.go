package config

import "github.com/spf13/viper"

type Config struct {
    PORT      string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path)
    viper.SetConfigFile(".env")
    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)
    return
}