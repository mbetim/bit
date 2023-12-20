package config

import (
	"github.com/keybase/go-keychain"
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
