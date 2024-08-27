package settings

import (
	"github.com/charmbracelet/log"
	"semaphore/internal/mail"
	"sync"
)

// GetConversations gets all conversations from all accounts
func (s *Settings) GetConversations() []mail.Conversation {
	log.Info("settings.GetConversations")
	var conversations []mail.Conversation
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(s.Account))
	for _, a := range s.Account {
		log.Info("Getting conversations", "account", a.DisplayName)
		go func(a *mail.Account, mutex *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			if c, err := a.GetConversations(); err != nil {
				log.Error("failed to get conversations for account "+a.DisplayName, "err", err)
				return
			} else {
				mutex.Lock()
				defer mutex.Unlock()
				conversations = append(conversations, c...)
			}
		}(a, &mutex, &wg)
	}
	wg.Wait()
	return conversations
}
