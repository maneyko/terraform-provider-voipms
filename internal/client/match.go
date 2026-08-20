package client

import (
	"fmt"
	"strings"
)

func matchUnique[T any](items []T, query, kind string, keys func(*T) []string) (*T, error) {
	q := strings.TrimSpace(query)
	if q == "" {
		return nil, fmt.Errorf("%w: empty %s query", ErrNotFound, kind)
	}
	var matches []*T
	for i := range items {
		item := &items[i]
		for _, key := range keys(item) {
			if key != "" && strings.EqualFold(key, q) {
				matches = append(matches, item)
				break
			}
		}
	}
	if len(matches) == 0 {
		return nil, fmt.Errorf("%w: %s %q", ErrNotFound, kind, query)
	}
	if len(matches) > 1 {
		return nil, fmt.Errorf("%s %q matches more than one object", kind, query)
	}
	return matches[0], nil
}

// MatchVoicemail finds a mailbox by number or display name.
func MatchVoicemail(items []Voicemail, query string) (*Voicemail, error) {
	return matchUnique(items, query, "voicemail", func(v *Voicemail) []string {
		return []string{v.Mailbox.String(), v.Name.String()}
	})
}

// MatchForwarding finds a forwarding by id or description.
func MatchForwarding(items []Forwarding, query string) (*Forwarding, error) {
	return matchUnique(items, query, "forwarding", func(f *Forwarding) []string {
		return []string{f.Forwarding.String(), f.Description.String()}
	})
}

// MatchCallback finds a callback by id or description.
func MatchCallback(items []Callback, query string) (*Callback, error) {
	return matchUnique(items, query, "callback", func(c *Callback) []string {
		return []string{c.Callback.String(), c.Description.String()}
	})
}

// MatchPhonebookGroup finds a group by id or name.
func MatchPhonebookGroup(items []PhonebookGroup, query string) (*PhonebookGroup, error) {
	return matchUnique(items, query, "phonebook group", func(g *PhonebookGroup) []string {
		return []string{g.PhonebookGroup.String(), g.Name.String()}
	})
}

// MatchPhonebookEntry finds an entry by id or contact name.
func MatchPhonebookEntry(items []PhonebookEntry, query string) (*PhonebookEntry, error) {
	return matchUnique(items, query, "phonebook entry", func(p *PhonebookEntry) []string {
		return []string{p.Phonebook.String(), p.Name.String()}
	})
}

// NameForMailbox returns the display name for a mailbox number.
func NameForMailbox(items []Voicemail, mailbox string) string {
	for i := range items {
		if items[i].Mailbox.String() == mailbox {
			return items[i].Name.String()
		}
	}
	return ""
}
