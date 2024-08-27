package settings

import (
	"semaphore/internal/mail"
	"sync"
)

// GetConversations gets all conversations from all accounts
func (s *Settings) GetConversations() []mail.Conversation {
	var conversations []mail.Conversation
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(s.Account))
	for _, a := range s.Account {
		go func(a *mail.Account, mutex *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			conversations = append(conversations, a.GetConversations()...)
		}(a, &mutex, &wg)
	}
	wg.Wait()
	return conversations
}
