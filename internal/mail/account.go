package mail

import (
	"errors"
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"net"
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

func (a *Account) Close() error {
	if a.c == nil {
		return nil
	}
	if a.c.State()&imap.LogoutState == 0 {
		if err := a.c.Logout(); err != nil {
			return errors.Join(fmt.Errorf("failed to logout from %s", a.DisplayName), err)
		}
	}
	if err := a.c.Close(); err != nil {
		if errors.Is(err, net.ErrClosed) || err.Error() == "imap: connection closed" {
			// ignore
			return nil
		}
		return errors.Join(fmt.Errorf("failed to close connection to %s", a.DisplayName), err)
	}
	return nil
}
