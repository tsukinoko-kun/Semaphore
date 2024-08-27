package mail

type (
	Conversation struct {
		Account            string   `json:"account"`
		Subject            string   `json:"subject"`
		Participant        []string `json:"participants"`
		LastMessagePreview string   `json:"lastMessage"`
	}
)

// GetConversations gets all conversations from the account
func (a *Account) GetConversations() []Conversation {
	return []Conversation{
		{
			Account:            a.DisplayName,
			Subject:            "This new email client is awesome!",
			Participant:        []string{"john@example.com", "holger@example.de"},
			LastMessagePreview: `<span class="sender">john@example.com</span><span class="content">I tried it out and it's really cool! Can't wait to see your thoughts on it.</span>`,
		},
		{
			Account:            a.DisplayName,
			Subject:            "Meeting next week",
			Participant:        []string{"john@example.com"},
			LastMessagePreview: `<span class="sender">john@example.com</span><span class="content">Hey, I'm planning a meeting next week. Are you available? I need your input on the new project.</span>`,
		},
	}
}
