package mail

import (
	"errors"
	"github.com/charmbracelet/log"
	"github.com/emersion/go-imap"
	"github.com/k3a/html2text"
	"semaphore/internal/util"
	"strings"
)

type (
	Conversation struct {
		Account            string   `json:"account"`
		Subject            string   `json:"subject"`
		Participant        []string `json:"participants"`
		LastMessagePreview string   `json:"lastMessage"`
	}
)

const maxFetchCount = 16

// GetConversations gets all conversations from the account
func (a *Account) GetConversations() ([]Conversation, error) {
	log.Info("mail.GetConversations", "account", a.DisplayName)
	c, err := a.Client()
	if err != nil {
		return nil, errors.Join(errors.New("failed to get client"), err)
	}
	log.Info("Got client")
	mbox, err := c.Select("INBOX", false)
	if mbox == nil {
		return nil, errors.Join(errors.New("failed to select INBOX"), err)
	}
	log.Info("Selected INBOX")

	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > maxFetchCount {
		from = mbox.Messages - maxFetchCount
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)
	log.Info("Fetching messages", "from", from, "to", to)

	section := &imap.BodySectionName{
		BodyPartName: imap.BodyPartName{Specifier: imap.TextSpecifier},
		Peek:         true,
	}
	items := []imap.FetchItem{section.FetchItem(), imap.FetchEnvelope}

	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, items, messages)
	}()

	conversations := make([]Conversation, 0, to-from+1)
	for msg := range messages {
		participants := util.NewSet[string]()
		for _, a := range msg.Envelope.From {
			participants.Add(a.Address())
		}
		for _, a := range msg.Envelope.To {
			participants.Add(a.Address())
		}
		for _, a := range msg.Envelope.Cc {
			participants.Add(a.Address())
		}
		for _, a := range msg.Envelope.Bcc {
			participants.Add(a.Address())
		}
		for _, a := range msg.Envelope.ReplyTo {
			participants.Add(a.Address())
		}
		for _, a := range msg.Envelope.Sender {
			participants.Add(a.Address())
		}

		sb := strings.Builder{}
		for section, l := range msg.Body {
			b := make([]byte, l.Len())
			if _, err := l.Read(b); err != nil {
				log.Error("failed to read message body", "section", section, "err", err)
			} else {
				s := string(b)
				if len(s) > 100 {
					s = s[:100]
				}
				log.Info("message body", "section", section, "l", s)
				sb.Write(b)
				sb.WriteString("\n\n")
			}
		}
		body := sb.String()
		body = strings.TrimSpace(html2text.HTML2Text(body))

		if len(body) > 200 {
			body = body[:200]
		}

		conversations = append(conversations, Conversation{
			Account:            a.DisplayName,
			Subject:            msg.Envelope.Subject,
			Participant:        participants.Values(),
			LastMessagePreview: body,
		})
	}

	if err := <-done; err != nil {
		return nil, errors.Join(errors.New("failed to fetch messages"), err)
	}

	log.Info("mail.GetConversations done")

	return conversations, nil
}
