package mail

import (
	"errors"
	"fmt"
	"github.com/emersion/go-imap/client"
)

type Account struct {
	Server      string         `toml:"server"`
	Port        int            `toml:"port"`
	Username    string         `toml:"username"`
	Password    string         `toml:"-"`
	DisplayName string         `toml:"display_name"`
	c           *client.Client `toml:"-"`
}

func (a *Account) Client() (*client.Client, error) {
	if a.c != nil {
		return a.c, nil
	}

	var err error
	a.c, err = Dial(a.Address())
	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to dial %s", a.DisplayName), err)
	}

	if err := a.c.Login(a.Username, a.Password); err != nil {
		return nil, errors.Join(fmt.Errorf("failed to login to %s", a.DisplayName), err)
	}

	return a.c, nil
}

func (a *Account) Address() string {
	return fmt.Sprintf("%s:%d", a.Server, a.Port)
}
