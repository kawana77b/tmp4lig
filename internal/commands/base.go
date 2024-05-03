package commands

import (
	"embed"
)

type CommandProps struct {
	Fs      embed.FS
	SaveDir string
}

type ICommand interface {
	Set(props CommandProps)
	Execute() error
}
