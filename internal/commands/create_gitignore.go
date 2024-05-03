package commands

import (
	"fmt"
	"os"
	"path/filepath"

	gitignoreio "github.com/kawana77b/tmp4lig/internal/gitignore_io"
	"github.com/kawana77b/tmp4lig/internal/prompt"
	"github.com/logrusorgru/aurora"
)

type CreateGitIgnoreCommand struct {
	saveDir string
}

func (c *CreateGitIgnoreCommand) Set(props CommandProps) {
	c.saveDir = props.SaveDir
}

func (c *CreateGitIgnoreCommand) Execute() error {
	gio := gitignoreio.NewGitignoreIO()
	list, err := gio.GetList()
	if err != nil {
		return err
	}

	answers, err := prompt.NewMultiSelectQuestion().Ask("Select a gitignore", list)
	if err != nil {
		// if the user cancels the prompt, just exit
		return nil
	}

	data, err := gio.GetGitignoreData(answers...)
	if err != nil {
		return err
	}

	savePath := filepath.Join(c.saveDir, ".gitignore")
	os.WriteFile(savePath, data, 0644)

	fmt.Printf("%s %s\n", aurora.Cyan("created:"), savePath)

	return nil
}
