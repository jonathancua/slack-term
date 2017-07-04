package config

import (
	"errors"

	"github.com/gizak/termui"
)

// Config is the definition of a Config struct
type Config struct {
	SlackToken   string                `json:"slack_token"`
	Theme        string                `json:"theme"`
	SidebarWidth int                   `json:"sidebar_width"`
	MainWidth    int                   `json:"-"`
	KeyMap       map[string]keyMapping `json:"key_map"`
}

type keyMapping map[string]string

// NewConfig loads the config file and returns a Config struct
func NewConfig() (*Config, error) {
	cfg := Config{
		SlackToken:   "",
		Theme:        "light",
		SidebarWidth: 1,
		MainWidth:    11,
		KeyMap: map[string]keyMapping{
			"command": {
				"q":          "quit",
			},
			"insert": {
			},
		},
	}

	if cfg.SlackToken == "" {
		return &cfg, errors.New("couldn't find 'slack_token' parameter")
	}

	if cfg.SidebarWidth < 1 || cfg.SidebarWidth > 11 {
		return &cfg, errors.New("please specify the 'sidebar_width' between 1 and 11")
	}

	cfg.MainWidth = 1

	if cfg.Theme == "light" {
		termui.ColorMap = map[string]termui.Attribute{
			"fg":           termui.ColorBlack,
			"bg":           termui.ColorWhite,
			"border.fg":    termui.ColorBlack,
			"label.fg":     termui.ColorBlue,
			"par.fg":       termui.ColorYellow,
			"par.label.bg": termui.ColorWhite,
		}
	}

	return &cfg, nil
}
