package settings

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/emersion/go-imap"
	"github.com/zalando/go-keyring"
	"net"
	"os"
	"path/filepath"
	"semaphore/internal/mail"
)

const (
	KeyringService = "github.com/tsukinoko-kun/semaphore"
)

type Settings struct {
	Account     []*mail.Account `toml:"account"`
	showWelcome bool            `toml:"-"`
}

func LoadSettings() (*Settings, error) {
	loc, err := location()
	if err != nil {
		return nil, errors.Join(errors.New("failed to get location of settings file"), err)
	}

	settings := &Settings{}

	if _, err := os.Stat(loc); os.IsNotExist(err) {
		settings.showWelcome = true
		return settings, nil
	}

	b, err := os.ReadFile(loc)
	if err != nil {
		return nil, errors.Join(errors.New("failed to read settings file"), err)
	}

	if err := toml.Unmarshal(b, settings); err != nil {
		return nil, errors.Join(errors.New("failed to unmarshal settings"), err)
	}

	for _, a := range settings.Account {
		if pwd, err := keyring.Get(KeyringService, a.Username); err != nil {
			return nil, errors.Join(fmt.Errorf("failed to get password for %s from keyring", a.Username), err)
		} else if pwd == "" {
			return nil, fmt.Errorf("no password found for %s in keyring", a.Username)
		} else {
			a.Password = pwd
		}
		if _, err := a.Client(); err != nil {
			return nil, errors.Join(fmt.Errorf("failed to receive imap client for %s", a.DisplayName), err)
		}
	}

	return settings, nil
}

func (s *Settings) ShowWelcome() bool {
	return s.showWelcome
}

func (s *Settings) AddAccount(a *mail.Account) error {
	if _, err := a.Client(); err != nil {
		return errors.Join(fmt.Errorf("failed to receive imap client for %s", a.DisplayName), err)
	}
	s.Account = append(s.Account, a)
	if err := s.Save(); err != nil {
		return errors.Join(errors.New("failed to save settings"), err)
	}
	if err := keyring.Set(KeyringService, a.Username, a.Password); err != nil {
		return errors.Join(fmt.Errorf("failed to save password for %s to keyring", a.Username), err)
	}
	return nil
}

func (s *Settings) Save() error {
	b, err := toml.Marshal(s)
	if err != nil {
		return errors.Join(errors.New("failed to marshal settings"), err)
	}
	loc, err := location()
	if err != nil {
		return errors.Join(errors.New("failed to get location of settings file"), err)
	}
	if err := os.WriteFile(loc, b, 0600); err != nil {
		return errors.Join(errors.New("failed to write settings file"), err)
	}
	return nil
}

func location() (string, error) {
	confDir, err := os.UserConfigDir()
	if err != nil {
		return "", errors.Join(errors.New("failed to determine user config directory"), err)
	}
	dir := filepath.Join(confDir, "semaphore")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", errors.Join(errors.New("failed to create config directory for semaphore"), err)
	}
	return filepath.Join(dir, "config.toml"), nil
}

func (s *Settings) HasAccount() bool {
	return len(s.Account) > 0
}

func (s *Settings) Quit() error {
	if !s.HasAccount() {
		return nil
	}
	if err := s.Save(); err != nil {
		return err
	}
	for _, a := range s.Account {
		c, _ := a.Client()
		if c.State()&imap.LogoutState == 0 {
			if err := c.Logout(); err != nil {
				return errors.Join(fmt.Errorf("failed to logout from %s", a.DisplayName), err)
			}
		}
		if err := c.Close(); err != nil {
			if errors.Is(err, net.ErrClosed) || err.Error() == "imap: connection closed" {
				// ignore
				continue
			}
			return errors.Join(fmt.Errorf("failed to close connection to %s", a.DisplayName), err)
		}
	}
	return nil
}
