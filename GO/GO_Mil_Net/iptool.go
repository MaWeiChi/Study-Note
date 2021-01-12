package main

import (
	log "github.com/sirupsen/logrus"
)

// IPToolsInterface ...
type IPToolsInterface interface {
	AddRoute(dest, dev, src string) error
	DelRoute(network string)
	ShowRoute(network string) string
}

// IPTools ...
type IPTools struct {
	cmd    CommanderInterface
	logger *log.Logger
}

// NewIPTools ...
func NewIPTools(cmd CommanderInterface, logger *log.Logger) *IPTools {
	if cmd == nil {
		cmd = CommanderNew(logger)
	}
	return &IPTools{
		cmd,
		logger,
	}
}

// AddRoute : Add a routing rule.
// Args:
//     dest: destination for the routing rule, default for default route.
//     dev: routing device, could be empty
//     src: source for the routing rule, fill "gateway" if dest is "default"
func (t *IPTools) AddRoute(dest string, dev string, src string) error {
	if len(src) <= 0 {
		_, _, err := t.cmd.Run(true, "ip", "route", "add", dest, "dev", dev)
		return err
	}

	if dest == "default" && len(dev) > 0 {
		_, _, err := t.cmd.Run(true, "ip", "route", "add", dest, "dev", dev, "via", src)
		return err
	}

	if dest == "default" {
		_, _, err := t.cmd.Run(true, "ip", "route", "add", dest, "via", src)
		return err
	}

	_, _, err := t.cmd.Run(true, "ip", "route", "add", dest, "dev", dev, "proto", "kernel", "scope", "link", "src", src)
	return err
}

// DelRoute ...
func (t *IPTools) DelRoute(network string) {
	t.cmd.Run(true, "ip", "route", "del", network)
}

// ShowRoute ...
func (t *IPTools) ShowRoute(network string) string {
	out, _, _ := t.cmd.Run(true, "ip", "route", "show", network)
	return out
}
