package domain

import "fmt"

type Message struct {
	Author  string
	Message string
}

func (m *Message) Validate() error {
	if m.Author == "" {
		return fmt.Errorf("missing author")
	}

	if m.Message == "" {
		return fmt.Errorf("missing message")
	}
	return nil
}
