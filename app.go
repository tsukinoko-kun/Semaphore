package main

import (
	"context"
	"fmt"
	"semaphore/internal/mail"
	"semaphore/internal/settings"
	"strconv"
)

// App struct
type App struct {
	ctx context.Context
	s   *settings.Settings
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

// shutdown is called after the frontend has been destroyed,
// just before the application terminates.
func (a *App) shutdown(ctx context.Context) {
	a.s.Quit()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) FirstPage() string {
	if a.s.ShowWelcome() {
		return "/welcome"
	} else if a.s.HasAccount() {
		return "/inbox"
	} else {
		return "/login"
	}
}

func (a *App) AddAccount(displayName, email, password, server, port string) string {
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
	return a.s.GetConversations()
}
