package main

import (
	"context"
	"fmt"
	"semaphore/internal/mail"
	"semaphore/internal/settings"
	"strconv"
	"sync"
)

// App struct
type App struct {
	ctx context.Context
	s   *settings.Settings
	mu  sync.Mutex
}

// NewApp creates a new App application struct
func NewApp() *App {
	s, err := settings.LoadSettings()
	if err != nil {
		panic(err)
	}
	return &App{
		s: s,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) FirstPage() string {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.s.ShowWelcome() {
		return "/welcome"
	} else if a.s.HasAccount() {
		return "/inbox"
	} else {
		return "/login"
	}
}

func (a *App) AddAccount(displayName, email, password, server, port string) string {
	a.mu.Lock()
	defer a.mu.Unlock()

	portNum, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Sprintf("failed to convert port to number: %w", err)
	}
	if portNum < 1 || portNum > 65535 {
		return fmt.Sprintf("port number must be between 1 and 65535")
	}

	acc := mail.Account{
		Server:      server,
		Port:        portNum,
		Username:    email,
		Password:    password,
		DisplayName: displayName,
	}

	if err := a.s.AddAccount(&acc); err != nil {
		return err.Error()
	} else {
		return ""
	}
}

func (a *App) GetConversations() []mail.Conversation {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.s.GetConversations()
}
