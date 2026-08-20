package client

import (
	"context"
	"fmt"
	"strings"
)

type serversResponse struct {
	Status  string   `json:"status"`
	Servers []Server `json:"servers"`
}

// GetServersInfo lists VoIP.ms POPs. pop, if set, filters to one server_pop.
func (c *Client) GetServersInfo(ctx context.Context, pop string) ([]Server, error) {
	params := map[string]string{}
	if pop != "" {
		params["server_pop"] = pop
	}
	var resp serversResponse
	err := c.Call(ctx, "getServersInfo", params, &resp)
	if err != nil {
		if emptyResult(err) {
			return []Server{}, nil
		}
		return nil, err
	}
	return resp.Servers, nil
}

// GetServer returns one POP by server_pop id.
func (c *Client) GetServer(ctx context.Context, pop string) (*Server, error) {
	items, err := c.GetServersInfo(ctx, pop)
	if err != nil {
		return nil, err
	}
	for i := range items {
		if items[i].POP.String() == pop {
			return &items[i], nil
		}
	}
	if len(items) == 1 && pop != "" {
		return &items[0], nil
	}
	return nil, fmt.Errorf("%w: server pop %s", ErrNotFound, pop)
}

// FindServer returns one POP by id, hostname, display name, or "Name, Country".
func (c *Client) FindServer(ctx context.Context, query string) (*Server, error) {
	items, err := c.GetServersInfo(ctx, "")
	if err != nil {
		return nil, err
	}
	return MatchServer(items, query)
}

// MatchServer finds a POP in a catalog list. query may be a numeric pop id,
// SIP hostname (newyork7.voip.ms), display name (New York 7), or "Name, Country".
func MatchServer(servers []Server, query string) (*Server, error) {
	q := strings.TrimSpace(query)
	if q == "" {
		return nil, fmt.Errorf("%w: empty POP query", ErrNotFound)
	}
	var matches []Server
	for i := range servers {
		if serverMatchesQuery(servers[i], q) {
			matches = append(matches, servers[i])
		}
	}
	if len(matches) == 0 {
		return nil, fmt.Errorf("%w: POP %q", ErrNotFound, query)
	}
	if len(matches) > 1 {
		ids := make([]string, 0, len(matches))
		for _, m := range matches {
			ids = append(ids, m.POP.String())
		}
		return nil, fmt.Errorf("POP %q matches more than one server (%s)", query, strings.Join(ids, ", "))
	}
	return &matches[0], nil
}

func serverMatchesQuery(s Server, query string) bool {
	if strings.EqualFold(s.POP.String(), query) {
		return true
	}
	if strings.EqualFold(s.Hostname.String(), query) {
		return true
	}
	if strings.EqualFold(s.Name.String(), query) {
		return true
	}
	if s.Shortname.String() != "" && strings.EqualFold(s.Shortname.String(), query) {
		return true
	}
	if s.Country.String() != "" {
		labeled := s.Name.String() + ", " + s.Country.String()
		if strings.EqualFold(labeled, query) {
			return true
		}
	}
	return false
}

// HostnameForPOP returns the SIP hostname for a numeric pop id.
func HostnameForPOP(servers []Server, pop int64) string {
	want := fmt.Sprintf("%d", pop)
	for i := range servers {
		if servers[i].POP.String() == want {
			return servers[i].Hostname.String()
		}
	}
	return ""
}
