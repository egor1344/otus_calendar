package config

import (
	"github.com/spf13/viper"
)

// ReadConfigFile - чтение конфинг файла
func ReadConfigFile(name, path string) error {
	viper.SetConfigName(name)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

// ReadConfigEnv - чтение конфинга из env
func ReadConfigEnv() error {
	viper.AutomaticEnv()
	return nil
}
