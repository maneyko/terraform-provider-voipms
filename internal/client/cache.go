package client

import "sync"

// listCache holds unfiltered account lists for one Terraform run so DID
// Read/ModifyPlan do not hit VoIP.ms once per object (API rate limit).
type listCache struct {
	mu           sync.Mutex
	servers      []Server
	forwardings  []Forwarding
	voicemails   []Voicemail
	subaccounts  []SubAccount
	haveServers  bool
	haveFwds     bool
	haveVMs      bool
	haveAccounts bool
}

func (c *Client) cachedServers(load func() ([]Server, error)) ([]Server, error) {
	c.cache.mu.Lock()
	defer c.cache.mu.Unlock()
	if c.cache.haveServers {
		return c.cache.servers, nil
	}
	items, err := load()
	if err != nil {
		return nil, err
	}
	c.cache.servers = items
	c.cache.haveServers = true
	return items, nil
}

func (c *Client) cachedForwardings(load func() ([]Forwarding, error)) ([]Forwarding, error) {
	c.cache.mu.Lock()
	defer c.cache.mu.Unlock()
	if c.cache.haveFwds {
		return c.cache.forwardings, nil
	}
	items, err := load()
	if err != nil {
		return nil, err
	}
	c.cache.forwardings = items
	c.cache.haveFwds = true
	return items, nil
}

func (c *Client) cachedVoicemails(load func() ([]Voicemail, error)) ([]Voicemail, error) {
	c.cache.mu.Lock()
	defer c.cache.mu.Unlock()
	if c.cache.haveVMs {
		return c.cache.voicemails, nil
	}
	items, err := load()
	if err != nil {
		return nil, err
	}
	c.cache.voicemails = items
	c.cache.haveVMs = true
	return items, nil
}

func (c *Client) cachedSubAccounts(load func() ([]SubAccount, error)) ([]SubAccount, error) {
	c.cache.mu.Lock()
	defer c.cache.mu.Unlock()
	if c.cache.haveAccounts {
		return c.cache.subaccounts, nil
	}
	items, err := load()
	if err != nil {
		return nil, err
	}
	c.cache.subaccounts = items
	c.cache.haveAccounts = true
	return items, nil
}
