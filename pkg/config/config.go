package config

import (
	"os"
	"path/filepath"

	"github.com/keybase/go-keychain"
	"github.com/spf13/viper"
)

const (
	service     = "bit"
	account     = "access-token"
	accessGroup = "github.com/mbetim/bit"
)

func deleteCurrentToken() error {
	err := keychain.DeleteGenericPasswordItem(service, account)

	if err != nil && err != keychain.ErrorItemNotFound {
		return err
	}

	return nil
}

func StoreToken(token string) error {
	err := deleteCurrentToken()
	if err != nil {
		return err
	}

	item := keychain.NewGenericPassword(service, account, "Bit token", []byte(token), accessGroup)
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)

	return keychain.AddItem(item)
}

func GetToken() (string, error) {
	item, err := keychain.GetGenericPassword(service, account, "Bit token", accessGroup)

	return string(item), err
}

type Config struct {
	DefaultWorkspace string `mapstructure:"default_workspace"`
}

func StartConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configBasePath := filepath.Join(homeDir, ".config", service)
	configFullPath := filepath.Join(configBasePath, "config.yaml")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configBasePath)

	if _, err := os.Stat(configBasePath); os.IsNotExist(err) {
		if err := os.Mkdir(configBasePath, os.ModePerm); err != nil {
			return err
		}
	}

	if _, err := os.Stat(configFullPath); os.IsNotExist(err) {
		if _, err := os.Create(configFullPath); err != nil {
			return err
		}
	}

	return nil
}

func AddConfig(key string, value string) error {
	viper.Set(key, value)

	return viper.WriteConfig()
}

func GetConfig() (Config, error) {
	var config Config

	viper.ReadInConfig()

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
