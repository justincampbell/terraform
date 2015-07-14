package namecheap

import namecheap "github.com/billputer/go-namecheap"

type Config struct {
	User  string
	Token string
}

// Client() returns a new client for accessing Namecheap
func (c *Config) Client() (*namecheap.Client, error) {
	return namecheap.NewClient(c.User, c.Token, c.User), nil
}
