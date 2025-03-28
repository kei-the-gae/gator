package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, fn func(*state, command) error) {
	c.registeredCommands[name] = fn
}

func (c *commands) run(s *state, cmd command) error {
	if fn, ok := c.registeredCommands[cmd.Name]; ok {
		return fn(s, cmd)
	}
	return errors.New("command not found")
}
