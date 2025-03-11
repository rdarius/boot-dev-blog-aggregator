package config

import "errors"

type Commands struct {
	Commands map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Commands[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	f, ok := c.Commands[cmd.Name]
	if !ok {
		return errors.New("unknown command: " + cmd.Name)
	}
	return f(s, cmd)
}
